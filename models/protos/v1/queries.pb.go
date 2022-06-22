// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: queries.proto

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

type QueryOrderBy int32

const (
	QueryOrderBy_DEFAULT      QueryOrderBy = 0
	QueryOrderBy_GTFS_STOP_ID QueryOrderBy = 1
	QueryOrderBy_STATION_ID   QueryOrderBy = 2
	QueryOrderBy_COMPLEX_ID   QueryOrderBy = 3
)

// Enum value maps for QueryOrderBy.
var (
	QueryOrderBy_name = map[int32]string{
		0: "DEFAULT",
		1: "GTFS_STOP_ID",
		2: "STATION_ID",
		3: "COMPLEX_ID",
	}
	QueryOrderBy_value = map[string]int32{
		"DEFAULT":      0,
		"GTFS_STOP_ID": 1,
		"STATION_ID":   2,
		"COMPLEX_ID":   3,
	}
)

func (x QueryOrderBy) Enum() *QueryOrderBy {
	p := new(QueryOrderBy)
	*p = x
	return p
}

func (x QueryOrderBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryOrderBy) Descriptor() protoreflect.EnumDescriptor {
	return file_queries_proto_enumTypes[0].Descriptor()
}

func (QueryOrderBy) Type() protoreflect.EnumType {
	return &file_queries_proto_enumTypes[0]
}

func (x QueryOrderBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryOrderBy.Descriptor instead.
func (QueryOrderBy) EnumDescriptor() ([]byte, []int) {
	return file_queries_proto_rawDescGZIP(), []int{0}
}

type UpcomingTrainsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StationId  string `protobuf:"bytes,1,opt,name=station_id,json=stationId,proto3" json:"station_id,omitempty"`
	GtfsStopId string `protobuf:"bytes,2,opt,name=gtfs_stop_id,json=gtfsStopId,proto3" json:"gtfs_stop_id,omitempty"`
	Route      Route  `protobuf:"varint,3,opt,name=route,proto3,enum=mtadata.v1.Route" json:"route,omitempty"`
}

func (x *UpcomingTrainsRequest) Reset() {
	*x = UpcomingTrainsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpcomingTrainsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpcomingTrainsRequest) ProtoMessage() {}

func (x *UpcomingTrainsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpcomingTrainsRequest.ProtoReflect.Descriptor instead.
func (*UpcomingTrainsRequest) Descriptor() ([]byte, []int) {
	return file_queries_proto_rawDescGZIP(), []int{0}
}

func (x *UpcomingTrainsRequest) GetStationId() string {
	if x != nil {
		return x.StationId
	}
	return ""
}

func (x *UpcomingTrainsRequest) GetGtfsStopId() string {
	if x != nil {
		return x.GtfsStopId
	}
	return ""
}

func (x *UpcomingTrainsRequest) GetRoute() Route {
	if x != nil {
		return x.Route
	}
	return Route_UNKNOWN_ROUTE
}

type StationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StationId  string `protobuf:"bytes,1,opt,name=station_id,json=stationId,proto3" json:"station_id,omitempty"`
	GtfsStopId string `protobuf:"bytes,2,opt,name=gtfs_stop_id,json=gtfsStopId,proto3" json:"gtfs_stop_id,omitempty"`
}

func (x *StationRequest) Reset() {
	*x = StationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queries_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationRequest) ProtoMessage() {}

