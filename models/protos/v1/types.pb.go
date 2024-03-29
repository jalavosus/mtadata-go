// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: types.proto

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

type DirectionLabels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	North string `protobuf:"bytes,1,opt,name=north,proto3" json:"north,omitempty"`
	South string `protobuf:"bytes,2,opt,name=south,proto3" json:"south,omitempty"`
}

func (x *DirectionLabels) Reset() {
	*x = DirectionLabels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectionLabels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectionLabels) ProtoMessage() {}

func (x *DirectionLabels) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectionLabels.ProtoReflect.Descriptor instead.
func (*DirectionLabels) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *DirectionLabels) GetNorth() string {
	if x != nil {
		return x.North
	}
	return ""
}

func (x *DirectionLabels) GetSouth() string {
	if x != nil {
		return x.South
	}
	return ""
}

type GtfsLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *GtfsLocation) Reset() {
	*x = GtfsLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GtfsLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GtfsLocation) ProtoMessage() {}

func (x *GtfsLocation) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GtfsLocation.ProtoReflect.Descriptor instead.
func (*GtfsLocation) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

func (x *GtfsLocation) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *GtfsLocation) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type StationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StationId  int64  `protobuf:"varint,1,opt,name=station_id,json=stationId,proto3" json:"station_id,omitempty"`
	GtfsStopId string `protobuf:"bytes,2,opt,name=gtfs_stop_id,json=gtfsStopId,proto3" json:"gtfs_stop_id,omitempty"`
}

func (x *StationInfo) Reset() {
	*x = StationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationInfo) ProtoMessage() {}

func (x *StationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationInfo.ProtoReflect.Descriptor instead.
func (*StationInfo) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{2}
}

func (x *StationInfo) GetStationId() int64 {
	if x != nil {
		return x.StationId
	}
	return 0
}

func (x *StationInfo) GetGtfsStopId() string {
	if x != nil {
		return x.GtfsStopId
	}
	return ""
}

type Station struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DirectionLabels *DirectionLabels `protobuf:"bytes,1,opt,name=direction_labels,json=directionLabels,proto3" json:"direction_labels,omitempty"`
	GtfsStopId      string           `protobuf:"bytes,2,opt,name=gtfs_stop_id,json=gtfsStopId,proto3" json:"gtfs_stop_id,omitempty"`
	StopName        string           `protobuf:"bytes,3,opt,name=stop_name,json=stopName,proto3" json:"stop_name,omitempty"`
	Line            string           `protobuf:"bytes,4,opt,name=line,proto3" json:"line,omitempty"`
	Division        Division         `protobuf:"varint,5,opt,name=division,proto3,enum=mtadata.v1.Division" json:"division,omitempty"`
	Borough         Borough          `protobuf:"varint,6,opt,name=borough,proto3,enum=mtadata.v1.Borough" json:"borough,omitempty"`
	Structure       Structure        `protobuf:"varint,7,opt,name=structure,proto3,enum=mtadata.v1.Structure" json:"structure,omitempty"`
	StationId       int64            `protobuf:"varint,8,opt,name=station_id,json=stationId,proto3" json:"station_id,omitempty"`
	ComplexId       int64            `protobuf:"varint,9,opt,name=complex_id,json=complexId,proto3" json:"complex_id,omitempty"`
	DaytimeRoutes   []Route          `protobuf:"varint,10,rep,packed,name=daytime_routes,json=daytimeRoutes,proto3,enum=mtadata.v1.Route" json:"daytime_routes,omitempty"`
	GtfsLocation    *GtfsLocation    `protobuf:"bytes,11,opt,name=gtfs_location,json=gtfsLocation,proto3" json:"gtfs_location,omitempty"`
}

func (x *Station) Reset() {
	*x = Station{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Station) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Station) ProtoMessage() {}

func (x *Station) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Station.ProtoReflect.Descriptor instead.
func (*Station) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{3}
}

func (x *Station) GetDirectionLabels() *DirectionLabels {
	if x != nil {
		return x.DirectionLabels
	}
	return nil
}

func (x *Station) GetGtfsStopId() string {
	if x != nil {
		return x.GtfsStopId
	}
	return ""
}

func (x *Station) GetStopName() string {
	if x != nil {
		return x.StopName
	}
	return ""
}

func (x *Station) GetLine() string {
	if x != nil {
		return x.Line
	}
	return ""
}

func (x *Station) GetDivision() Division {
	if x != nil {
		return x.Division
	}
	return Division_UNKNOWN_DIVISION
}

func (x *Station) GetBorough() Borough {
	if x != nil {
		return x.Borough
	}
	return Borough_UNKNOWN_BOROUGH
}

