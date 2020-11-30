// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: pbft.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/relab/gorums"
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

type Pm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PP *PrePrepareArgs `protobuf:"bytes,1,opt,name=PP,proto3" json:"PP,omitempty"`
	P  []*PrepareArgs  `protobuf:"bytes,2,rep,name=P,proto3" json:"P,omitempty"`
}

func (x *Pm) Reset() {
	*x = Pm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pm) ProtoMessage() {}

func (x *Pm) ProtoReflect() protoreflect.Message {
	mi := &file_pbft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pm.ProtoReflect.Descriptor instead.
func (*Pm) Descriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{0}
}

func (x *Pm) GetPP() *PrePrepareArgs {
	if x != nil {
		return x.PP
	}
	return nil
}

func (x *Pm) GetP() []*PrepareArgs {
	if x != nil {
		return x.P
	}
	return nil
}

type CheckPointArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq    uint32 `protobuf:"varint,1,opt,name=Seq,proto3" json:"Seq,omitempty"`
	Digest []byte `protobuf:"bytes,2,opt,name=Digest,proto3" json:"Digest,omitempty"`
	Rid    uint32 `protobuf:"varint,3,opt,name=Rid,proto3" json:"Rid,omitempty"`
}

func (x *CheckPointArgs) Reset() {
	*x = CheckPointArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPointArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPointArgs) ProtoMessage() {}

func (x *CheckPointArgs) ProtoReflect() protoreflect.Message {
	mi := &file_pbft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPointArgs.ProtoReflect.Descriptor instead.
func (*CheckPointArgs) Descriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{1}
}

func (x *CheckPointArgs) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *CheckPointArgs) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *CheckPointArgs) GetRid() uint32 {
	if x != nil {
		return x.Rid
	}
	return 0
}

type CheckPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq    uint32            `protobuf:"varint,1,opt,name=Seq,proto3" json:"Seq,omitempty"`
	Stable bool              `protobuf:"varint,2,opt,name=Stable,proto3" json:"Stable,omitempty"`
	State  []byte            `protobuf:"bytes,3,opt,name=State,proto3" json:"State,omitempty"`
	Proof  []*CheckPointArgs `protobuf:"bytes,4,rep,name=Proof,proto3" json:"Proof,omitempty"`
}

func (x *CheckPoint) Reset() {
	*x = CheckPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPoint) ProtoMessage() {}

func (x *CheckPoint) ProtoReflect() protoreflect.Message {
	mi := &file_pbft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPoint.ProtoReflect.Descriptor instead.
func (*CheckPoint) Descriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{2}
}

func (x *CheckPoint) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *CheckPoint) GetStable() bool {
	if x != nil {
		return x.Stable
	}
	return false
}

func (x *CheckPoint) GetState() []byte {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *CheckPoint) GetProof() []*CheckPointArgs {
	if x != nil {
		return x.Proof
	}
	return nil
}

type ViewChangeArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	View uint32      `protobuf:"varint,1,opt,name=View,proto3" json:"View,omitempty"`
	CP   *CheckPoint `protobuf:"bytes,2,opt,name=CP,proto3" json:"CP,omitempty"`
	P    []*Pm       `protobuf:"bytes,3,rep,name=P,proto3" json:"P,omitempty"`
	Rid  uint32      `protobuf:"varint,4,opt,name=Rid,proto3" json:"Rid,omitempty"`
}

func (x *ViewChangeArgs) Reset() {
	*x = ViewChangeArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewChangeArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewChangeArgs) ProtoMessage() {}

func (x *ViewChangeArgs) ProtoReflect() protoreflect.Message {
	mi := &file_pbft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewChangeArgs.ProtoReflect.Descriptor instead.
func (*ViewChangeArgs) Descriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{3}
}

func (x *ViewChangeArgs) GetView() uint32 {
	if x != nil {
		return x.View
	}
	return 0
}

func (x *ViewChangeArgs) GetCP() *CheckPoint {
	if x != nil {
		return x.CP
	}
	return nil
}

func (x *ViewChangeArgs) GetP() []*Pm {
	if x != nil {
		return x.P
	}
	return nil
}

func (x *ViewChangeArgs) GetRid() uint32 {
	if x != nil {
		return x.Rid
	}
	return 0
}

type NewViewArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	View uint32            `protobuf:"varint,1,opt,name=View,proto3" json:"View,omitempty"`
	V    []*ViewChangeArgs `protobuf:"bytes,2,rep,name=V,proto3" json:"V,omitempty"`
	O    []*PrePrepareArgs `protobuf:"bytes,3,rep,name=O,proto3" json:"O,omitempty"`
}

func (x *NewViewArgs) Reset() {
	*x = NewViewArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbft_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewViewArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewViewArgs) ProtoMessage() {}

func (x *NewViewArgs) ProtoReflect() protoreflect.Message {
	mi := &file_pbft_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewViewArgs.ProtoReflect.Descriptor instead.
func (*NewViewArgs) Descriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{4}
}

func (x *NewViewArgs) GetView() uint32 {
	if x != nil {
		return x.View
	}
	return 0
}

