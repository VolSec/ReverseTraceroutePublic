// Code generated by protoc-gen-go.
// source: github.com/NEU-SNS/ReverseTraceroute/plcontroller/pb/plcontrollerapi.proto
// DO NOT EDIT!

/*
Package plcontrollerapi is a generated protocol buffer package.

It is generated from these files:
	github.com/NEU-SNS/ReverseTraceroute/plcontroller/pb/plcontrollerapi.proto

It has these top-level messages:
*/
package plcontrollerapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import datamodel1 "github.com/NEU-SNS/ReverseTraceroute/datamodel"
import datamodel2 "github.com/NEU-SNS/ReverseTraceroute/datamodel"
import datamodel3 "github.com/NEU-SNS/ReverseTraceroute/datamodel"
import datamodel4 "github.com/NEU-SNS/ReverseTraceroute/datamodel"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for PLController service

type PLControllerClient interface {
	Ping(ctx context.Context, opts ...grpc.CallOption) (PLController_PingClient, error)
	Traceroute(ctx context.Context, opts ...grpc.CallOption) (PLController_TracerouteClient, error)
	ReceiveSpoof(ctx context.Context, in *datamodel4.RecSpoof, opts ...grpc.CallOption) (PLController_ReceiveSpoofClient, error)
	GetVPs(ctx context.Context, in *datamodel3.VPRequest, opts ...grpc.CallOption) (PLController_GetVPsClient, error)
	AcceptProbes(ctx context.Context, in *datamodel4.SpoofedProbes, opts ...grpc.CallOption) (*datamodel4.SpoofedProbesResponse, error)
}

type pLControllerClient struct {
	cc *grpc.ClientConn
}

func NewPLControllerClient(cc *grpc.ClientConn) PLControllerClient {
	return &pLControllerClient{cc}
}

func (c *pLControllerClient) Ping(ctx context.Context, opts ...grpc.CallOption) (PLController_PingClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PLController_serviceDesc.Streams[0], c.cc, "/.PLController/Ping", opts...)
	if err != nil {
		return nil, err
	}
	x := &pLControllerPingClient{stream}
	return x, nil
}

type PLController_PingClient interface {
	Send(*datamodel1.PingArg) error
	Recv() (*datamodel1.Ping, error)
	grpc.ClientStream
}

type pLControllerPingClient struct {
	grpc.ClientStream
}

func (x *pLControllerPingClient) Send(m *datamodel1.PingArg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pLControllerPingClient) Recv() (*datamodel1.Ping, error) {
	m := new(datamodel1.Ping)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pLControllerClient) Traceroute(ctx context.Context, opts ...grpc.CallOption) (PLController_TracerouteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PLController_serviceDesc.Streams[1], c.cc, "/.PLController/Traceroute", opts...)
	if err != nil {
		return nil, err
	}
	x := &pLControllerTracerouteClient{stream}
	return x, nil
}

type PLController_TracerouteClient interface {
	Send(*datamodel2.TracerouteArg) error
	Recv() (*datamodel2.Traceroute, error)
	grpc.ClientStream
}

type pLControllerTracerouteClient struct {
	grpc.ClientStream
}

func (x *pLControllerTracerouteClient) Send(m *datamodel2.TracerouteArg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pLControllerTracerouteClient) Recv() (*datamodel2.Traceroute, error) {
	m := new(datamodel2.Traceroute)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pLControllerClient) ReceiveSpoof(ctx context.Context, in *datamodel4.RecSpoof, opts ...grpc.CallOption) (PLController_ReceiveSpoofClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PLController_serviceDesc.Streams[2], c.cc, "/.PLController/ReceiveSpoof", opts...)
	if err != nil {
		return nil, err
	}
	x := &pLControllerReceiveSpoofClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PLController_ReceiveSpoofClient interface {
	Recv() (*datamodel4.NotifyRecSpoofResponse, error)
	grpc.ClientStream
}

type pLControllerReceiveSpoofClient struct {
	grpc.ClientStream
}

