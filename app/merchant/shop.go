package merchant

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	merchpb "merchant-service/proto"
)

// AddShop adds a shop to the directory.
func (d Directory) AddShop(ctx context.Context, req *merchpb.AddShopRequest) (*merchpb.Shop, error) {
	var merchantID pgtype.UUID
	err := merchantID.Set(req.GetMerchantId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid merchant UUID provided")
	}
	pgShop, err := d.querier.AddShop(ctx, AddShopParams{
		FullName:   req.FullName,
		Url:        req.Url,
		MerchantID: merchantID,
		Login:      req.Login,
		Password:   req.Password,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error adding shop: %s", err.Error())
	}
	return shopPostgresToProto(pgShop)
}

// DeleteShop deletes the shop, if found.
func (d Directory) DeleteShop(ctx context.Context, req *merchpb.DeleteShopRequest) (*merchpb.Shop, error) {
	var shopID pgtype.UUID
	err := shopID.Set(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid shop UUID provided")
	}
	pgShop, err := d.querier.DeleteShop(ctx, shopID)
	if err != nil {
		return nil, err
	}
	return shopPostgresToProto(pgShop)
}

// ListPartners lists shops in the directory, subject to the request filters.
func (d Directory) ListShops(req *merchpb.ListShopsRequest, srv merchpb.MerchantService_ListShopsServer) (retErr error) {
	q := d.sb.Select(
		"id",
		"create_time",
		"full_name",
		"url",
		"merchant_id",
		"login",
		"password",
	).From(
		"shops",
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

	if req.GetMerchantId() != "" {
		var merchantID pgtype.UUID
		err := merchantID.Set(req.GetMerchantId())
		if err != nil {
			return status.Error(codes.InvalidArgument, "invalid merchant UUID provided")
		}
		q = q.Where(
			squirrel.Expr(
				"merchant_id = ?", merchantID,
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
		var pgShop Shop
		err := rows.Scan(
			&pgShop.ID,
			&pgShop.CreateTime,
			&pgShop.FullName,
			&pgShop.Url,
			&pgShop.MerchantID,
			&pgShop.Password,
			&pgShop.Login,
		)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		protoShop, err := shopPostgresToProto(pgShop)
		if err != nil {
			return err
		}
		err = srv.Send(protoShop)
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
