// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: responses.proto

package mtadatav1

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

type UpcomingTrainsResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpcomingTrains []*TrainTimeUpdate `protobuf:"bytes,1,rep,name=upcoming_trains,json=upcomingTrains,proto3" json:"upcoming_trains,omitempty"`
	Error          *Error             `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *UpcomingTrainsResult) Reset() {
	*x = UpcomingTrainsResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpcomingTrainsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpcomingTrainsResult) ProtoMessage() {}

func (x *UpcomingTrainsResult) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpcomingTrainsResult.ProtoReflect.Descriptor instead.
func (*UpcomingTrainsResult) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{0}
}

func (x *UpcomingTrainsResult) GetUpcomingTrains() []*TrainTimeUpdate {
	if x != nil {
		return x.UpcomingTrains
	}
	return nil
}

func (x *UpcomingTrainsResult) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type StationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Station *Station `protobuf:"bytes,1,opt,name=station,proto3" json:"station,omitempty"`
	Error   *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *StationResult) Reset() {
	*x = StationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationResult) ProtoMessage() {}

func (x *StationResult) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationResult.ProtoReflect.Descriptor instead.
func (*StationResult) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{1}
}

func (x *StationResult) GetStation() *Station {
	if x != nil {
		return x.Station
	}
	return nil
}

func (x *StationResult) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type StationsResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stations []*Station `protobuf:"bytes,1,rep,name=stations,proto3" json:"stations,omitempty"`
	Error    *Error     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *StationsResult) Reset() {
	*x = StationsResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationsResult) ProtoMessage() {}

func (x *StationsResult) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationsResult.ProtoReflect.Descriptor instead.
func (*StationsResult) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{2}
}

func (x *StationsResult) GetStations() []*Station {
	if x != nil {
		return x.Stations
	}
	return nil
}

func (x *StationsResult) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type StationComplexResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StationComplex *StationComplex `protobuf:"bytes,1,opt,name=station_complex,json=stationComplex,proto3" json:"station_complex,omitempty"`
	Error          *Error          `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *StationComplexResult) Reset() {
	*x = StationComplexResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationComplexResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationComplexResult) ProtoMessage() {}

func (x *StationComplexResult) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationComplexResult.ProtoReflect.Descriptor instead.
func (*StationComplexResult) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{3}
}

func (x *StationComplexResult) GetStationComplex() *StationComplex {
	if x != nil {
		return x.StationComplex
	}
	return nil
}

