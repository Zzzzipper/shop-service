// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: proto/merchant.proto

package proto

import (
	api "gitlab.mapcard.pro/external-map-team/api-proto/merchant/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_merchant_proto protoreflect.FileDescriptor

var file_proto_merchant_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x68, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xd7, 0x07, 0x0a, 0x0f, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x22, 0x00, 0x12,
	0x4c, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x12, 0x22, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x22, 0x00, 0x12, 0x4c, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x21, 0x2e,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x49, 0x0a, 0x0b, 0x41,
	0x64, 0x64, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3d, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x53,
	0x68, 0x6f, 0x70, 0x12, 0x1c, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x68, 0x6f, 0x70, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x1f, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x09,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x12, 0x1e, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x6f,
	0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x49, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c,
	0x12, 0x20, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x64, 0x64, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x23,
	0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x22, 0x00, 0x12, 0x4f, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x12, 0x22,
	0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3d,
	0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x19, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x43, 0x5a,
	0x41, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6d, 0x61, 0x70, 0x63, 0x61, 0x72, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2d, 0x6d, 0x61, 0x70,
	0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_merchant_proto_goTypes = []interface{}{
	(*api.AddPartnerRequest)(nil),     // 0: merchant.api.AddPartnerRequest
	(*api.DeletePartnerRequest)(nil),  // 1: merchant.api.DeletePartnerRequest
	(*api.ListPartnersRequest)(nil),   // 2: merchant.api.ListPartnersRequest
	(*api.AddMerchantRequest)(nil),    // 3: merchant.api.AddMerchantRequest
	(*api.DeleteMerchantRequest)(nil), // 4: merchant.api.DeleteMerchantRequest
	(*api.ListMerchantsRequest)(nil),  // 5: merchant.api.ListMerchantsRequest
	(*api.AddShopRequest)(nil),        // 6: merchant.api.AddShopRequest
	(*api.DeleteShopRequest)(nil),     // 7: merchant.api.DeleteShopRequest
	(*api.ListShopsRequest)(nil),      // 8: merchant.api.ListShopsRequest
	(*api.AddTerminalRequest)(nil),    // 9: merchant.api.AddTerminalRequest
	(*api.DeleteTerminalRequest)(nil), // 10: merchant.api.DeleteTerminalRequest
	(*api.ListTerminalsRequest)(nil),  // 11: merchant.api.ListTerminalsRequest
	(*api.AuthRequest)(nil),           // 12: merchant.api.AuthRequest
	(*api.Partner)(nil),               // 13: merchant.api.Partner
	(*api.Merchant)(nil),              // 14: merchant.api.Merchant
	(*api.Shop)(nil),                  // 15: merchant.api.Shop
	(*api.Terminal)(nil),              // 16: merchant.api.Terminal
	(*api.SellerInfo)(nil),            // 17: merchant.api.SellerInfo
}
var file_proto_merchant_proto_depIdxs = []int32{
	0,  // 0: proto.MerchantService.AddPartner:input_type -> merchant.api.AddPartnerRequest
	1,  // 1: proto.MerchantService.DeletePartner:input_type -> merchant.api.DeletePartnerRequest
	2,  // 2: proto.MerchantService.ListPartners:input_type -> merchant.api.ListPartnersRequest
	3,  // 3: proto.MerchantService.AddMerchant:input_type -> merchant.api.AddMerchantRequest
	4,  // 4: proto.MerchantService.DeleteMerchant:input_type -> merchant.api.DeleteMerchantRequest
	5,  // 5: proto.MerchantService.ListMerchants:input_type -> merchant.api.ListMerchantsRequest
	6,  // 6: proto.MerchantService.AddShop:input_type -> merchant.api.AddShopRequest
	7,  // 7: proto.MerchantService.DeleteShop:input_type -> merchant.api.DeleteShopRequest
	8,  // 8: proto.MerchantService.ListShops:input_type -> merchant.api.ListShopsRequest
	9,  // 9: proto.MerchantService.AddTerminal:input_type -> merchant.api.AddTerminalRequest
	10, // 10: proto.MerchantService.DeleteTerminal:input_type -> merchant.api.DeleteTerminalRequest
	11, // 11: proto.MerchantService.ListTerminals:input_type -> merchant.api.ListTerminalsRequest
	12, // 12: proto.MerchantService.Auth:input_type -> merchant.api.AuthRequest
	13, // 13: proto.MerchantService.AddPartner:output_type -> merchant.api.Partner
	13, // 14: proto.MerchantService.DeletePartner:output_type -> merchant.api.Partner
	13, // 15: proto.MerchantService.ListPartners:output_type -> merchant.api.Partner
	14, // 16: proto.MerchantService.AddMerchant:output_type -> merchant.api.Merchant
	14, // 17: proto.MerchantService.DeleteMerchant:output_type -> merchant.api.Merchant
	14, // 18: proto.MerchantService.ListMerchants:output_type -> merchant.api.Merchant
	15, // 19: proto.MerchantService.AddShop:output_type -> merchant.api.Shop
	15, // 20: proto.MerchantService.DeleteShop:output_type -> merchant.api.Shop
	15, // 21: proto.MerchantService.ListShops:output_type -> merchant.api.Shop
	16, // 22: proto.MerchantService.AddTerminal:output_type -> merchant.api.Terminal
	16, // 23: proto.MerchantService.DeleteTerminal:output_type -> merchant.api.Terminal
	16, // 24: proto.MerchantService.ListTerminals:output_type -> merchant.api.Terminal
	17, // 25: proto.MerchantService.Auth:output_type -> merchant.api.SellerInfo
	13, // [13:26] is the sub-list for method output_type
	0,  // [0:13] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_merchant_proto_init() }
func file_proto_merchant_proto_init() {
	if File_proto_merchant_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_merchant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_merchant_proto_goTypes,
		DependencyIndexes: file_proto_merchant_proto_depIdxs,
	}.Build()
	File_proto_merchant_proto = out.File
	file_proto_merchant_proto_rawDesc = nil
	file_proto_merchant_proto_goTypes = nil
	file_proto_merchant_proto_depIdxs = nil
}
