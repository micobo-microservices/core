// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: webhook/v1/webhook.proto

package webhookv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v11 "github.com/micobo-microservices/core/domain/common/v1"
	v1 "github.com/micobo-microservices/core/domain/provider/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/httpbody"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IDnowAbortedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantName string                  `protobuf:"bytes,1,opt,name=tenant_name,json=tenantName,proto3" json:"tenant_name,omitempty"`
	Request    *v1.IDnowAbortedPayload `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *IDnowAbortedRequest) Reset() {
	*x = IDnowAbortedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_v1_webhook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDnowAbortedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDnowAbortedRequest) ProtoMessage() {}

func (x *IDnowAbortedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_v1_webhook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDnowAbortedRequest.ProtoReflect.Descriptor instead.
func (*IDnowAbortedRequest) Descriptor() ([]byte, []int) {
	return file_webhook_v1_webhook_proto_rawDescGZIP(), []int{0}
}

func (x *IDnowAbortedRequest) GetTenantName() string {
	if x != nil {
		return x.TenantName
	}
	return ""
}

func (x *IDnowAbortedRequest) GetRequest() *v1.IDnowAbortedPayload {
	if x != nil {
		return x.Request
	}
	return nil
}

type IDnowCanceledRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantName string                   `protobuf:"bytes,1,opt,name=tenant_name,json=tenantName,proto3" json:"tenant_name,omitempty"`
	Request    *v1.IDnowCanceledPayload `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *IDnowCanceledRequest) Reset() {
	*x = IDnowCanceledRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_v1_webhook_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDnowCanceledRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDnowCanceledRequest) ProtoMessage() {}

func (x *IDnowCanceledRequest) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_v1_webhook_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDnowCanceledRequest.ProtoReflect.Descriptor instead.
func (*IDnowCanceledRequest) Descriptor() ([]byte, []int) {
	return file_webhook_v1_webhook_proto_rawDescGZIP(), []int{1}
}

func (x *IDnowCanceledRequest) GetTenantName() string {
	if x != nil {
		return x.TenantName
	}
	return ""
}

func (x *IDnowCanceledRequest) GetRequest() *v1.IDnowCanceledPayload {
	if x != nil {
		return x.Request
	}
	return nil
}

type IDnowFraudSuspicionConfirmedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantName string                                  `protobuf:"bytes,1,opt,name=tenant_name,json=tenantName,proto3" json:"tenant_name,omitempty"`
	Request    *v1.IDnowFraudSuspicionConfirmedPayload `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *IDnowFraudSuspicionConfirmedRequest) Reset() {
	*x = IDnowFraudSuspicionConfirmedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_v1_webhook_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDnowFraudSuspicionConfirmedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDnowFraudSuspicionConfirmedRequest) ProtoMessage() {}

func (x *IDnowFraudSuspicionConfirmedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_v1_webhook_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDnowFraudSuspicionConfirmedRequest.ProtoReflect.Descriptor instead.
func (*IDnowFraudSuspicionConfirmedRequest) Descriptor() ([]byte, []int) {
	return file_webhook_v1_webhook_proto_rawDescGZIP(), []int{2}
}

func (x *IDnowFraudSuspicionConfirmedRequest) GetTenantName() string {
	if x != nil {
		return x.TenantName
	}
	return ""
}

func (x *IDnowFraudSuspicionConfirmedRequest) GetRequest() *v1.IDnowFraudSuspicionConfirmedPayload {
	if x != nil {
		return x.Request
	}
	return nil
}

type IDnowSuccessDataChangedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantName string                             `protobuf:"bytes,1,opt,name=tenant_name,json=tenantName,proto3" json:"tenant_name,omitempty"`
	Request    *v1.IDnowSuccessDataChangedPayload `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *IDnowSuccessDataChangedRequest) Reset() {
	*x = IDnowSuccessDataChangedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_v1_webhook_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDnowSuccessDataChangedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDnowSuccessDataChangedRequest) ProtoMessage() {}

func (x *IDnowSuccessDataChangedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_v1_webhook_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDnowSuccessDataChangedRequest.ProtoReflect.Descriptor instead.
func (*IDnowSuccessDataChangedRequest) Descriptor() ([]byte, []int) {
	return file_webhook_v1_webhook_proto_rawDescGZIP(), []int{3}
}

func (x *IDnowSuccessDataChangedRequest) GetTenantName() string {
	if x != nil {
		return x.TenantName
	}
	return ""
}

func (x *IDnowSuccessDataChangedRequest) GetRequest() *v1.IDnowSuccessDataChangedPayload {
	if x != nil {
		return x.Request
	}
	return nil
}

type IDnowFraudSuspicionPendingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantName string                                `protobuf:"bytes,1,opt,name=tenant_name,json=tenantName,proto3" json:"tenant_name,omitempty"`
	Request    *v1.IDnowFraudSuspicionPendingPayload `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *IDnowFraudSuspicionPendingRequest) Reset() {
	*x = IDnowFraudSuspicionPendingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_v1_webhook_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDnowFraudSuspicionPendingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDnowFraudSuspicionPendingRequest) ProtoMessage() {}

func (x *IDnowFraudSuspicionPendingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_v1_webhook_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDnowFraudSuspicionPendingRequest.ProtoReflect.Descriptor instead.
func (*IDnowFraudSuspicionPendingRequest) Descriptor() ([]byte, []int) {
	return file_webhook_v1_webhook_proto_rawDescGZIP(), []int{4}
}

func (x *IDnowFraudSuspicionPendingRequest) GetTenantName() string {
	if x != nil {
		return x.TenantName
	}
	return ""
}

func (x *IDnowFraudSuspicionPendingRequest) GetRequest() *v1.IDnowFraudSuspicionPendingPayload {
	if x != nil {
		return x.Request
	}
	return nil
}

type IDnowReviewPendingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantName string                        `protobuf:"bytes,1,opt,name=tenant_name,json=tenantName,proto3" json:"tenant_name,omitempty"`
	Request    *v1.IDnowReviewPendingPayload `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *IDnowReviewPendingRequest) Reset() {
	*x = IDnowReviewPendingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_v1_webhook_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDnowReviewPendingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDnowReviewPendingRequest) ProtoMessage() {}

