// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/placement-service.proto

package documents

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
type PlacementHttpMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int32 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Total   int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PlacementHttpMeta) Reset() {
	*x = PlacementHttpMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlacementHttpMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlacementHttpMeta) ProtoMessage() {}

func (x *PlacementHttpMeta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlacementHttpMeta.ProtoReflect.Descriptor instead.
func (*PlacementHttpMeta) Descriptor() ([]byte, []int) {
	return file_proto_placement_service_proto_rawDescGZIP(), []int{0}
}

func (x *PlacementHttpMeta) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PlacementHttpMeta) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *PlacementHttpMeta) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

// =================
// Parameter Request
// =================
type PlacementDocParameterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page        int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage     int32  `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	OrderBy     string `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	OrderMethod string `protobuf:"bytes,4,opt,name=order_method,json=orderMethod,proto3" json:"order_method,omitempty"`
	DateStart   string `protobuf:"bytes,5,opt,name=date_start,json=dateStart,proto3" json:"date_start,omitempty"`
	DateEnd     string `protobuf:"bytes,6,opt,name=date_end,json=dateEnd,proto3" json:"date_end,omitempty"`
	Status      string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PlacementDocParameterRequest) Reset() {
	*x = PlacementDocParameterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlacementDocParameterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlacementDocParameterRequest) ProtoMessage() {}

func (x *PlacementDocParameterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlacementDocParameterRequest.ProtoReflect.Descriptor instead.
func (*PlacementDocParameterRequest) Descriptor() ([]byte, []int) {
	return file_proto_placement_service_proto_rawDescGZIP(), []int{1}
}

func (x *PlacementDocParameterRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PlacementDocParameterRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *PlacementDocParameterRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *PlacementDocParameterRequest) GetOrderMethod() string {
	if x != nil {
		return x.OrderMethod
	}
	return ""
}

func (x *PlacementDocParameterRequest) GetDateStart() string {
	if x != nil {
		return x.DateStart
	}
	return ""
}

func (x *PlacementDocParameterRequest) GetDateEnd() string {
	if x != nil {
		return x.DateEnd
	}
	return ""
}

func (x *PlacementDocParameterRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type PlacementDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Channel         string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	ReffNo          string `protobuf:"bytes,3,opt,name=reff_no,json=reffNo,proto3" json:"reff_no,omitempty"`
	SignType        string `protobuf:"bytes,4,opt,name=sign_type,json=signType,proto3" json:"sign_type,omitempty"`
	DocFlow         string `protobuf:"bytes,5,opt,name=doc_flow,json=docFlow,proto3" json:"doc_flow,omitempty"`
	DocCategory     string `protobuf:"bytes,6,opt,name=doc_category,json=docCategory,proto3" json:"doc_category,omitempty"`
	DocTitle        string `protobuf:"bytes,7,opt,name=doc_title,json=docTitle,proto3" json:"doc_title,omitempty"`
	DocStatus       string `protobuf:"bytes,8,opt,name=doc_status,json=docStatus,proto3" json:"doc_status,omitempty"`
	Uploader        string `protobuf:"bytes,9,opt,name=uploader,proto3" json:"uploader,omitempty"`
	UploaderPrivyId string `protobuf:"bytes,10,opt,name=uploader_privy_id,json=uploaderPrivyId,proto3" json:"uploader_privy_id,omitempty"`
	ESign           string `protobuf:"bytes,11,opt,name=e_sign,json=eSign,proto3" json:"e_sign,omitempty"`
	EMaterai        string `protobuf:"bytes,12,opt,name=e_materai,json=eMaterai,proto3" json:"e_materai,omitempty"`
	CreatedAt       string `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       string `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *PlacementDocument) Reset() {
	*x = PlacementDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlacementDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlacementDocument) ProtoMessage() {}

func (x *PlacementDocument) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlacementDocument.ProtoReflect.Descriptor instead.
func (*PlacementDocument) Descriptor() ([]byte, []int) {
	return file_proto_placement_service_proto_rawDescGZIP(), []int{2}
}

func (x *PlacementDocument) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlacementDocument) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *PlacementDocument) GetReffNo() string {
	if x != nil {
		return x.ReffNo
	}
	return ""
}

func (x *PlacementDocument) GetSignType() string {
	if x != nil {
		return x.SignType
	}
	return ""
}

func (x *PlacementDocument) GetDocFlow() string {
	if x != nil {
		return x.DocFlow
	}
	return ""
}

