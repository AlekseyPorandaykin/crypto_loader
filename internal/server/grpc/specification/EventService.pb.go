// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: api/grpc/EventService.proto

package specification

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SymbolPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exchange string                 `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Symbol   string                 `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Price    float32                `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Date     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *SymbolPrice) Reset() {
	*x = SymbolPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_EventService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SymbolPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SymbolPrice) ProtoMessage() {}

func (x *SymbolPrice) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_EventService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SymbolPrice.ProtoReflect.Descriptor instead.
func (*SymbolPrice) Descriptor() ([]byte, []int) {
	return file_api_grpc_EventService_proto_rawDescGZIP(), []int{0}
}

func (x *SymbolPrice) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *SymbolPrice) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *SymbolPrice) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SymbolPrice) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type SymbolPrices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prices []*SymbolPrice `protobuf:"bytes,1,rep,name=prices,proto3" json:"prices,omitempty"`
}

func (x *SymbolPrices) Reset() {
	*x = SymbolPrices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_EventService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SymbolPrices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SymbolPrices) ProtoMessage() {}

func (x *SymbolPrices) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_EventService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SymbolPrices.ProtoReflect.Descriptor instead.
func (*SymbolPrices) Descriptor() ([]byte, []int) {
	return file_api_grpc_EventService_proto_rawDescGZIP(), []int{1}
}

func (x *SymbolPrices) GetPrices() []*SymbolPrice {
	if x != nil {
		return x.Prices
	}
	return nil
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_EventService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_EventService_proto_msgTypes[2]
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
	return file_api_grpc_EventService_proto_rawDescGZIP(), []int{2}
}

type DurationSeconds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Second int64 `protobuf:"varint,1,opt,name=second,proto3" json:"second,omitempty"`
}

func (x *DurationSeconds) Reset() {
	*x = DurationSeconds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_EventService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationSeconds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationSeconds) ProtoMessage() {}

func (x *DurationSeconds) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_EventService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationSeconds.ProtoReflect.Descriptor instead.
func (*DurationSeconds) Descriptor() ([]byte, []int) {
	return file_api_grpc_EventService_proto_rawDescGZIP(), []int{3}
}

func (x *DurationSeconds) GetSecond() int64 {
	if x != nil {
		return x.Second
	}
	return 0
}

var File_api_grpc_EventService_proto protoreflect.FileDescriptor

var file_api_grpc_EventService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x0b, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22,
	0x3a, 0x0a, 0x0c, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x2a, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x0f, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x32, 0x81, 0x01, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x13, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x53,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x1a, 0x13, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x30, 0x01, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f,
	0x3b, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_EventService_proto_rawDescOnce sync.Once
	file_api_grpc_EventService_proto_rawDescData = file_api_grpc_EventService_proto_rawDesc
)

func file_api_grpc_EventService_proto_rawDescGZIP() []byte {
	file_api_grpc_EventService_proto_rawDescOnce.Do(func() {
		file_api_grpc_EventService_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_EventService_proto_rawDescData)
	})
	return file_api_grpc_EventService_proto_rawDescData
}

var file_api_grpc_EventService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_grpc_EventService_proto_goTypes = []interface{}{
	(*SymbolPrice)(nil),           // 0: event.SymbolPrice
	(*SymbolPrices)(nil),          // 1: event.SymbolPrices
	(*EmptyRequest)(nil),          // 2: event.EmptyRequest
	(*DurationSeconds)(nil),       // 3: event.DurationSeconds
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_api_grpc_EventService_proto_depIdxs = []int32{
	4, // 0: event.SymbolPrice.date:type_name -> google.protobuf.Timestamp
	0, // 1: event.SymbolPrices.prices:type_name -> event.SymbolPrice
	2, // 2: event.EventService.Prices:input_type -> event.EmptyRequest
	3, // 3: event.EventService.TickerPrices:input_type -> event.DurationSeconds
	1, // 4: event.EventService.Prices:output_type -> event.SymbolPrices
	1, // 5: event.EventService.TickerPrices:output_type -> event.SymbolPrices
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_grpc_EventService_proto_init() }
func file_api_grpc_EventService_proto_init() {
	if File_api_grpc_EventService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_EventService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SymbolPrice); i {
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
		file_api_grpc_EventService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SymbolPrices); i {
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
		file_api_grpc_EventService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_grpc_EventService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationSeconds); i {
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
			RawDescriptor: file_api_grpc_EventService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_EventService_proto_goTypes,
		DependencyIndexes: file_api_grpc_EventService_proto_depIdxs,
		MessageInfos:      file_api_grpc_EventService_proto_msgTypes,
	}.Build()
	File_api_grpc_EventService_proto = out.File
	file_api_grpc_EventService_proto_rawDesc = nil
	file_api_grpc_EventService_proto_goTypes = nil
	file_api_grpc_EventService_proto_depIdxs = nil
}