func (x *StationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queries_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationRequest.ProtoReflect.Descriptor instead.
func (*StationRequest) Descriptor() ([]byte, []int) {
	return file_queries_proto_rawDescGZIP(), []int{1}
}

func (x *StationRequest) GetStationId() string {
	if x != nil {
		return x.StationId
	}
	return ""
}

func (x *StationRequest) GetGtfsStopId() string {
	if x != nil {
		return x.GtfsStopId
	}
	return ""
}

type BaseQueryParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Route    *Route        `protobuf:"varint,1,opt,name=route,proto3,enum=mtadata.v1.Route,oneof" json:"route,omitempty"`
	Borough  *Borough      `protobuf:"varint,2,opt,name=borough,proto3,enum=mtadata.v1.Borough,oneof" json:"borough,omitempty"`
	Division *Division     `protobuf:"varint,3,opt,name=division,proto3,enum=mtadata.v1.Division,oneof" json:"division,omitempty"`
	OrderBy  *QueryOrderBy `protobuf:"varint,4,opt,name=order_by,json=orderBy,proto3,enum=mtadata.v1.QueryOrderBy,oneof" json:"order_by,omitempty"`
	Limit    *int32        `protobuf:"varint,5,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *BaseQueryParams) Reset() {
	*x = BaseQueryParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queries_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseQueryParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseQueryParams) ProtoMessage() {}

func (x *BaseQueryParams) ProtoReflect() protoreflect.Message {
	mi := &file_queries_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseQueryParams.ProtoReflect.Descriptor instead.
func (*BaseQueryParams) Descriptor() ([]byte, []int) {
	return file_queries_proto_rawDescGZIP(), []int{2}
}

func (x *BaseQueryParams) GetRoute() Route {
	if x != nil && x.Route != nil {
		return *x.Route
	}
	return Route_UNKNOWN_ROUTE
}

func (x *BaseQueryParams) GetBorough() Borough {
	if x != nil && x.Borough != nil {
		return *x.Borough
	}
	return Borough_UNKNOWN_BOROUGH
}

func (x *BaseQueryParams) GetDivision() Division {
	if x != nil && x.Division != nil {
		return *x.Division
	}
	return Division_UNKNOWN_DIVISION
}

func (x *BaseQueryParams) GetOrderBy() QueryOrderBy {
	if x != nil && x.OrderBy != nil {
		return *x.OrderBy
	}
	return QueryOrderBy_DEFAULT
}

func (x *BaseQueryParams) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type StationsQueryParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseParams *BaseQueryParams `protobuf:"bytes,1,opt,name=base_params,json=baseParams,proto3,oneof" json:"base_params,omitempty"`
	StationId  *string          `protobuf:"bytes,2,opt,name=station_id,json=stationId,proto3,oneof" json:"station_id,omitempty"`
	ComplexId  *string          `protobuf:"bytes,3,opt,name=complex_id,json=complexId,proto3,oneof" json:"complex_id,omitempty"`
	GtfsStopId *string          `protobuf:"bytes,4,opt,name=gtfs_stop_id,json=gtfsStopId,proto3,oneof" json:"gtfs_stop_id,omitempty"`
}

func (x *StationsQueryParams) Reset() {
	*x = StationsQueryParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queries_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationsQueryParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationsQueryParams) ProtoMessage() {}

func (x *StationsQueryParams) ProtoReflect() protoreflect.Message {
	mi := &file_queries_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationsQueryParams.ProtoReflect.Descriptor instead.
func (*StationsQueryParams) Descriptor() ([]byte, []int) {
	return file_queries_proto_rawDescGZIP(), []int{3}
}

func (x *StationsQueryParams) GetBaseParams() *BaseQueryParams {
	if x != nil {
		return x.BaseParams
	}
	return nil
}

func (x *StationsQueryParams) GetStationId() string {
	if x != nil && x.StationId != nil {
		return *x.StationId
	}
	return ""
}

func (x *StationsQueryParams) GetComplexId() string {
	if x != nil && x.ComplexId != nil {
		return *x.ComplexId
	}
	return ""
}

func (x *StationsQueryParams) GetGtfsStopId() string {
	if x != nil && x.GtfsStopId != nil {
		return *x.GtfsStopId
	}
	return ""
}

type StationsQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryParams *StationsQueryParams `protobuf:"bytes,1,opt,name=query_params,json=queryParams,proto3" json:"query_params,omitempty"`
}

func (x *StationsQuery) Reset() {
	*x = StationsQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queries_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationsQuery) ProtoMessage() {}

func (x *StationsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_queries_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationsQuery.ProtoReflect.Descriptor instead.
func (*StationsQuery) Descriptor() ([]byte, []int) {
	return file_queries_proto_rawDescGZIP(), []int{4}
}

func (x *StationsQuery) GetQueryParams() *StationsQueryParams {
	if x != nil {
		return x.QueryParams
	}
	return nil
}

type StationComplexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComplexId string `protobuf:"bytes,1,opt,name=complex_id,json=complexId,proto3" json:"complex_id,omitempty"`
	Verbose   bool   `protobuf:"varint,2,opt,name=verbose,proto3" json:"verbose,omitempty"`
}

func (x *StationComplexRequest) Reset() {
	*x = StationComplexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queries_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationComplexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationComplexRequest) ProtoMessage() {}

func (x *StationComplexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queries_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationComplexRequest.ProtoReflect.Descriptor instead.
func (*StationComplexRequest) Descriptor() ([]byte, []int) {
	return file_queries_proto_rawDescGZIP(), []int{5}
}

func (x *StationComplexRequest) GetComplexId() string {
	if x != nil {
		return x.ComplexId
	}
	return ""
}

func (x *StationComplexRequest) GetVerbose() bool {
	if x != nil {
		return x.Verbose
	}
	return false
}

type StationComplexesQueryParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseParams *BaseQueryParams `protobuf:"bytes,1,opt,name=base_params,json=baseParams,proto3,oneof" json:"base_params,omitempty"`
	ComplexId  *string          `protobuf:"bytes,2,opt,name=complex_id,json=complexId,proto3,oneof" json:"complex_id,omitempty"`
}

func (x *StationComplexesQueryParams) Reset() {
	*x = StationComplexesQueryParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queries_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationComplexesQueryParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationComplexesQueryParams) ProtoMessage() {}

func (x *StationComplexesQueryParams) ProtoReflect() protoreflect.Message {
	mi := &file_queries_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationComplexesQueryParams.ProtoReflect.Descriptor instead.
func (*StationComplexesQueryParams) Descriptor() ([]byte, []int) {
	return file_queries_proto_rawDescGZIP(), []int{6}
}

func (x *StationComplexesQueryParams) GetBaseParams() *BaseQueryParams {
	if x != nil {
		return x.BaseParams
	}
	return nil
}

func (x *StationComplexesQueryParams) GetComplexId() string {
	if x != nil && x.ComplexId != nil {
		return *x.ComplexId
	}
	return ""
}

type StationComplexesQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryParams *StationComplexesQueryParams `protobuf:"bytes,1,opt,name=query_params,json=queryParams,proto3" json:"query_params,omitempty"`
	Verbose     bool                         `protobuf:"varint,2,opt,name=verbose,proto3" json:"verbose,omitempty"`
}

func (x *StationComplexesQuery) Reset() {
	*x = StationComplexesQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queries_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StationComplexesQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StationComplexesQuery) ProtoMessage() {}

func (x *StationComplexesQuery) ProtoReflect() protoreflect.Message {
	mi := &file_queries_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StationComplexesQuery.ProtoReflect.Descriptor instead.
func (*StationComplexesQuery) Descriptor() ([]byte, []int) {
	return file_queries_proto_rawDescGZIP(), []int{7}
}

func (x *StationComplexesQuery) GetQueryParams() *StationComplexesQueryParams {
	if x != nil {
		return x.QueryParams
	}
	return nil
}

func (x *StationComplexesQuery) GetVerbose() bool {
	if x != nil {
		return x.Verbose
	}
	return false
}

var File_queries_proto protoreflect.FileDescriptor

var file_queries_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x0b, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x63,
	0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x74, 0x66, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x74, 0x66, 0x73, 0x53, 0x74, 0x6f,
	0x70, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x51, 0x0a, 0x0e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0c, 0x67, 0x74, 0x66, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x74, 0x66, 0x73, 0x53, 0x74, 0x6f, 0x70, 0x49, 0x64, 0x22,
	0xb9, 0x02, 0x0a, 0x0f, 0x42, 0x61, 0x73, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x48, 0x00, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x32, 0x0a, 0x07, 0x62, 0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x48, 0x01, 0x52, 0x07, 0x62, 0x6f, 0x72, 0x6f, 0x75,
	0x67, 0x68, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x02, 0x52,
	0x08, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x48, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x62, 0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x69, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62,
	0x79, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x86, 0x02, 0x0a, 0x13,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x41, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x25,
	0x0a, 0x0c, 0x67, 0x74, 0x66, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0a, 0x67, 0x74, 0x66, 0x73, 0x53, 0x74, 0x6f, 0x70,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78,
	0x5f, 0x69, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x67, 0x74, 0x66, 0x73, 0x5f, 0x73, 0x74, 0x6f,
	0x70, 0x5f, 0x69, 0x64, 0x22, 0x53, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x42, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0b, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x50, 0x0a, 0x15, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x1b,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x73,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x41, 0x0a, 0x0b, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x00, 0x52,
	0x0a, 0x62, 0x61, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x88, 0x01, 0x01, 0x12, 0x22,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x69,
	0x64, 0x22, 0x7d, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x78, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x4a, 0x0a, 0x0c, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x73, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65,
	0x2a, 0x4d, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x47, 0x54, 0x46, 0x53, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x49, 0x44, 0x10, 0x01, 0x12,
	0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x02, 0x12,
	0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x5f, 0x49, 0x44, 0x10, 0x03, 0x42,
	0x92, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x42, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a,
	0x61, 0x6c, 0x61, 0x76, 0x6f, 0x73, 0x75, 0x73, 0x2f, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2d, 0x67, 0x6f, 0x3b, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x4d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0a, 0x4d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16,
	0x4d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x4d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_queries_proto_rawDescOnce sync.Once
	file_queries_proto_rawDescData = file_queries_proto_rawDesc
)

func file_queries_proto_rawDescGZIP() []byte {
	file_queries_proto_rawDescOnce.Do(func() {
		file_queries_proto_rawDescData = protoimpl.X.CompressGZIP(file_queries_proto_rawDescData)
	})
	return file_queries_proto_rawDescData
}

var file_queries_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_queries_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_queries_proto_goTypes = []interface{}{
	(QueryOrderBy)(0),                   // 0: mtadata.v1.QueryOrderBy
	(*UpcomingTrainsRequest)(nil),       // 1: mtadata.v1.UpcomingTrainsRequest
	(*StationRequest)(nil),              // 2: mtadata.v1.StationRequest
	(*BaseQueryParams)(nil),             // 3: mtadata.v1.BaseQueryParams
	(*StationsQueryParams)(nil),         // 4: mtadata.v1.StationsQueryParams
	(*StationsQuery)(nil),               // 5: mtadata.v1.StationsQuery
	(*StationComplexRequest)(nil),       // 6: mtadata.v1.StationComplexRequest
	(*StationComplexesQueryParams)(nil), // 7: mtadata.v1.StationComplexesQueryParams
	(*StationComplexesQuery)(nil),       // 8: mtadata.v1.StationComplexesQuery
	(Route)(0),                          // 9: mtadata.v1.Route
	(Borough)(0),                        // 10: mtadata.v1.Borough
	(Division)(0),                       // 11: mtadata.v1.Division
}
var file_queries_proto_depIdxs = []int32{
	9,  // 0: mtadata.v1.UpcomingTrainsRequest.route:type_name -> mtadata.v1.Route
	9,  // 1: mtadata.v1.BaseQueryParams.route:type_name -> mtadata.v1.Route
	10, // 2: mtadata.v1.BaseQueryParams.borough:type_name -> mtadata.v1.Borough
	11, // 3: mtadata.v1.BaseQueryParams.division:type_name -> mtadata.v1.Division
	0,  // 4: mtadata.v1.BaseQueryParams.order_by:type_name -> mtadata.v1.QueryOrderBy
	3,  // 5: mtadata.v1.StationsQueryParams.base_params:type_name -> mtadata.v1.BaseQueryParams
	4,  // 6: mtadata.v1.StationsQuery.query_params:type_name -> mtadata.v1.StationsQueryParams
	3,  // 7: mtadata.v1.StationComplexesQueryParams.base_params:type_name -> mtadata.v1.BaseQueryParams
	7,  // 8: mtadata.v1.StationComplexesQuery.query_params:type_name -> mtadata.v1.StationComplexesQueryParams
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_queries_proto_init() }
func file_queries_proto_init() {
	if File_queries_proto != nil {
		return
	}
	file_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_queries_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpcomingTrainsRequest); i {
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
		file_queries_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationRequest); i {
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
		file_queries_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseQueryParams); i {
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
		file_queries_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationsQueryParams); i {
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
		file_queries_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationsQuery); i {
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
		file_queries_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationComplexRequest); i {
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
		file_queries_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationComplexesQueryParams); i {
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
		file_queries_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StationComplexesQuery); i {
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
	file_queries_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_queries_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_queries_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_queries_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_queries_proto_goTypes,
		DependencyIndexes: file_queries_proto_depIdxs,
		EnumInfos:         file_queries_proto_enumTypes,
		MessageInfos:      file_queries_proto_msgTypes,
	}.Build()
	File_queries_proto = out.File
	file_queries_proto_rawDesc = nil
	file_queries_proto_goTypes = nil
	file_queries_proto_depIdxs = nil
}