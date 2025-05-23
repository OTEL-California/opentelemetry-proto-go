// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.21.6
// source: opentelemetry/proto/collector/profiles/v1development/profiles_service.proto

package v1development

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProfilesServiceClient is the client API for ProfilesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfilesServiceClient interface {
	Export(ctx context.Context, in *ExportProfilesServiceRequest, opts ...grpc.CallOption) (*ExportProfilesServiceResponse, error)
}

type profilesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfilesServiceClient(cc grpc.ClientConnInterface) ProfilesServiceClient {
	return &profilesServiceClient{cc}
}

func (c *profilesServiceClient) Export(ctx context.Context, in *ExportProfilesServiceRequest, opts ...grpc.CallOption) (*ExportProfilesServiceResponse, error) {
	out := new(ExportProfilesServiceResponse)
	err := c.cc.Invoke(ctx, "/opentelemetry.proto.collector.profiles.v1development.ProfilesService/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfilesServiceServer is the server API for ProfilesService service.
// All implementations must embed UnimplementedProfilesServiceServer
// for forward compatibility
type ProfilesServiceServer interface {
	Export(context.Context, *ExportProfilesServiceRequest) (*ExportProfilesServiceResponse, error)
	mustEmbedUnimplementedProfilesServiceServer()
}

// UnimplementedProfilesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfilesServiceServer struct {
}

func (UnimplementedProfilesServiceServer) Export(context.Context, *ExportProfilesServiceRequest) (*ExportProfilesServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}
func (UnimplementedProfilesServiceServer) mustEmbedUnimplementedProfilesServiceServer() {}

// UnsafeProfilesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfilesServiceServer will
// result in compilation errors.
type UnsafeProfilesServiceServer interface {
	mustEmbedUnimplementedProfilesServiceServer()
}

func RegisterProfilesServiceServer(s grpc.ServiceRegistrar, srv ProfilesServiceServer) {
	s.RegisterService(&ProfilesService_ServiceDesc, srv)
}

func _ProfilesService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportProfilesServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opentelemetry.proto.collector.profiles.v1development.ProfilesService/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServiceServer).Export(ctx, req.(*ExportProfilesServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfilesService_ServiceDesc is the grpc.ServiceDesc for ProfilesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfilesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "opentelemetry.proto.collector.profiles.v1development.ProfilesService",
	HandlerType: (*ProfilesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler:    _ProfilesService_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opentelemetry/proto/collector/profiles/v1development/profiles_service.proto",
}
