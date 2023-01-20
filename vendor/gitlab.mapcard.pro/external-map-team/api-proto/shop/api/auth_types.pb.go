// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: shop/api/auth_types.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ShopInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Индекс магазина
	ShopId uint64 `protobuf:"varint,1,opt,name=ShopId,proto3" json:"ShopId,omitempty"`
	// Индекс мерчанта
	MerchantId uint64 `protobuf:"varint,2,opt,name=MerchantId,proto3" json:"MerchantId,omitempty"`
	// Индекс партнера
	PartnerId uint64 `protobuf:"varint,3,opt,name=PartnerId,proto3" json:"PartnerId,omitempty"`
	// Параметры провайдера
	Settings *Settings `protobuf:"bytes,4,opt,name=Settings,proto3" json:"Settings,omitempty"`
}

func (x *ShopInfo) Reset() {
	*x = ShopInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_api_auth_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopInfo) ProtoMessage() {}

func (x *ShopInfo) ProtoReflect() protoreflect.Message {
	mi := &file_shop_api_auth_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopInfo.ProtoReflect.Descriptor instead.
func (*ShopInfo) Descriptor() ([]byte, []int) {
	return file_shop_api_auth_types_proto_rawDescGZIP(), []int{0}
}

func (x *ShopInfo) GetShopId() uint64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *ShopInfo) GetMerchantId() uint64 {
	if x != nil {
		return x.MerchantId
	}
	return 0
}

func (x *ShopInfo) GetPartnerId() uint64 {
	if x != nil {
		return x.PartnerId
	}
	return 0
}

func (x *ShopInfo) GetSettings() *Settings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Логин
	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	// Пароль авторизации
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Токен доступа
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_api_auth_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_api_auth_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_shop_api_auth_types_proto_rawDescGZIP(), []int{1}
}

func (x *AuthRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *AuthRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AuthRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Статус операции: true/false
	Status bool `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	// Текстовый код ошибки
	ErrCode string `protobuf:"bytes,2,opt,name=ErrCode,proto3" json:"ErrCode,omitempty"`
	// Описание ошибки.
	ErrMessage string `protobuf:"bytes,3,opt,name=ErrMessage,proto3" json:"ErrMessage,omitempty"`
	// Объект ShopInfo
	ShopInfo *ShopInfo `protobuf:"bytes,4,opt,name=ShopInfo,proto3" json:"ShopInfo,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_api_auth_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_api_auth_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_shop_api_auth_types_proto_rawDescGZIP(), []int{2}
}

func (x *AuthResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *AuthResponse) GetErrCode() string {
	if x != nil {
		return x.ErrCode
	}
	return ""
}

func (x *AuthResponse) GetErrMessage() string {
	if x != nil {
		return x.ErrMessage
	}
	return ""
}

func (x *AuthResponse) GetShopInfo() *ShopInfo {
	if x != nil {
		return x.ShopInfo
	}
	return nil
}

var File_shop_api_auth_types_proto protoreflect.FileDescriptor

var file_shop_api_auth_types_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x19, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x68, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x90, 0x01, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x53,
	0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x4d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x22, 0x55, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x90, 0x01, 0x0a, 0x0c, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x45, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x45, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a,
	0x08, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x3d, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6d, 0x61, 0x70, 0x63, 0x61, 0x72, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2d, 0x6d, 0x61, 0x70,
	0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shop_api_auth_types_proto_rawDescOnce sync.Once
	file_shop_api_auth_types_proto_rawDescData = file_shop_api_auth_types_proto_rawDesc
)

func file_shop_api_auth_types_proto_rawDescGZIP() []byte {
	file_shop_api_auth_types_proto_rawDescOnce.Do(func() {
		file_shop_api_auth_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_api_auth_types_proto_rawDescData)
	})
	return file_shop_api_auth_types_proto_rawDescData
}

var file_shop_api_auth_types_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_shop_api_auth_types_proto_goTypes = []interface{}{
	(*ShopInfo)(nil),     // 0: shop.api.ShopInfo
	(*AuthRequest)(nil),  // 1: shop.api.AuthRequest
	(*AuthResponse)(nil), // 2: shop.api.AuthResponse
	(*Settings)(nil),     // 3: shop.api.Settings
}
var file_shop_api_auth_types_proto_depIdxs = []int32{
	3, // 0: shop.api.ShopInfo.Settings:type_name -> shop.api.Settings
	0, // 1: shop.api.AuthResponse.ShopInfo:type_name -> shop.api.ShopInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_shop_api_auth_types_proto_init() }
func file_shop_api_auth_types_proto_init() {
	if File_shop_api_auth_types_proto != nil {
		return
	}
	file_shop_api_shop_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_shop_api_auth_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shop_api_auth_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shop_api_auth_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shop_api_auth_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shop_api_auth_types_proto_goTypes,
		DependencyIndexes: file_shop_api_auth_types_proto_depIdxs,
		MessageInfos:      file_shop_api_auth_types_proto_msgTypes,
	}.Build()
	File_shop_api_auth_types_proto = out.File
	file_shop_api_auth_types_proto_rawDesc = nil
	file_shop_api_auth_types_proto_goTypes = nil
	file_shop_api_auth_types_proto_depIdxs = nil
}
