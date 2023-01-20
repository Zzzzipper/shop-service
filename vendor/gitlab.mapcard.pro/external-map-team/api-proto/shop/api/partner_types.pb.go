// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: shop/api/partner_types.proto

package api

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

//
// Partners
//
type Role int32

const (
	Role_BASE_PARTNER Role = 0
	Role_PARTNER      Role = 1
	Role_ADMIN        Role = 2
	Role_GUEST        Role = 3
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "BASE_PARTNER",
		1: "PARTNER",
		2: "ADMIN",
		3: "GUEST",
	}
	Role_value = map[string]int32{
		"BASE_PARTNER": 0,
		"PARTNER":      1,
		"ADMIN":        2,
		"GUEST":        3,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_shop_api_partner_types_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_shop_api_partner_types_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_shop_api_partner_types_proto_rawDescGZIP(), []int{0}
}

type Partner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Индекс партнера
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Время создания записи
	CreateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Полное наименование
	FullName string `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	// URL площадки
	Url string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	// Токен для доступа (резерв)
	ApiToken string `protobuf:"bytes,5,opt,name=api_token,json=apiToken,proto3" json:"api_token,omitempty"`
	// Роль
	Role Role `protobuf:"varint,6,opt,name=role,proto3,enum=shop.api.Role" json:"role,omitempty"`
}

func (x *Partner) Reset() {
	*x = Partner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_api_partner_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Partner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Partner) ProtoMessage() {}

func (x *Partner) ProtoReflect() protoreflect.Message {
	mi := &file_shop_api_partner_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Partner.ProtoReflect.Descriptor instead.
func (*Partner) Descriptor() ([]byte, []int) {
	return file_shop_api_partner_types_proto_rawDescGZIP(), []int{0}
}

func (x *Partner) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Partner) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Partner) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Partner) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Partner) GetApiToken() string {
	if x != nil {
		return x.ApiToken
	}
	return ""
}

func (x *Partner) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_BASE_PARTNER
}

type AddPartnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Полное наименование
	FullName string `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	// URL площадки
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// Роль
	Role Role `protobuf:"varint,3,opt,name=role,proto3,enum=shop.api.Role" json:"role,omitempty"`
}

func (x *AddPartnerRequest) Reset() {
	*x = AddPartnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_api_partner_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPartnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPartnerRequest) ProtoMessage() {}

func (x *AddPartnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_api_partner_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPartnerRequest.ProtoReflect.Descriptor instead.
func (*AddPartnerRequest) Descriptor() ([]byte, []int) {
	return file_shop_api_partner_types_proto_rawDescGZIP(), []int{1}
}

func (x *AddPartnerRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *AddPartnerRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AddPartnerRequest) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_BASE_PARTNER
}

type AddPartnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Статус операции: true/false
	Status bool `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	// Текстовый код ошибки
	ErrCode string `protobuf:"bytes,2,opt,name=ErrCode,proto3" json:"ErrCode,omitempty"`
	// Описание ошибки.
	ErrMessage string `protobuf:"bytes,3,opt,name=ErrMessage,proto3" json:"ErrMessage,omitempty"`
	// Объект партнер
	Partner *Partner `protobuf:"bytes,4,opt,name=Partner,proto3" json:"Partner,omitempty"`
}

func (x *AddPartnerResponse) Reset() {
	*x = AddPartnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_api_partner_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPartnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPartnerResponse) ProtoMessage() {}

func (x *AddPartnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_api_partner_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPartnerResponse.ProtoReflect.Descriptor instead.
func (*AddPartnerResponse) Descriptor() ([]byte, []int) {
	return file_shop_api_partner_types_proto_rawDescGZIP(), []int{2}
}

func (x *AddPartnerResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *AddPartnerResponse) GetErrCode() string {
	if x != nil {
		return x.ErrCode
	}
	return ""
}

func (x *AddPartnerResponse) GetErrMessage() string {
	if x != nil {
		return x.ErrMessage
	}
	return ""
}

func (x *AddPartnerResponse) GetPartner() *Partner {
	if x != nil {
		return x.Partner
	}
	return nil
}

var File_shop_api_partner_types_proto protoreflect.FileDescriptor

var file_shop_api_partner_types_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x07, 0x50, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x22, 0x66, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x12, 0x41,
	0x64, 0x64, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x72, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x72, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x07, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x2a, 0x3b, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x41, 0x53, 0x45,
	0x5f, 0x50, 0x41, 0x52, 0x54, 0x4e, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x41,
	0x52, 0x54, 0x4e, 0x45, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x55, 0x45, 0x53, 0x54, 0x10, 0x03, 0x42, 0x3d, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6d, 0x61, 0x70, 0x63, 0x61, 0x72, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2d, 0x6d, 0x61, 0x70,
	0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shop_api_partner_types_proto_rawDescOnce sync.Once
	file_shop_api_partner_types_proto_rawDescData = file_shop_api_partner_types_proto_rawDesc
)

func file_shop_api_partner_types_proto_rawDescGZIP() []byte {
	file_shop_api_partner_types_proto_rawDescOnce.Do(func() {
		file_shop_api_partner_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_api_partner_types_proto_rawDescData)
	})
	return file_shop_api_partner_types_proto_rawDescData
}

var file_shop_api_partner_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shop_api_partner_types_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_shop_api_partner_types_proto_goTypes = []interface{}{
	(Role)(0),                   // 0: shop.api.Role
	(*Partner)(nil),             // 1: shop.api.Partner
	(*AddPartnerRequest)(nil),   // 2: shop.api.AddPartnerRequest
	(*AddPartnerResponse)(nil),  // 3: shop.api.AddPartnerResponse
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_shop_api_partner_types_proto_depIdxs = []int32{
	4, // 0: shop.api.Partner.create_time:type_name -> google.protobuf.Timestamp
	0, // 1: shop.api.Partner.role:type_name -> shop.api.Role
	0, // 2: shop.api.AddPartnerRequest.role:type_name -> shop.api.Role
	1, // 3: shop.api.AddPartnerResponse.Partner:type_name -> shop.api.Partner
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_shop_api_partner_types_proto_init() }
func file_shop_api_partner_types_proto_init() {
	if File_shop_api_partner_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shop_api_partner_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Partner); i {
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
		file_shop_api_partner_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPartnerRequest); i {
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
		file_shop_api_partner_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPartnerResponse); i {
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
			RawDescriptor: file_shop_api_partner_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shop_api_partner_types_proto_goTypes,
		DependencyIndexes: file_shop_api_partner_types_proto_depIdxs,
		EnumInfos:         file_shop_api_partner_types_proto_enumTypes,
		MessageInfos:      file_shop_api_partner_types_proto_msgTypes,
	}.Build()
	File_shop_api_partner_types_proto = out.File
	file_shop_api_partner_types_proto_rawDesc = nil
	file_shop_api_partner_types_proto_goTypes = nil
	file_shop_api_partner_types_proto_depIdxs = nil
}
