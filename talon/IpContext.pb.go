// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.13.0
// source: IpContext.proto

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

type IpContextMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddress       *string  `protobuf:"bytes,1,opt,name=ipAddress" json:"ipAddress,omitempty"`              // the IPv4 address of the event
	IpNum           *int64   `protobuf:"varint,2,opt,name=ipNum" json:"ipNum,omitempty"`                     // an integer that represents the IP address
	Timestamp       *int64   `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`             // the record timestamp
	Hostname        *string  `protobuf:"bytes,4,opt,name=hostname" json:"hostname,omitempty"`                // the host name
	MacAddress      *string  `protobuf:"bytes,5,opt,name=macAddress" json:"macAddress,omitempty"`            // the media access control address
	SiteCode        *string  `protobuf:"bytes,6,opt,name=siteCode" json:"siteCode,omitempty"`                // the site code that is associated with the network
	Country         *string  `protobuf:"bytes,7,opt,name=country" json:"country,omitempty"`                  // the name of the country that is associated with the network
	Region          *string  `protobuf:"bytes,8,opt,name=region" json:"region,omitempty"`                    // the region that is associated with the network
	City            *string  `protobuf:"bytes,9,opt,name=city" json:"city,omitempty"`                        // the city that is associated with the network
	Latitude        *float64 `protobuf:"fixed64,10,opt,name=latitude" json:"latitude,omitempty"`             // the latitude that is associated with the network
	Longitude       *float64 `protobuf:"fixed64,11,opt,name=longitude" json:"longitude,omitempty"`           // the longitude that is associated with the network
	TzOffset        *int32   `protobuf:"varint,12,opt,name=tzOffset" json:"tzOffset,omitempty"`              // the time zone offset of the network
	NetworkName     *string  `protobuf:"bytes,13,opt,name=networkName" json:"networkName,omitempty"`         // the name of the network
	NetworkType     *string  `protobuf:"bytes,14,opt,name=networkType" json:"networkType,omitempty"`         // the type of the network
	NetworkScope    *int32   `protobuf:"varint,15,opt,name=networkScope" json:"networkScope,omitempty"`      // the network scope
	PeerGroup       *string  `protobuf:"bytes,16,opt,name=peerGroup" json:"peerGroup,omitempty"`             // the assigned peer group that is associated with the authentication event
	DeviceType      *string  `protobuf:"bytes,17,opt,name=deviceType" json:"deviceType,omitempty"`           // the device type that is associated with the authentication event
	DevicePeerGroup *string  `protobuf:"bytes,18,opt,name=devicePeerGroup" json:"devicePeerGroup,omitempty"` // the assigned peer group that is associated with the device
	IsUserDevice    *int32   `protobuf:"varint,19,opt,name=isUserDevice" json:"isUserDevice,omitempty"`      // a flag that indicates whether the device is a user device
	HVE             *int32   `protobuf:"varint,20,opt,name=HVE" json:"HVE,omitempty"`                        // the high-value entity value that is associated with the device
	BizProc         *string  `protobuf:"bytes,21,opt,name=bizProc" json:"bizProc,omitempty"`                 // the business process that is associated with the device
	UserId          *string  `protobuf:"bytes,22,opt,name=userId" json:"userId,omitempty"`                   // the unique identifier of the user that is associated with the authentication event
	Email           *string  `protobuf:"bytes,23,opt,name=email" json:"email,omitempty"`                     // the email address of the user that is associated with the authentication event
	Lname           *string  `protobuf:"bytes,24,opt,name=lname" json:"lname,omitempty"`                     // the last name of the user that is associated with the authentication event
	Fname           *string  `protobuf:"bytes,25,opt,name=fname" json:"fname,omitempty"`                     // the first name of the user that is associated with the authentication event
	Division        *string  `protobuf:"bytes,26,opt,name=division" json:"division,omitempty"`               // the division of the user that is associated with the authentication event
	Department      *string  `protobuf:"bytes,27,opt,name=department" json:"department,omitempty"`           // the department of the user that is associated with the authentication event
	UserCity        *string  `protobuf:"bytes,28,opt,name=userCity" json:"userCity,omitempty"`               // the city of the user that is associated with the authentication event
	UserRegion      *string  `protobuf:"bytes,29,opt,name=userRegion" json:"userRegion,omitempty"`           // the region of the user that is associated with the authentication event
	UserCountry     *string  `protobuf:"bytes,30,opt,name=userCountry" json:"userCountry,omitempty"`         // the country of the user that is associated with the authentication event
	UserPeerGroup   *string  `protobuf:"bytes,31,opt,name=userPeerGroup" json:"userPeerGroup,omitempty"`     // the assigned peer group of the user that is associated with the authentication event
	UserHVE         *int32   `protobuf:"varint,32,opt,name=userHVE" json:"userHVE,omitempty"`                // the high-value entity value that is associated with the user
	UserBizProc     *string  `protobuf:"bytes,33,opt,name=userBizProc" json:"userBizProc,omitempty"`         // the business process that is associated with the user
}

