package merchant

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	merchpb "merchant-service/proto"
)

// Directory stores a directory of partners.
type Directory struct {
	logger  *logrus.Logger
	db      *sql.DB
	sb      squirrel.StatementBuilderType
	querier Querier
}

// NewDirectory creates a new Directory, connecting it to the postgres server on
// the URL provided.
func NewDirectory(logger *logrus.Logger, pgURL *url.URL) (*Directory, error) {
	connURL := *pgURL
	if connURL.Scheme == "cockroachdb" {
		// Overwrite the scheme before parsing with pgx, since
		// it doesn't support the "cockroachdb" scheme.
		connURL.Scheme = "merchant"
	}
	c, err := pgx.ParseConfig(connURL.String())
	if err != nil {
		return nil, fmt.Errorf("parsing merchant URI: %w", err)
	}

	c.Logger = logrusadapter.NewLogger(logger)
	db := stdlib.OpenDB(*c)

	err = validateSchema(db, pgURL.Scheme)
	if err != nil {
		return nil, fmt.Errorf("validating schema: %w", err)
	}

	return &Directory{
		logger:  logger,
		db:      db,
		sb:      squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(db),
		querier: New(db),
	}, nil
}

// Close releases any resources.
func (d Directory) Close() error {
	return d.db.Close()
}

// AddPartner adds a partner to the directory.
func (d Directory) AddPartner(ctx context.Context, req *merchpb.AddPartnerRequest) (*merchpb.Partner, error) {
	pgRole, err := roleProtoToPostgres(req.Role)
	if err != nil {
		return nil, err
	}
	pgPartner, err := d.querier.AddPartner(ctx, AddPartnerParams{
		FullName: req.FullName,
		Url:      req.Url,
		Role:     pgRole,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error adding partner: %s", err.Error())
	}
	return partnerPostgresToProto(pgPartner)
}

// DeletePartner deletes the partner, if found.
func (d Directory) DeletePartner(ctx context.Context, req *merchpb.DeletePartnerRequest) (*merchpb.Partner, error) {
	var partnerID pgtype.UUID
	err := partnerID.Set(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid UUID provided")
	}
	pgPartner, err := d.querier.DeletePartner(ctx, partnerID)
	if err != nil {
		return nil, err
	}
	return partnerPostgresToProto(pgPartner)
}

// ListPartners lists partners in the directory, subject to the request filters.
func (d Directory) ListPartners(req *merchpb.ListPartnersRequest, srv merchpb.MerchantService_ListPartnersServer) (retErr error) {
	q := d.sb.Select(
		"id",
		"create_time",
		"full_name",
		"url",
		"api_token",
		"role",
	).From(
		"partners",
	).OrderBy(
		"create_time ASC",
	)

	if req.GetCreatedSince() != nil {
		var pgTime pgtype.Timestamptz
		err := pgTime.Set(req.GetCreatedSince().AsTime())
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "invalid timestamp: %s", err.Error())
		}
		q = q.Where(squirrel.Gt{
			"create_time": pgTime,
		})
	}

	if req.GetOlderThan() != nil {
		var pgInterval pgtype.Interval
		err := pgInterval.Set(req.GetOlderThan().AsDuration())
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "invalid duration: %s", err.Error())
		}
		q = q.Where(
			squirrel.Expr(
				"CURRENT_TIMESTAMP - create_time > ?", pgInterval,
			),
		)
	}

	rows, retErr := q.QueryContext(srv.Context())
	if retErr != nil {
		return status.Error(codes.Internal, retErr.Error())
	}
	defer func() {
		cerr := rows.Close()
		if retErr == nil && cerr != nil {
			retErr = status.Error(codes.Internal, cerr.Error())
		}
	}()

	for rows.Next() {
		var pgPartner Partner
		err := rows.Scan(
			&pgPartner.ID,
			&pgPartner.CreateTime,
			&pgPartner.FullName,
			&pgPartner.Url,
			&pgPartner.ApiToken,
			&pgPartner.Role,
		)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		protoPartner, err := partnerPostgresToProto(pgPartner)
		if err != nil {
			return err
		}
		err = srv.Send(protoPartner)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	retErr = rows.Err()
	if retErr != nil {
		return status.Error(codes.Internal, retErr.Error())
	}

	return nil
}

// AddMerchant adds a merchant to the directory.
func (d Directory) AddMerchant(ctx context.Context, req *merchpb.AddMerchantRequest) (*merchpb.Merchant, error) {
	var partnerID pgtype.UUID
	err := partnerID.Set(req.GetPartnerId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid partner UUID provided")
	}
	pgMerchant, err := d.querier.AddMerchant(ctx, AddMerchantParams{
		FullName:  req.FullName,
		Url:       req.Url,
		PartnerID: partnerID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error adding merchant: %s", err.Error())
	}
	return merchantPostgresToProto(pgMerchant)
}

// DeleteMerchant deletes the merchant, if found.
func (d Directory) DeleteMerchant(ctx context.Context, req *merchpb.DeleteMerchantRequest) (*merchpb.Merchant, error) {
	var merchantID pgtype.UUID
	err := merchantID.Set(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid merchant UUID provided")
	}
	pgMerchant, err := d.querier.DeleteMerchant(ctx, merchantID)
	if err != nil {
		return nil, err
	}
	return merchantPostgresToProto(pgMerchant)
}

// ListPartners lists partners in the directory, subject to the request filters.
func (d Directory) ListMerchants(req *merchpb.ListMerchantsRequest, srv merchpb.MerchantService_ListMerchantsServer) (retErr error) {
	q := d.sb.Select(
		"id",
		"create_time",
		"full_name",
		"url",
		"partner_id",
	).From(
		"merchants",
	).OrderBy(
		"create_time ASC",
	)

	if req.GetCreatedSince() != nil {
		var pgTime pgtype.Timestamptz
		err := pgTime.Set(req.GetCreatedSince().AsTime())
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "invalid timestamp: %s", err.Error())
		}
		q = q.Where(squirrel.Gt{
			"create_time": pgTime,
		})
	}

	if req.GetOlderThan() != nil {
		var pgInterval pgtype.Interval
		err := pgInterval.Set(req.GetOlderThan().AsDuration())
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "invalid duration: %s", err.Error())
		}
		q = q.Where(
			squirrel.Expr(
				"CURRENT_TIMESTAMP - create_time > ?", pgInterval,
			),
		)
	}

	if req.GetPartnerId() != "" {
		var merchantID pgtype.UUID
		err := merchantID.Set(req.GetPartnerId())
		if err != nil {
			return status.Error(codes.InvalidArgument, "invalid partner UUID provided")
		}
		q = q.Where(
			squirrel.Expr(
				"partner_id = ?", merchantID,
			),
		)
	}

	rows, retErr := q.QueryContext(srv.Context())
	if retErr != nil {
		return status.Error(codes.Internal, retErr.Error())
	}
	defer func() {
		cerr := rows.Close()
		if retErr == nil && cerr != nil {
			retErr = status.Error(codes.Internal, cerr.Error())
		}
	}()

	for rows.Next() {
		var pgMerchant Merchant
		err := rows.Scan(
			&pgMerchant.ID,
			&pgMerchant.CreateTime,
			&pgMerchant.FullName,
			&pgMerchant.Url,
			&pgMerchant.PartnerID,
		)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		protoMerchant, err := merchantPostgresToProto(pgMerchant)
		if err != nil {
			return err
		}
		err = srv.Send(protoMerchant)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	retErr = rows.Err()
	if retErr != nil {
		return status.Error(codes.Internal, retErr.Error())
	}

	return nil
}
