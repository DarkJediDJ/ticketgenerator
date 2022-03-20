// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: internal/proto/generator.proto

package ticketgenerator

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

type TicketRequset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time  string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Price int64  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	Seat  int64  `protobuf:"varint,3,opt,name=seat,proto3" json:"seat,omitempty"`
	Id    int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *TicketRequset) Reset() {
	*x = TicketRequset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_generator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketRequset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketRequset) ProtoMessage() {}

func (x *TicketRequset) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_generator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketRequset.ProtoReflect.Descriptor instead.
func (*TicketRequset) Descriptor() ([]byte, []int) {
	return file_internal_proto_generator_proto_rawDescGZIP(), []int{0}
}

func (x *TicketRequset) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *TicketRequset) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *TicketRequset) GetSeat() int64 {
	if x != nil {
		return x.Seat
	}
	return 0
}

func (x *TicketRequset) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TicketRequset) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type IDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *IDReply) Reset() {
	*x = IDReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_generator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDReply) ProtoMessage() {}

func (x *IDReply) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_generator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDReply.ProtoReflect.Descriptor instead.
func (*IDReply) Descriptor() ([]byte, []int) {
	return file_internal_proto_generator_proto_rawDescGZIP(), []int{1}
}

func (x *IDReply) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

var File_internal_proto_generator_proto protoreflect.FileDescriptor

var file_internal_proto_generator_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x73,
	0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x73, 0x65, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x65, 0x61, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x22, 0x19, 0x0a, 0x07, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x32, 0x54,
	0x0a, 0x0f, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x41, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x73, 0x65, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x72, 0x6b, 0x6a, 0x65, 0x64, 0x69, 0x64, 0x6a, 0x2f, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_generator_proto_rawDescOnce sync.Once
	file_internal_proto_generator_proto_rawDescData = file_internal_proto_generator_proto_rawDesc
)

func file_internal_proto_generator_proto_rawDescGZIP() []byte {
	file_internal_proto_generator_proto_rawDescOnce.Do(func() {
		file_internal_proto_generator_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_generator_proto_rawDescData)
	})
	return file_internal_proto_generator_proto_rawDescData
}

var file_internal_proto_generator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_proto_generator_proto_goTypes = []interface{}{
	(*TicketRequset)(nil), // 0: transactions.TicketRequset
	(*IDReply)(nil),       // 1: transactions.IDReply
}
var file_internal_proto_generator_proto_depIdxs = []int32{
	0, // 0: transactions.TicketGenerator.GetTicket:input_type -> transactions.TicketRequset
	1, // 1: transactions.TicketGenerator.GetTicket:output_type -> transactions.IDReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_generator_proto_init() }
func file_internal_proto_generator_proto_init() {
	if File_internal_proto_generator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_generator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketRequset); i {
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
		file_internal_proto_generator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDReply); i {
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
			RawDescriptor: file_internal_proto_generator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_generator_proto_goTypes,
		DependencyIndexes: file_internal_proto_generator_proto_depIdxs,
		MessageInfos:      file_internal_proto_generator_proto_msgTypes,
	}.Build()
	File_internal_proto_generator_proto = out.File
	file_internal_proto_generator_proto_rawDesc = nil
	file_internal_proto_generator_proto_goTypes = nil
	file_internal_proto_generator_proto_depIdxs = nil
}