func (x *IpContextMsg) Reset() {
	*x = IpContextMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IpContext_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpContextMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpContextMsg) ProtoMessage() {}

func (x *IpContextMsg) ProtoReflect() protoreflect.Message {
	mi := &file_IpContext_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpContextMsg.ProtoReflect.Descriptor instead.
func (*IpContextMsg) Descriptor() ([]byte, []int) {
	return file_IpContext_proto_rawDescGZIP(), []int{0}
}

func (x *IpContextMsg) GetIpAddress() string {
	if x != nil && x.IpAddress != nil {
		return *x.IpAddress
	}
	return ""
}

func (x *IpContextMsg) GetIpNum() int64 {
	if x != nil && x.IpNum != nil {
		return *x.IpNum
	}
	return 0
}

func (x *IpContextMsg) GetTimestamp() int64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *IpContextMsg) GetHostname() string {
	if x != nil && x.Hostname != nil {
		return *x.Hostname
	}
	return ""
}

func (x *IpContextMsg) GetMacAddress() string {
	if x != nil && x.MacAddress != nil {
		return *x.MacAddress
	}
	return ""
}

func (x *IpContextMsg) GetSiteCode() string {
	if x != nil && x.SiteCode != nil {
		return *x.SiteCode
	}
	return ""
}

func (x *IpContextMsg) GetCountry() string {
	if x != nil && x.Country != nil {
		return *x.Country
	}
	return ""
}

func (x *IpContextMsg) GetRegion() string {
	if x != nil && x.Region != nil {
		return *x.Region
	}
	return ""
}

func (x *IpContextMsg) GetCity() string {
	if x != nil && x.City != nil {
		return *x.City
	}
	return ""
}

func (x *IpContextMsg) GetLatitude() float64 {
	if x != nil && x.Latitude != nil {
		return *x.Latitude
	}
	return 0
}

func (x *IpContextMsg) GetLongitude() float64 {
	if x != nil && x.Longitude != nil {
		return *x.Longitude
	}
	return 0
}

func (x *IpContextMsg) GetTzOffset() int32 {
	if x != nil && x.TzOffset != nil {
		return *x.TzOffset
	}
	return 0
}

func (x *IpContextMsg) GetNetworkName() string {
	if x != nil && x.NetworkName != nil {
		return *x.NetworkName
	}
	return ""
}

func (x *IpContextMsg) GetNetworkType() string {
	if x != nil && x.NetworkType != nil {
		return *x.NetworkType
	}
	return ""
}

func (x *IpContextMsg) GetNetworkScope() int32 {
	if x != nil && x.NetworkScope != nil {
		return *x.NetworkScope
	}
	return 0
}

func (x *IpContextMsg) GetPeerGroup() string {
	if x != nil && x.PeerGroup != nil {
		return *x.PeerGroup
	}
	return ""
}

func (x *IpContextMsg) GetDeviceType() string {
	if x != nil && x.DeviceType != nil {
		return *x.DeviceType
	}
	return ""
}

func (x *IpContextMsg) GetDevicePeerGroup() string {
	if x != nil && x.DevicePeerGroup != nil {
		return *x.DevicePeerGroup
	}
	return ""
}

func (x *IpContextMsg) GetIsUserDevice() int32 {
	if x != nil && x.IsUserDevice != nil {
		return *x.IsUserDevice
	}
	return 0
}