func (x *IDnowReviewPendingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_v1_webhook_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDnowReviewPendingRequest.ProtoReflect.Descriptor instead.
func (*IDnowReviewPendingRequest) Descriptor() ([]byte, []int) {
	return file_webhook_v1_webhook_proto_rawDescGZIP(), []int{5}
}

func (x *IDnowReviewPendingRequest) GetTenantName() string {
	if x != nil {
		return x.TenantName
	}
	return ""
}

func (x *IDnowReviewPendingRequest) GetRequest() *v1.IDnowReviewPendingPayload {
	if x != nil {
		return x.Request
	}
	return nil
}

type IDnowSuccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantName string                  `protobuf:"bytes,1,opt,name=tenant_name,json=tenantName,proto3" json:"tenant_name,omitempty"`
	Request    *v1.IDnowSuccessPayload `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *IDnowSuccessRequest) Reset() {
	*x = IDnowSuccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_v1_webhook_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDnowSuccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDnowSuccessRequest) ProtoMessage() {}

func (x *IDnowSuccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_v1_webhook_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDnowSuccessRequest.ProtoReflect.Descriptor instead.
func (*IDnowSuccessRequest) Descriptor() ([]byte, []int) {
	return file_webhook_v1_webhook_proto_rawDescGZIP(), []int{6}
}

func (x *IDnowSuccessRequest) GetTenantName() string {
	if x != nil {
		return x.TenantName
	}
	return ""
}

func (x *IDnowSuccessRequest) GetRequest() *v1.IDnowSuccessPayload {
	if x != nil {
		return x.Request
	}
	return nil
}

var File_webhook_v1_webhook_proto protoreflect.FileDescriptor

var file_webhook_v1_webhook_proto_rawDesc = []byte{
	0x0a, 0x18, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x77, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62,
	0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x6e, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x72, 0x0a, 0x13, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x41, 0x62, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x74, 0x0a, 0x14, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x44, 0x6e, 0x6f,
	0x77, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x92, 0x01, 0x0a, 0x23, 0x49, 0x44,
	0x6e, 0x6f, 0x77, 0x46, 0x72, 0x61, 0x75, 0x64, 0x53, 0x75, 0x73, 0x70, 0x69, 0x63, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x46, 0x72, 0x61, 0x75, 0x64, 0x53, 0x75, 0x73, 0x70,
	0x69, 0x63, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x88,
	0x01, 0x0a, 0x1e, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x45, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x21, 0x49, 0x44,
	0x6e, 0x6f, 0x77, 0x46, 0x72, 0x61, 0x75, 0x64, 0x53, 0x75, 0x73, 0x70, 0x69, 0x63, 0x69, 0x6f,
	0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x48, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x44, 0x6e, 0x6f, 0x77, 0x46, 0x72, 0x61, 0x75, 0x64, 0x53, 0x75, 0x73, 0x70, 0x69, 0x63,
	0x69, 0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7e, 0x0a, 0x19, 0x49, 0x44,
	0x6e, 0x6f, 0x77, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x72, 0x0a, 0x13, 0x49, 0x44,
	0x6e, 0x6f, 0x77, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xa5,
	0x08, 0x0a, 0x0e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x7f, 0x0a, 0x0c, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x1f, 0x2e, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x44, 0x6e, 0x6f, 0x77, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2f, 0x3a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x24, 0x2f, 0x7b,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x77, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x2f, 0x69, 0x64, 0x6e, 0x6f, 0x77, 0x2f, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x92, 0x01, 0x0a, 0x12, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x2e, 0x77, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x36, 0x3a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b, 0x2f, 0x7b, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x2f, 0x69, 0x64, 0x6e, 0x6f, 0x77, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f,
	0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0xa5, 0x01, 0x0a, 0x1a, 0x49, 0x44, 0x6e, 0x6f,
	0x77, 0x46, 0x72, 0x61, 0x75, 0x64, 0x53, 0x75, 0x73, 0x70, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x50,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x2d, 0x2e, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x46, 0x72, 0x61, 0x75, 0x64, 0x53, 0x75,
	0x73, 0x70, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x3a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x2e, 0x2f, 0x7b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2f, 0x69, 0x64, 0x6e, 0x6f, 0x77, 0x2f, 0x73, 0x75,
	0x73, 0x70, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x9a, 0x01, 0x0a, 0x17, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x2a, 0x2e, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x3a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x29, 0x2f, 0x7b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2f, 0x69, 0x64, 0x6e, 0x6f, 0x77, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0xb1, 0x01, 0x0a,
	0x1c, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x46, 0x72, 0x61, 0x75, 0x64, 0x53, 0x75, 0x73, 0x70, 0x69,
	0x63, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x2f, 0x2e,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x44, 0x6e, 0x6f, 0x77,
	0x46, 0x72, 0x61, 0x75, 0x64, 0x53, 0x75, 0x73, 0x70, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x41, 0x3a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x36, 0x2f, 0x7b, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x2f, 0x69, 0x64, 0x6e, 0x6f, 0x77, 0x2f, 0x66, 0x72, 0x61, 0x75, 0x64, 0x5f, 0x73, 0x75, 0x73,
	0x70, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64,
	0x12, 0x82, 0x01, 0x0a, 0x0d, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x65, 0x64, 0x12, 0x20, 0x2e, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x44, 0x6e, 0x6f, 0x77, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x30, 0x3a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25,
	0x2f, 0x7b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x77,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2f, 0x69, 0x64, 0x6e, 0x6f, 0x77, 0x2f, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x65, 0x64, 0x12, 0x7f, 0x0a, 0x0c, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x41, 0x62,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x2e, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x44, 0x6e, 0x6f, 0x77, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x24, 0x2f, 0x7b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2f, 0x69, 0x64, 0x6e, 0x6f, 0x77, 0x2f, 0x61,
	0x62, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x42, 0xa9, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x77,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x57, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x6f, 0x62, 0x6f, 0x2d, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2f, 0x76,
	0x31, 0x3b, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x57, 0x58,
	0x58, 0xaa, 0x02, 0x0a, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0a, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x57, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_webhook_v1_webhook_proto_rawDescOnce sync.Once
	file_webhook_v1_webhook_proto_rawDescData = file_webhook_v1_webhook_proto_rawDesc
)

func file_webhook_v1_webhook_proto_rawDescGZIP() []byte {
	file_webhook_v1_webhook_proto_rawDescOnce.Do(func() {
		file_webhook_v1_webhook_proto_rawDescData = protoimpl.X.CompressGZIP(file_webhook_v1_webhook_proto_rawDescData)
	})
	return file_webhook_v1_webhook_proto_rawDescData
}

var file_webhook_v1_webhook_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_webhook_v1_webhook_proto_goTypes = []interface{}{
	(*IDnowAbortedRequest)(nil),                    // 0: webhook.v1.IDnowAbortedRequest
	(*IDnowCanceledRequest)(nil),                   // 1: webhook.v1.IDnowCanceledRequest
	(*IDnowFraudSuspicionConfirmedRequest)(nil),    // 2: webhook.v1.IDnowFraudSuspicionConfirmedRequest
	(*IDnowSuccessDataChangedRequest)(nil),         // 3: webhook.v1.IDnowSuccessDataChangedRequest
	(*IDnowFraudSuspicionPendingRequest)(nil),      // 4: webhook.v1.IDnowFraudSuspicionPendingRequest
	(*IDnowReviewPendingRequest)(nil),              // 5: webhook.v1.IDnowReviewPendingRequest
	(*IDnowSuccessRequest)(nil),                    // 6: webhook.v1.IDnowSuccessRequest
	(*v1.IDnowAbortedPayload)(nil),                 // 7: provider.v1.IDnowAbortedPayload
	(*v1.IDnowCanceledPayload)(nil),                // 8: provider.v1.IDnowCanceledPayload
	(*v1.IDnowFraudSuspicionConfirmedPayload)(nil), // 9: provider.v1.IDnowFraudSuspicionConfirmedPayload
	(*v1.IDnowSuccessDataChangedPayload)(nil),      // 10: provider.v1.IDnowSuccessDataChangedPayload
	(*v1.IDnowFraudSuspicionPendingPayload)(nil),   // 11: provider.v1.IDnowFraudSuspicionPendingPayload
	(*v1.IDnowReviewPendingPayload)(nil),           // 12: provider.v1.IDnowReviewPendingPayload
	(*v1.IDnowSuccessPayload)(nil),                 // 13: provider.v1.IDnowSuccessPayload
	(*v11.BoolResponse)(nil),                       // 14: common.v1.BoolResponse
}
var file_webhook_v1_webhook_proto_depIdxs = []int32{
	7,  // 0: webhook.v1.IDnowAbortedRequest.request:type_name -> provider.v1.IDnowAbortedPayload
	8,  // 1: webhook.v1.IDnowCanceledRequest.request:type_name -> provider.v1.IDnowCanceledPayload
	9,  // 2: webhook.v1.IDnowFraudSuspicionConfirmedRequest.request:type_name -> provider.v1.IDnowFraudSuspicionConfirmedPayload
	10, // 3: webhook.v1.IDnowSuccessDataChangedRequest.request:type_name -> provider.v1.IDnowSuccessDataChangedPayload
	11, // 4: webhook.v1.IDnowFraudSuspicionPendingRequest.request:type_name -> provider.v1.IDnowFraudSuspicionPendingPayload
	12, // 5: webhook.v1.IDnowReviewPendingRequest.request:type_name -> provider.v1.IDnowReviewPendingPayload
	13, // 6: webhook.v1.IDnowSuccessRequest.request:type_name -> provider.v1.IDnowSuccessPayload
	6,  // 7: webhook.v1.WebhookService.IDnowSuccess:input_type -> webhook.v1.IDnowSuccessRequest
	5,  // 8: webhook.v1.WebhookService.IDnowReviewPending:input_type -> webhook.v1.IDnowReviewPendingRequest
	4,  // 9: webhook.v1.WebhookService.IDnowFraudSuspicionPending:input_type -> webhook.v1.IDnowFraudSuspicionPendingRequest
	3,  // 10: webhook.v1.WebhookService.IDnowSuccessDataChanged:input_type -> webhook.v1.IDnowSuccessDataChangedRequest
	2,  // 11: webhook.v1.WebhookService.IDnowFraudSuspicionConfirmed:input_type -> webhook.v1.IDnowFraudSuspicionConfirmedRequest
	1,  // 12: webhook.v1.WebhookService.IDnowCanceled:input_type -> webhook.v1.IDnowCanceledRequest
	0,  // 13: webhook.v1.WebhookService.IDnowAborted:input_type -> webhook.v1.IDnowAbortedRequest
	14, // 14: webhook.v1.WebhookService.IDnowSuccess:output_type -> common.v1.BoolResponse
	14, // 15: webhook.v1.WebhookService.IDnowReviewPending:output_type -> common.v1.BoolResponse
	14, // 16: webhook.v1.WebhookService.IDnowFraudSuspicionPending:output_type -> common.v1.BoolResponse
	14, // 17: webhook.v1.WebhookService.IDnowSuccessDataChanged:output_type -> common.v1.BoolResponse
	14, // 18: webhook.v1.WebhookService.IDnowFraudSuspicionConfirmed:output_type -> common.v1.BoolResponse
	14, // 19: webhook.v1.WebhookService.IDnowCanceled:output_type -> common.v1.BoolResponse
	14, // 20: webhook.v1.WebhookService.IDnowAborted:output_type -> common.v1.BoolResponse
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_webhook_v1_webhook_proto_init() }
func file_webhook_v1_webhook_proto_init() {
	if File_webhook_v1_webhook_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_webhook_v1_webhook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDnowAbortedRequest); i {
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
		file_webhook_v1_webhook_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDnowCanceledRequest); i {
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
		file_webhook_v1_webhook_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDnowFraudSuspicionConfirmedRequest); i {
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
		file_webhook_v1_webhook_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDnowSuccessDataChangedRequest); i {
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
		file_webhook_v1_webhook_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDnowFraudSuspicionPendingRequest); i {
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
		file_webhook_v1_webhook_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDnowReviewPendingRequest); i {
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
		file_webhook_v1_webhook_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDnowSuccessRequest); i {
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
			RawDescriptor: file_webhook_v1_webhook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_webhook_v1_webhook_proto_goTypes,
		DependencyIndexes: file_webhook_v1_webhook_proto_depIdxs,
		MessageInfos:      file_webhook_v1_webhook_proto_msgTypes,
	}.Build()
	File_webhook_v1_webhook_proto = out.File
	file_webhook_v1_webhook_proto_rawDesc = nil
	file_webhook_v1_webhook_proto_goTypes = nil
	file_webhook_v1_webhook_proto_depIdxs = nil
}