func (x *PlacementDocument) GetDocCategory() string {
	if x != nil {
		return x.DocCategory
	}
	return ""
}

func (x *PlacementDocument) GetDocTitle() string {
	if x != nil {
		return x.DocTitle
	}
	return ""
}

func (x *PlacementDocument) GetDocStatus() string {
	if x != nil {
		return x.DocStatus
	}
	return ""
}

func (x *PlacementDocument) GetUploader() string {
	if x != nil {
		return x.Uploader
	}
	return ""
}

func (x *PlacementDocument) GetUploaderPrivyId() string {
	if x != nil {
		return x.UploaderPrivyId
	}
	return ""
}

func (x *PlacementDocument) GetESign() string {
	if x != nil {
		return x.ESign
	}
	return ""
}

func (x *PlacementDocument) GetEMaterai() string {
	if x != nil {
		return x.EMaterai
	}
	return ""
}

func (x *PlacementDocument) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PlacementDocument) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type PlacementDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32              `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string             `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *PlacementDocument `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Meta    *PlacementHttpMeta `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *PlacementDocumentResponse) Reset() {
	*x = PlacementDocumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlacementDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlacementDocumentResponse) ProtoMessage() {}

func (x *PlacementDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlacementDocumentResponse.ProtoReflect.Descriptor instead.
func (*PlacementDocumentResponse) Descriptor() ([]byte, []int) {
	return file_proto_placement_service_proto_rawDescGZIP(), []int{3}
}

func (x *PlacementDocumentResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PlacementDocumentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PlacementDocumentResponse) GetData() *PlacementDocument {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PlacementDocumentResponse) GetMeta() *PlacementHttpMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type ListPlacementDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*PlacementDocument `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Meta    *PlacementHttpMeta   `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *ListPlacementDocumentResponse) Reset() {
	*x = ListPlacementDocumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlacementDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlacementDocumentResponse) ProtoMessage() {}

