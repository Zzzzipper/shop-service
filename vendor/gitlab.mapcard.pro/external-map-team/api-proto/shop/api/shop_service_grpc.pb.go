// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: shop/api/shop_service.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopServiceClient interface {
	AddPartner(ctx context.Context, in *AddPartnerRequest, opts ...grpc.CallOption) (*Partner, error)
	AddMerchant(ctx context.Context, in *AddMerchantRequest, opts ...grpc.CallOption) (*Merchant, error)
	AddShop(ctx context.Context, in *AddShopRequest, opts ...grpc.CallOption) (*Shop, error)
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*ShopInfo, error)
}

type shopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShopServiceClient(cc grpc.ClientConnInterface) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) AddPartner(ctx context.Context, in *AddPartnerRequest, opts ...grpc.CallOption) (*Partner, error) {
	out := new(Partner)
	err := c.cc.Invoke(ctx, "/shop.api.ShopService/AddPartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) AddMerchant(ctx context.Context, in *AddMerchantRequest, opts ...grpc.CallOption) (*Merchant, error) {
	out := new(Merchant)
	err := c.cc.Invoke(ctx, "/shop.api.ShopService/AddMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) AddShop(ctx context.Context, in *AddShopRequest, opts ...grpc.CallOption) (*Shop, error) {
	out := new(Shop)
	err := c.cc.Invoke(ctx, "/shop.api.ShopService/AddShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*ShopInfo, error) {
	out := new(ShopInfo)
	err := c.cc.Invoke(ctx, "/shop.api.ShopService/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
// All implementations must embed UnimplementedShopServiceServer
// for forward compatibility
type ShopServiceServer interface {
	AddPartner(context.Context, *AddPartnerRequest) (*Partner, error)
	AddMerchant(context.Context, *AddMerchantRequest) (*Merchant, error)
	AddShop(context.Context, *AddShopRequest) (*Shop, error)
	Auth(context.Context, *AuthRequest) (*ShopInfo, error)
	mustEmbedUnimplementedShopServiceServer()
}

// UnimplementedShopServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShopServiceServer struct {
}

func (UnimplementedShopServiceServer) AddPartner(context.Context, *AddPartnerRequest) (*Partner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPartner not implemented")
}
func (UnimplementedShopServiceServer) AddMerchant(context.Context, *AddMerchantRequest) (*Merchant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMerchant not implemented")
}
func (UnimplementedShopServiceServer) AddShop(context.Context, *AddShopRequest) (*Shop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddShop not implemented")
}
func (UnimplementedShopServiceServer) Auth(context.Context, *AuthRequest) (*ShopInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedShopServiceServer) mustEmbedUnimplementedShopServiceServer() {}

// UnsafeShopServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopServiceServer will
// result in compilation errors.
type UnsafeShopServiceServer interface {
	mustEmbedUnimplementedShopServiceServer()
}

func RegisterShopServiceServer(s grpc.ServiceRegistrar, srv ShopServiceServer) {
	s.RegisterService(&ShopService_ServiceDesc, srv)
}

func _ShopService_AddPartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPartnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).AddPartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.api.ShopService/AddPartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).AddPartner(ctx, req.(*AddPartnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_AddMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMerchantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).AddMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.api.ShopService/AddMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).AddMerchant(ctx, req.(*AddMerchantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_AddShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).AddShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.api.ShopService/AddShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).AddShop(ctx, req.(*AddShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.api.ShopService/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopService_ServiceDesc is the grpc.ServiceDesc for ShopService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.api.ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPartner",
			Handler:    _ShopService_AddPartner_Handler,
		},
		{
			MethodName: "AddMerchant",
			Handler:    _ShopService_AddMerchant_Handler,
		},
		{
			MethodName: "AddShop",
			Handler:    _ShopService_AddShop_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _ShopService_Auth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop/api/shop_service.proto",
}