func (x *NewViewArgs) GetV() []*ViewChangeArgs {
	if x != nil {
		return x.V
	}
	return nil
}

func (x *NewViewArgs) GetO() []*PrePrepareArgs {
	if x != nil {
		return x.O
	}
	return nil
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbft_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_pbft_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{5}
}

func (x *Command) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PrePrepareArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	View     uint32     `protobuf:"varint,1,opt,name=View,proto3" json:"View,omitempty"`
	Seq      uint32     `protobuf:"varint,2,opt,name=Seq,proto3" json:"Seq,omitempty"`
	Commands []*Command `protobuf:"bytes,3,rep,name=Commands,proto3" json:"Commands,omitempty"`
}

func (x *PrePrepareArgs) Reset() {
	*x = PrePrepareArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbft_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrePrepareArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrePrepareArgs) ProtoMessage() {}

func (x *PrePrepareArgs) ProtoReflect() protoreflect.Message {
	mi := &file_pbft_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrePrepareArgs.ProtoReflect.Descriptor instead.
func (*PrePrepareArgs) Descriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{6}
}

func (x *PrePrepareArgs) GetView() uint32 {
	if x != nil {
		return x.View
	}
	return 0
}

func (x *PrePrepareArgs) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *PrePrepareArgs) GetCommands() []*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

type PrepareArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	View   uint32 `protobuf:"varint,1,opt,name=View,proto3" json:"View,omitempty"`
	Seq    uint32 `protobuf:"varint,2,opt,name=Seq,proto3" json:"Seq,omitempty"`
	Digest []byte `protobuf:"bytes,3,opt,name=Digest,proto3" json:"Digest,omitempty"`
}

func (x *PrepareArgs) Reset() {
	*x = PrepareArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbft_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareArgs) ProtoMessage() {}

func (x *PrepareArgs) ProtoReflect() protoreflect.Message {
	mi := &file_pbft_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareArgs.ProtoReflect.Descriptor instead.
func (*PrepareArgs) Descriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{7}
}

func (x *PrepareArgs) GetView() uint32 {
	if x != nil {
		return x.View
	}
	return 0
}

func (x *PrepareArgs) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *PrepareArgs) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

type CommitArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	View   uint32 `protobuf:"varint,1,opt,name=View,proto3" json:"View,omitempty"`
	Seq    uint32 `protobuf:"varint,2,opt,name=Seq,proto3" json:"Seq,omitempty"`
	Digest []byte `protobuf:"bytes,3,opt,name=Digest,proto3" json:"Digest,omitempty"`
}

func (x *CommitArgs) Reset() {
	*x = CommitArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbft_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitArgs) ProtoMessage() {}

func (x *CommitArgs) ProtoReflect() protoreflect.Message {
	mi := &file_pbft_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitArgs.ProtoReflect.Descriptor instead.
func (*CommitArgs) Descriptor() ([]byte, []int) {
	return file_pbft_proto_rawDescGZIP(), []int{8}
}

func (x *CommitArgs) GetView() uint32 {
	if x != nil {
		return x.View
	}
	return 0
}

func (x *CommitArgs) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *CommitArgs) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

var File_pbft_proto protoreflect.FileDescriptor

var file_pbft_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x62, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6f, 0x72, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d,
	0x0a, 0x02, 0x50, 0x6d, 0x12, 0x25, 0x0a, 0x02, 0x50, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x65, 0x50, 0x72, 0x65, 0x70,
	0x61, 0x72, 0x65, 0x41, 0x72, 0x67, 0x73, 0x52, 0x02, 0x50, 0x50, 0x12, 0x20, 0x0a, 0x01, 0x50,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x41, 0x72, 0x67, 0x73, 0x52, 0x01, 0x50, 0x22, 0x4c, 0x0a,
	0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x53, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x53, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x52, 0x69, 0x64, 0x22, 0x79, 0x0a, 0x0a, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x71,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x53, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x41, 0x72, 0x67, 0x73, 0x52,
	0x05, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x72, 0x0a, 0x0e, 0x56, 0x69, 0x65, 0x77, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x21, 0x0a, 0x02,
	0x43, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x02, 0x43, 0x50, 0x12,
	0x17, 0x0a, 0x01, 0x50, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x6d, 0x52, 0x01, 0x50, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x52, 0x69, 0x64, 0x22, 0x6b, 0x0a, 0x0b, 0x4e, 0x65,
	0x77, 0x56, 0x69, 0x65, 0x77, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x56, 0x69, 0x65,
	0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x23, 0x0a,
	0x01, 0x56, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x69, 0x65, 0x77, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x72, 0x67, 0x73, 0x52,
	0x01, 0x56, 0x12, 0x23, 0x0a, 0x01, 0x4f, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x65, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65,
	0x41, 0x72, 0x67, 0x73, 0x52, 0x01, 0x4f, 0x22, 0x1d, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x62, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x50, 0x72, 0x65,
	0x70, 0x61, 0x72, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x10, 0x0a, 0x03,
	0x53, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x53, 0x65, 0x71, 0x12, 0x2a,
	0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x4b, 0x0a, 0x0b, 0x50, 0x72,
	0x65, 0x70, 0x61, 0x72, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x56, 0x69, 0x65,
	0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x10, 0x0a,
	0x03, 0x53, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x53, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x06, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x71,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x53, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x44,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x44, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x32, 0xc1, 0x02, 0x0a, 0x04, 0x50, 0x42, 0x46, 0x54, 0x12, 0x41, 0x0a, 0x0a,
	0x50, 0x72, 0x65, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x72, 0x65, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x41, 0x72, 0x67,
	0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0x98, 0xb5, 0x18, 0x01, 0x12,
	0x3b, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0x98, 0xb5, 0x18, 0x01, 0x12, 0x39, 0x0a, 0x06,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x04, 0x98, 0xb5, 0x18, 0x01, 0x12, 0x41, 0x0a, 0x0a, 0x56, 0x69, 0x65, 0x77, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69,
	0x65, 0x77, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0x98, 0xb5, 0x18, 0x01, 0x12, 0x3b, 0x0a, 0x07, 0x4e, 0x65,
	0x77, 0x56, 0x69, 0x65, 0x77, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65,
	0x77, 0x56, 0x69, 0x65, 0x77, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x04, 0x98, 0xb5, 0x18, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x65, 0x2d, 0x7a, 0x78, 0x68, 0x2f, 0x70, 0x62,
	0x66, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbft_proto_rawDescOnce sync.Once
	file_pbft_proto_rawDescData = file_pbft_proto_rawDesc
)