func (x *pLControllerReceiveSpoofClient) Recv() (*datamodel4.NotifyRecSpoofResponse, error) {
	m := new(datamodel4.NotifyRecSpoofResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pLControllerClient) GetVPs(ctx context.Context, in *datamodel3.VPRequest, opts ...grpc.CallOption) (PLController_GetVPsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PLController_serviceDesc.Streams[3], c.cc, "/.PLController/GetVPs", opts...)
	if err != nil {
		return nil, err
	}
	x := &pLControllerGetVPsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PLController_GetVPsClient interface {
	Recv() (*datamodel3.VPReturn, error)
	grpc.ClientStream
}

type pLControllerGetVPsClient struct {
	grpc.ClientStream
}

func (x *pLControllerGetVPsClient) Recv() (*datamodel3.VPReturn, error) {
	m := new(datamodel3.VPReturn)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pLControllerClient) AcceptProbes(ctx context.Context, in *datamodel4.SpoofedProbes, opts ...grpc.CallOption) (*datamodel4.SpoofedProbesResponse, error) {
	out := new(datamodel4.SpoofedProbesResponse)
	err := grpc.Invoke(ctx, "/.PLController/AcceptProbes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PLController service

type PLControllerServer interface {
	Ping(PLController_PingServer) error
	Traceroute(PLController_TracerouteServer) error
	ReceiveSpoof(*datamodel4.RecSpoof, PLController_ReceiveSpoofServer) error
	GetVPs(*datamodel3.VPRequest, PLController_GetVPsServer) error
	AcceptProbes(context.Context, *datamodel4.SpoofedProbes) (*datamodel4.SpoofedProbesResponse, error)
}

func RegisterPLControllerServer(s *grpc.Server, srv PLControllerServer) {
	s.RegisterService(&_PLController_serviceDesc, srv)
}

func _PLController_Ping_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PLControllerServer).Ping(&pLControllerPingServer{stream})
}

type PLController_PingServer interface {
	Send(*datamodel1.Ping) error
	Recv() (*datamodel1.PingArg, error)
	grpc.ServerStream
}

type pLControllerPingServer struct {
	grpc.ServerStream
}

func (x *pLControllerPingServer) Send(m *datamodel1.Ping) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pLControllerPingServer) Recv() (*datamodel1.PingArg, error) {
	m := new(datamodel1.PingArg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PLController_Traceroute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PLControllerServer).Traceroute(&pLControllerTracerouteServer{stream})
}

type PLController_TracerouteServer interface {
	Send(*datamodel2.Traceroute) error
	Recv() (*datamodel2.TracerouteArg, error)
	grpc.ServerStream
}

type pLControllerTracerouteServer struct {
	grpc.ServerStream
}

func (x *pLControllerTracerouteServer) Send(m *datamodel2.Traceroute) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pLControllerTracerouteServer) Recv() (*datamodel2.TracerouteArg, error) {
	m := new(datamodel2.TracerouteArg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PLController_ReceiveSpoof_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(datamodel4.RecSpoof)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PLControllerServer).ReceiveSpoof(m, &pLControllerReceiveSpoofServer{stream})
}

type PLController_ReceiveSpoofServer interface {
	Send(*datamodel4.NotifyRecSpoofResponse) error
	grpc.ServerStream
}

type pLControllerReceiveSpoofServer struct {
	grpc.ServerStream
}

func (x *pLControllerReceiveSpoofServer) Send(m *datamodel4.NotifyRecSpoofResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PLController_GetVPs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(datamodel3.VPRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PLControllerServer).GetVPs(m, &pLControllerGetVPsServer{stream})
}

type PLController_GetVPsServer interface {
	Send(*datamodel3.VPReturn) error
	grpc.ServerStream
}

type pLControllerGetVPsServer struct {
	grpc.ServerStream
}

func (x *pLControllerGetVPsServer) Send(m *datamodel3.VPReturn) error {
	return x.ServerStream.SendMsg(m)
}

func _PLController_AcceptProbes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(datamodel4.SpoofedProbes)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PLControllerServer).AcceptProbes(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _PLController_serviceDesc = grpc.ServiceDesc{
	ServiceName: ".PLController",
	HandlerType: (*PLControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AcceptProbes",
			Handler:    _PLController_AcceptProbes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Ping",
			Handler:       _PLController_Ping_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Traceroute",
			Handler:       _PLController_Traceroute_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ReceiveSpoof",
			Handler:       _PLController_ReceiveSpoof_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetVPs",
			Handler:       _PLController_GetVPs_Handler,
			ServerStreams: true,
		},
	},
}