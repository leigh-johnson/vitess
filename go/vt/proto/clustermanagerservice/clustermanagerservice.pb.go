// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clustermanagerservice.proto

package clustermanagerservice // import "vitess.io/vitess/go/vt/proto/clustermanagerservice"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import clustermanagerdata "vitess.io/vitess/go/vt/proto/clustermanagerdata"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VtClusterManager service

type VtClusterManagerClient interface {
	SetShardConfig(ctx context.Context, in *clustermanagerdata.SetShardConfigRequest, opts ...grpc.CallOption) (*clustermanagerdata.ClusterConfig, error)
	// GetClusterConfig defines api to get configuration about
	// cell / keyspace / shard / tablet_type.
	GetClusterConfig(ctx context.Context, in *clustermanagerdata.GetClusterConfigRequest, opts ...grpc.CallOption) (*clustermanagerdata.ClusterConfig, error)
	// DeleteShardConfig removes configuration from cluster manager
	DeleteShardConfig(ctx context.Context, in *clustermanagerdata.DeleteShardConfigRequest, opts ...grpc.CallOption) (*clustermanagerdata.DeleteShardConfigResponse, error)
}

type vtClusterManagerClient struct {
	cc *grpc.ClientConn
}

func NewVtClusterManagerClient(cc *grpc.ClientConn) VtClusterManagerClient {
	return &vtClusterManagerClient{cc}
}

func (c *vtClusterManagerClient) SetShardConfig(ctx context.Context, in *clustermanagerdata.SetShardConfigRequest, opts ...grpc.CallOption) (*clustermanagerdata.ClusterConfig, error) {
	out := new(clustermanagerdata.ClusterConfig)
	err := grpc.Invoke(ctx, "/clustermanagerservice.VtClusterManager/SetShardConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vtClusterManagerClient) GetClusterConfig(ctx context.Context, in *clustermanagerdata.GetClusterConfigRequest, opts ...grpc.CallOption) (*clustermanagerdata.ClusterConfig, error) {
	out := new(clustermanagerdata.ClusterConfig)
	err := grpc.Invoke(ctx, "/clustermanagerservice.VtClusterManager/GetClusterConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vtClusterManagerClient) DeleteShardConfig(ctx context.Context, in *clustermanagerdata.DeleteShardConfigRequest, opts ...grpc.CallOption) (*clustermanagerdata.DeleteShardConfigResponse, error) {
	out := new(clustermanagerdata.DeleteShardConfigResponse)
	err := grpc.Invoke(ctx, "/clustermanagerservice.VtClusterManager/DeleteShardConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VtClusterManager service

type VtClusterManagerServer interface {
	SetShardConfig(context.Context, *clustermanagerdata.SetShardConfigRequest) (*clustermanagerdata.ClusterConfig, error)
	// GetClusterConfig defines api to get configuration about
	// cell / keyspace / shard / tablet_type.
	GetClusterConfig(context.Context, *clustermanagerdata.GetClusterConfigRequest) (*clustermanagerdata.ClusterConfig, error)
	// DeleteShardConfig removes configuration from cluster manager
	DeleteShardConfig(context.Context, *clustermanagerdata.DeleteShardConfigRequest) (*clustermanagerdata.DeleteShardConfigResponse, error)
}

func RegisterVtClusterManagerServer(s *grpc.Server, srv VtClusterManagerServer) {
	s.RegisterService(&_VtClusterManager_serviceDesc, srv)
}

func _VtClusterManager_SetShardConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clustermanagerdata.SetShardConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VtClusterManagerServer).SetShardConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clustermanagerservice.VtClusterManager/SetShardConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VtClusterManagerServer).SetShardConfig(ctx, req.(*clustermanagerdata.SetShardConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VtClusterManager_GetClusterConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clustermanagerdata.GetClusterConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VtClusterManagerServer).GetClusterConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clustermanagerservice.VtClusterManager/GetClusterConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VtClusterManagerServer).GetClusterConfig(ctx, req.(*clustermanagerdata.GetClusterConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VtClusterManager_DeleteShardConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clustermanagerdata.DeleteShardConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VtClusterManagerServer).DeleteShardConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clustermanagerservice.VtClusterManager/DeleteShardConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VtClusterManagerServer).DeleteShardConfig(ctx, req.(*clustermanagerdata.DeleteShardConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VtClusterManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clustermanagerservice.VtClusterManager",
	HandlerType: (*VtClusterManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetShardConfig",
			Handler:    _VtClusterManager_SetShardConfig_Handler,
		},
		{
			MethodName: "GetClusterConfig",
			Handler:    _VtClusterManager_GetClusterConfig_Handler,
		},
		{
			MethodName: "DeleteShardConfig",
			Handler:    _VtClusterManager_DeleteShardConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clustermanagerservice.proto",
}

func init() {
	proto.RegisterFile("clustermanagerservice.proto", fileDescriptor_clustermanagerservice_c8be03c48046dd05)
}

var fileDescriptor_clustermanagerservice_c8be03c48046dd05 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xce, 0x29, 0x2d,
	0x2e, 0x49, 0x2d, 0xca, 0x4d, 0xcc, 0x4b, 0x4c, 0x4f, 0x2d, 0x2a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc5, 0x2a, 0x29, 0x25, 0x81, 0x2a, 0x9c,
	0x92, 0x58, 0x92, 0x08, 0xd1, 0x60, 0x74, 0x8e, 0x89, 0x4b, 0x20, 0xac, 0xc4, 0x19, 0x22, 0xed,
	0x0b, 0x91, 0x16, 0x4a, 0xe0, 0xe2, 0x0b, 0x4e, 0x2d, 0x09, 0xce, 0x48, 0x2c, 0x4a, 0x71, 0xce,
	0xcf, 0x4b, 0xcb, 0x4c, 0x17, 0xd2, 0xd4, 0xc3, 0x62, 0x02, 0xaa, 0x9a, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x29, 0x45, 0x6c, 0x4a, 0xa1, 0x16, 0x40, 0x54, 0x2a, 0x31, 0x08, 0xa5, 0x70,
	0x09, 0xb8, 0xa7, 0x96, 0xa0, 0x88, 0x0a, 0x69, 0x63, 0xd3, 0x88, 0xae, 0x8a, 0x24, 0x5b, 0x8a,
	0xb8, 0x04, 0x5d, 0x52, 0x73, 0x52, 0x4b, 0x52, 0x91, 0xbd, 0xa2, 0x83, 0x4d, 0x27, 0x86, 0x32,
	0x98, 0x3d, 0xba, 0x44, 0xaa, 0x2e, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x55, 0x62, 0x70, 0x72, 0xe5,
	0xe2, 0xcf, 0xcc, 0xd7, 0x2b, 0xcb, 0x2c, 0x49, 0x2d, 0x2e, 0x86, 0x84, 0x71, 0x94, 0x11, 0x94,
	0x97, 0x99, 0xaf, 0x0f, 0x61, 0xe9, 0xa7, 0xe7, 0xeb, 0x97, 0x95, 0xe8, 0x83, 0x65, 0xf5, 0xb1,
	0xc6, 0x58, 0x12, 0x1b, 0x58, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x66, 0x02, 0xe8,
	0xee, 0x01, 0x00, 0x00,
}
