// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/notification/proto/device_token.proto

package __

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

// =================
// Data
// =================
type DeviceToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeviceToken string `protobuf:"bytes,2,opt,name=device_token,json=deviceToken,proto3" json:"device_token,omitempty"`
	PrivyId     string `protobuf:"bytes,3,opt,name=privy_id,json=privyId,proto3" json:"privy_id,omitempty"`
}

func (x *DeviceToken) Reset() {
	*x = DeviceToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notification_proto_device_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceToken) ProtoMessage() {}

func (x *DeviceToken) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_proto_device_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceToken.ProtoReflect.Descriptor instead.
func (*DeviceToken) Descriptor() ([]byte, []int) {
	return file_api_notification_proto_device_token_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceToken) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeviceToken) GetDeviceToken() string {
	if x != nil {
		return x.DeviceToken
	}
	return ""
}

func (x *DeviceToken) GetPrivyId() string {
	if x != nil {
		return x.PrivyId
	}
	return ""
}

// =================
// Parameter Request
// =================
type CreateDeviceTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceToken string `protobuf:"bytes,1,opt,name=device_token,json=deviceToken,proto3" json:"device_token,omitempty"`
	PrivyId     string `protobuf:"bytes,2,opt,name=privy_id,json=privyId,proto3" json:"privy_id,omitempty"`
}

func (x *CreateDeviceTokenRequest) Reset() {
	*x = CreateDeviceTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notification_proto_device_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeviceTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeviceTokenRequest) ProtoMessage() {}

func (x *CreateDeviceTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_proto_device_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeviceTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateDeviceTokenRequest) Descriptor() ([]byte, []int) {
	return file_api_notification_proto_device_token_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDeviceTokenRequest) GetDeviceToken() string {
	if x != nil {
		return x.DeviceToken
	}
	return ""
}

func (x *CreateDeviceTokenRequest) GetPrivyId() string {
	if x != nil {
		return x.PrivyId
	}
	return ""
}

type FindDeviceTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivyId string `protobuf:"bytes,1,opt,name=privy_id,json=privyId,proto3" json:"privy_id,omitempty"`
}

func (x *FindDeviceTokensRequest) Reset() {
	*x = FindDeviceTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notification_proto_device_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindDeviceTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindDeviceTokensRequest) ProtoMessage() {}

func (x *FindDeviceTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_proto_device_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindDeviceTokensRequest.ProtoReflect.Descriptor instead.
func (*FindDeviceTokensRequest) Descriptor() ([]byte, []int) {
	return file_api_notification_proto_device_token_proto_rawDescGZIP(), []int{2}
}

func (x *FindDeviceTokensRequest) GetPrivyId() string {
	if x != nil {
		return x.PrivyId
	}
	return ""
}

type CreateDeviceTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32        `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *DeviceToken `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateDeviceTokenResponse) Reset() {
	*x = CreateDeviceTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notification_proto_device_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeviceTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeviceTokenResponse) ProtoMessage() {}

func (x *CreateDeviceTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_proto_device_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeviceTokenResponse.ProtoReflect.Descriptor instead.
func (*CreateDeviceTokenResponse) Descriptor() ([]byte, []int) {
	return file_api_notification_proto_device_token_proto_rawDescGZIP(), []int{3}
}

func (x *CreateDeviceTokenResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateDeviceTokenResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateDeviceTokenResponse) GetData() *DeviceToken {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListDeviceTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32          `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*DeviceToken `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Meta    *Pagination    `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *ListDeviceTokenResponse) Reset() {
	*x = ListDeviceTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notification_proto_device_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeviceTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeviceTokenResponse) ProtoMessage() {}

func (x *ListDeviceTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_proto_device_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeviceTokenResponse.ProtoReflect.Descriptor instead.
func (*ListDeviceTokenResponse) Descriptor() ([]byte, []int) {
	return file_api_notification_proto_device_token_proto_rawDescGZIP(), []int{4}
}

func (x *ListDeviceTokenResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ListDeviceTokenResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListDeviceTokenResponse) GetData() []*DeviceToken {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListDeviceTokenResponse) GetMeta() *Pagination {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_api_notification_proto_device_token_proto protoreflect.FileDescriptor

var file_api_notification_proto_device_token_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x76, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x79, 0x49,
	0x64, 0x22, 0x58, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x76, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x79, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x17, 0x46,
	0x69, 0x6e, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x76, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x79, 0x49,
	0x64, 0x22, 0x7c, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0xa0, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x32, 0x67, 0x0a, 0x12, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x32, 0x42, 0x03, 0x5a, 0x01, 0x2e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_notification_proto_device_token_proto_rawDescOnce sync.Once
	file_api_notification_proto_device_token_proto_rawDescData = file_api_notification_proto_device_token_proto_rawDesc
)

func file_api_notification_proto_device_token_proto_rawDescGZIP() []byte {
	file_api_notification_proto_device_token_proto_rawDescOnce.Do(func() {
		file_api_notification_proto_device_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_notification_proto_device_token_proto_rawDescData)
	})
	return file_api_notification_proto_device_token_proto_rawDescData
}

var file_api_notification_proto_device_token_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_notification_proto_device_token_proto_goTypes = []interface{}{
	(*DeviceToken)(nil),               // 0: device_token.DeviceToken
	(*CreateDeviceTokenRequest)(nil),  // 1: device_token.CreateDeviceTokenRequest
	(*FindDeviceTokensRequest)(nil),   // 2: device_token.FindDeviceTokensRequest
	(*CreateDeviceTokenResponse)(nil), // 3: device_token.CreateDeviceTokenResponse
	(*ListDeviceTokenResponse)(nil),   // 4: device_token.ListDeviceTokenResponse
	(*Pagination)(nil),                // 5: meta.Pagination
	(*BaseResponseV2)(nil),            // 6: meta.BaseResponseV2
}
var file_api_notification_proto_device_token_proto_depIdxs = []int32{
	0, // 0: device_token.CreateDeviceTokenResponse.data:type_name -> device_token.DeviceToken
	0, // 1: device_token.ListDeviceTokenResponse.data:type_name -> device_token.DeviceToken
	5, // 2: device_token.ListDeviceTokenResponse.meta:type_name -> meta.Pagination
	1, // 3: device_token.DeviceTokenService.CreateDeviceToken:input_type -> device_token.CreateDeviceTokenRequest
	6, // 4: device_token.DeviceTokenService.CreateDeviceToken:output_type -> meta.BaseResponseV2
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_notification_proto_device_token_proto_init() }
func file_api_notification_proto_device_token_proto_init() {
	if File_api_notification_proto_device_token_proto != nil {
		return
	}
	file_api_meta_proto_meta_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_notification_proto_device_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceToken); i {
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
		file_api_notification_proto_device_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeviceTokenRequest); i {
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
		file_api_notification_proto_device_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindDeviceTokensRequest); i {
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
		file_api_notification_proto_device_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeviceTokenResponse); i {
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
		file_api_notification_proto_device_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeviceTokenResponse); i {
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
			RawDescriptor: file_api_notification_proto_device_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_notification_proto_device_token_proto_goTypes,
		DependencyIndexes: file_api_notification_proto_device_token_proto_depIdxs,
		MessageInfos:      file_api_notification_proto_device_token_proto_msgTypes,
	}.Build()
	File_api_notification_proto_device_token_proto = out.File
	file_api_notification_proto_device_token_proto_rawDesc = nil
	file_api_notification_proto_device_token_proto_goTypes = nil
	file_api_notification_proto_device_token_proto_depIdxs = nil
}
