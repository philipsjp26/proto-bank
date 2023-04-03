// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: proto/log_activity.proto

package activitypb

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
// Metadata
// =================
type ActivityMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int32 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Total   int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ActivityMeta) Reset() {
	*x = ActivityMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_activity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityMeta) ProtoMessage() {}

func (x *ActivityMeta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_activity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityMeta.ProtoReflect.Descriptor instead.
func (*ActivityMeta) Descriptor() ([]byte, []int) {
	return file_proto_log_activity_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityMeta) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ActivityMeta) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *ActivityMeta) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

// =================
// Parameter Request
// =================
type ActivityLogParameterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page        int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage     int32  `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	OrderBy     string `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	OrderMethod string `protobuf:"bytes,4,opt,name=order_method,json=orderMethod,proto3" json:"order_method,omitempty"`
	DateStart   string `protobuf:"bytes,5,opt,name=date_start,json=dateStart,proto3" json:"date_start,omitempty"`
	DateEnd     string `protobuf:"bytes,6,opt,name=date_end,json=dateEnd,proto3" json:"date_end,omitempty"`
}

func (x *ActivityLogParameterRequest) Reset() {
	*x = ActivityLogParameterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_activity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityLogParameterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityLogParameterRequest) ProtoMessage() {}

