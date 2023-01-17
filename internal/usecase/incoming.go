package usecase

import (
	"context"
	"time"

	"gitlab.mapcard.pro/external-map-team/shop-service/internal/entity"
	"gitlab.mapcard.pro/external-map-team/shop-service/pkg/logger"
	"gitlab.mapcard.pro/external-map-team/shop-service/pkg/metrics"
)

type ShopUseCase struct {
	logger *logger.Logger
	repo   ShopRepo
}

func New(logger *logger.Logger, r ShopRepo) *ShopUseCase {
	return &ShopUseCase{
		logger: logger,
		repo:   r,
	}
}

func (u *ShopUseCase) AddPartner(ctx context.Context, partner *entity.Partner) (uint64, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("AddPartner", float64(time.Since(beginTime).Milliseconds()))
		u.logger.Infof("AddPartner time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	return u.repo.StorePartner(ctx, partner)
}

func (u *ShopUseCase) AddMerchant(ctx context.Context, merchant *entity.Merchant) (uint64, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("AddMerchant", float64(time.Since(beginTime).Milliseconds()))
		u.logger.Infof("AddMerchant time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	return u.repo.StoreMerchant(ctx, merchant)
}

func (u *ShopUseCase) AddShop(ctx context.Context, shop *entity.Shop) (uint64, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("AddShop", float64(time.Since(beginTime).Milliseconds()))
		u.logger.Infof("AddShop time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	return u.repo.StoreShop(ctx, shop)
}

func (u *ShopUseCase) Auth(ctx context.Context, login, password string) (*entity.ShopInfo, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("Auth", float64(time.Since(beginTime).Milliseconds()))
		u.logger.Infof("Auth time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	return u.repo.SelectAuth(ctx, login, password)
}
