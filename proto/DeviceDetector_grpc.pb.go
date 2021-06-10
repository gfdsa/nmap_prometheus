// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// HomeDetectorClient is the client API for HomeDetector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HomeDetectorClient interface {
	// Sends a greeting
	Ack(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*Reply, error)
	Address(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*Reply, error)
	Addresses(ctx context.Context, in *AddressesRequest, opts ...grpc.CallOption) (*Reply, error)
	ListTimedCommands(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TCsResponse, error)
	ListCommandQueue(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CQsResponse, error)
	ListDevices(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DevicesResponse, error)
	UpdateDevice(ctx context.Context, in *Devices, opts ...grpc.CallOption) (*Reply, error)
	DeleteDevice(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*Reply, error)
	DeleteCommandQueue(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*Reply, error)
	DeleteTimedCommand(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*Reply, error)
	CompleteTimedCommands(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*Reply, error)
	CompleteTimedCommand(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*Reply, error)
}

type homeDetectorClient struct {
	cc grpc.ClientConnInterface
}

func NewHomeDetectorClient(cc grpc.ClientConnInterface) HomeDetectorClient {
	return &homeDetectorClient{cc}
}

func (c *homeDetectorClient) Ack(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.HomeDetector/Ack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeDetectorClient) Address(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.HomeDetector/Address", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeDetectorClient) Addresses(ctx context.Context, in *AddressesRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.HomeDetector/Addresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeDetectorClient) ListTimedCommands(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TCsResponse, error) {
	out := new(TCsResponse)
	err := c.cc.Invoke(ctx, "/proto.HomeDetector/ListTimedCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeDetectorClient) ListCommandQueue(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CQsResponse, error) {
	out := new(CQsResponse)
	err := c.cc.Invoke(ctx, "/proto.HomeDetector/ListCommandQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeDetectorClient) ListDevices(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DevicesResponse, error) {
	out := new(DevicesResponse)
	err := c.cc.Invoke(ctx, "/proto.HomeDetector/ListDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeDetectorClient) UpdateDevice(ctx context.Context, in *Devices, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.HomeDetector/UpdateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeDetectorClient) DeleteDevice(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.HomeDetector/DeleteDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeDetectorClient) DeleteCommandQueue(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.HomeDetector/DeleteCommandQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeDetectorClient) DeleteTimedCommand(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.HomeDetector/DeleteTimedCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeDetectorClient) CompleteTimedCommands(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.HomeDetector/CompleteTimedCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeDetectorClient) CompleteTimedCommand(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.HomeDetector/CompleteTimedCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeDetectorServer is the server API for HomeDetector service.
// All implementations must embed UnimplementedHomeDetectorServer
// for forward compatibility
type HomeDetectorServer interface {
	// Sends a greeting
	Ack(context.Context, *StringRequest) (*Reply, error)
	Address(context.Context, *AddressRequest) (*Reply, error)
	Addresses(context.Context, *AddressesRequest) (*Reply, error)
	ListTimedCommands(context.Context, *empty.Empty) (*TCsResponse, error)
	ListCommandQueue(context.Context, *empty.Empty) (*CQsResponse, error)
	ListDevices(context.Context, *empty.Empty) (*DevicesResponse, error)
	UpdateDevice(context.Context, *Devices) (*Reply, error)
	DeleteDevice(context.Context, *StringRequest) (*Reply, error)
	DeleteCommandQueue(context.Context, *StringRequest) (*Reply, error)
	DeleteTimedCommand(context.Context, *StringRequest) (*Reply, error)
	CompleteTimedCommands(context.Context, *StringRequest) (*Reply, error)
	CompleteTimedCommand(context.Context, *StringRequest) (*Reply, error)
	mustEmbedUnimplementedHomeDetectorServer()
}

// UnimplementedHomeDetectorServer must be embedded to have forward compatible implementations.
type UnimplementedHomeDetectorServer struct {
}

func (UnimplementedHomeDetectorServer) Ack(context.Context, *StringRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ack not implemented")
}
func (UnimplementedHomeDetectorServer) Address(context.Context, *AddressRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Address not implemented")
}
func (UnimplementedHomeDetectorServer) Addresses(context.Context, *AddressesRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Addresses not implemented")
}
func (UnimplementedHomeDetectorServer) ListTimedCommands(context.Context, *empty.Empty) (*TCsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTimedCommands not implemented")
}
func (UnimplementedHomeDetectorServer) ListCommandQueue(context.Context, *empty.Empty) (*CQsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCommandQueue not implemented")
}
func (UnimplementedHomeDetectorServer) ListDevices(context.Context, *empty.Empty) (*DevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDevices not implemented")
}
func (UnimplementedHomeDetectorServer) UpdateDevice(context.Context, *Devices) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDevice not implemented")
}
func (UnimplementedHomeDetectorServer) DeleteDevice(context.Context, *StringRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevice not implemented")
}
func (UnimplementedHomeDetectorServer) DeleteCommandQueue(context.Context, *StringRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommandQueue not implemented")
}
func (UnimplementedHomeDetectorServer) DeleteTimedCommand(context.Context, *StringRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTimedCommand not implemented")
}
func (UnimplementedHomeDetectorServer) CompleteTimedCommands(context.Context, *StringRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTimedCommands not implemented")
}
func (UnimplementedHomeDetectorServer) CompleteTimedCommand(context.Context, *StringRequest) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTimedCommand not implemented")
}
func (UnimplementedHomeDetectorServer) mustEmbedUnimplementedHomeDetectorServer() {}

// UnsafeHomeDetectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomeDetectorServer will
// result in compilation errors.
type UnsafeHomeDetectorServer interface {
	mustEmbedUnimplementedHomeDetectorServer()
}

func RegisterHomeDetectorServer(s grpc.ServiceRegistrar, srv HomeDetectorServer) {
	s.RegisterService(&_HomeDetector_serviceDesc, srv)
}

func _HomeDetector_Ack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeDetectorServer).Ack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HomeDetector/Ack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeDetectorServer).Ack(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeDetector_Address_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeDetectorServer).Address(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HomeDetector/Address",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeDetectorServer).Address(ctx, req.(*AddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeDetector_Addresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeDetectorServer).Addresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HomeDetector/Addresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeDetectorServer).Addresses(ctx, req.(*AddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeDetector_ListTimedCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeDetectorServer).ListTimedCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HomeDetector/ListTimedCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeDetectorServer).ListTimedCommands(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeDetector_ListCommandQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeDetectorServer).ListCommandQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HomeDetector/ListCommandQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeDetectorServer).ListCommandQueue(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeDetector_ListDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeDetectorServer).ListDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HomeDetector/ListDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeDetectorServer).ListDevices(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeDetector_UpdateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Devices)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeDetectorServer).UpdateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HomeDetector/UpdateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeDetectorServer).UpdateDevice(ctx, req.(*Devices))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeDetector_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeDetectorServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HomeDetector/DeleteDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeDetectorServer).DeleteDevice(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeDetector_DeleteCommandQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeDetectorServer).DeleteCommandQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HomeDetector/DeleteCommandQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeDetectorServer).DeleteCommandQueue(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeDetector_DeleteTimedCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeDetectorServer).DeleteTimedCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HomeDetector/DeleteTimedCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeDetectorServer).DeleteTimedCommand(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeDetector_CompleteTimedCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeDetectorServer).CompleteTimedCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HomeDetector/CompleteTimedCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeDetectorServer).CompleteTimedCommands(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeDetector_CompleteTimedCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeDetectorServer).CompleteTimedCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HomeDetector/CompleteTimedCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeDetectorServer).CompleteTimedCommand(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HomeDetector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HomeDetector",
	HandlerType: (*HomeDetectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ack",
			Handler:    _HomeDetector_Ack_Handler,
		},
		{
			MethodName: "Address",
			Handler:    _HomeDetector_Address_Handler,
		},
		{
			MethodName: "Addresses",
			Handler:    _HomeDetector_Addresses_Handler,
		},
		{
			MethodName: "ListTimedCommands",
			Handler:    _HomeDetector_ListTimedCommands_Handler,
		},
		{
			MethodName: "ListCommandQueue",
			Handler:    _HomeDetector_ListCommandQueue_Handler,
		},
		{
			MethodName: "ListDevices",
			Handler:    _HomeDetector_ListDevices_Handler,
		},
		{
			MethodName: "UpdateDevice",
			Handler:    _HomeDetector_UpdateDevice_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _HomeDetector_DeleteDevice_Handler,
		},
		{
			MethodName: "DeleteCommandQueue",
			Handler:    _HomeDetector_DeleteCommandQueue_Handler,
		},
		{
			MethodName: "DeleteTimedCommand",
			Handler:    _HomeDetector_DeleteTimedCommand_Handler,
		},
		{
			MethodName: "CompleteTimedCommands",
			Handler:    _HomeDetector_CompleteTimedCommands_Handler,
		},
		{
			MethodName: "CompleteTimedCommand",
			Handler:    _HomeDetector_CompleteTimedCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "DeviceDetector.proto",
}
