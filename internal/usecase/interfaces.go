package usecase

import (
	"context"

	"gitlab.mapcard.pro/external-map-team/shop-service/internal/entity"
)

type (
	Shop interface {
		AddPartner(ctx context.Context, req *entity.Partner) (uint64, error)
		AddMerchant(ctx context.Context, req *entity.Merchant) (uint64, error)
		AddShop(ctx context.Context, req *entity.Shop) (uint64, error)
		Auth(ctx context.Context, login, password string) (*entity.ShopInfo, error)
	}

	ShopRepo interface {
		StorePartner(ctx context.Context, req *entity.Partner) (uint64, error)
		StoreMerchant(ctx context.Context, req *entity.Merchant) (uint64, error)
		StoreShop(ctx context.Context, req *entity.Shop) (uint64, error)
		SelectAuth(ctx context.Context, login, password string) (*entity.ShopInfo, error)
	}
)
