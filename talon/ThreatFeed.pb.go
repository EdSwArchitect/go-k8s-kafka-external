// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.13.0
// source: ThreatFeed.proto

package talon

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

type ThreatFeedMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddress *string  `protobuf:"bytes,1,opt,name=ipAddress" json:"ipAddress,omitempty"`   // the IP address of the device that is associated with a threat feed correlation
	Score     *float64 `protobuf:"fixed64,2,opt,name=score" json:"score,omitempty"`         // the risk score of the IP address that is associated with the threat feed correlation
	Latitude  *float64 `protobuf:"fixed64,3,opt,name=latitude" json:"latitude,omitempty"`   // the latitude of the IP address that is associated with the threat feed correlation
	Longitude *float64 `protobuf:"fixed64,4,opt,name=longitude" json:"longitude,omitempty"` // the longitude of the IP address that is associated with the threat feed correlation
	Country   *string  `protobuf:"bytes,5,opt,name=country" json:"country,omitempty"`       // the country of the IP address that is associated with the threat feed correlation
	Protocol  *string  `protobuf:"bytes,6,opt,name=protocol" json:"protocol,omitempty"`     // the protocol of the threat feed correlation that is associated with the IP address
	Category  *string  `protobuf:"bytes,7,opt,name=category" json:"category,omitempty"`     // the category of the threat feed correlation that is associated with the IP address
}

func (x *ThreatFeedMsg) Reset() {
	*x = ThreatFeedMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ThreatFeed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThreatFeedMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreatFeedMsg) ProtoMessage() {}

func (x *ThreatFeedMsg) ProtoReflect() protoreflect.Message {
	mi := &file_ThreatFeed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreatFeedMsg.ProtoReflect.Descriptor instead.
func (*ThreatFeedMsg) Descriptor() ([]byte, []int) {
	return file_ThreatFeed_proto_rawDescGZIP(), []int{0}
}

func (x *ThreatFeedMsg) GetIpAddress() string {
	if x != nil && x.IpAddress != nil {
		return *x.IpAddress
	}
	return ""
}

func (x *ThreatFeedMsg) GetScore() float64 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

func (x *ThreatFeedMsg) GetLatitude() float64 {
	if x != nil && x.Latitude != nil {
		return *x.Latitude
	}
	return 0
}

func (x *ThreatFeedMsg) GetLongitude() float64 {
	if x != nil && x.Longitude != nil {
		return *x.Longitude
	}
	return 0
}

func (x *ThreatFeedMsg) GetCountry() string {
	if x != nil && x.Country != nil {
		return *x.Country
	}
	return ""
}

func (x *ThreatFeedMsg) GetProtocol() string {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return ""
}

func (x *ThreatFeedMsg) GetCategory() string {
	if x != nil && x.Category != nil {
		return *x.Category
	}
	return ""
}

var File_ThreatFeed_proto protoreflect.FileDescriptor

var file_ThreatFeed_proto_rawDesc = []byte{
	0x0a, 0x10, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x46, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x0d, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x46, 0x65, 0x65,
	0x64, 0x4d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x42, 0x28, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x73, 0x2e,
	0x63, 0x79, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x42, 0x0a, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x46, 0x65, 0x65, 0x64,
}

var (
	file_ThreatFeed_proto_rawDescOnce sync.Once
	file_ThreatFeed_proto_rawDescData = file_ThreatFeed_proto_rawDesc
)

func file_ThreatFeed_proto_rawDescGZIP() []byte {
	file_ThreatFeed_proto_rawDescOnce.Do(func() {
		file_ThreatFeed_proto_rawDescData = protoimpl.X.CompressGZIP(file_ThreatFeed_proto_rawDescData)
	})
	return file_ThreatFeed_proto_rawDescData
}

var file_ThreatFeed_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ThreatFeed_proto_goTypes = []interface{}{
	(*ThreatFeedMsg)(nil), // 0: ThreatFeedMsg
}
var file_ThreatFeed_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ThreatFeed_proto_init() }
func file_ThreatFeed_proto_init() {
	if File_ThreatFeed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ThreatFeed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThreatFeedMsg); i {
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
			RawDescriptor: file_ThreatFeed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ThreatFeed_proto_goTypes,
		DependencyIndexes: file_ThreatFeed_proto_depIdxs,
		MessageInfos:      file_ThreatFeed_proto_msgTypes,
	}.Build()
	File_ThreatFeed_proto = out.File
	file_ThreatFeed_proto_rawDesc = nil
	file_ThreatFeed_proto_goTypes = nil
	file_ThreatFeed_proto_depIdxs = nil
}
