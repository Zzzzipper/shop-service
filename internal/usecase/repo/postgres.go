package repo

import (
	"context"
	"fmt"
	"time"

	"gitlab.mapcard.pro/external-map-team/shop-service/pkg/custom_error"
	"gitlab.mapcard.pro/external-map-team/shop-service/pkg/logger"
	"gitlab.mapcard.pro/external-map-team/shop-service/pkg/metrics"
	"gorm.io/gorm"

	"gitlab.mapcard.pro/external-map-team/shop-service/internal/entity"
	"gitlab.mapcard.pro/external-map-team/shop-service/pkg/postgres"
)

type ShopRepo struct {
	*postgres.Postgres
	l *logger.Logger
}

func New(pg *postgres.Postgres, l *logger.Logger) *ShopRepo {
	return &ShopRepo{pg, l}
}

func (r *ShopRepo) StorePartner(ctx context.Context, partner *entity.Partner) (uint64, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("StorePartner", float64(time.Since(beginTime).Milliseconds()))
		r.l.Infof("StorePartner time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	err := r.DB.Table("partner").Transaction(func(tx *gorm.DB) error {
		return r.DB.Table("partner").Create(&partner).Error
	})

	if err != nil {
		return 0, fmt.Errorf("ShopRepo - StoreMerchant - error: %w", err)
	}

	return partner.ID, nil
}

func (r *ShopRepo) StoreMerchant(ctx context.Context, merchant *entity.Merchant) (uint64, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("StoreMerchant", float64(time.Since(beginTime).Milliseconds()))
		r.l.Infof("StoreMerchant time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	err := r.DB.Table("merchant").Transaction(func(tx *gorm.DB) error {
		return r.DB.Table("merchant").Create(&merchant).Error
	})

	if err != nil {
		return 0, fmt.Errorf("ShopRepo - StoreMerchant - error: %w", err)
	}

	return merchant.ID, nil
}

func (r *ShopRepo) StoreShop(ctx context.Context, shop *entity.Shop) (uint64, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("StoreShop", float64(time.Since(beginTime).Milliseconds()))
		r.l.Infof("StoreShop time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	err := r.DB.Table("shop").Transaction(func(tx *gorm.DB) error {
		return r.DB.Table("shop").Create(&shop).Error
	})

	if err != nil {
		return 0, fmt.Errorf("ShopRepo - StoreShop - error: %w", err)
	}

	return shop.ID, nil
}

func (r *ShopRepo) SelectAuth(ctx context.Context, login, password string) (*entity.ShopInfo, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("SelectAuth", float64(time.Since(beginTime).Milliseconds()))
		r.l.Infof("SelectAuth time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	var shopInfo entity.ShopInfo

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		var count int64
		err := r.DB.
			Select(`p.id as partner_id, m.id as merchant_id, s.id as shop_id, s.settings as settings`).
			Table(`shop s, merchant m, partner p`).
			Where(`s.merchant_id=m.id AND m.partner_id=p.id AND s.login=? AND s.password=?`, login, password).
			Find(&shopInfo).Count(&count).Error
		if count == 0 {
			return custom_error.New("INVALID_CREDENTIALS", "Профиль не найден")
		}
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("ShopRepo - SelectAuth - error: %w", err)
	}

	return &shopInfo, nil
}