func (x *ListPlacementDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlacementDocumentResponse.ProtoReflect.Descriptor instead.
func (*ListPlacementDocumentResponse) Descriptor() ([]byte, []int) {
	return file_proto_placement_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListPlacementDocumentResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListPlacementDocumentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListPlacementDocumentResponse) GetData() []*PlacementDocument {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListPlacementDocumentResponse) GetMeta() *PlacementHttpMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type UpdatePlacementDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel         string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	ReffNo          string `protobuf:"bytes,2,opt,name=reff_no,json=reffNo,proto3" json:"reff_no,omitempty"`
	SignType        string `protobuf:"bytes,3,opt,name=sign_type,json=signType,proto3" json:"sign_type,omitempty"`
	DocFlow         string `protobuf:"bytes,4,opt,name=doc_flow,json=docFlow,proto3" json:"doc_flow,omitempty"`
	DocCategory     string `protobuf:"bytes,5,opt,name=doc_category,json=docCategory,proto3" json:"doc_category,omitempty"`
	DocTitle        string `protobuf:"bytes,6,opt,name=doc_title,json=docTitle,proto3" json:"doc_title,omitempty"`
	DocStatus       string `protobuf:"bytes,7,opt,name=doc_status,json=docStatus,proto3" json:"doc_status,omitempty"`
	Uploader        string `protobuf:"bytes,8,opt,name=uploader,proto3" json:"uploader,omitempty"`
	UploaderPrivyId string `protobuf:"bytes,9,opt,name=uploader_privy_id,json=uploaderPrivyId,proto3" json:"uploader_privy_id,omitempty"`
	ESign           string `protobuf:"bytes,10,opt,name=e_sign,json=eSign,proto3" json:"e_sign,omitempty"`
	EMaterai        string `protobuf:"bytes,11,opt,name=e_materai,json=eMaterai,proto3" json:"e_materai,omitempty"`
	UpdatedAt       string `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UpdatePlacementDocumentRequest) Reset() {
	*x = UpdatePlacementDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlacementDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlacementDocumentRequest) ProtoMessage() {}

func (x *UpdatePlacementDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlacementDocumentRequest.ProtoReflect.Descriptor instead.
func (*UpdatePlacementDocumentRequest) Descriptor() ([]byte, []int) {
	return file_proto_placement_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePlacementDocumentRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *UpdatePlacementDocumentRequest) GetReffNo() string {
	if x != nil {
		return x.ReffNo
	}
	return ""
}

func (x *UpdatePlacementDocumentRequest) GetSignType() string {
	if x != nil {
		return x.SignType
	}
	return ""
}

func (x *UpdatePlacementDocumentRequest) GetDocFlow() string {
	if x != nil {
		return x.DocFlow
	}
	return ""
}

func (x *UpdatePlacementDocumentRequest) GetDocCategory() string {
	if x != nil {
		return x.DocCategory
	}
	return ""
}

func (x *UpdatePlacementDocumentRequest) GetDocTitle() string {
	if x != nil {
		return x.DocTitle
	}
	return ""
}

func (x *UpdatePlacementDocumentRequest) GetDocStatus() string {
	if x != nil {
		return x.DocStatus
	}
	return ""
}

func (x *UpdatePlacementDocumentRequest) GetUploader() string {
	if x != nil {
		return x.Uploader
	}
	return ""
}

func (x *UpdatePlacementDocumentRequest) GetUploaderPrivyId() string {
	if x != nil {
		return x.UploaderPrivyId
	}
	return ""
}

func (x *UpdatePlacementDocumentRequest) GetESign() string {
	if x != nil {
		return x.ESign
	}
	return ""
}

func (x *UpdatePlacementDocumentRequest) GetEMaterai() string {
	if x != nil {
		return x.EMaterai
	}
	return ""
}

func (x *UpdatePlacementDocumentRequest) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type RevokePlacementDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReffNo    string `protobuf:"bytes,1,opt,name=reff_no,json=reffNo,proto3" json:"reff_no,omitempty"`
	DocStatus string `protobuf:"bytes,2,opt,name=doc_status,json=docStatus,proto3" json:"doc_status,omitempty"`
}

func (x *RevokePlacementDocumentRequest) Reset() {
	*x = RevokePlacementDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokePlacementDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokePlacementDocumentRequest) ProtoMessage() {}

func (x *RevokePlacementDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokePlacementDocumentRequest.ProtoReflect.Descriptor instead.
func (*RevokePlacementDocumentRequest) Descriptor() ([]byte, []int) {
	return file_proto_placement_service_proto_rawDescGZIP(), []int{6}
}

func (x *RevokePlacementDocumentRequest) GetReffNo() string {
	if x != nil {
		return x.ReffNo
	}
	return ""
}

func (x *RevokePlacementDocumentRequest) GetDocStatus() string {
	if x != nil {
		return x.DocStatus
	}
	return ""
}

type DetailRevokedDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReffNo string `protobuf:"bytes,1,opt,name=reff_no,json=reffNo,proto3" json:"reff_no,omitempty"`
}

func (x *DetailRevokedDocumentRequest) Reset() {
	*x = DetailRevokedDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_placement_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailRevokedDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailRevokedDocumentRequest) ProtoMessage() {}

func (x *DetailRevokedDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_placement_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailRevokedDocumentRequest.ProtoReflect.Descriptor instead.
func (*DetailRevokedDocumentRequest) Descriptor() ([]byte, []int) {
	return file_proto_placement_service_proto_rawDescGZIP(), []int{7}
}

func (x *DetailRevokedDocumentRequest) GetReffNo() string {
	if x != nil {
		return x.ReffNo
	}
	return ""
}

var File_proto_placement_service_proto protoreflect.FileDescriptor

var file_proto_placement_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x22, 0x58, 0x0a, 0x11,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xdd, 0x01, 0x0a, 0x1c, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70,
	0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa7, 0x03, 0x0a, 0x11, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x66, 0x5f, 0x6e,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x66, 0x66, 0x4e, 0x6f, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x64, 0x6f, 0x63, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x6f, 0x63, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x5f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x6f, 0x63, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f,
	0x63, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x63, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x6f, 0x63,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x69, 0x76, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x76, 0x79, 0x49, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x65, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x61, 0x69, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x61, 0x69, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xb1, 0x01, 0x0a, 0x19, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x32, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x22, 0xb5, 0x01, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x74,
	0x74, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x85, 0x03, 0x0a,
	0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x66,
	0x66, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x66, 0x66,
	0x4e, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x64, 0x6f, 0x63, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f,
	0x63, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x6f, 0x63, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x6f, 0x63, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x6f, 0x63, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x6f,
	0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x76, 0x79, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x5f, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x61, 0x69, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x4d, 0x61,
	0x74, 0x65, 0x72, 0x61, 0x69, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x58, 0x0a, 0x1e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x66, 0x5f, 0x6e,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x66, 0x66, 0x4e, 0x6f, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x37,
	0x0a, 0x1c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x65, 0x66, 0x66, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x66, 0x66, 0x4e, 0x6f, 0x32, 0xf8, 0x02, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x1c,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x7a, 0x0a, 0x21, 0x53, 0x65, 0x72, 0x76, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f,
	0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a,
	0x1c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x2e,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_placement_service_proto_rawDescOnce sync.Once
	file_proto_placement_service_proto_rawDescData = file_proto_placement_service_proto_rawDesc
)

func file_proto_placement_service_proto_rawDescGZIP() []byte {
	file_proto_placement_service_proto_rawDescOnce.Do(func() {
		file_proto_placement_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_placement_service_proto_rawDescData)
	})
	return file_proto_placement_service_proto_rawDescData
}

var file_proto_placement_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_placement_service_proto_goTypes = []interface{}{
	(*PlacementHttpMeta)(nil),              // 0: placementpb.PlacementHttpMeta
	(*PlacementDocParameterRequest)(nil),   // 1: placementpb.PlacementDocParameterRequest
	(*PlacementDocument)(nil),              // 2: placementpb.PlacementDocument
	(*PlacementDocumentResponse)(nil),      // 3: placementpb.PlacementDocumentResponse
	(*ListPlacementDocumentResponse)(nil),  // 4: placementpb.ListPlacementDocumentResponse
	(*UpdatePlacementDocumentRequest)(nil), // 5: placementpb.UpdatePlacementDocumentRequest
	(*RevokePlacementDocumentRequest)(nil), // 6: placementpb.RevokePlacementDocumentRequest
	(*DetailRevokedDocumentRequest)(nil),   // 7: placementpb.DetailRevokedDocumentRequest
}
var file_proto_placement_service_proto_depIdxs = []int32{
	2, // 0: placementpb.PlacementDocumentResponse.data:type_name -> placementpb.PlacementDocument
	0, // 1: placementpb.PlacementDocumentResponse.meta:type_name -> placementpb.PlacementHttpMeta
	2, // 2: placementpb.ListPlacementDocumentResponse.data:type_name -> placementpb.PlacementDocument
	0, // 3: placementpb.ListPlacementDocumentResponse.meta:type_name -> placementpb.PlacementHttpMeta
	6, // 4: placementpb.PlacementService.ServeRevokePlacementDocument:input_type -> placementpb.RevokePlacementDocumentRequest
	1, // 5: placementpb.PlacementService.ServeListRevokedPlacementDocument:input_type -> placementpb.PlacementDocParameterRequest
	5, // 6: placementpb.PlacementService.ServeUpdatePlacementDocument:input_type -> placementpb.UpdatePlacementDocumentRequest
	3, // 7: placementpb.PlacementService.ServeRevokePlacementDocument:output_type -> placementpb.PlacementDocumentResponse
	4, // 8: placementpb.PlacementService.ServeListRevokedPlacementDocument:output_type -> placementpb.ListPlacementDocumentResponse
	3, // 9: placementpb.PlacementService.ServeUpdatePlacementDocument:output_type -> placementpb.PlacementDocumentResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_placement_service_proto_init() }
func file_proto_placement_service_proto_init() {
	if File_proto_placement_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_placement_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlacementHttpMeta); i {
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
		file_proto_placement_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlacementDocParameterRequest); i {
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
		file_proto_placement_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlacementDocument); i {
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
		file_proto_placement_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlacementDocumentResponse); i {
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
		file_proto_placement_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPlacementDocumentResponse); i {
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
		file_proto_placement_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlacementDocumentRequest); i {
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
		file_proto_placement_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokePlacementDocumentRequest); i {
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
		file_proto_placement_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailRevokedDocumentRequest); i {
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
			RawDescriptor: file_proto_placement_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_placement_service_proto_goTypes,
		DependencyIndexes: file_proto_placement_service_proto_depIdxs,
		MessageInfos:      file_proto_placement_service_proto_msgTypes,
	}.Build()
	File_proto_placement_service_proto = out.File
	file_proto_placement_service_proto_rawDesc = nil
	file_proto_placement_service_proto_goTypes = nil
	file_proto_placement_service_proto_depIdxs = nil
}
