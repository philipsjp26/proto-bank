// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/email/proto/email.proto

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

type EmailData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BalanceTypeName string `protobuf:"bytes,1,opt,name=balance_type_name,json=balanceTypeName,proto3" json:"balance_type_name,omitempty"`
	Price           int32  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	Quantity        int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Discount        int32  `protobuf:"varint,4,opt,name=discount,proto3" json:"discount,omitempty"`
	Tax             int32  `protobuf:"varint,5,opt,name=tax,proto3" json:"tax,omitempty"`
	Total           int32  `protobuf:"varint,6,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *EmailData) Reset() {
	*x = EmailData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_email_proto_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailData) ProtoMessage() {}

func (x *EmailData) ProtoReflect() protoreflect.Message {
	mi := &file_api_email_proto_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailData.ProtoReflect.Descriptor instead.
func (*EmailData) Descriptor() ([]byte, []int) {
	return file_api_email_proto_email_proto_rawDescGZIP(), []int{0}
}

func (x *EmailData) GetBalanceTypeName() string {
	if x != nil {
		return x.BalanceTypeName
	}
	return ""
}

func (x *EmailData) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *EmailData) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *EmailData) GetDiscount() int32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *EmailData) GetTax() int32 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *EmailData) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type EmailContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvoiceNumber  string       `protobuf:"bytes,1,opt,name=invoice_number,json=invoiceNumber,proto3" json:"invoice_number,omitempty"`
	PaymentDueDate string       `protobuf:"bytes,2,opt,name=payment_due_date,json=paymentDueDate,proto3" json:"payment_due_date,omitempty"`
	SubTotal       int32        `protobuf:"varint,3,opt,name=sub_total,json=subTotal,proto3" json:"sub_total,omitempty"`
	TotalPaid      int32        `protobuf:"varint,4,opt,name=total_paid,json=totalPaid,proto3" json:"total_paid,omitempty"`
	AmountDue      int32        `protobuf:"varint,5,opt,name=amount_due,json=amountDue,proto3" json:"amount_due,omitempty"`
	Data           []*EmailData `protobuf:"bytes,6,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *EmailContent) Reset() {
	*x = EmailContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_email_proto_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailContent) ProtoMessage() {}

func (x *EmailContent) ProtoReflect() protoreflect.Message {
	mi := &file_api_email_proto_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailContent.ProtoReflect.Descriptor instead.
func (*EmailContent) Descriptor() ([]byte, []int) {
	return file_api_email_proto_email_proto_rawDescGZIP(), []int{1}
}

func (x *EmailContent) GetInvoiceNumber() string {
	if x != nil {
		return x.InvoiceNumber
	}
	return ""
}

func (x *EmailContent) GetPaymentDueDate() string {
	if x != nil {
		return x.PaymentDueDate
	}
	return ""
}

func (x *EmailContent) GetSubTotal() int32 {
	if x != nil {
		return x.SubTotal
	}
	return 0
}

func (x *EmailContent) GetTotalPaid() int32 {
	if x != nil {
		return x.TotalPaid
	}
	return 0
}

func (x *EmailContent) GetAmountDue() int32 {
	if x != nil {
		return x.AmountDue
	}
	return 0
}

func (x *EmailContent) GetData() []*EmailData {
	if x != nil {
		return x.Data
	}
	return nil
}

type EmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To       []string      `protobuf:"bytes,1,rep,name=to,proto3" json:"to,omitempty"`
	Cc       []string      `protobuf:"bytes,2,rep,name=cc,proto3" json:"cc,omitempty"`
	Bcc      []string      `protobuf:"bytes,3,rep,name=bcc,proto3" json:"bcc,omitempty"`
	Subject  string        `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Template string        `protobuf:"bytes,5,opt,name=template,proto3" json:"template,omitempty"`
	Content  *EmailContent `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *EmailRequest) Reset() {
	*x = EmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_email_proto_email_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailRequest) ProtoMessage() {}

