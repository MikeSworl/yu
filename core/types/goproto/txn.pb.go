// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: txn.proto

package goproto

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

type UnsignedTxn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Caller    []byte `protobuf:"bytes,1,opt,name=caller,proto3" json:"caller,omitempty"`
	Ecall     *Ecall `protobuf:"bytes,2,opt,name=ecall,proto3" json:"ecall,omitempty"`
	Timestamp uint64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *UnsignedTxn) Reset() {
	*x = UnsignedTxn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsignedTxn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsignedTxn) ProtoMessage() {}

func (x *UnsignedTxn) ProtoReflect() protoreflect.Message {
	mi := &file_txn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsignedTxn.ProtoReflect.Descriptor instead.
func (*UnsignedTxn) Descriptor() ([]byte, []int) {
	return file_txn_proto_rawDescGZIP(), []int{0}
}

func (x *UnsignedTxn) GetCaller() []byte {
	if x != nil {
		return x.Caller
	}
	return nil
}

func (x *UnsignedTxn) GetEcall() *Ecall {
	if x != nil {
		return x.Ecall
	}
	return nil
}

func (x *UnsignedTxn) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type SignedTxn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Raw       *UnsignedTxn `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	TxnHash   []byte       `protobuf:"bytes,2,opt,name=txn_hash,json=txnHash,proto3" json:"txn_hash,omitempty"`
	Pubkey    []byte       `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Signature []byte       `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignedTxn) Reset() {
	*x = SignedTxn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedTxn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedTxn) ProtoMessage() {}

func (x *SignedTxn) ProtoReflect() protoreflect.Message {
	mi := &file_txn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedTxn.ProtoReflect.Descriptor instead.
func (*SignedTxn) Descriptor() ([]byte, []int) {
	return file_txn_proto_rawDescGZIP(), []int{1}
}

func (x *SignedTxn) GetRaw() *UnsignedTxn {
	if x != nil {
		return x.Raw
	}
	return nil
}

func (x *SignedTxn) GetTxnHash() []byte {
	if x != nil {
		return x.TxnHash
	}
	return nil
}

func (x *SignedTxn) GetPubkey() []byte {
	if x != nil {
		return x.Pubkey
	}
	return nil
}

func (x *SignedTxn) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type SignedTxns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txns []*SignedTxn `protobuf:"bytes,1,rep,name=txns,proto3" json:"txns,omitempty"`
}

func (x *SignedTxns) Reset() {
	*x = SignedTxns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedTxns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedTxns) ProtoMessage() {}

func (x *SignedTxns) ProtoReflect() protoreflect.Message {
	mi := &file_txn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedTxns.ProtoReflect.Descriptor instead.
func (*SignedTxns) Descriptor() ([]byte, []int) {
	return file_txn_proto_rawDescGZIP(), []int{2}
}

func (x *SignedTxns) GetTxns() []*SignedTxn {
	if x != nil {
		return x.Txns
	}
	return nil
}

type Ecall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TripodName string `protobuf:"bytes,1,opt,name=tripod_name,json=tripodName,proto3" json:"tripod_name,omitempty"`
	ExecName   string `protobuf:"bytes,2,opt,name=exec_name,json=execName,proto3" json:"exec_name,omitempty"`
	Params     string `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	LeiPrice   uint64 `protobuf:"varint,4,opt,name=lei_price,json=leiPrice,proto3" json:"lei_price,omitempty"`
}

func (x *Ecall) Reset() {
	*x = Ecall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ecall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ecall) ProtoMessage() {}

func (x *Ecall) ProtoReflect() protoreflect.Message {
	mi := &file_txn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ecall.ProtoReflect.Descriptor instead.
func (*Ecall) Descriptor() ([]byte, []int) {
	return file_txn_proto_rawDescGZIP(), []int{3}
}

func (x *Ecall) GetTripodName() string {
	if x != nil {
		return x.TripodName
	}
	return ""
}

func (x *Ecall) GetExecName() string {
	if x != nil {
		return x.ExecName
	}
	return ""
}

func (x *Ecall) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

func (x *Ecall) GetLeiPrice() uint64 {
	if x != nil {
		return x.LeiPrice
	}
	return 0
}

type Qcall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TripodName string `protobuf:"bytes,1,opt,name=tripod_name,json=tripodName,proto3" json:"tripod_name,omitempty"`
	ExecName   string `protobuf:"bytes,2,opt,name=exec_name,json=execName,proto3" json:"exec_name,omitempty"`
	Params     string `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	BlockHash  []byte `protobuf:"bytes,4,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
}

func (x *Qcall) Reset() {
	*x = Qcall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Qcall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Qcall) ProtoMessage() {}

