// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: arise.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SenderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *SenderRequest) Reset() {
	*x = SenderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arise_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SenderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SenderRequest) ProtoMessage() {}

func (x *SenderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_arise_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SenderRequest.ProtoReflect.Descriptor instead.
func (*SenderRequest) Descriptor() ([]byte, []int) {
	return file_arise_proto_rawDescGZIP(), []int{0}
}

func (x *SenderRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SenderRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type RecieverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *RecieverRequest) Reset() {
	*x = RecieverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arise_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecieverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecieverRequest) ProtoMessage() {}

func (x *RecieverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_arise_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecieverRequest.ProtoReflect.Descriptor instead.
func (*RecieverRequest) Descriptor() ([]byte, []int) {
	return file_arise_proto_rawDescGZIP(), []int{1}
}

func (x *RecieverRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type SenderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *SenderInfo) Reset() {
	*x = SenderInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arise_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SenderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SenderInfo) ProtoMessage() {}

func (x *SenderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_arise_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SenderInfo.ProtoReflect.Descriptor instead.
func (*SenderInfo) Descriptor() ([]byte, []int) {
	return file_arise_proto_rawDescGZIP(), []int{2}
}

func (x *SenderInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type RecieverInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *RecieverInfo) Reset() {
	*x = RecieverInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arise_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecieverInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecieverInfo) ProtoMessage() {}

func (x *RecieverInfo) ProtoReflect() protoreflect.Message {
	mi := &file_arise_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecieverInfo.ProtoReflect.Descriptor instead.
func (*RecieverInfo) Descriptor() ([]byte, []int) {
	return file_arise_proto_rawDescGZIP(), []int{3}
}

func (x *RecieverInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type SenderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *SenderResponse) Reset() {
	*x = SenderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arise_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SenderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SenderResponse) ProtoMessage() {}

func (x *SenderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_arise_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SenderResponse.ProtoReflect.Descriptor instead.
func (*SenderResponse) Descriptor() ([]byte, []int) {
	return file_arise_proto_rawDescGZIP(), []int{4}
}

func (x *SenderResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type RecieverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *RecieverResponse) Reset() {
	*x = RecieverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arise_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecieverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecieverResponse) ProtoMessage() {}

func (x *RecieverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_arise_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecieverResponse.ProtoReflect.Descriptor instead.
func (*RecieverResponse) Descriptor() ([]byte, []int) {
	return file_arise_proto_rawDescGZIP(), []int{5}
}

func (x *RecieverResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RecieverResponse) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arise_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_arise_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_arise_proto_rawDescGZIP(), []int{6}
}

func (x *SendResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Code) Reset() {
	*x = Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arise_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code) ProtoMessage() {}

func (x *Code) ProtoReflect() protoreflect.Message {
	mi := &file_arise_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code.ProtoReflect.Descriptor instead.
func (*Code) Descriptor() ([]byte, []int) {
	return file_arise_proto_rawDescGZIP(), []int{7}
}

func (x *Code) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arise_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_arise_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_arise_proto_rawDescGZIP(), []int{8}
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Code    string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arise_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_arise_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_arise_proto_rawDescGZIP(), []int{9}
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Chunk) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type RecieveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *RecieveResponse) Reset() {
	*x = RecieveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arise_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecieveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecieveResponse) ProtoMessage() {}

func (x *RecieveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_arise_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecieveResponse.ProtoReflect.Descriptor instead.
func (*RecieveResponse) Descriptor() ([]byte, []int) {
	return file_arise_proto_rawDescGZIP(), []int{10}
}

func (x *RecieveResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *PublicKey) Reset() {
	*x = PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arise_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey) ProtoMessage() {}

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_arise_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey.ProtoReflect.Descriptor instead.
func (*PublicKey) Descriptor() ([]byte, []int) {
	return file_arise_proto_rawDescGZIP(), []int{11}
}

func (x *PublicKey) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *PublicKey) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type PublicKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PublicKeyResponse) Reset() {
	*x = PublicKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arise_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKeyResponse) ProtoMessage() {}

func (x *PublicKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_arise_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKeyResponse.ProtoReflect.Descriptor instead.
func (*PublicKeyResponse) Descriptor() ([]byte, []int) {
	return file_arise_proto_rawDescGZIP(), []int{12}
}

func (x *PublicKeyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EncryptionKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *EncryptionKey) Reset() {
	*x = EncryptionKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arise_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptionKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptionKey) ProtoMessage() {}

func (x *EncryptionKey) ProtoReflect() protoreflect.Message {
	mi := &file_arise_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptionKey.ProtoReflect.Descriptor instead.
func (*EncryptionKey) Descriptor() ([]byte, []int) {
	return file_arise_proto_rawDescGZIP(), []int{13}
}

func (x *EncryptionKey) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *EncryptionKey) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type EncryptionKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EncryptionKeyResponse) Reset() {
	*x = EncryptionKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arise_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptionKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptionKeyResponse) ProtoMessage() {}

func (x *EncryptionKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_arise_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptionKeyResponse.ProtoReflect.Descriptor instead.
func (*EncryptionKeyResponse) Descriptor() ([]byte, []int) {
	return file_arise_proto_rawDescGZIP(), []int{14}
}

func (x *EncryptionKeyResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_arise_proto protoreflect.FileDescriptor

var file_arise_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x72, 0x69, 0x73, 0x65, 0x22, 0x37, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x25, 0x0a,
	0x0f, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x1c, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x22, 0x1e, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x22, 0x24, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x69,
	0x65, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x22, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x1a, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x35, 0x0a,
	0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x31, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x35, 0x0a, 0x0d, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x31, 0x0a, 0x15, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc1, 0x04,
	0x0a, 0x05, 0x41, 0x72, 0x69, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x14, 0x2e, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x08, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x61, 0x72, 0x69,
	0x73, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x65,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x0c, 0x2e, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x13, 0x2e, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x3f, 0x0a, 0x0b,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x72,
	0x69, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x69,
	0x65, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x33, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0b, 0x2e, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0x13, 0x2e,
	0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0b, 0x2e, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x1a, 0x11, 0x2e, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x12, 0x0b, 0x2e, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x1a, 0x10, 0x2e, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x12, 0x3c, 0x0a, 0x0e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x2e, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x1a, 0x18, 0x2e, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x12, 0x0b, 0x2e, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x1a, 0x14, 0x2e, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x48, 0x0a, 0x12, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x2e,
	0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x4b, 0x65, 0x79, 0x1a, 0x1c, 0x2e, 0x61, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_arise_proto_rawDescOnce sync.Once
	file_arise_proto_rawDescData = file_arise_proto_rawDesc
)

func file_arise_proto_rawDescGZIP() []byte {
	file_arise_proto_rawDescOnce.Do(func() {
		file_arise_proto_rawDescData = protoimpl.X.CompressGZIP(file_arise_proto_rawDescData)
	})
	return file_arise_proto_rawDescData
}

var file_arise_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_arise_proto_goTypes = []interface{}{
	(*SenderRequest)(nil),         // 0: arise.SenderRequest
	(*RecieverRequest)(nil),       // 1: arise.RecieverRequest
	(*SenderInfo)(nil),            // 2: arise.SenderInfo
	(*RecieverInfo)(nil),          // 3: arise.RecieverInfo
	(*SenderResponse)(nil),        // 4: arise.SenderResponse
	(*RecieverResponse)(nil),      // 5: arise.RecieverResponse
	(*SendResponse)(nil),          // 6: arise.SendResponse
	(*Code)(nil),                  // 7: arise.Code
	(*Empty)(nil),                 // 8: arise.Empty
	(*Chunk)(nil),                 // 9: arise.Chunk
	(*RecieveResponse)(nil),       // 10: arise.RecieveResponse
	(*PublicKey)(nil),             // 11: arise.PublicKey
	(*PublicKeyResponse)(nil),     // 12: arise.PublicKeyResponse
	(*EncryptionKey)(nil),         // 13: arise.EncryptionKey
	(*EncryptionKeyResponse)(nil), // 14: arise.EncryptionKeyResponse
}
var file_arise_proto_depIdxs = []int32{
	0,  // 0: arise.Arise.Sender:input_type -> arise.SenderRequest
	1,  // 1: arise.Arise.Reciever:input_type -> arise.RecieverRequest
	9,  // 2: arise.Arise.DataSend:input_type -> arise.Chunk
	1,  // 3: arise.Arise.DataRecieve:input_type -> arise.RecieverRequest
	7,  // 4: arise.Arise.GetRecieverInfo:input_type -> arise.Code
	7,  // 5: arise.Arise.GetSenderInfo:input_type -> arise.Code
	7,  // 6: arise.Arise.GetPublicKey:input_type -> arise.Code
	11, // 7: arise.Arise.SharePublicKey:input_type -> arise.PublicKey
	7,  // 8: arise.Arise.GetEncryptionKey:input_type -> arise.Code
	13, // 9: arise.Arise.ShareEncryptionKey:input_type -> arise.EncryptionKey
	4,  // 10: arise.Arise.Sender:output_type -> arise.SenderResponse
	5,  // 11: arise.Arise.Reciever:output_type -> arise.RecieverResponse
	6,  // 12: arise.Arise.DataSend:output_type -> arise.SendResponse
	10, // 13: arise.Arise.DataRecieve:output_type -> arise.RecieveResponse
	3,  // 14: arise.Arise.GetRecieverInfo:output_type -> arise.RecieverInfo
	2,  // 15: arise.Arise.GetSenderInfo:output_type -> arise.SenderInfo
	11, // 16: arise.Arise.GetPublicKey:output_type -> arise.PublicKey
	12, // 17: arise.Arise.SharePublicKey:output_type -> arise.PublicKeyResponse
	13, // 18: arise.Arise.GetEncryptionKey:output_type -> arise.EncryptionKey
	14, // 19: arise.Arise.ShareEncryptionKey:output_type -> arise.EncryptionKeyResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_arise_proto_init() }
func file_arise_proto_init() {
	if File_arise_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_arise_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SenderRequest); i {
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
		file_arise_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecieverRequest); i {
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
		file_arise_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SenderInfo); i {
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
		file_arise_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecieverInfo); i {
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
		file_arise_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SenderResponse); i {
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
		file_arise_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecieverResponse); i {
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
		file_arise_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResponse); i {
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
		file_arise_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code); i {
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
		file_arise_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_arise_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_arise_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecieveResponse); i {
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
		file_arise_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicKey); i {
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
		file_arise_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicKeyResponse); i {
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
		file_arise_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptionKey); i {
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
		file_arise_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptionKeyResponse); i {
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
			RawDescriptor: file_arise_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_arise_proto_goTypes,
		DependencyIndexes: file_arise_proto_depIdxs,
		MessageInfos:      file_arise_proto_msgTypes,
	}.Build()
	File_arise_proto = out.File
	file_arise_proto_rawDesc = nil
	file_arise_proto_goTypes = nil
	file_arise_proto_depIdxs = nil
}
