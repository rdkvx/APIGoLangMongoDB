// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1--rc1
// source: desafiotecnico.proto

package proto

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

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desafiotecnico_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_desafiotecnico_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_desafiotecnico_proto_rawDescGZIP(), []int{0}
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desafiotecnico_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_desafiotecnico_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_desafiotecnico_proto_rawDescGZIP(), []int{1}
}

type Cryptocurrency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol    string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Votes     int32  `protobuf:"varint,4,opt,name=votes,proto3" json:"votes,omitempty"`
	Createdat string `protobuf:"bytes,5,opt,name=createdat,proto3" json:"createdat,omitempty"`
	Updatedat  string `protobuf:"bytes,6,opt,name=Updatedat,proto3" json:"Updatedat,omitempty"`
}

func (x *Cryptocurrency) Reset() {
	*x = Cryptocurrency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desafiotecnico_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cryptocurrency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cryptocurrency) ProtoMessage() {}

func (x *Cryptocurrency) ProtoReflect() protoreflect.Message {
	mi := &file_desafiotecnico_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cryptocurrency.ProtoReflect.Descriptor instead.
func (*Cryptocurrency) Descriptor() ([]byte, []int) {
	return file_desafiotecnico_proto_rawDescGZIP(), []int{2}
}

func (x *Cryptocurrency) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cryptocurrency) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cryptocurrency) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Cryptocurrency) GetVotes() int32 {
	if x != nil {
		return x.Votes
	}
	return 0
}

func (x *Cryptocurrency) GetCreatedat() string {
	if x != nil {
		return x.Createdat
	}
	return ""
}

func (x *Cryptocurrency) GetUpdateat() string {
	if x != nil {
		return x.Updatedat
	}
	return ""
}

type ListCryptosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sortparam string `protobuf:"bytes,1,opt,name=sortparam,proto3" json:"sortparam,omitempty"`
	Ascending bool   `protobuf:"varint,2,opt,name=ascending,proto3" json:"ascending,omitempty"`
}

func (x *ListCryptosRequest) Reset() {
	*x = ListCryptosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desafiotecnico_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCryptosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCryptosRequest) ProtoMessage() {}

func (x *ListCryptosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_desafiotecnico_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCryptosRequest.ProtoReflect.Descriptor instead.
func (*ListCryptosRequest) Descriptor() ([]byte, []int) {
	return file_desafiotecnico_proto_rawDescGZIP(), []int{3}
}

func (x *ListCryptosRequest) GetSortparam() string {
	if x != nil {
		return x.Sortparam
	}
	return ""
}

func (x *ListCryptosRequest) GetAscending() bool {
	if x != nil {
		return x.Ascending
	}
	return false
}

type ListCryptosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crypto []*Cryptocurrency `protobuf:"bytes,1,rep,name=crypto,proto3" json:"crypto,omitempty"`
}

func (x *ListCryptosResponse) Reset() {
	*x = ListCryptosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desafiotecnico_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCryptosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCryptosResponse) ProtoMessage() {}

func (x *ListCryptosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_desafiotecnico_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCryptosResponse.ProtoReflect.Descriptor instead.
func (*ListCryptosResponse) Descriptor() ([]byte, []int) {
	return file_desafiotecnico_proto_rawDescGZIP(), []int{4}
}

func (x *ListCryptosResponse) GetCrypto() []*Cryptocurrency {
	if x != nil {
		return x.Crypto
	}
	return nil
}

type NewCryptoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *NewCryptoRequest) Reset() {
	*x = NewCryptoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desafiotecnico_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCryptoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCryptoRequest) ProtoMessage() {}