func (x *Station) GetStructure() Structure {
	if x != nil {
		return x.Structure
	}
	return Structure_UNKNOWN_STRUCTURE
}

func (x *Station) GetStationId() int64 {
	if x != nil {
		return x.StationId
	}
	return 0
}

func (x *Station) GetComplexId() int64 {
	if x != nil {
		return x.ComplexId
	}
	return 0
}

func (x *Station) GetDaytimeRoutes() []Route {
	if x != nil {
		return x.DaytimeRoutes
	}
	return nil
}

func (x *Station) GetGtfsLocation() *GtfsLocation {
	if x != nil {
		return x.GtfsLocation
	}
	return nil
}

type StationComplex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComplexId     int64          `protobuf:"varint,1,opt,name=complex_id,json=complexId,proto3" json:"complex_id,omitempty"`
	Borough       Borough        `protobuf:"varint,2,opt,name=borough,proto3,enum=mtadata.v1.Borough" json:"borough,omitempty"`
	Divisions     []Division     `protobuf:"varint,3,rep,packed,name=divisions,proto3,enum=mtadata.v1.Division" json:"divisions,omitempty"`
	DaytimeRoutes []Route        `protobuf:"varint,4,rep,packed,name=daytime_routes,json=daytimeRoutes,proto3,enum=mtadata.v1.Route" json:"daytime_routes,omitempty"`
	StationInfos  []*StationInfo `protobuf:"bytes,5,rep,name=station_infos,json=stationInfos,proto3" json:"station_infos,omitempty"`
	Stations      *Stations      `protobuf:"bytes,6,opt,name=stations,proto3,oneof" json:"stations,omitempty"`
}

func (x *StationComplex) Reset() {
	*x = StationComplex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationComplex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationComplex) ProtoMessage() {}

func (x *StationComplex) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationComplex.ProtoReflect.Descriptor instead.
func (*StationComplex) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{4}
}

func (x *StationComplex) GetComplexId() int64 {
	if x != nil {
		return x.ComplexId
	}
	return 0
}

func (x *StationComplex) GetBorough() Borough {
	if x != nil {
		return x.Borough
	}
	return Borough_UNKNOWN_BOROUGH
}

func (x *StationComplex) GetDivisions() []Division {
	if x != nil {
		return x.Divisions
	}
	return nil
}

func (x *StationComplex) GetDaytimeRoutes() []Route {
	if x != nil {
		return x.DaytimeRoutes
	}
	return nil
}

func (x *StationComplex) GetStationInfos() []*StationInfo {
	if x != nil {
		return x.StationInfos
	}
	return nil
}

func (x *StationComplex) GetStations() *Stations {
	if x != nil {
		return x.Stations
	}
	return nil
}

type Stations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stations []*Station `protobuf:"bytes,1,rep,name=stations,proto3" json:"stations,omitempty"`
}

func (x *Stations) Reset() {
	*x = Stations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stations) ProtoMessage() {}

func (x *Stations) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stations.ProtoReflect.Descriptor instead.
func (*Stations) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{5}
}

func (x *Stations) GetStations() []*Station {
	if x != nil {
		return x.Stations
	}
	return nil
}

type TrainTimeUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GtfsStopId    string   `protobuf:"bytes,1,opt,name=gtfs_stop_id,json=gtfsStopId,proto3" json:"gtfs_stop_id,omitempty"`
	Station       *Station `protobuf:"bytes,2,opt,name=station,proto3" json:"station,omitempty"`
	Route         Route    `protobuf:"varint,3,opt,name=route,proto3,enum=mtadata.v1.Route" json:"route,omitempty"`
	Direction     string   `protobuf:"bytes,4,opt,name=direction,proto3" json:"direction,omitempty"`
	DepartureTime int64    `protobuf:"varint,5,opt,name=departure_time,json=departureTime,proto3" json:"departure_time,omitempty"`
}

func (x *TrainTimeUpdate) Reset() {
	*x = TrainTimeUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrainTimeUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainTimeUpdate) ProtoMessage() {}

func (x *TrainTimeUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainTimeUpdate.ProtoReflect.Descriptor instead.
func (*TrainTimeUpdate) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{6}
}

func (x *TrainTimeUpdate) GetGtfsStopId() string {
	if x != nil {
		return x.GtfsStopId
	}
	return ""
}

func (x *TrainTimeUpdate) GetStation() *Station {
	if x != nil {
		return x.Station
	}
	return nil
}

func (x *TrainTimeUpdate) GetRoute() Route {
	if x != nil {
		return x.Route
	}
	return Route_UNKNOWN_ROUTE
}

