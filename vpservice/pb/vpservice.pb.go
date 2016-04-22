// Code generated by protoc-gen-go.
// source: github.com/NEU-SNS/ReverseTraceroute/vpservice/pb/vpservice.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	github.com/NEU-SNS/ReverseTraceroute/vpservice/pb/vpservice.proto

It has these top-level messages:
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import datamodel "github.com/NEU-SNS/ReverseTraceroute/datamodel"

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
const _ = proto.ProtoPackageIsVersion1

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for VPService service

type VPServiceClient interface {
	GetVPs(ctx context.Context, in *datamodel.VPRequest, opts ...grpc.CallOption) (*datamodel.VPReturn, error)
	GetRRSpoofers(ctx context.Context, in *datamodel.RRSpooferRequest, opts ...grpc.CallOption) (*datamodel.RRSpooferResponse, error)
	GetTSSpoofers(ctx context.Context, in *datamodel.TSSpooferRequest, opts ...grpc.CallOption) (*datamodel.TSSpooferResponse, error)
}

type vPServiceClient struct {
	cc *grpc.ClientConn
}

func NewVPServiceClient(cc *grpc.ClientConn) VPServiceClient {
	return &vPServiceClient{cc}
}

func (c *vPServiceClient) GetVPs(ctx context.Context, in *datamodel.VPRequest, opts ...grpc.CallOption) (*datamodel.VPReturn, error) {
	out := new(datamodel.VPReturn)
	err := grpc.Invoke(ctx, "/pb.VPService/GetVPs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vPServiceClient) GetRRSpoofers(ctx context.Context, in *datamodel.RRSpooferRequest, opts ...grpc.CallOption) (*datamodel.RRSpooferResponse, error) {
	out := new(datamodel.RRSpooferResponse)
	err := grpc.Invoke(ctx, "/pb.VPService/GetRRSpoofers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vPServiceClient) GetTSSpoofers(ctx context.Context, in *datamodel.TSSpooferRequest, opts ...grpc.CallOption) (*datamodel.TSSpooferResponse, error) {
	out := new(datamodel.TSSpooferResponse)
	err := grpc.Invoke(ctx, "/pb.VPService/GetTSSpoofers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VPService service

type VPServiceServer interface {
	GetVPs(context.Context, *datamodel.VPRequest) (*datamodel.VPReturn, error)
	GetRRSpoofers(context.Context, *datamodel.RRSpooferRequest) (*datamodel.RRSpooferResponse, error)
	GetTSSpoofers(context.Context, *datamodel.TSSpooferRequest) (*datamodel.TSSpooferResponse, error)
}

func RegisterVPServiceServer(s *grpc.Server, srv VPServiceServer) {
	s.RegisterService(&_VPService_serviceDesc, srv)
}

func _VPService_GetVPs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(datamodel.VPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPServiceServer).GetVPs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VPService/GetVPs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPServiceServer).GetVPs(ctx, req.(*datamodel.VPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VPService_GetRRSpoofers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(datamodel.RRSpooferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPServiceServer).GetRRSpoofers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VPService/GetRRSpoofers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPServiceServer).GetRRSpoofers(ctx, req.(*datamodel.RRSpooferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VPService_GetTSSpoofers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(datamodel.TSSpooferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPServiceServer).GetTSSpoofers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VPService/GetTSSpoofers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPServiceServer).GetTSSpoofers(ctx, req.(*datamodel.TSSpooferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VPService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.VPService",
	HandlerType: (*VPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVPs",
			Handler:    _VPService_GetVPs_Handler,
		},
		{
			MethodName: "GetRRSpoofers",
			Handler:    _VPService_GetRRSpoofers_Handler,
		},
		{
			MethodName: "GetTSSpoofers",
			Handler:    _VPService_GetTSSpoofers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0xcf, 0xb1, 0x4e, 0xc5, 0x20,
	0x14, 0x06, 0x60, 0x75, 0xb8, 0x89, 0x24, 0x2e, 0xe8, 0x74, 0x75, 0x72, 0x17, 0x12, 0x8d, 0x0f,
	0xe0, 0x60, 0x5c, 0xcc, 0x4d, 0x03, 0xb5, 0x3b, 0xb4, 0xc7, 0xda, 0xc4, 0x72, 0x10, 0x0e, 0x3c,
	0xae, 0xcf, 0x62, 0x63, 0x53, 0x6c, 0x4c, 0x87, 0x3b, 0xf2, 0xff, 0xfc, 0x1f, 0x81, 0x3d, 0xf5,
	0x03, 0x7d, 0x24, 0x2b, 0x5a, 0x1c, 0xe5, 0xe1, 0xf9, 0xed, 0x4e, 0x1f, 0xb4, 0x54, 0x90, 0x21,
	0x44, 0xa8, 0x83, 0x69, 0x21, 0x60, 0x22, 0x90, 0xd9, 0x47, 0x08, 0x79, 0x68, 0x41, 0x7a, 0xfb,
	0x77, 0x10, 0x3e, 0x20, 0x21, 0x3f, 0xf3, 0x76, 0x7f, 0x1c, 0xd3, 0x19, 0x32, 0x23, 0x76, 0xf0,
	0x29, 0xb3, 0x71, 0x64, 0x7a, 0xf0, 0x38, 0x38, 0x9a, 0x99, 0xfb, 0xef, 0x53, 0x76, 0xde, 0x54,
	0x7a, 0xa6, 0xf9, 0x23, 0xdb, 0xbd, 0x00, 0x35, 0x55, 0xe4, 0x57, 0xa2, 0xcc, 0x44, 0x53, 0x29,
	0xf8, 0x4a, 0x10, 0x69, 0x7f, 0xf9, 0x2f, 0xa5, 0x14, 0xdc, 0xed, 0x09, 0x7f, 0x65, 0x17, 0xd3,
	0x4c, 0x29, 0xed, 0x11, 0xdf, 0xa7, 0xd7, 0xf9, 0xf5, 0xea, 0x5e, 0x89, 0x17, 0xe4, 0x66, 0xbb,
	0x8c, 0x1e, 0x5d, 0x84, 0xa2, 0xd5, 0x7a, 0x53, 0x2b, 0xf1, 0x96, 0xb6, 0x2a, 0x17, 0xcd, 0xee,
	0x7e, 0xff, 0xf9, 0xf0, 0x13, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xe5, 0x09, 0x2e, 0x73, 0x01, 0x00,
	0x00,
}