func (x *Qcall) ProtoReflect() protoreflect.Message {
	mi := &file_txn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Qcall.ProtoReflect.Descriptor instead.
func (*Qcall) Descriptor() ([]byte, []int) {
	return file_txn_proto_rawDescGZIP(), []int{4}
}

func (x *Qcall) GetTripodName() string {
	if x != nil {
		return x.TripodName
	}
	return ""
}

func (x *Qcall) GetExecName() string {
	if x != nil {
		return x.ExecName
	}
	return ""
}

func (x *Qcall) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

func (x *Qcall) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

type TxnsHashes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hashes [][]byte `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
}

func (x *TxnsHashes) Reset() {
	*x = TxnsHashes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txn_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxnsHashes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxnsHashes) ProtoMessage() {}

func (x *TxnsHashes) ProtoReflect() protoreflect.Message {
	mi := &file_txn_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxnsHashes.ProtoReflect.Descriptor instead.
func (*TxnsHashes) Descriptor() ([]byte, []int) {
	return file_txn_proto_rawDescGZIP(), []int{5}
}

func (x *TxnsHashes) GetHashes() [][]byte {
	if x != nil {
		return x.Hashes
	}
	return nil
}

type BatchSignedTxns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txns []*SignedTxn `protobuf:"bytes,1,rep,name=txns,proto3" json:"txns,omitempty"`
}

func (x *BatchSignedTxns) Reset() {
	*x = BatchSignedTxns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txn_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchSignedTxns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchSignedTxns) ProtoMessage() {}

func (x *BatchSignedTxns) ProtoReflect() protoreflect.Message {
	mi := &file_txn_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchSignedTxns.ProtoReflect.Descriptor instead.
func (*BatchSignedTxns) Descriptor() ([]byte, []int) {
	return file_txn_proto_rawDescGZIP(), []int{6}
}

func (x *BatchSignedTxns) GetTxns() []*SignedTxn {
	if x != nil {
		return x.Txns
	}
	return nil
}

type TxnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txn   *SignedTxn `protobuf:"bytes,1,opt,name=txn,proto3" json:"txn,omitempty"`
	Error string     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *TxnResponse) Reset() {
	*x = TxnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txn_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxnResponse) ProtoMessage() {}

func (x *TxnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_txn_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxnResponse.ProtoReflect.Descriptor instead.
func (*TxnResponse) Descriptor() ([]byte, []int) {
	return file_txn_proto_rawDescGZIP(), []int{7}
}

func (x *TxnResponse) GetTxn() *SignedTxn {
	if x != nil {
		return x.Txn
	}
	return nil
}

func (x *TxnResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type TxnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txn *SignedTxn `protobuf:"bytes,1,opt,name=txn,proto3" json:"txn,omitempty"`
}

func (x *TxnRequest) Reset() {
	*x = TxnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txn_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxnRequest) ProtoMessage() {}

func (x *TxnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_txn_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxnRequest.ProtoReflect.Descriptor instead.
func (*TxnRequest) Descriptor() ([]byte, []int) {
	return file_txn_proto_rawDescGZIP(), []int{8}
}

func (x *TxnRequest) GetTxn() *SignedTxn {
	if x != nil {
		return x.Txn
	}
	return nil
}

type TxnsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash []byte       `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	Txns      []*SignedTxn `protobuf:"bytes,2,rep,name=txns,proto3" json:"txns,omitempty"`
}

func (x *TxnsRequest) Reset() {
	*x = TxnsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txn_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxnsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxnsRequest) ProtoMessage() {}

func (x *TxnsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_txn_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxnsRequest.ProtoReflect.Descriptor instead.
func (*TxnsRequest) Descriptor() ([]byte, []int) {
	return file_txn_proto_rawDescGZIP(), []int{9}
}

func (x *TxnsRequest) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *TxnsRequest) GetTxns() []*SignedTxn {
	if x != nil {
		return x.Txns
	}
	return nil
}

type TxnsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txns  []*SignedTxn `protobuf:"bytes,1,rep,name=txns,proto3" json:"txns,omitempty"`
	Error string       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *TxnsResponse) Reset() {
	*x = TxnsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txn_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxnsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxnsResponse) ProtoMessage() {}