func (x *TrainTimeUpdate) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *TrainTimeUpdate) GetDepartureTime() int64 {
	if x != nil {
		return x.DepartureTime
	}
	return 0
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0f, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x72,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x72, 0x74, 0x68, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x6f, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x6f, 0x75, 0x74, 0x68, 0x22, 0x48, 0x0a, 0x0c, 0x47, 0x74, 0x66, 0x73, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22,
	0x4e, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0c, 0x67, 0x74, 0x66, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x74, 0x66, 0x73, 0x53, 0x74, 0x6f, 0x70, 0x49, 0x64, 0x22,
	0xf1, 0x03, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x10, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x52, 0x0f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x74, 0x66, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x74, 0x66, 0x73, 0x53,
	0x74, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x07, 0x62, 0x6f, 0x72, 0x6f,
	0x75, 0x67, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x52, 0x07,
	0x62, 0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x12, 0x33, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6d, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x09, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0e, 0x64, 0x61,
	0x79, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x0d, 0x64, 0x61, 0x79, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0d, 0x67, 0x74, 0x66, 0x73, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x74, 0x66, 0x73, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x67, 0x74, 0x66, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xce, 0x02, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x78, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x07, 0x62, 0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x52, 0x07, 0x62, 0x6f, 0x72,
	0x6f, 0x75, 0x67, 0x68, 0x12, 0x32, 0x0a, 0x09, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64,
	0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x0e, 0x64, 0x61, 0x79, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x52, 0x0d, 0x64, 0x61, 0x79, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x35, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3b, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x2f, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0xd0, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x74, 0x66, 0x73, 0x5f, 0x73, 0x74,
	0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x74, 0x66,
	0x73, 0x53, 0x74, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a,
	0x0e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x42, 0x90, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6a, 0x61, 0x6c, 0x61, 0x76, 0x6f, 0x73, 0x75, 0x73, 0x2f, 0x6d, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2d, 0x67, 0x6f, 0x3b, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x4d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x4d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x16, 0x4d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x4d, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_types_proto_goTypes = []interface{}{
	(*DirectionLabels)(nil), // 0: mtadata.v1.DirectionLabels
	(*GtfsLocation)(nil),    // 1: mtadata.v1.GtfsLocation
	(*StationInfo)(nil),     // 2: mtadata.v1.StationInfo
	(*Station)(nil),         // 3: mtadata.v1.Station
	(*StationComplex)(nil),  // 4: mtadata.v1.StationComplex
	(*Stations)(nil),        // 5: mtadata.v1.Stations
	(*TrainTimeUpdate)(nil), // 6: mtadata.v1.TrainTimeUpdate
	(Division)(0),           // 7: mtadata.v1.Division
	(Borough)(0),            // 8: mtadata.v1.Borough
	(Structure)(0),          // 9: mtadata.v1.Structure
	(Route)(0),              // 10: mtadata.v1.Route
}
var file_types_proto_depIdxs = []int32{
	0,  // 0: mtadata.v1.Station.direction_labels:type_name -> mtadata.v1.DirectionLabels
	7,  // 1: mtadata.v1.Station.division:type_name -> mtadata.v1.Division
	8,  // 2: mtadata.v1.Station.borough:type_name -> mtadata.v1.Borough
	9,  // 3: mtadata.v1.Station.structure:type_name -> mtadata.v1.Structure
	10, // 4: mtadata.v1.Station.daytime_routes:type_name -> mtadata.v1.Route
	1,  // 5: mtadata.v1.Station.gtfs_location:type_name -> mtadata.v1.GtfsLocation
	8,  // 6: mtadata.v1.StationComplex.borough:type_name -> mtadata.v1.Borough
	7,  // 7: mtadata.v1.StationComplex.divisions:type_name -> mtadata.v1.Division
	10, // 8: mtadata.v1.StationComplex.daytime_routes:type_name -> mtadata.v1.Route
	2,  // 9: mtadata.v1.StationComplex.station_infos:type_name -> mtadata.v1.StationInfo
	5,  // 10: mtadata.v1.StationComplex.stations:type_name -> mtadata.v1.Stations
	3,  // 11: mtadata.v1.Stations.stations:type_name -> mtadata.v1.Station
	3,  // 12: mtadata.v1.TrainTimeUpdate.station:type_name -> mtadata.v1.Station
	10, // 13: mtadata.v1.TrainTimeUpdate.route:type_name -> mtadata.v1.Route
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	file_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectionLabels); i {
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
		file_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GtfsLocation); i {
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
		file_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationInfo); i {
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
		file_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Station); i {
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
		file_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationComplex); i {
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
		file_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stations); i {
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
		file_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrainTimeUpdate); i {
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
	file_types_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}
