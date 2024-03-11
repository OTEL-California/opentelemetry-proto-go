// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.6
// source: opentelemetry/proto/collector/metrics/v1/metrics_service.proto

package v1

import (
	v1 "go.opentelemetry.io/proto/slim/otlp/metrics/v1"
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

type ExportMetricsServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An array of ResourceMetrics.
	// For data coming from a single resource this array will typically contain one
	// element. Intermediary nodes (such as OpenTelemetry Collector) that receive
	// data from multiple origins typically batch the data before forwarding further and
	// in that case this array will contain multiple elements.
	ResourceMetrics []*v1.ResourceMetrics `protobuf:"bytes,1,rep,name=resource_metrics,json=resourceMetrics,proto3" json:"resource_metrics,omitempty"`
}

func (x *ExportMetricsServiceRequest) Reset() {
	*x = ExportMetricsServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportMetricsServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportMetricsServiceRequest) ProtoMessage() {}

func (x *ExportMetricsServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportMetricsServiceRequest.ProtoReflect.Descriptor instead.
func (*ExportMetricsServiceRequest) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_rawDescGZIP(), []int{0}
}

func (x *ExportMetricsServiceRequest) GetResourceMetrics() []*v1.ResourceMetrics {
	if x != nil {
		return x.ResourceMetrics
	}
	return nil
}

type ExportMetricsServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The details of a partially successful export request.
	//
	// If the request is only partially accepted
	// (i.e. when the server accepts only parts of the data and rejects the rest)
	// the server MUST initialize the `partial_success` field and MUST
	// set the `rejected_<signal>` with the number of items it rejected.
	//
	// Servers MAY also make use of the `partial_success` field to convey
	// warnings/suggestions to senders even when the request was fully accepted.
	// In such cases, the `rejected_<signal>` MUST have a value of `0` and
	// the `error_message` MUST be non-empty.
	//
	// A `partial_success` message with an empty value (rejected_<signal> = 0 and
	// `error_message` = "") is equivalent to it not being set/present. Senders
	// SHOULD interpret it the same way as in the full success case.
	PartialSuccess *ExportMetricsPartialSuccess `protobuf:"bytes,1,opt,name=partial_success,json=partialSuccess,proto3" json:"partial_success,omitempty"`
}

func (x *ExportMetricsServiceResponse) Reset() {
	*x = ExportMetricsServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportMetricsServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportMetricsServiceResponse) ProtoMessage() {}

func (x *ExportMetricsServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportMetricsServiceResponse.ProtoReflect.Descriptor instead.
func (*ExportMetricsServiceResponse) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_rawDescGZIP(), []int{1}
}

func (x *ExportMetricsServiceResponse) GetPartialSuccess() *ExportMetricsPartialSuccess {
	if x != nil {
		return x.PartialSuccess
	}
	return nil
}

type ExportMetricsPartialSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of rejected data points.
	//
	// A `rejected_<signal>` field holding a `0` value indicates that the
	// request was fully accepted.
	RejectedDataPoints int64 `protobuf:"varint,1,opt,name=rejected_data_points,json=rejectedDataPoints,proto3" json:"rejected_data_points,omitempty"`
	// A developer-facing human-readable message in English. It should be used
	// either to explain why the server rejected parts of the data during a partial
	// success or to convey warnings/suggestions during a full success. The message
	// should offer guidance on how users can address such issues.
	//
	// error_message is an optional field. An error_message with an empty value
	// is equivalent to it not being set.
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *ExportMetricsPartialSuccess) Reset() {
	*x = ExportMetricsPartialSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportMetricsPartialSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportMetricsPartialSuccess) ProtoMessage() {}

func (x *ExportMetricsPartialSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportMetricsPartialSuccess.ProtoReflect.Descriptor instead.
func (*ExportMetricsPartialSuccess) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_rawDescGZIP(), []int{2}
}

func (x *ExportMetricsPartialSuccess) GetRejectedDataPoints() int64 {
	if x != nil {
		return x.RejectedDataPoints
	}
	return 0
}

func (x *ExportMetricsPartialSuccess) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_opentelemetry_proto_collector_metrics_v1_metrics_service_proto protoreflect.FileDescriptor

var file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x28, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2c, 0x6f, 0x70, 0x65, 0x6e,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x1b, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5a, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x1c, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x74, 0x0a, 0x1b, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x12, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xac, 0x01, 0x0a, 0x0e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x99, 0x01,
	0x0a, 0x06, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x45, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x46, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xa9, 0x01, 0x0a, 0x2b, 0x69, 0x6f,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x38, 0x67, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6c, 0x69, 0x6d,
	0x2f, 0x6f, 0x74, 0x6c, 0x70, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0xaa, 0x02, 0x28, 0x4f, 0x70, 0x65,
	0x6e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_rawDescOnce sync.Once
	file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_rawDescData = file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_rawDesc
)

func file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_rawDescGZIP() []byte {
	file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_rawDescOnce.Do(func() {
		file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_rawDescData)
	})
	return file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_rawDescData
}

var file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_goTypes = []interface{}{
	(*ExportMetricsServiceRequest)(nil),  // 0: opentelemetry.proto.collector.metrics.v1.ExportMetricsServiceRequest
	(*ExportMetricsServiceResponse)(nil), // 1: opentelemetry.proto.collector.metrics.v1.ExportMetricsServiceResponse
	(*ExportMetricsPartialSuccess)(nil),  // 2: opentelemetry.proto.collector.metrics.v1.ExportMetricsPartialSuccess
	(*v1.ResourceMetrics)(nil),           // 3: opentelemetry.proto.metrics.v1.ResourceMetrics
}
var file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_depIdxs = []int32{
	3, // 0: opentelemetry.proto.collector.metrics.v1.ExportMetricsServiceRequest.resource_metrics:type_name -> opentelemetry.proto.metrics.v1.ResourceMetrics
	2, // 1: opentelemetry.proto.collector.metrics.v1.ExportMetricsServiceResponse.partial_success:type_name -> opentelemetry.proto.collector.metrics.v1.ExportMetricsPartialSuccess
	0, // 2: opentelemetry.proto.collector.metrics.v1.MetricsService.Export:input_type -> opentelemetry.proto.collector.metrics.v1.ExportMetricsServiceRequest
	1, // 3: opentelemetry.proto.collector.metrics.v1.MetricsService.Export:output_type -> opentelemetry.proto.collector.metrics.v1.ExportMetricsServiceResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_init() }
func file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_init() {
	if File_opentelemetry_proto_collector_metrics_v1_metrics_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportMetricsServiceRequest); i {
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
		file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportMetricsServiceResponse); i {
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
		file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportMetricsPartialSuccess); i {
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
			RawDescriptor: file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_goTypes,
		DependencyIndexes: file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_depIdxs,
		MessageInfos:      file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_msgTypes,
	}.Build()
	File_opentelemetry_proto_collector_metrics_v1_metrics_service_proto = out.File
	file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_rawDesc = nil
	file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_goTypes = nil
	file_opentelemetry_proto_collector_metrics_v1_metrics_service_proto_depIdxs = nil
}
