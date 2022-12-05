package merchant

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	merchpb "gitlab.mapcard.pro/external-map-team/merchant-service/app/proto"

	"gitlab.mapcard.pro/external-map-team/api-proto/merchant/api"
)

// AddTerminal adds a terminal to the directory.
func (d Directory) AddTerminal(ctx context.Context, req *api.AddTerminalRequest) (*api.Terminal, error) {
	var shopID pgtype.UUID
	err := shopID.Set(req.GetShopId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid shop UUID provided")
	}
	pgTerminal, err := d.querier.AddTerminal(ctx, AddTerminalParams{
		FullName: req.FullName,
		Url:      req.Url,
		ShopID:   shopID,
		Login:    req.Login,
		Password: req.Password,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error adding terminal: %s", err.Error())
	}
	return terminalPostgresToProto(pgTerminal)
}

// DeleteTerminal deletes the terminal, if found.
func (d Directory) DeleteTerminal(ctx context.Context, req *api.DeleteTerminalRequest) (*api.Terminal, error) {
	var terminalID pgtype.UUID
	err := terminalID.Set(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid terminal UUID provided")
	}
	pgTerminal, err := d.querier.DeleteTerminal(ctx, terminalID)
	if err != nil {
		return nil, err
	}
	return terminalPostgresToProto(pgTerminal)
}

// ListPartners lists terminals in the directory, subject to the request filters.
func (d Directory) ListTerminals(req *api.ListTerminalsRequest, srv merchpb.MerchantService_ListTerminalsServer) (retErr error) {
	q := d.sb.Select(
		"id",
		"create_time",
		"full_name",
		"url",
		"shop_id",
		"login",
		"password",
	).From(
		"terminals",
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

	if req.GetShopId() != "" {
		var merchantID pgtype.UUID
		err := merchantID.Set(req.GetShopId())
		if err != nil {
			return status.Error(codes.InvalidArgument, "invalid shop UUID provided")
		}
		q = q.Where(
			squirrel.Expr(
				"shop_id = ?", merchantID,
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
		var pgTerminal Terminal
		err := rows.Scan(
			&pgTerminal.ID,
			&pgTerminal.CreateTime,
			&pgTerminal.FullName,
			&pgTerminal.Url,
			&pgTerminal.ShopID,
			&pgTerminal.Password,
			&pgTerminal.Login,
		)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		protoTerminal, err := terminalPostgresToProto(pgTerminal)
		if err != nil {
			return err
		}
		err = srv.Send(protoTerminal)
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
