// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: services.proto

package mtadatav1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_proto protoreflect.FileDescriptor

var file_services_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x0b, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd2, 0x07, 0x0a, 0x0e, 0x4d, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x21, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x68, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19,
	0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x1a, 0x2e, 0x6d, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x0c, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x0c, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x55, 0x70, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x12,
	0x21, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x63,
	0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x75, 0x70, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67,
	0x12, 0x7c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x12, 0x21, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x65,
	0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x81,
	0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x22, 0x2e, 0x6d, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x23, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x65, 0x73, 0x3a, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x51, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x6d, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x57, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42,
	0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x17, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c,
	0x42, 0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x73, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e,
	0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x73, 0x12, 0x5a,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x6d, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5d, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x42, 0x93, 0x01, 0x0a, 0x0e, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x6c, 0x61, 0x76, 0x6f,
	0x73, 0x75, 0x73, 0x2f, 0x6d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x67, 0x6f, 0x3b, 0x6d,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02,
	0x0a, 0x4d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x4d, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x4d, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0b, 0x4d, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_services_proto_goTypes = []interface{}{
	(*StationRequest)(nil),         // 0: mtadata.v1.StationRequest
	(*StationsQuery)(nil),          // 1: mtadata.v1.StationsQuery
	(*UpcomingTrainsRequest)(nil),  // 2: mtadata.v1.UpcomingTrainsRequest
	(*StationComplexRequest)(nil),  // 3: mtadata.v1.StationComplexRequest
	(*StationComplexesQuery)(nil),  // 4: mtadata.v1.StationComplexesQuery
	(*emptypb.Empty)(nil),          // 5: google.protobuf.Empty
	(*StationResult)(nil),          // 6: mtadata.v1.StationResult
	(*StationsResult)(nil),         // 7: mtadata.v1.StationsResult
	(*UpcomingTrainsResult)(nil),   // 8: mtadata.v1.UpcomingTrainsResult
	(*StationComplexResult)(nil),   // 9: mtadata.v1.StationComplexResult
	(*StationComplexesResult)(nil), // 10: mtadata.v1.StationComplexesResult
	(*AllRoutes)(nil),              // 11: mtadata.v1.AllRoutes
	(*AllBoroughs)(nil),            // 12: mtadata.v1.AllBoroughs
	(*AllDivisions)(nil),           // 13: mtadata.v1.AllDivisions
	(*AllStructures)(nil),          // 14: mtadata.v1.AllStructures
}
var file_services_proto_depIdxs = []int32{
	0,  // 0: mtadata.v1.MtaDataService.GetStation:input_type -> mtadata.v1.StationRequest
	1,  // 1: mtadata.v1.MtaDataService.GetStations:input_type -> mtadata.v1.StationsQuery
	2,  // 2: mtadata.v1.MtaDataService.GetUpcomingTrains:input_type -> mtadata.v1.UpcomingTrainsRequest
	3,  // 3: mtadata.v1.MtaDataService.GetStationComplex:input_type -> mtadata.v1.StationComplexRequest
	4,  // 4: mtadata.v1.MtaDataService.GetStationComplexes:input_type -> mtadata.v1.StationComplexesQuery
	5,  // 5: mtadata.v1.MtaDataService.GetAllRoutes:input_type -> google.protobuf.Empty
	5,  // 6: mtadata.v1.MtaDataService.GetAllBoroughs:input_type -> google.protobuf.Empty
	5,  // 7: mtadata.v1.MtaDataService.GetAllDivisions:input_type -> google.protobuf.Empty
	5,  // 8: mtadata.v1.MtaDataService.GetAllStructures:input_type -> google.protobuf.Empty
	6,  // 9: mtadata.v1.MtaDataService.GetStation:output_type -> mtadata.v1.StationResult
	7,  // 10: mtadata.v1.MtaDataService.GetStations:output_type -> mtadata.v1.StationsResult
	8,  // 11: mtadata.v1.MtaDataService.GetUpcomingTrains:output_type -> mtadata.v1.UpcomingTrainsResult
	9,  // 12: mtadata.v1.MtaDataService.GetStationComplex:output_type -> mtadata.v1.StationComplexResult
	10, // 13: mtadata.v1.MtaDataService.GetStationComplexes:output_type -> mtadata.v1.StationComplexesResult
	11, // 14: mtadata.v1.MtaDataService.GetAllRoutes:output_type -> mtadata.v1.AllRoutes
	12, // 15: mtadata.v1.MtaDataService.GetAllBoroughs:output_type -> mtadata.v1.AllBoroughs
	13, // 16: mtadata.v1.MtaDataService.GetAllDivisions:output_type -> mtadata.v1.AllDivisions
	14, // 17: mtadata.v1.MtaDataService.GetAllStructures:output_type -> mtadata.v1.AllStructures
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_proto_init() }
func file_services_proto_init() {
	if File_services_proto != nil {
		return
	}
	file_types_proto_init()
	file_queries_proto_init()
	file_responses_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_proto_goTypes,
		DependencyIndexes: file_services_proto_depIdxs,
	}.Build()
	File_services_proto = out.File
	file_services_proto_rawDesc = nil
	file_services_proto_goTypes = nil
	file_services_proto_depIdxs = nil
}
