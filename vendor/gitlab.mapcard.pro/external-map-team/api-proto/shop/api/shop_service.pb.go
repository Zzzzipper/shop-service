// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: shop/api/shop_service.proto

package api

import (
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

var File_shop_api_shop_service_proto protoreflect.FileDescriptor

var file_shop_api_shop_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x68, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa1, 0x02, 0x0a, 0x0b, 0x53,
	0x68, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x41, 0x64,
	0x64, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x41, 0x64, 0x64, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x64, 0x64, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64,
	0x64, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x18,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x68, 0x6f,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x15, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3d,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6d, 0x61, 0x70, 0x63, 0x61, 0x72, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2d, 0x6d, 0x61,
	0x70, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_shop_api_shop_service_proto_goTypes = []interface{}{
	(*AddPartnerRequest)(nil),   // 0: shop.api.AddPartnerRequest
	(*AddMerchantRequest)(nil),  // 1: shop.api.AddMerchantRequest
	(*AddShopRequest)(nil),      // 2: shop.api.AddShopRequest
	(*AuthRequest)(nil),         // 3: shop.api.AuthRequest
	(*AddPartnerResponse)(nil),  // 4: shop.api.AddPartnerResponse
	(*AddMerchantResponse)(nil), // 5: shop.api.AddMerchantResponse
	(*AddShopResponse)(nil),     // 6: shop.api.AddShopResponse
	(*AuthResponse)(nil),        // 7: shop.api.AuthResponse
}
var file_shop_api_shop_service_proto_depIdxs = []int32{
	0, // 0: shop.api.ShopService.AddPartner:input_type -> shop.api.AddPartnerRequest
	1, // 1: shop.api.ShopService.AddMerchant:input_type -> shop.api.AddMerchantRequest
	2, // 2: shop.api.ShopService.AddShop:input_type -> shop.api.AddShopRequest
	3, // 3: shop.api.ShopService.Auth:input_type -> shop.api.AuthRequest
	4, // 4: shop.api.ShopService.AddPartner:output_type -> shop.api.AddPartnerResponse
	5, // 5: shop.api.ShopService.AddMerchant:output_type -> shop.api.AddMerchantResponse
	6, // 6: shop.api.ShopService.AddShop:output_type -> shop.api.AddShopResponse
	7, // 7: shop.api.ShopService.Auth:output_type -> shop.api.AuthResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shop_api_shop_service_proto_init() }
func file_shop_api_shop_service_proto_init() {
	if File_shop_api_shop_service_proto != nil {
		return
	}
	file_shop_api_partner_types_proto_init()
	file_shop_api_merchant_types_proto_init()
	file_shop_api_shop_types_proto_init()
	file_shop_api_auth_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shop_api_shop_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shop_api_shop_service_proto_goTypes,
		DependencyIndexes: file_shop_api_shop_service_proto_depIdxs,
	}.Build()
	File_shop_api_shop_service_proto = out.File
	file_shop_api_shop_service_proto_rawDesc = nil
	file_shop_api_shop_service_proto_goTypes = nil
	file_shop_api_shop_service_proto_depIdxs = nil
}