func (x *EmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_email_proto_email_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailRequest.ProtoReflect.Descriptor instead.
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return file_api_email_proto_email_proto_rawDescGZIP(), []int{2}
}

func (x *EmailRequest) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *EmailRequest) GetCc() []string {
	if x != nil {
		return x.Cc
	}
	return nil
}

func (x *EmailRequest) GetBcc() []string {
	if x != nil {
		return x.Bcc
	}
	return nil
}

func (x *EmailRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *EmailRequest) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

func (x *EmailRequest) GetContent() *EmailContent {
	if x != nil {
		return x.Content
	}
	return nil
}

type EmailReesponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *BaseResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Message  string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EmailReesponse) Reset() {
	*x = EmailReesponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_email_proto_email_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailReesponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailReesponse) ProtoMessage() {}

func (x *EmailReesponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_email_proto_email_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailReesponse.ProtoReflect.Descriptor instead.
func (*EmailReesponse) Descriptor() ([]byte, []int) {
	return file_api_email_proto_email_proto_rawDescGZIP(), []int{3}
}

func (x *EmailReesponse) GetResponse() *BaseResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *EmailReesponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_email_proto_email_proto protoreflect.FileDescriptor

var file_api_email_proto_email_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xad, 0x01, 0x0a, 0x09, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a,
	0x0a, 0x11, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x78, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x61, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0xe1, 0x01, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x64, 0x75, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x75, 0x65, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x75, 0x62, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x75, 0x65, 0x12, 0x25, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xa6, 0x01, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x63, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x02, 0x63, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x63, 0x63, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x03, 0x62, 0x63, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x5a, 0x0a,
	0x0e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x47, 0x0a, 0x0c, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x53, 0x65, 0x6e,
	0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_email_proto_email_proto_rawDescOnce sync.Once
	file_api_email_proto_email_proto_rawDescData = file_api_email_proto_email_proto_rawDesc
)

func file_api_email_proto_email_proto_rawDescGZIP() []byte {
	file_api_email_proto_email_proto_rawDescOnce.Do(func() {
		file_api_email_proto_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_email_proto_email_proto_rawDescData)
	})
	return file_api_email_proto_email_proto_rawDescData
}

var file_api_email_proto_email_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_email_proto_email_proto_goTypes = []interface{}{
	(*EmailData)(nil),      // 0: status.EmailData
	(*EmailContent)(nil),   // 1: status.EmailContent
	(*EmailRequest)(nil),   // 2: status.EmailRequest
	(*EmailReesponse)(nil), // 3: status.EmailReesponse
	(*BaseResponse)(nil),   // 4: meta.BaseResponse
}
var file_api_email_proto_email_proto_depIdxs = []int32{
	0, // 0: status.EmailContent.data:type_name -> status.EmailData
	1, // 1: status.EmailRequest.content:type_name -> status.EmailContent
	4, // 2: status.EmailReesponse.response:type_name -> meta.BaseResponse
	2, // 3: status.EmailService.SendEmail:input_type -> status.EmailRequest
	4, // 4: status.EmailService.SendEmail:output_type -> meta.BaseResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_email_proto_email_proto_init() }
func file_api_email_proto_email_proto_init() {
	if File_api_email_proto_email_proto != nil {
		return
	}
	file_api_meta_proto_meta_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_email_proto_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailData); i {
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
		file_api_email_proto_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailContent); i {
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
		file_api_email_proto_email_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailRequest); i {
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
		file_api_email_proto_email_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailReesponse); i {
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
			RawDescriptor: file_api_email_proto_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_email_proto_email_proto_goTypes,
		DependencyIndexes: file_api_email_proto_email_proto_depIdxs,
		MessageInfos:      file_api_email_proto_email_proto_msgTypes,
	}.Build()
	File_api_email_proto_email_proto = out.File
	file_api_email_proto_email_proto_rawDesc = nil
	file_api_email_proto_email_proto_goTypes = nil
	file_api_email_proto_email_proto_depIdxs = nil
}
