// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// ColorsApiClient is the client API for ColorsApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ColorsApiClient interface {
	GetColors(ctx context.Context, in *ColorsRequest, opts ...grpc.CallOption) (ColorsApi_GetColorsClient, error)
	CreateColor(ctx context.Context, in *NewColorRequest, opts ...grpc.CallOption) (*NewColorResponse, error)
	GetMwColors(ctx context.Context, in *ManufacturersWColorsRequest, opts ...grpc.CallOption) (ColorsApi_GetMwColorsClient, error)
}

type colorsApiClient struct {
	cc grpc.ClientConnInterface
}

func NewColorsApiClient(cc grpc.ClientConnInterface) ColorsApiClient {
	return &colorsApiClient{cc}
}

func (c *colorsApiClient) GetColors(ctx context.Context, in *ColorsRequest, opts ...grpc.CallOption) (ColorsApi_GetColorsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ColorsApi_ServiceDesc.Streams[0], "/wood.color.ColorsApi/GetColors", opts...)
	if err != nil {
		return nil, err
	}
	x := &colorsApiGetColorsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ColorsApi_GetColorsClient interface {
	Recv() (*ColorsResponse, error)
	grpc.ClientStream
}

type colorsApiGetColorsClient struct {
	grpc.ClientStream
}

func (x *colorsApiGetColorsClient) Recv() (*ColorsResponse, error) {
	m := new(ColorsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *colorsApiClient) CreateColor(ctx context.Context, in *NewColorRequest, opts ...grpc.CallOption) (*NewColorResponse, error) {
	out := new(NewColorResponse)
	err := c.cc.Invoke(ctx, "/wood.color.ColorsApi/CreateColor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *colorsApiClient) GetMwColors(ctx context.Context, in *ManufacturersWColorsRequest, opts ...grpc.CallOption) (ColorsApi_GetMwColorsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ColorsApi_ServiceDesc.Streams[1], "/wood.color.ColorsApi/GetMwColors", opts...)
	if err != nil {
		return nil, err
	}
	x := &colorsApiGetMwColorsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ColorsApi_GetMwColorsClient interface {
	Recv() (*ManufacturersWColorsResponse, error)
	grpc.ClientStream
}

type colorsApiGetMwColorsClient struct {
	grpc.ClientStream
}

func (x *colorsApiGetMwColorsClient) Recv() (*ManufacturersWColorsResponse, error) {
	m := new(ManufacturersWColorsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ColorsApiServer is the server API for ColorsApi service.
// All implementations must embed UnimplementedColorsApiServer
// for forward compatibility
type ColorsApiServer interface {
	GetColors(*ColorsRequest, ColorsApi_GetColorsServer) error
	CreateColor(context.Context, *NewColorRequest) (*NewColorResponse, error)
	GetMwColors(*ManufacturersWColorsRequest, ColorsApi_GetMwColorsServer) error
	mustEmbedUnimplementedColorsApiServer()
}

// UnimplementedColorsApiServer must be embedded to have forward compatible implementations.
type UnimplementedColorsApiServer struct {
}

func (UnimplementedColorsApiServer) GetColors(*ColorsRequest, ColorsApi_GetColorsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetColors not implemented")
}
func (UnimplementedColorsApiServer) CreateColor(context.Context, *NewColorRequest) (*NewColorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateColor not implemented")
}
func (UnimplementedColorsApiServer) GetMwColors(*ManufacturersWColorsRequest, ColorsApi_GetMwColorsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMwColors not implemented")
}
func (UnimplementedColorsApiServer) mustEmbedUnimplementedColorsApiServer() {}

// UnsafeColorsApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ColorsApiServer will
// result in compilation errors.
type UnsafeColorsApiServer interface {
	mustEmbedUnimplementedColorsApiServer()
}

func RegisterColorsApiServer(s grpc.ServiceRegistrar, srv ColorsApiServer) {
	s.RegisterService(&ColorsApi_ServiceDesc, srv)
}

func _ColorsApi_GetColors_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ColorsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ColorsApiServer).GetColors(m, &colorsApiGetColorsServer{stream})
}

type ColorsApi_GetColorsServer interface {
	Send(*ColorsResponse) error
	grpc.ServerStream
}

type colorsApiGetColorsServer struct {
	grpc.ServerStream
}

func (x *colorsApiGetColorsServer) Send(m *ColorsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ColorsApi_CreateColor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewColorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColorsApiServer).CreateColor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wood.color.ColorsApi/CreateColor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColorsApiServer).CreateColor(ctx, req.(*NewColorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ColorsApi_GetMwColors_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ManufacturersWColorsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ColorsApiServer).GetMwColors(m, &colorsApiGetMwColorsServer{stream})
}

type ColorsApi_GetMwColorsServer interface {
	Send(*ManufacturersWColorsResponse) error
	grpc.ServerStream
}

type colorsApiGetMwColorsServer struct {
	grpc.ServerStream
}

func (x *colorsApiGetMwColorsServer) Send(m *ManufacturersWColorsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ColorsApi_ServiceDesc is the grpc.ServiceDesc for ColorsApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ColorsApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wood.color.ColorsApi",
	HandlerType: (*ColorsApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateColor",
			Handler:    _ColorsApi_CreateColor_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetColors",
			Handler:       _ColorsApi_GetColors_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMwColors",
			Handler:       _ColorsApi_GetMwColors_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "wood-colors.proto",
}
