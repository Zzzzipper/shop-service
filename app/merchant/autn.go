package merchant

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"

	"gitlab.mapcard.pro/external-map-team/api-proto/merchant/api"
)

// Auth get seller info
func (d Directory) Auth(ctx context.Context, req *api.AuthRequest) (*api.SellerInfo, error) {
	authRequest := AuthParams{
		Login:    req.Login,
		Password: req.Password,
	}
	sellerInfo, err := d.querier.Auth(ctx, authRequest)
	if err != nil {
		Log().Format("Error Auth: %s\n", err.Error())
		if strings.Contains(err.Error(), "no rows in result set") {
			return authinfoPostgresToProto(AuthRow{})
		} else {
			return nil, Log().StatusErrorf(codes.Internal, "Unexpected error select auth: %s", err.Error())
		}
	}

	return authinfoPostgresToProto(sellerInfo)

}
