package merchant

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	merchpb "merchant-service/proto"

	"gitlab.mapcard.pro/external-map-team/api-proto/merchant/api"
)

// AddMerchant adds a merchant to the directory.
func (d Directory) AddMerchant(ctx context.Context, req *api.AddMerchantRequest) (*api.Merchant, error) {
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
func (d Directory) DeleteMerchant(ctx context.Context, req *api.DeleteMerchantRequest) (*api.Merchant, error) {
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
func (d Directory) ListMerchants(req *api.ListMerchantsRequest, srv merchpb.MerchantService_ListMerchantsServer) (retErr error) {
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
