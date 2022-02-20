// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: blockbase.proto

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

type EventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *EventsRequest) Reset() {
	*x = EventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockbase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsRequest) ProtoMessage() {}

func (x *EventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blockbase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsRequest.ProtoReflect.Descriptor instead.
func (*EventsRequest) Descriptor() ([]byte, []int) {
	return file_blockbase_proto_rawDescGZIP(), []int{0}
}

func (x *EventsRequest) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type EventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	Error  string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *EventsResponse) Reset() {
	*x = EventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockbase_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsResponse) ProtoMessage() {}

func (x *EventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blockbase_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsResponse.ProtoReflect.Descriptor instead.
func (*EventsResponse) Descriptor() ([]byte, []int) {
	return file_blockbase_proto_rawDescGZIP(), []int{1}
}

func (x *EventsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *EventsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ErrorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []*Error `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	Error  string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ErrorsResponse) Reset() {
	*x = ErrorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockbase_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorsResponse) ProtoMessage() {}

func (x *ErrorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blockbase_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorsResponse.ProtoReflect.Descriptor instead.
func (*ErrorsResponse) Descriptor() ([]byte, []int) {
	return file_blockbase_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorsResponse) GetErrors() []*Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *ErrorsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_blockbase_proto protoreflect.FileDescriptor

var file_blockbase_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x09, 0x74, 0x78, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0d,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x46, 0x0a,
	0x0e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x46, 0x0a, 0x0e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x9f, 0x02,
	0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x61, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x54, 0x78, 0x6e, 0x12, 0x08, 0x2e, 0x54, 0x78, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x1a,
	0x0c, 0x2e, 0x54, 0x78, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x06, 0x53, 0x65, 0x74, 0x54, 0x78, 0x6e, 0x12, 0x0a, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x54, 0x78, 0x6e, 0x1a, 0x04, 0x2e, 0x45, 0x72, 0x72, 0x12, 0x24, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x54, 0x78, 0x6e, 0x73, 0x12, 0x0a, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68,
	0x1a, 0x0d, 0x2e, 0x54, 0x78, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x54, 0x78, 0x6e, 0x73, 0x12, 0x0c, 0x2e, 0x54, 0x78, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x04, 0x2e, 0x45, 0x72, 0x72, 0x12, 0x28,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0a, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x0f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x04, 0x2e, 0x45, 0x72, 0x72, 0x12, 0x28, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x0a, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x61, 0x73, 0x68, 0x1a, 0x0f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x04, 0x2e, 0x45, 0x72, 0x72, 0x42,
	0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockbase_proto_rawDescOnce sync.Once
	file_blockbase_proto_rawDescData = file_blockbase_proto_rawDesc
)

func file_blockbase_proto_rawDescGZIP() []byte {
	file_blockbase_proto_rawDescOnce.Do(func() {
		file_blockbase_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockbase_proto_rawDescData)
	})
	return file_blockbase_proto_rawDescData
}

var file_blockbase_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_blockbase_proto_goTypes = []interface{}{
	(*EventsRequest)(nil),  // 0: EventsRequest
	(*EventsResponse)(nil), // 1: EventsResponse
	(*ErrorsResponse)(nil), // 2: ErrorsResponse
	(*Event)(nil),          // 3: Event
	(*Error)(nil),          // 4: Error
	(*TxnHash)(nil),        // 5: TxnHash
	(*SignedTxn)(nil),      // 6: SignedTxn
	(*BlockHash)(nil),      // 7: BlockHash
	(*TxnsRequest)(nil),    // 8: TxnsRequest
	(*TxnResponse)(nil),    // 9: TxnResponse
	(*Err)(nil),            // 10: Err
	(*TxnsResponse)(nil),   // 11: TxnsResponse
}
var file_blockbase_proto_depIdxs = []int32{
	3,  // 0: EventsRequest.events:type_name -> Event
	3,  // 1: EventsResponse.events:type_name -> Event
	4,  // 2: ErrorsResponse.errors:type_name -> Error
	5,  // 3: YuDB.GetTxn:input_type -> TxnHash
	6,  // 4: YuDB.SetTxn:input_type -> SignedTxn
	7,  // 5: YuDB.GetTxns:input_type -> BlockHash
	8,  // 6: YuDB.SetTxns:input_type -> TxnsRequest
	7,  // 7: YuDB.GetEvents:input_type -> BlockHash
	0,  // 8: YuDB.SetEvents:input_type -> EventsRequest
	7,  // 9: YuDB.GetErrors:input_type -> BlockHash
	4,  // 10: YuDB.SetError:input_type -> Error
	9,  // 11: YuDB.GetTxn:output_type -> TxnResponse
	10, // 12: YuDB.SetTxn:output_type -> Err
	11, // 13: YuDB.GetTxns:output_type -> TxnsResponse
	10, // 14: YuDB.SetTxns:output_type -> Err
	1,  // 15: YuDB.GetEvents:output_type -> EventsResponse
	10, // 16: YuDB.SetEvents:output_type -> Err
	2,  // 17: YuDB.GetErrors:output_type -> ErrorsResponse
	10, // 18: YuDB.SetError:output_type -> Err
	11, // [11:19] is the sub-list for method output_type
	3,  // [3:11] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_blockbase_proto_init() }
func file_blockbase_proto_init() {
	if File_blockbase_proto != nil {
		return
	}
	file_txn_proto_init()
	file_result_proto_init()
	file_base_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_blockbase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventsRequest); i {
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
		file_blockbase_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventsResponse); i {
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
		file_blockbase_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorsResponse); i {
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
			RawDescriptor: file_blockbase_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blockbase_proto_goTypes,
		DependencyIndexes: file_blockbase_proto_depIdxs,
		MessageInfos:      file_blockbase_proto_msgTypes,
	}.Build()
	File_blockbase_proto = out.File
	file_blockbase_proto_rawDesc = nil
	file_blockbase_proto_goTypes = nil
	file_blockbase_proto_depIdxs = nil
}
