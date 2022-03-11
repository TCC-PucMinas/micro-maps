// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: communicate/map.proto

package communicate

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GelocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street   string `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	District string `protobuf:"bytes,2,opt,name=district,proto3" json:"district,omitempty"`
	City     string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country  string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	State    string `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	Number   string `protobuf:"bytes,6,opt,name=number,proto3" json:"number,omitempty"`
	ZipCode  string `protobuf:"bytes,7,opt,name=ZipCode,proto3" json:"ZipCode,omitempty"`
}

func (x *GelocationRequest) Reset() {
	*x = GelocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communicate_map_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GelocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GelocationRequest) ProtoMessage() {}

func (x *GelocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_communicate_map_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GelocationRequest.ProtoReflect.Descriptor instead.
func (*GelocationRequest) Descriptor() ([]byte, []int) {
	return file_communicate_map_proto_rawDescGZIP(), []int{0}
}

func (x *GelocationRequest) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *GelocationRequest) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *GelocationRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *GelocationRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GelocationRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *GelocationRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *GelocationRequest) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

type GelocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat string `protobuf:"bytes,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng string `protobuf:"bytes,2,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *GelocationResponse) Reset() {
	*x = GelocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communicate_map_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GelocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GelocationResponse) ProtoMessage() {}

func (x *GelocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communicate_map_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GelocationResponse.ProtoReflect.Descriptor instead.
func (*GelocationResponse) Descriptor() ([]byte, []int) {
	return file_communicate_map_proto_rawDescGZIP(), []int{1}
}

func (x *GelocationResponse) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *GelocationResponse) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

type LatAndLng struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat string `protobuf:"bytes,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng string `protobuf:"bytes,2,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *LatAndLng) Reset() {
	*x = LatAndLng{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communicate_map_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatAndLng) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatAndLng) ProtoMessage() {}

func (x *LatAndLng) ProtoReflect() protoreflect.Message {
	mi := &file_communicate_map_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatAndLng.ProtoReflect.Descriptor instead.
func (*LatAndLng) Descriptor() ([]byte, []int) {
	return file_communicate_map_proto_rawDescGZIP(), []int{2}
}

func (x *LatAndLng) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *LatAndLng) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

type DirectionLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin  *LatAndLng `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Destiny *LatAndLng `protobuf:"bytes,2,opt,name=destiny,proto3" json:"destiny,omitempty"`
}

func (x *DirectionLocationRequest) Reset() {
	*x = DirectionLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communicate_map_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectionLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectionLocationRequest) ProtoMessage() {}

func (x *DirectionLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_communicate_map_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectionLocationRequest.ProtoReflect.Descriptor instead.
func (*DirectionLocationRequest) Descriptor() ([]byte, []int) {
	return file_communicate_map_proto_rawDescGZIP(), []int{3}
}

func (x *DirectionLocationRequest) GetOrigin() *LatAndLng {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *DirectionLocationRequest) GetDestiny() *LatAndLng {
	if x != nil {
		return x.Destiny
	}
	return nil
}

type DirectionLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin        *LatAndLng `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Destiny       *LatAndLng `protobuf:"bytes,2,opt,name=destiny,proto3" json:"destiny,omitempty"`
	HumanReadable string     `protobuf:"bytes,3,opt,name=HumanReadable,proto3" json:"HumanReadable,omitempty"`
	Meters        int64      `protobuf:"varint,4,opt,name=Meters,proto3" json:"Meters,omitempty"`
}

func (x *DirectionLocationResponse) Reset() {
	*x = DirectionLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communicate_map_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectionLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectionLocationResponse) ProtoMessage() {}

func (x *DirectionLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_communicate_map_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectionLocationResponse.ProtoReflect.Descriptor instead.
func (*DirectionLocationResponse) Descriptor() ([]byte, []int) {
	return file_communicate_map_proto_rawDescGZIP(), []int{4}
}

func (x *DirectionLocationResponse) GetOrigin() *LatAndLng {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *DirectionLocationResponse) GetDestiny() *LatAndLng {
	if x != nil {
		return x.Destiny
	}
	return nil
}

func (x *DirectionLocationResponse) GetHumanReadable() string {
	if x != nil {
		return x.HumanReadable
	}
	return ""
}

func (x *DirectionLocationResponse) GetMeters() int64 {
	if x != nil {
		return x.Meters
	}
	return 0
}

var File_communicate_map_proto protoreflect.FileDescriptor

var file_communicate_map_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2f, 0x6d, 0x61,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x5a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x5a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x38, 0x0a, 0x12, 0x47, 0x65, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6e,
	0x67, 0x22, 0x2f, 0x0a, 0x09, 0x4c, 0x61, 0x74, 0x41, 0x6e, 0x64, 0x4c, 0x6e, 0x67, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c,
	0x6e, 0x67, 0x22, 0x64, 0x0a, 0x18, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x4c, 0x61, 0x74, 0x41, 0x6e, 0x64, 0x4c, 0x6e, 0x67, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x12, 0x24, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4c, 0x61, 0x74, 0x41, 0x6e, 0x64, 0x4c, 0x6e, 0x67, 0x52,
	0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x79, 0x22, 0xa3, 0x01, 0x0a, 0x19, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4c, 0x61, 0x74, 0x41, 0x6e, 0x64, 0x4c,
	0x6e, 0x67, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x07, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4c, 0x61,
	0x74, 0x41, 0x6e, 0x64, 0x4c, 0x6e, 0x67, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x79,
	0x12, 0x24, 0x0a, 0x0d, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65,
	0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x32, 0x9f,
	0x01, 0x0a, 0x15, 0x47, 0x65, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x11, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_communicate_map_proto_rawDescOnce sync.Once
	file_communicate_map_proto_rawDescData = file_communicate_map_proto_rawDesc
)