func file_pbft_proto_rawDescGZIP() []byte {
	file_pbft_proto_rawDescOnce.Do(func() {
		file_pbft_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbft_proto_rawDescData)
	})
	return file_pbft_proto_rawDescData
}

var file_pbft_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pbft_proto_goTypes = []interface{}{
	(*Pm)(nil),             // 0: proto.Pm
	(*CheckPointArgs)(nil), // 1: proto.CheckPointArgs
	(*CheckPoint)(nil),     // 2: proto.CheckPoint
	(*ViewChangeArgs)(nil), // 3: proto.ViewChangeArgs
	(*NewViewArgs)(nil),    // 4: proto.NewViewArgs
	(*Command)(nil),        // 5: proto.Command
	(*PrePrepareArgs)(nil), // 6: proto.PrePrepareArgs
	(*PrepareArgs)(nil),    // 7: proto.PrepareArgs
	(*CommitArgs)(nil),     // 8: proto.CommitArgs
	(*empty.Empty)(nil),    // 9: google.protobuf.Empty
}
var file_pbft_proto_depIdxs = []int32{
	6,  // 0: proto.Pm.PP:type_name -> proto.PrePrepareArgs
	7,  // 1: proto.Pm.P:type_name -> proto.PrepareArgs
	1,  // 2: proto.CheckPoint.Proof:type_name -> proto.CheckPointArgs
	2,  // 3: proto.ViewChangeArgs.CP:type_name -> proto.CheckPoint
	0,  // 4: proto.ViewChangeArgs.P:type_name -> proto.Pm
	3,  // 5: proto.NewViewArgs.V:type_name -> proto.ViewChangeArgs
	6,  // 6: proto.NewViewArgs.O:type_name -> proto.PrePrepareArgs
	5,  // 7: proto.PrePrepareArgs.Commands:type_name -> proto.Command
	6,  // 8: proto.PBFT.PrePrepare:input_type -> proto.PrePrepareArgs
	7,  // 9: proto.PBFT.Prepare:input_type -> proto.PrepareArgs
	8,  // 10: proto.PBFT.Commit:input_type -> proto.CommitArgs
	3,  // 11: proto.PBFT.ViewChange:input_type -> proto.ViewChangeArgs
	4,  // 12: proto.PBFT.NewView:input_type -> proto.NewViewArgs
	9,  // 13: proto.PBFT.PrePrepare:output_type -> google.protobuf.Empty
	9,  // 14: proto.PBFT.Prepare:output_type -> google.protobuf.Empty
	9,  // 15: proto.PBFT.Commit:output_type -> google.protobuf.Empty
	9,  // 16: proto.PBFT.ViewChange:output_type -> google.protobuf.Empty
	9,  // 17: proto.PBFT.NewView:output_type -> google.protobuf.Empty
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_pbft_proto_init() }
func file_pbft_proto_init() {
	if File_pbft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pm); i {
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
		file_pbft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPointArgs); i {
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
		file_pbft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPoint); i {
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
		file_pbft_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewChangeArgs); i {
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
		file_pbft_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewViewArgs); i {
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
		file_pbft_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_pbft_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrePrepareArgs); i {
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
		file_pbft_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareArgs); i {
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
		file_pbft_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitArgs); i {
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
			RawDescriptor: file_pbft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pbft_proto_goTypes,
		DependencyIndexes: file_pbft_proto_depIdxs,
		MessageInfos:      file_pbft_proto_msgTypes,
	}.Build()
	File_pbft_proto = out.File
	file_pbft_proto_rawDesc = nil
	file_pbft_proto_goTypes = nil
	file_pbft_proto_depIdxs = nil
}
