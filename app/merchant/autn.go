package merchant

import (
	"context"
	"fmt"
	"strings"

	merchpb "merchant-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Auth get seller info
func (d Directory) Auth(ctx context.Context, req *merchpb.AuthRequest) (*merchpb.SellerInfo, error) {
	authRequest := AuthParams{
		Login:    req.Login,
		Password: req.Password,
	}
	sellerInfo, err := d.querier.Auth(ctx, authRequest)
	if err != nil {
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), "no rows in result set") {
			return authinfoPostgresToProto(AuthRow{})
		} else {
			return nil, status.Errorf(codes.Internal, "unexpected error select auth: %s", err.Error())
		}
	}

	return authinfoPostgresToProto(sellerInfo)

}