func file_communicate_map_proto_rawDescGZIP() []byte {
	file_communicate_map_proto_rawDescOnce.Do(func() {
		file_communicate_map_proto_rawDescData = protoimpl.X.CompressGZIP(file_communicate_map_proto_rawDescData)
	})
	return file_communicate_map_proto_rawDescData
}

var file_communicate_map_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_communicate_map_proto_goTypes = []interface{}{
	(*GelocationRequest)(nil),         // 0: GelocationRequest
	(*GelocationResponse)(nil),        // 1: GelocationResponse
	(*LatAndLng)(nil),                 // 2: LatAndLng
	(*DirectionLocationRequest)(nil),  // 3: DirectionLocationRequest
	(*DirectionLocationResponse)(nil), // 4: DirectionLocationResponse
}
var file_communicate_map_proto_depIdxs = []int32{
	2, // 0: DirectionLocationRequest.origin:type_name -> LatAndLng
	2, // 1: DirectionLocationRequest.destiny:type_name -> LatAndLng
	2, // 2: DirectionLocationResponse.origin:type_name -> LatAndLng
	2, // 3: DirectionLocationResponse.destiny:type_name -> LatAndLng
	0, // 4: GelocationCommunicate.GetLocation:input_type -> GelocationRequest
	3, // 5: GelocationCommunicate.DirectionLocation:input_type -> DirectionLocationRequest
	1, // 6: GelocationCommunicate.GetLocation:output_type -> GelocationResponse
	4, // 7: GelocationCommunicate.DirectionLocation:output_type -> DirectionLocationResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_communicate_map_proto_init() }
func file_communicate_map_proto_init() {
	if File_communicate_map_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_communicate_map_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GelocationRequest); i {
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
		file_communicate_map_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GelocationResponse); i {
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
		file_communicate_map_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatAndLng); i {
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
		file_communicate_map_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectionLocationRequest); i {
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
		file_communicate_map_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectionLocationResponse); i {
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
			RawDescriptor: file_communicate_map_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_communicate_map_proto_goTypes,
		DependencyIndexes: file_communicate_map_proto_depIdxs,
		MessageInfos:      file_communicate_map_proto_msgTypes,
	}.Build()
	File_communicate_map_proto = out.File
	file_communicate_map_proto_rawDesc = nil
	file_communicate_map_proto_goTypes = nil
	file_communicate_map_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GelocationCommunicateClient is the client API for GelocationCommunicate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GelocationCommunicateClient interface {
	GetLocation(ctx context.Context, in *GelocationRequest, opts ...grpc.CallOption) (*GelocationResponse, error)
	DirectionLocation(ctx context.Context, in *DirectionLocationRequest, opts ...grpc.CallOption) (*DirectionLocationResponse, error)
}

type gelocationCommunicateClient struct {
	cc grpc.ClientConnInterface
}

func NewGelocationCommunicateClient(cc grpc.ClientConnInterface) GelocationCommunicateClient {
	return &gelocationCommunicateClient{cc}
}

func (c *gelocationCommunicateClient) GetLocation(ctx context.Context, in *GelocationRequest, opts ...grpc.CallOption) (*GelocationResponse, error) {
	out := new(GelocationResponse)
	err := c.cc.Invoke(ctx, "/GelocationCommunicate/GetLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gelocationCommunicateClient) DirectionLocation(ctx context.Context, in *DirectionLocationRequest, opts ...grpc.CallOption) (*DirectionLocationResponse, error) {
	out := new(DirectionLocationResponse)
	err := c.cc.Invoke(ctx, "/GelocationCommunicate/DirectionLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GelocationCommunicateServer is the server API for GelocationCommunicate service.
type GelocationCommunicateServer interface {
	GetLocation(context.Context, *GelocationRequest) (*GelocationResponse, error)
	DirectionLocation(context.Context, *DirectionLocationRequest) (*DirectionLocationResponse, error)
}

// UnimplementedGelocationCommunicateServer can be embedded to have forward compatible implementations.
type UnimplementedGelocationCommunicateServer struct {
}

func (*UnimplementedGelocationCommunicateServer) GetLocation(context.Context, *GelocationRequest) (*GelocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocation not implemented")
}
func (*UnimplementedGelocationCommunicateServer) DirectionLocation(context.Context, *DirectionLocationRequest) (*DirectionLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DirectionLocation not implemented")
}

func RegisterGelocationCommunicateServer(s *grpc.Server, srv GelocationCommunicateServer) {
	s.RegisterService(&_GelocationCommunicate_serviceDesc, srv)
}

func _GelocationCommunicate_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GelocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GelocationCommunicateServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GelocationCommunicate/GetLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GelocationCommunicateServer).GetLocation(ctx, req.(*GelocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GelocationCommunicate_DirectionLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DirectionLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GelocationCommunicateServer).DirectionLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GelocationCommunicate/DirectionLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GelocationCommunicateServer).DirectionLocation(ctx, req.(*DirectionLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GelocationCommunicate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GelocationCommunicate",
	HandlerType: (*GelocationCommunicateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLocation",
			Handler:    _GelocationCommunicate_GetLocation_Handler,
		},
		{
			MethodName: "DirectionLocation",
			Handler:    _GelocationCommunicate_DirectionLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "communicate/map.proto",
}