func (x *StationComplexResult) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type StationComplexesResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StationComplexes []*StationComplex `protobuf:"bytes,1,rep,name=station_complexes,json=stationComplexes,proto3" json:"station_complexes,omitempty"`
	Error            *Error            `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *StationComplexesResult) Reset() {
	*x = StationComplexesResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationComplexesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationComplexesResult) ProtoMessage() {}

func (x *StationComplexesResult) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationComplexesResult.ProtoReflect.Descriptor instead.
func (*StationComplexesResult) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{4}
}

func (x *StationComplexesResult) GetStationComplexes() []*StationComplex {
	if x != nil {
		return x.StationComplexes
	}
	return nil
}

func (x *StationComplexesResult) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type AllRoutes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routes []Route `protobuf:"varint,1,rep,packed,name=routes,proto3,enum=mtadata.v1.Route" json:"routes,omitempty"`
}

func (x *AllRoutes) Reset() {
	*x = AllRoutes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllRoutes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllRoutes) ProtoMessage() {}

func (x *AllRoutes) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllRoutes.ProtoReflect.Descriptor instead.
func (*AllRoutes) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{5}
}

func (x *AllRoutes) GetRoutes() []Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

type AllBoroughs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Boroughs []Borough `protobuf:"varint,1,rep,packed,name=boroughs,proto3,enum=mtadata.v1.Borough" json:"boroughs,omitempty"`
}

func (x *AllBoroughs) Reset() {
	*x = AllBoroughs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllBoroughs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllBoroughs) ProtoMessage() {}

func (x *AllBoroughs) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllBoroughs.ProtoReflect.Descriptor instead.
func (*AllBoroughs) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{6}
}

func (x *AllBoroughs) GetBoroughs() []Borough {
	if x != nil {
		return x.Boroughs
	}
	return nil
}

type AllDivisions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Divisions []Division `protobuf:"varint,1,rep,packed,name=divisions,proto3,enum=mtadata.v1.Division" json:"divisions,omitempty"`
}

func (x *AllDivisions) Reset() {
	*x = AllDivisions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllDivisions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllDivisions) ProtoMessage() {}

func (x *AllDivisions) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllDivisions.ProtoReflect.Descriptor instead.
func (*AllDivisions) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{7}
}

func (x *AllDivisions) GetDivisions() []Division {
	if x != nil {
		return x.Divisions
	}
	return nil
}

type AllStructures struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Structures []Structure `protobuf:"varint,1,rep,packed,name=structures,proto3,enum=mtadata.v1.Structure" json:"structures,omitempty"`
}

func (x *AllStructures) Reset() {
	*x = AllStructures{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllStructures) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllStructures) ProtoMessage() {}

func (x *AllStructures) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllStructures.ProtoReflect.Descriptor instead.
func (*AllStructures) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{8}
}

func (x *AllStructures) GetStructures() []Structure {
	if x != nil {
		return x.Structures
	}
	return nil
}

var File_responses_proto protoreflect.FileDescriptor

var file_responses_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x0b, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x63, 0x6f, 0x6d, 0x69,
	0x6e, 0x67, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x44,
	0x0a, 0x0f, 0x75, 0x70, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x0e, 0x75, 0x70, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x54, 0x72,
	0x61, 0x69, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88,
	0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x67, 0x0a, 0x0d,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2d, 0x0a,
	0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x6a, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x84, 0x01, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x43, 0x0a, 0x0f, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x52,
	0x0e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x12,
	0x27, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x8a, 0x01, 0x0a, 0x16, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x47, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x52, 0x10, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x36, 0x0a, 0x09, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x29, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0x3e, 0x0a,
	0x0b, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x73, 0x12, 0x2f, 0x0a, 0x08,
	0x62, 0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x72, 0x6f,
	0x75, 0x67, 0x68, 0x52, 0x08, 0x62, 0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x73, 0x22, 0x42, 0x0a,
	0x0c, 0x41, 0x6c, 0x6c, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x32, 0x0a,
	0x09, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x46, 0x0a, 0x0d, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x12, 0x35, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0a, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x42, 0x94, 0x01, 0x0a, 0x0e, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x6c, 0x61, 0x76,
	0x6f, 0x73, 0x75, 0x73, 0x2f, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x67, 0x6f, 0x3b,
	0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa,
	0x02, 0x0a, 0x4d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x4d,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x4d, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0b, 0x4d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_responses_proto_rawDescOnce sync.Once
	file_responses_proto_rawDescData = file_responses_proto_rawDesc
)

func file_responses_proto_rawDescGZIP() []byte {
	file_responses_proto_rawDescOnce.Do(func() {
		file_responses_proto_rawDescData = protoimpl.X.CompressGZIP(file_responses_proto_rawDescData)
	})
	return file_responses_proto_rawDescData
}

var file_responses_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_responses_proto_goTypes = []interface{}{
	(*UpcomingTrainsResult)(nil),   // 0: mtadata.v1.UpcomingTrainsResult
	(*StationResult)(nil),          // 1: mtadata.v1.StationResult
	(*StationsResult)(nil),         // 2: mtadata.v1.StationsResult
	(*StationComplexResult)(nil),   // 3: mtadata.v1.StationComplexResult
	(*StationComplexesResult)(nil), // 4: mtadata.v1.StationComplexesResult
	(*AllRoutes)(nil),              // 5: mtadata.v1.AllRoutes
	(*AllBoroughs)(nil),            // 6: mtadata.v1.AllBoroughs
	(*AllDivisions)(nil),           // 7: mtadata.v1.AllDivisions
	(*AllStructures)(nil),          // 8: mtadata.v1.AllStructures
	(*TrainTimeUpdate)(nil),        // 9: mtadata.v1.TrainTimeUpdate
	(*Error)(nil),                  // 10: mtadata.v1.Error
	(*Station)(nil),                // 11: mtadata.v1.Station
	(*StationComplex)(nil),         // 12: mtadata.v1.StationComplex
	(Route)(0),                     // 13: mtadata.v1.Route
	(Borough)(0),                   // 14: mtadata.v1.Borough
	(Division)(0),                  // 15: mtadata.v1.Division
	(Structure)(0),                 // 16: mtadata.v1.Structure
}
var file_responses_proto_depIdxs = []int32{
	9,  // 0: mtadata.v1.UpcomingTrainsResult.upcoming_trains:type_name -> mtadata.v1.TrainTimeUpdate
	10, // 1: mtadata.v1.UpcomingTrainsResult.error:type_name -> mtadata.v1.Error
	11, // 2: mtadata.v1.StationResult.station:type_name -> mtadata.v1.Station
	10, // 3: mtadata.v1.StationResult.error:type_name -> mtadata.v1.Error
	11, // 4: mtadata.v1.StationsResult.stations:type_name -> mtadata.v1.Station
	10, // 5: mtadata.v1.StationsResult.error:type_name -> mtadata.v1.Error
	12, // 6: mtadata.v1.StationComplexResult.station_complex:type_name -> mtadata.v1.StationComplex
	10, // 7: mtadata.v1.StationComplexResult.error:type_name -> mtadata.v1.Error
	12, // 8: mtadata.v1.StationComplexesResult.station_complexes:type_name -> mtadata.v1.StationComplex
	10, // 9: mtadata.v1.StationComplexesResult.error:type_name -> mtadata.v1.Error
	13, // 10: mtadata.v1.AllRoutes.routes:type_name -> mtadata.v1.Route
	14, // 11: mtadata.v1.AllBoroughs.boroughs:type_name -> mtadata.v1.Borough
	15, // 12: mtadata.v1.AllDivisions.divisions:type_name -> mtadata.v1.Division
	16, // 13: mtadata.v1.AllStructures.structures:type_name -> mtadata.v1.Structure
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_responses_proto_init() }
func file_responses_proto_init() {
	if File_responses_proto != nil {
		return
	}
	file_enums_proto_init()
	file_types_proto_init()
	file_errors_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_responses_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpcomingTrainsResult); i {
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
		file_responses_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationResult); i {
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
		file_responses_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationsResult); i {
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
		file_responses_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationComplexResult); i {
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
		file_responses_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationComplexesResult); i {
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
		file_responses_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllRoutes); i {
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
		file_responses_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllBoroughs); i {
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
		file_responses_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllDivisions); i {
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
		file_responses_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllStructures); i {
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
	file_responses_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_responses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_responses_proto_goTypes,
		DependencyIndexes: file_responses_proto_depIdxs,
		MessageInfos:      file_responses_proto_msgTypes,
	}.Build()
	File_responses_proto = out.File
	file_responses_proto_rawDesc = nil
	file_responses_proto_goTypes = nil
	file_responses_proto_depIdxs = nil
}
