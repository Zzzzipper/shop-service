package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	shop_api "gitlab.mapcard.pro/external-map-team/api-proto/shop/api"
	"gitlab.mapcard.pro/external-map-team/shop-service/internal/entity"
	"gitlab.mapcard.pro/external-map-team/shop-service/internal/usecase"
	"gitlab.mapcard.pro/external-map-team/shop-service/pkg/custom_error"
	"gitlab.mapcard.pro/external-map-team/shop-service/pkg/logger"
	"gitlab.mapcard.pro/external-map-team/shop-service/pkg/metrics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	shop_api.UnimplementedShopServiceServer
	logger *logger.Logger
	shop   usecase.Shop
}

func NewServer(logger *logger.Logger, shop usecase.Shop) *Server {
	return &Server{
		logger: logger,
		shop:   shop,
	}
}

func (s *Server) Start(port string) error {
	addr := fmt.Sprintf(":%s", port)

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		s.logger.Errorf(err)
		return err
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	shop_api.RegisterShopServiceServer(grpcServer, s)

	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	s.logger.Infof("Start serve GRPC at :%s", port)

	go grpcServer.Serve(listener)

	return nil
}

func (s *Server) AddPartner(ctx context.Context, req *shop_api.AddPartnerRequest) (*shop_api.AddPartnerResponse, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("AddPartner GRPC", float64(time.Since(beginTime).Milliseconds()))
		s.logger.Infof("AddPartner GRPC time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	s.logger.Infof("got grpc query AddPartner")

	partner := &entity.Partner{
		FullName: req.FullName,
		Url:      req.Url,
		Role:     int32(req.Role),
	}

	newId, err := s.shop.AddPartner(ctx, partner)

	if err != nil {
		s.logger.Errorf("s.shop.AddPartner(ctx, partner) - %w", err)
		return &shop_api.AddPartnerResponse{
			Status:     false,
			ErrCode:    "INTERNAL_ERROR",
			ErrMessage: err.Error(),
		}, nil
	}

	response := shop_api.AddPartnerResponse{
		Status:  true,
		ErrCode: "OK",
		Partner: &shop_api.Partner{
			Id:         newId,
			CreateTime: timestamppb.New(partner.CreateTime),
			FullName:   partner.FullName,
			Url:        partner.Url,
			ApiToken:   partner.ApiToken,
			Role:       shop_api.Role(partner.Role),
		},
	}

	return &response, nil
}

func (s *Server) AddMerchant(ctx context.Context, req *shop_api.AddMerchantRequest) (*shop_api.AddMerchantResponse, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("AddMerchant GRPC", float64(time.Since(beginTime).Milliseconds()))
		s.logger.Infof("AddMerchant GRPC time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	s.logger.Infof("got grpc query AddMerchant")

	merchant := &entity.Merchant{
		FullName:  req.FullName,
		Url:       req.Url,
		PartnerID: req.PartnerId,
	}

	newId, err := s.shop.AddMerchant(ctx, merchant)

	if err != nil {
		s.logger.Errorf("s.shop.AddMerchant(ctx, merchant) - %w", err)
		return &shop_api.AddMerchantResponse{
			Status:     false,
			ErrCode:    "INTERNAL_ERROR",
			ErrMessage: err.Error(),
		}, nil
	}

	response := shop_api.AddMerchantResponse{
		Status:  true,
		ErrCode: "OK",
		Merchant: &shop_api.Merchant{
			Id:         newId,
			CreateTime: timestamppb.New(merchant.CreateTime),
			FullName:   merchant.FullName,
			Url:        merchant.Url,
			PartnerId:  merchant.PartnerID,
		},
	}

	return &response, nil

}

func (s *Server) AddShop(ctx context.Context, req *shop_api.AddShopRequest) (*shop_api.AddShopResponse, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("AddShop GRPC", float64(time.Since(beginTime).Milliseconds()))
		s.logger.Infof("AddShop GRPC time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	s.logger.Infof("got grpc query AddShop")

	settingsBytes, err := json.Marshal(req.Settings)
	if err != nil {
		s.logger.Errorf("json.Marshal(req.Settings) - %w", err)
		return &shop_api.AddShopResponse{
			Status:     false,
			ErrCode:    "INTERNAL_ERROR",
			ErrMessage: err.Error(),
		}, nil
	}

	shop := &entity.Shop{
		FullName:   req.FullName,
		Url:        req.Url,
		MerchantID: req.MerchantId,
		Settings:   string(settingsBytes),
		Login:      req.Login,
		Password:   req.Password,
	}

	newId, err := s.shop.AddShop(ctx, shop)

	if err != nil {
		s.logger.Errorf("s.shop.AddShop(ctx, shop) - %w", err)
		return &shop_api.AddShopResponse{
			Status:     false,
			ErrCode:    "INTERNAL_ERROR",
			ErrMessage: err.Error(),
		}, nil
	}

	var settings shop_api.Settings
	err = json.Unmarshal([]byte(shop.Settings), &settings)

	if err != nil {
		s.logger.Errorf("json.Unmarshal([]byte(shop.Settings), &settings) - %w", err)
		return &shop_api.AddShopResponse{
			Status:     false,
			ErrCode:    "INTERNAL_ERROR",
			ErrMessage: err.Error(),
		}, nil
	}

	response := shop_api.AddShopResponse{
		Status:  true,
		ErrCode: "OK",
		Shop: &shop_api.Shop{
			Id:         newId,
			CreateTime: timestamppb.New(shop.CreateTime),
			FullName:   shop.FullName,
			Url:        shop.Url,
			MerchantId: shop.MerchantID,
			Login:      shop.Login,
			Password:   shop.Password,
			Settings:   &settings,
		},
	}

	return &response, nil
}

func (s *Server) Auth(ctx context.Context, req *shop_api.AuthRequest) (*shop_api.AuthResponse, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("Auth GRPC", float64(time.Since(beginTime).Milliseconds()))
		s.logger.Infof("Auth GRPC time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	s.logger.Infof("got grpc query Auth")

	shopInfo, err := s.shop.Auth(ctx, req.Login, req.Password)

	if err != nil {
		s.logger.Errorf("Auth: s.shop.Auth(ctx, req.Login, req.Password) - %w", err)
		switch e := err.(type) {
		case *custom_error.CustomError:
			return &shop_api.AuthResponse{
				Status:     false,
				ErrCode:    e.ErrCode,
				ErrMessage: e.ErrMessage,
			}, nil
		default:
			return &shop_api.AuthResponse{
				Status:     false,
				ErrCode:    "INTERNAL_ERROR",
				ErrMessage: e.Error(),
			}, nil
		}
	}

	var settings shop_api.Settings
	err = json.Unmarshal([]byte(shopInfo.Settings), &settings)

	if err != nil {
		s.logger.Errorf("Auth: json.Unmarshal([]byte(shopInfo.Settings), &settings) - %w", err)
		return &shop_api.AuthResponse{
			Status:     false,
			ErrCode:    "INTERNAL_ERROR",
			ErrMessage: err.Error(),
		}, nil
	}

	response := shop_api.AuthResponse{
		Status:  true,
		ErrCode: "OK",
		ShopInfo: &shop_api.ShopInfo{
			ShopId:     shopInfo.ShopId,
			MerchantId: shopInfo.MerchantId,
			PartnerId:  shopInfo.MerchantId,
			Settings:   &settings,
		},
	}

	return &response, nil
}
