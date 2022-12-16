package merchant

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"google.golang.org/grpc/codes"

	merchpb "gitlab.mapcard.pro/external-map-team/merchant-service/app/proto"

	"gitlab.mapcard.pro/external-map-team/api-proto/merchant/api"
)

// AddPartner adds a partner to the directory.
func (d Directory) AddPartner(ctx context.Context, req *api.AddPartnerRequest) (*api.Partner, error) {
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
		return nil, Log().StatusErrorf(codes.Internal, "Unexpected error adding partner: %s", err.Error())
	}
	return partnerPostgresToProto(pgPartner)
}

// DeletePartner deletes the partner, if found.
func (d Directory) DeletePartner(ctx context.Context, req *api.DeletePartnerRequest) (*api.Partner, error) {
	var partnerID pgtype.UUID
	err := partnerID.Set(req.GetId())
	if err != nil {
		return nil, Log().StatusError(codes.InvalidArgument, "Invalid UUID provided")
	}
	pgPartner, err := d.querier.DeletePartner(ctx, partnerID)
	if err != nil {
		return nil, err
	}
	return partnerPostgresToProto(pgPartner)
}

// ListPartners lists partners in the directory, subject to the request filters.
func (d Directory) ListPartners(req *api.ListPartnersRequest, srv merchpb.MerchantService_ListPartnersServer) (retErr error) {
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
			return Log().StatusErrorf(codes.InvalidArgument, "Invalid timestamp: %s", err.Error())
		}
		q = q.Where(squirrel.Gt{
			"create_time": pgTime,
		})
	}

	if req.GetOlderThan() != nil {
		var pgInterval pgtype.Interval
		err := pgInterval.Set(req.GetOlderThan().AsDuration())
		if err != nil {
			return Log().StatusErrorf(codes.InvalidArgument, "Invalid duration: %s", err.Error())
		}
		q = q.Where(
			squirrel.Expr(
				"CURRENT_TIMESTAMP - create_time > ?", pgInterval,
			),
		)
	}

	rows, retErr := q.QueryContext(srv.Context())
	if retErr != nil {
		return Log().StatusError(codes.Internal, retErr.Error())
	}
	defer func() {
		cerr := rows.Close()
		if retErr == nil && cerr != nil {
			retErr = Log().StatusError(codes.Internal, cerr.Error())
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
			return Log().StatusError(codes.Internal, err.Error())
		}
		protoPartner, err := partnerPostgresToProto(pgPartner)
		if err != nil {
			return err
		}
		err = srv.Send(protoPartner)
		if err != nil {
			return Log().StatusError(codes.Internal, err.Error())
		}
	}

	retErr = rows.Err()
	if retErr != nil {
		return Log().StatusError(codes.Internal, retErr.Error())
	}

	return nil
}