func (x *ActivityLogParameterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_activity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityLogParameterRequest.ProtoReflect.Descriptor instead.
func (*ActivityLogParameterRequest) Descriptor() ([]byte, []int) {
	return file_proto_log_activity_proto_rawDescGZIP(), []int{1}
}

func (x *ActivityLogParameterRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ActivityLogParameterRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *ActivityLogParameterRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ActivityLogParameterRequest) GetOrderMethod() string {
	if x != nil {
		return x.OrderMethod
	}
	return ""
}

func (x *ActivityLogParameterRequest) GetDateStart() string {
	if x != nil {
		return x.DateStart
	}
	return ""
}

func (x *ActivityLogParameterRequest) GetDateEnd() string {
	if x != nil {
		return x.DateEnd
	}
	return ""
}

type ActivityLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PrivyId       string `protobuf:"bytes,2,opt,name=privy_id,json=privyId,proto3" json:"privy_id,omitempty"`
	Ip            string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Activity      string `protobuf:"bytes,4,opt,name=activity,proto3" json:"activity,omitempty"`
	Browser       string `protobuf:"bytes,5,opt,name=browser,proto3" json:"browser,omitempty"`
	ApplicationId int32  `protobuf:"varint,6,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	Url           string `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	Request       string `protobuf:"bytes,8,opt,name=request,proto3" json:"request,omitempty"`
	Response      string `protobuf:"bytes,9,opt,name=response,proto3" json:"response,omitempty"`
	Headers       string `protobuf:"bytes,10,opt,name=headers,proto3" json:"headers,omitempty"`
	ActionMethod  string `protobuf:"bytes,11,opt,name=action_method,json=actionMethod,proto3" json:"action_method,omitempty"`
	CanShowDetail bool   `protobuf:"varint,12,opt,name=can_show_detail,json=canShowDetail,proto3" json:"can_show_detail,omitempty"`
	BeforeRequest string `protobuf:"bytes,13,opt,name=before_request,json=beforeRequest,proto3" json:"before_request,omitempty"`
	CreatedAt     string `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ActivityLogResponse) Reset() {
	*x = ActivityLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_activity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityLogResponse) ProtoMessage() {}

func (x *ActivityLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_activity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityLogResponse.ProtoReflect.Descriptor instead.
func (*ActivityLogResponse) Descriptor() ([]byte, []int) {
	return file_proto_log_activity_proto_rawDescGZIP(), []int{2}
}

func (x *ActivityLogResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ActivityLogResponse) GetPrivyId() string {
	if x != nil {
		return x.PrivyId
	}
	return ""
}

func (x *ActivityLogResponse) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ActivityLogResponse) GetActivity() string {
	if x != nil {
		return x.Activity
	}
	return ""
}

func (x *ActivityLogResponse) GetBrowser() string {
	if x != nil {
		return x.Browser
	}
	return ""
}

func (x *ActivityLogResponse) GetApplicationId() int32 {
	if x != nil {
		return x.ApplicationId
	}
	return 0
}

func (x *ActivityLogResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ActivityLogResponse) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *ActivityLogResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *ActivityLogResponse) GetHeaders() string {
	if x != nil {
		return x.Headers
	}
	return ""
}

func (x *ActivityLogResponse) GetActionMethod() string {
	if x != nil {
		return x.ActionMethod
	}
	return ""
}

func (x *ActivityLogResponse) GetCanShowDetail() bool {
	if x != nil {
		return x.CanShowDetail
	}
	return false
}

func (x *ActivityLogResponse) GetBeforeRequest() string {
	if x != nil {
		return x.BeforeRequest
	}
	return ""
}

func (x *ActivityLogResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ActivityLogResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ActivityLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*ActivityLogResponse `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Meta    *ActivityMeta          `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *ActivityLogsResponse) Reset() {
	*x = ActivityLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_activity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityLogsResponse) ProtoMessage() {}

func (x *ActivityLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_activity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityLogsResponse.ProtoReflect.Descriptor instead.
func (*ActivityLogsResponse) Descriptor() ([]byte, []int) {
	return file_proto_log_activity_proto_rawDescGZIP(), []int{3}
}

func (x *ActivityLogsResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ActivityLogsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ActivityLogsResponse) GetData() []*ActivityLogResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ActivityLogsResponse) GetMeta() *ActivityMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type ActivityLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *ActivityLogResponse `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Meta    *ActivityMeta        `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *ActivityLog) Reset() {
	*x = ActivityLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_activity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityLog) ProtoMessage() {}

func (x *ActivityLog) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_activity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityLog.ProtoReflect.Descriptor instead.
func (*ActivityLog) Descriptor() ([]byte, []int) {
	return file_proto_log_activity_proto_rawDescGZIP(), []int{4}
}

func (x *ActivityLog) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ActivityLog) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ActivityLog) GetData() *ActivityLogResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ActivityLog) GetMeta() *ActivityMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_activity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_activity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_log_activity_proto_rawDescGZIP(), []int{5}
}

func (x *StatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type FindActivityLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindActivityLogRequest) Reset() {
	*x = FindActivityLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_activity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindActivityLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindActivityLogRequest) ProtoMessage() {}

func (x *FindActivityLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_activity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindActivityLogRequest.ProtoReflect.Descriptor instead.
func (*FindActivityLogRequest) Descriptor() ([]byte, []int) {
	return file_proto_log_activity_proto_rawDescGZIP(), []int{6}
}

func (x *FindActivityLogRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateActivityLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivyId       string `protobuf:"bytes,2,opt,name=privy_id,json=privyId,proto3" json:"privy_id,omitempty"`
	Ip            string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Activity      string `protobuf:"bytes,4,opt,name=activity,proto3" json:"activity,omitempty"`
	Browser       string `protobuf:"bytes,5,opt,name=browser,proto3" json:"browser,omitempty"`
	ApplicationId int32  `protobuf:"varint,6,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	Url           string `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	Request       string `protobuf:"bytes,8,opt,name=request,proto3" json:"request,omitempty"`
	Response      string `protobuf:"bytes,9,opt,name=response,proto3" json:"response,omitempty"`
	Headers       string `protobuf:"bytes,10,opt,name=headers,proto3" json:"headers,omitempty"`
	ActionMethod  string `protobuf:"bytes,11,opt,name=action_method,json=actionMethod,proto3" json:"action_method,omitempty"`
	CanShowDetail bool   `protobuf:"varint,12,opt,name=can_show_detail,json=canShowDetail,proto3" json:"can_show_detail,omitempty"`
	BeforeRequest string `protobuf:"bytes,13,opt,name=before_request,json=beforeRequest,proto3" json:"before_request,omitempty"`
	CreatedAt     string `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *CreateActivityLogRequest) Reset() {
	*x = CreateActivityLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_log_activity_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateActivityLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateActivityLogRequest) ProtoMessage() {}

func (x *CreateActivityLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_log_activity_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateActivityLogRequest.ProtoReflect.Descriptor instead.
func (*CreateActivityLogRequest) Descriptor() ([]byte, []int) {
	return file_proto_log_activity_proto_rawDescGZIP(), []int{7}
}

func (x *CreateActivityLogRequest) GetPrivyId() string {
	if x != nil {
		return x.PrivyId
	}
	return ""
}

func (x *CreateActivityLogRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *CreateActivityLogRequest) GetActivity() string {
	if x != nil {
		return x.Activity
	}
	return ""
}

func (x *CreateActivityLogRequest) GetBrowser() string {
	if x != nil {
		return x.Browser
	}
	return ""
}

func (x *CreateActivityLogRequest) GetApplicationId() int32 {
	if x != nil {
		return x.ApplicationId
	}
	return 0
}

func (x *CreateActivityLogRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreateActivityLogRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *CreateActivityLogRequest) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *CreateActivityLogRequest) GetHeaders() string {
	if x != nil {
		return x.Headers
	}
	return ""
}

func (x *CreateActivityLogRequest) GetActionMethod() string {
	if x != nil {
		return x.ActionMethod
	}
	return ""
}

func (x *CreateActivityLogRequest) GetCanShowDetail() bool {
	if x != nil {
		return x.CanShowDetail
	}
	return false
}

func (x *CreateActivityLogRequest) GetBeforeRequest() string {
	if x != nil {
		return x.BeforeRequest
	}
	return ""
}

func (x *CreateActivityLogRequest) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CreateActivityLogRequest) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_proto_log_activity_proto protoreflect.FileDescriptor

var file_proto_log_activity_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x70, 0x62, 0x22, 0x53, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65,
	0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65,
	0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xc4, 0x01, 0x0a, 0x1b,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x65, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x6e, 0x64, 0x22, 0xc1, 0x03, 0x0a, 0x13, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72,
	0x69, 0x76, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72,
	0x69, 0x76, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x61, 0x6e,
	0x5f, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x63, 0x61, 0x6e, 0x53, 0x68, 0x6f, 0x77, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa7, 0x01, 0x0a, 0x14, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x70, 0x62, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x22, 0x9e, 0x01, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x70, 0x62, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x22, 0x28, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x28, 0x0a, 0x16, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb6, 0x03, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x76, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x72, 0x6f,
	0x77, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x72, 0x6f, 0x77,
	0x73, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x61, 0x6e, 0x53,
	0x68, 0x6f, 0x77, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xb2,
	0x02, 0x0a, 0x12, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x65, 0x49, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x12,
	0x24, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x70, 0x62, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x17, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c,
	0x6f, 0x67, 0x12, 0x27, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x70, 0x62, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a,
	0x15, 0x53, 0x65, 0x72, 0x76, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x42, 0x79, 0x49, 0x44, 0x12, 0x22, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x4c, 0x6f, 0x67, 0x42, 0x16, 0x5a, 0x14, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_log_activity_proto_rawDescOnce sync.Once
	file_proto_log_activity_proto_rawDescData = file_proto_log_activity_proto_rawDesc
)

func file_proto_log_activity_proto_rawDescGZIP() []byte {
	file_proto_log_activity_proto_rawDescOnce.Do(func() {
		file_proto_log_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_log_activity_proto_rawDescData)
	})
	return file_proto_log_activity_proto_rawDescData
}

var file_proto_log_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_log_activity_proto_goTypes = []interface{}{
	(*ActivityMeta)(nil),                // 0: activitypb.ActivityMeta
	(*ActivityLogParameterRequest)(nil), // 1: activitypb.ActivityLogParameterRequest
	(*ActivityLogResponse)(nil),         // 2: activitypb.ActivityLogResponse
	(*ActivityLogsResponse)(nil),        // 3: activitypb.ActivityLogsResponse
	(*ActivityLog)(nil),                 // 4: activitypb.ActivityLog
	(*StatusResponse)(nil),              // 5: activitypb.StatusResponse
	(*FindActivityLogRequest)(nil),      // 6: activitypb.FindActivityLogRequest
	(*CreateActivityLogRequest)(nil),    // 7: activitypb.CreateActivityLogRequest
}
var file_proto_log_activity_proto_depIdxs = []int32{
	2, // 0: activitypb.ActivityLogsResponse.data:type_name -> activitypb.ActivityLogResponse
	0, // 1: activitypb.ActivityLogsResponse.meta:type_name -> activitypb.ActivityMeta
	2, // 2: activitypb.ActivityLog.data:type_name -> activitypb.ActivityLogResponse
	0, // 3: activitypb.ActivityLog.meta:type_name -> activitypb.ActivityMeta
	7, // 4: activitypb.ActivityLogService.ServeInsertActivityLog:input_type -> activitypb.CreateActivityLogRequest
	1, // 5: activitypb.ActivityLogService.ServeGetListActivityLog:input_type -> activitypb.ActivityLogParameterRequest
	6, // 6: activitypb.ActivityLogService.ServeFindActivityByID:input_type -> activitypb.FindActivityLogRequest
	3, // 7: activitypb.ActivityLogService.ServeInsertActivityLog:output_type -> activitypb.ActivityLogsResponse
	3, // 8: activitypb.ActivityLogService.ServeGetListActivityLog:output_type -> activitypb.ActivityLogsResponse
	4, // 9: activitypb.ActivityLogService.ServeFindActivityByID:output_type -> activitypb.ActivityLog
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_log_activity_proto_init() }
func file_proto_log_activity_proto_init() {
	if File_proto_log_activity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_log_activity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityMeta); i {
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
		file_proto_log_activity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityLogParameterRequest); i {
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
		file_proto_log_activity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityLogResponse); i {
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
		file_proto_log_activity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityLogsResponse); i {
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
		file_proto_log_activity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityLog); i {
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
		file_proto_log_activity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_proto_log_activity_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindActivityLogRequest); i {
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
		file_proto_log_activity_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateActivityLogRequest); i {
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
			RawDescriptor: file_proto_log_activity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_log_activity_proto_goTypes,
		DependencyIndexes: file_proto_log_activity_proto_depIdxs,
		MessageInfos:      file_proto_log_activity_proto_msgTypes,
	}.Build()
	File_proto_log_activity_proto = out.File
	file_proto_log_activity_proto_rawDesc = nil
	file_proto_log_activity_proto_goTypes = nil
	file_proto_log_activity_proto_depIdxs = nil
}