func (x *NewCryptoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_desafiotecnico_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCryptoRequest.ProtoReflect.Descriptor instead.
func (*NewCryptoRequest) Descriptor() ([]byte, []int) {
	return file_desafiotecnico_proto_rawDescGZIP(), []int{5}
}

func (x *NewCryptoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewCryptoRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type EditCryptoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *EditCryptoRequest) Reset() {
	*x = EditCryptoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desafiotecnico_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditCryptoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditCryptoRequest) ProtoMessage() {}

func (x *EditCryptoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_desafiotecnico_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditCryptoRequest.ProtoReflect.Descriptor instead.
func (*EditCryptoRequest) Descriptor() ([]byte, []int) {
	return file_desafiotecnico_proto_rawDescGZIP(), []int{6}
}

func (x *EditCryptoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EditCryptoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EditCryptoRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type DeleteCryptoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCryptoRequest) Reset() {
	*x = DeleteCryptoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desafiotecnico_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCryptoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCryptoRequest) ProtoMessage() {}

func (x *DeleteCryptoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_desafiotecnico_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCryptoRequest.ProtoReflect.Descriptor instead.
func (*DeleteCryptoRequest) Descriptor() ([]byte, []int) {
	return file_desafiotecnico_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCryptoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type VoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *VoteRequest) Reset() {
	*x = VoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desafiotecnico_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequest) ProtoMessage() {}

func (x *VoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_desafiotecnico_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequest.ProtoReflect.Descriptor instead.
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return file_desafiotecnico_proto_rawDescGZIP(), []int{8}
}

func (x *VoteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SubscriptionRequest) Reset() {
	*x = SubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desafiotecnico_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionRequest) ProtoMessage() {}

func (x *SubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_desafiotecnico_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_desafiotecnico_proto_rawDescGZIP(), []int{9}
}

func (x *SubscriptionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindRequest) Reset() {
	*x = FindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_desafiotecnico_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRequest) ProtoMessage() {}

func (x *FindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_desafiotecnico_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRequest.ProtoReflect.Descriptor instead.
func (*FindRequest) Descriptor() ([]byte, []int) {
	return file_desafiotecnico_proto_rawDescGZIP(), []int{10}
}

func (x *FindRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_desafiotecnico_proto protoreflect.FileDescriptor

var file_desafiotecnico_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x65, 0x73, 0x61, 0x66, 0x69, 0x6f, 0x74, 0x65, 0x63, 0x6e, 0x69, 0x63, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a,
	0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e,
	0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9c,
	0x01, 0x0a, 0x0e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x6f,
	0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x61,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x61, 0x74, 0x22, 0x50, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22,
	0x44, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x06, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x10, 0x4e, 0x65, 0x77, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x4f, 0x0a, 0x11, 0x45, 0x64, 0x69, 0x74, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a,
	0x0b, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x13,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x32, 0xec, 0x03, 0x0a, 0x0d, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x04, 0x45, 0x64, 0x69, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x64, 0x69, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x04, 0x46, 0x69, 0x6e,
	0x64, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x34, 0x0a, 0x06, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74,
	0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a,
	0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x16, 0x5a, 0x14, 0x44, 0x65, 0x73, 0x61, 0x66, 0x69, 0x6f, 0x54, 0x65, 0x63, 0x6e,
	0x69, 0x63, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_desafiotecnico_proto_rawDescOnce sync.Once
	file_desafiotecnico_proto_rawDescData = file_desafiotecnico_proto_rawDesc
)

func file_desafiotecnico_proto_rawDescGZIP() []byte {
	file_desafiotecnico_proto_rawDescOnce.Do(func() {
		file_desafiotecnico_proto_rawDescData = protoimpl.X.CompressGZIP(file_desafiotecnico_proto_rawDescData)
	})
	return file_desafiotecnico_proto_rawDescData
}

var file_desafiotecnico_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_desafiotecnico_proto_goTypes = []interface{}{
	(*EmptyResponse)(nil),       // 0: proto.EmptyResponse
	(*EmptyRequest)(nil),        // 1: proto.EmptyRequest
	(*Cryptocurrency)(nil),      // 2: proto.Cryptocurrency
	(*ListCryptosRequest)(nil),  // 3: proto.ListCryptosRequest
	(*ListCryptosResponse)(nil), // 4: proto.ListCryptosResponse
	(*NewCryptoRequest)(nil),    // 5: proto.NewCryptoRequest
	(*EditCryptoRequest)(nil),   // 6: proto.EditCryptoRequest
	(*DeleteCryptoRequest)(nil), // 7: proto.DeleteCryptoRequest
	(*VoteRequest)(nil),         // 8: proto.VoteRequest
	(*SubscriptionRequest)(nil), // 9: proto.SubscriptionRequest
	(*FindRequest)(nil),         // 10: proto.FindRequest
}
var file_desafiotecnico_proto_depIdxs = []int32{
	2,  // 0: proto.ListCryptosResponse.crypto:type_name -> proto.Cryptocurrency
	5,  // 1: proto.CryptoService.Create:input_type -> proto.NewCryptoRequest
	6,  // 2: proto.CryptoService.Edit:input_type -> proto.EditCryptoRequest
	7,  // 3: proto.CryptoService.Delete:input_type -> proto.DeleteCryptoRequest
	10, // 4: proto.CryptoService.Find:input_type -> proto.FindRequest
	3,  // 5: proto.CryptoService.List:input_type -> proto.ListCryptosRequest
	8,  // 6: proto.CryptoService.Upvote:input_type -> proto.VoteRequest
	8,  // 7: proto.CryptoService.Downvote:input_type -> proto.VoteRequest
	9,  // 8: proto.CryptoService.Subscribe:input_type -> proto.SubscriptionRequest
	2,  // 9: proto.CryptoService.Create:output_type -> proto.Cryptocurrency
	2,  // 10: proto.CryptoService.Edit:output_type -> proto.Cryptocurrency
	0,  // 11: proto.CryptoService.Delete:output_type -> proto.EmptyResponse
	2,  // 12: proto.CryptoService.Find:output_type -> proto.Cryptocurrency
	4,  // 13: proto.CryptoService.List:output_type -> proto.ListCryptosResponse
	0,  // 14: proto.CryptoService.Upvote:output_type -> proto.EmptyResponse
	0,  // 15: proto.CryptoService.Downvote:output_type -> proto.EmptyResponse
	2,  // 16: proto.CryptoService.Subscribe:output_type -> proto.Cryptocurrency
	9,  // [9:17] is the sub-list for method output_type
	1,  // [1:9] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_desafiotecnico_proto_init() }
func file_desafiotecnico_proto_init() {
	if File_desafiotecnico_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_desafiotecnico_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
		file_desafiotecnico_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_desafiotecnico_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cryptocurrency); i {
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
		file_desafiotecnico_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCryptosRequest); i {
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
		file_desafiotecnico_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCryptosResponse); i {
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
		file_desafiotecnico_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCryptoRequest); i {
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
		file_desafiotecnico_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditCryptoRequest); i {
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
		file_desafiotecnico_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCryptoRequest); i {
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
		file_desafiotecnico_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteRequest); i {
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
		file_desafiotecnico_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionRequest); i {
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
		file_desafiotecnico_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRequest); i {
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
			RawDescriptor: file_desafiotecnico_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_desafiotecnico_proto_goTypes,
		DependencyIndexes: file_desafiotecnico_proto_depIdxs,
		MessageInfos:      file_desafiotecnico_proto_msgTypes,
	}.Build()
	File_desafiotecnico_proto = out.File
	file_desafiotecnico_proto_rawDesc = nil
	file_desafiotecnico_proto_goTypes = nil
	file_desafiotecnico_proto_depIdxs = nil
}