func (x *IpContextMsg) GetHVE() int32 {
	if x != nil && x.HVE != nil {
		return *x.HVE
	}
	return 0
}

func (x *IpContextMsg) GetBizProc() string {
	if x != nil && x.BizProc != nil {
		return *x.BizProc
	}
	return ""
}

func (x *IpContextMsg) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *IpContextMsg) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *IpContextMsg) GetLname() string {
	if x != nil && x.Lname != nil {
		return *x.Lname
	}
	return ""
}

func (x *IpContextMsg) GetFname() string {
	if x != nil && x.Fname != nil {
		return *x.Fname
	}
	return ""
}

func (x *IpContextMsg) GetDivision() string {
	if x != nil && x.Division != nil {
		return *x.Division
	}
	return ""
}

func (x *IpContextMsg) GetDepartment() string {
	if x != nil && x.Department != nil {
		return *x.Department
	}
	return ""
}

func (x *IpContextMsg) GetUserCity() string {
	if x != nil && x.UserCity != nil {
		return *x.UserCity
	}
	return ""
}

func (x *IpContextMsg) GetUserRegion() string {
	if x != nil && x.UserRegion != nil {
		return *x.UserRegion
	}
	return ""
}

func (x *IpContextMsg) GetUserCountry() string {
	if x != nil && x.UserCountry != nil {
		return *x.UserCountry
	}
	return ""
}

func (x *IpContextMsg) GetUserPeerGroup() string {
	if x != nil && x.UserPeerGroup != nil {
		return *x.UserPeerGroup
	}
	return ""
}

func (x *IpContextMsg) GetUserHVE() int32 {
	if x != nil && x.UserHVE != nil {
		return *x.UserHVE
	}
	return 0
}

func (x *IpContextMsg) GetUserBizProc() string {
	if x != nil && x.UserBizProc != nil {
		return *x.UserBizProc
	}
	return ""
}

var File_IpContext_proto protoreflect.FileDescriptor

var file_IpContext_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x49, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xca, 0x07, 0x0a, 0x0c, 0x49, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4d,
	0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x70, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x69, 0x70, 0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x7a, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x74, 0x7a, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x65, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x65, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x65, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x69, 0x73, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x48, 0x56, 0x45, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x48, 0x56, 0x45, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x69, 0x7a, 0x50, 0x72, 0x6f, 0x63, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x69, 0x7a, 0x50, 0x72, 0x6f, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x1b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x43, 0x69, 0x74, 0x79, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x43, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73,
	0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x0d,
	0x75, 0x73, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x48, 0x56, 0x45, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x48, 0x56, 0x45, 0x12, 0x20, 0x0a, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x42, 0x69, 0x7a, 0x50, 0x72, 0x6f, 0x63, 0x18, 0x21, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x42, 0x69, 0x7a, 0x50, 0x72, 0x6f, 0x63, 0x42, 0x27,
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x73, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x09, 0x49, 0x70,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
}

var (
	file_IpContext_proto_rawDescOnce sync.Once
	file_IpContext_proto_rawDescData = file_IpContext_proto_rawDesc
)

func file_IpContext_proto_rawDescGZIP() []byte {
	file_IpContext_proto_rawDescOnce.Do(func() {
		file_IpContext_proto_rawDescData = protoimpl.X.CompressGZIP(file_IpContext_proto_rawDescData)
	})
	return file_IpContext_proto_rawDescData
}

var file_IpContext_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_IpContext_proto_goTypes = []interface{}{
	(*IpContextMsg)(nil), // 0: IpContextMsg
}
var file_IpContext_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_IpContext_proto_init() }
func file_IpContext_proto_init() {
	if File_IpContext_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_IpContext_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpContextMsg); i {
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
			RawDescriptor: file_IpContext_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_IpContext_proto_goTypes,
		DependencyIndexes: file_IpContext_proto_depIdxs,
		MessageInfos:      file_IpContext_proto_msgTypes,
	}.Build()
	File_IpContext_proto = out.File
	file_IpContext_proto_rawDesc = nil
	file_IpContext_proto_goTypes = nil
	file_IpContext_proto_depIdxs = nil
}