func (x *TxnsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_txn_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxnsResponse.ProtoReflect.Descriptor instead.
func (*TxnsResponse) Descriptor() ([]byte, []int) {
	return file_txn_proto_rawDescGZIP(), []int{10}
}

func (x *TxnsResponse) GetTxns() []*SignedTxn {
	if x != nil {
		return x.Txns
	}
	return nil
}

func (x *TxnsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_txn_proto protoreflect.FileDescriptor

var file_txn_proto_rawDesc = []byte{
	0x0a, 0x09, 0x74, 0x78, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0b, 0x55,
	0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x78, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x05, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x06, 0x2e, 0x45, 0x63, 0x61, 0x6c, 0x6c, 0x52, 0x05, 0x65, 0x63, 0x61, 0x6c, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x7c,
	0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x78, 0x6e, 0x12, 0x1e, 0x0a, 0x03, 0x72,
	0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x54, 0x78, 0x6e, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x78, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74,
	0x78, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x2c, 0x0a, 0x0a,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x78, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x78,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x54, 0x78, 0x6e, 0x52, 0x04, 0x74, 0x78, 0x6e, 0x73, 0x22, 0x7a, 0x0a, 0x05, 0x45, 0x63,
	0x61, 0x6c, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x69, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x69, 0x70, 0x6f, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x65, 0x63, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x69,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x65,
	0x69, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x7c, 0x0a, 0x05, 0x51, 0x63, 0x61, 0x6c, 0x6c, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x69, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x69, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x65, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x61, 0x73, 0x68, 0x22, 0x24, 0x0a, 0x0a, 0x54, 0x78, 0x6e, 0x73, 0x48, 0x61, 0x73, 0x68,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0x31, 0x0a, 0x0f, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x78, 0x6e, 0x73, 0x12, 0x1e, 0x0a,
	0x04, 0x74, 0x78, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x54, 0x78, 0x6e, 0x52, 0x04, 0x74, 0x78, 0x6e, 0x73, 0x22, 0x41, 0x0a,
	0x0b, 0x54, 0x78, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x03,
	0x74, 0x78, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x54, 0x78, 0x6e, 0x52, 0x03, 0x74, 0x78, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x2a, 0x0a, 0x0a, 0x54, 0x78, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x03, 0x74, 0x78, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x54, 0x78, 0x6e, 0x52, 0x03, 0x74, 0x78, 0x6e, 0x22, 0x4c, 0x0a, 0x0b,
	0x54, 0x78, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x78,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x54, 0x78, 0x6e, 0x52, 0x04, 0x74, 0x78, 0x6e, 0x73, 0x22, 0x44, 0x0a, 0x0c, 0x54, 0x78,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x78,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x54, 0x78, 0x6e, 0x52, 0x04, 0x74, 0x78, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_txn_proto_rawDescOnce sync.Once
	file_txn_proto_rawDescData = file_txn_proto_rawDesc
)

func file_txn_proto_rawDescGZIP() []byte {
	file_txn_proto_rawDescOnce.Do(func() {
		file_txn_proto_rawDescData = protoimpl.X.CompressGZIP(file_txn_proto_rawDescData)
	})
	return file_txn_proto_rawDescData
}

var file_txn_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_txn_proto_goTypes = []interface{}{
	(*UnsignedTxn)(nil),     // 0: UnsignedTxn
	(*SignedTxn)(nil),       // 1: SignedTxn
	(*SignedTxns)(nil),      // 2: SignedTxns
	(*Ecall)(nil),           // 3: Ecall
	(*Qcall)(nil),           // 4: Qcall
	(*TxnsHashes)(nil),      // 5: TxnsHashes
	(*BatchSignedTxns)(nil), // 6: BatchSignedTxns
	(*TxnResponse)(nil),     // 7: TxnResponse
	(*TxnRequest)(nil),      // 8: TxnRequest
	(*TxnsRequest)(nil),     // 9: TxnsRequest
	(*TxnsResponse)(nil),    // 10: TxnsResponse
}
var file_txn_proto_depIdxs = []int32{
	3, // 0: UnsignedTxn.ecall:type_name -> Ecall
	0, // 1: SignedTxn.raw:type_name -> UnsignedTxn
	1, // 2: SignedTxns.txns:type_name -> SignedTxn
	1, // 3: BatchSignedTxns.txns:type_name -> SignedTxn
	1, // 4: TxnResponse.txn:type_name -> SignedTxn
	1, // 5: TxnRequest.txn:type_name -> SignedTxn
	1, // 6: TxnsRequest.txns:type_name -> SignedTxn
	1, // 7: TxnsResponse.txns:type_name -> SignedTxn
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_txn_proto_init() }
func file_txn_proto_init() {
	if File_txn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_txn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsignedTxn); i {
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
		file_txn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedTxn); i {
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
		file_txn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedTxns); i {
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
		file_txn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ecall); i {
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
		file_txn_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Qcall); i {
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
		file_txn_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxnsHashes); i {
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
		file_txn_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchSignedTxns); i {
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
		file_txn_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxnResponse); i {
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
		file_txn_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxnRequest); i {
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
		file_txn_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxnsRequest); i {
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
		file_txn_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxnsResponse); i {
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
			RawDescriptor: file_txn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_txn_proto_goTypes,
		DependencyIndexes: file_txn_proto_depIdxs,
		MessageInfos:      file_txn_proto_msgTypes,
	}.Build()
	File_txn_proto = out.File
	file_txn_proto_rawDesc = nil
	file_txn_proto_goTypes = nil
	file_txn_proto_depIdxs = nil
}
