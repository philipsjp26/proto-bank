// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: api/logs/proto/log_http.proto

package httplogpb

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

// HttpLogServiceClient is the client API for HttpLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HttpLogServiceClient interface {
	ServInsertHttpLog(ctx context.Context, in *CreateHttpLogRequest, opts ...grpc.CallOption) (*HttpLogResponse, error)
	ServeGetListHttpLog(ctx context.Context, in *HttpLogParameterRequest, opts ...grpc.CallOption) (*ListHttpLogResponse, error)
	ServeFindHttpLogById(ctx context.Context, in *FindHttpLogRequest, opts ...grpc.CallOption) (*HttpLogResponse, error)
	ServeUpdateHttpLogById(ctx context.Context, in *UpdateHttpLogRequest, opts ...grpc.CallOption) (*HttpLogResponse, error)
}

type httpLogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHttpLogServiceClient(cc grpc.ClientConnInterface) HttpLogServiceClient {
	return &httpLogServiceClient{cc}
}

func (c *httpLogServiceClient) ServInsertHttpLog(ctx context.Context, in *CreateHttpLogRequest, opts ...grpc.CallOption) (*HttpLogResponse, error) {
	out := new(HttpLogResponse)
	err := c.cc.Invoke(ctx, "/httplogpb.HttpLogService/ServInsertHttpLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpLogServiceClient) ServeGetListHttpLog(ctx context.Context, in *HttpLogParameterRequest, opts ...grpc.CallOption) (*ListHttpLogResponse, error) {
	out := new(ListHttpLogResponse)
	err := c.cc.Invoke(ctx, "/httplogpb.HttpLogService/ServeGetListHttpLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpLogServiceClient) ServeFindHttpLogById(ctx context.Context, in *FindHttpLogRequest, opts ...grpc.CallOption) (*HttpLogResponse, error) {
	out := new(HttpLogResponse)
	err := c.cc.Invoke(ctx, "/httplogpb.HttpLogService/ServeFindHttpLogById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpLogServiceClient) ServeUpdateHttpLogById(ctx context.Context, in *UpdateHttpLogRequest, opts ...grpc.CallOption) (*HttpLogResponse, error) {
	out := new(HttpLogResponse)
	err := c.cc.Invoke(ctx, "/httplogpb.HttpLogService/ServeUpdateHttpLogById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HttpLogServiceServer is the server API for HttpLogService service.
// All implementations must embed UnimplementedHttpLogServiceServer
// for forward compatibility
type HttpLogServiceServer interface {
	ServInsertHttpLog(context.Context, *CreateHttpLogRequest) (*HttpLogResponse, error)
	ServeGetListHttpLog(context.Context, *HttpLogParameterRequest) (*ListHttpLogResponse, error)
	ServeFindHttpLogById(context.Context, *FindHttpLogRequest) (*HttpLogResponse, error)
	ServeUpdateHttpLogById(context.Context, *UpdateHttpLogRequest) (*HttpLogResponse, error)
	mustEmbedUnimplementedHttpLogServiceServer()
}

// UnimplementedHttpLogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHttpLogServiceServer struct {
}

func (UnimplementedHttpLogServiceServer) ServInsertHttpLog(context.Context, *CreateHttpLogRequest) (*HttpLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServInsertHttpLog not implemented")
}
func (UnimplementedHttpLogServiceServer) ServeGetListHttpLog(context.Context, *HttpLogParameterRequest) (*ListHttpLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServeGetListHttpLog not implemented")
}
func (UnimplementedHttpLogServiceServer) ServeFindHttpLogById(context.Context, *FindHttpLogRequest) (*HttpLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServeFindHttpLogById not implemented")
}
func (UnimplementedHttpLogServiceServer) ServeUpdateHttpLogById(context.Context, *UpdateHttpLogRequest) (*HttpLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServeUpdateHttpLogById not implemented")
}
func (UnimplementedHttpLogServiceServer) mustEmbedUnimplementedHttpLogServiceServer() {}

// UnsafeHttpLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HttpLogServiceServer will
// result in compilation errors.
type UnsafeHttpLogServiceServer interface {
	mustEmbedUnimplementedHttpLogServiceServer()
}

func RegisterHttpLogServiceServer(s grpc.ServiceRegistrar, srv HttpLogServiceServer) {
	s.RegisterService(&HttpLogService_ServiceDesc, srv)
}

func _HttpLogService_ServInsertHttpLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHttpLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpLogServiceServer).ServInsertHttpLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/httplogpb.HttpLogService/ServInsertHttpLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpLogServiceServer).ServInsertHttpLog(ctx, req.(*CreateHttpLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpLogService_ServeGetListHttpLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HttpLogParameterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpLogServiceServer).ServeGetListHttpLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/httplogpb.HttpLogService/ServeGetListHttpLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpLogServiceServer).ServeGetListHttpLog(ctx, req.(*HttpLogParameterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpLogService_ServeFindHttpLogById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindHttpLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpLogServiceServer).ServeFindHttpLogById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/httplogpb.HttpLogService/ServeFindHttpLogById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpLogServiceServer).ServeFindHttpLogById(ctx, req.(*FindHttpLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpLogService_ServeUpdateHttpLogById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHttpLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpLogServiceServer).ServeUpdateHttpLogById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/httplogpb.HttpLogService/ServeUpdateHttpLogById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpLogServiceServer).ServeUpdateHttpLogById(ctx, req.(*UpdateHttpLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HttpLogService_ServiceDesc is the grpc.ServiceDesc for HttpLogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HttpLogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "httplogpb.HttpLogService",
	HandlerType: (*HttpLogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServInsertHttpLog",
			Handler:    _HttpLogService_ServInsertHttpLog_Handler,
		},
		{
			MethodName: "ServeGetListHttpLog",
			Handler:    _HttpLogService_ServeGetListHttpLog_Handler,
		},
		{
			MethodName: "ServeFindHttpLogById",
			Handler:    _HttpLogService_ServeFindHttpLogById_Handler,
		},
		{
			MethodName: "ServeUpdateHttpLogById",
			Handler:    _HttpLogService_ServeUpdateHttpLogById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/logs/proto/log_http.proto",
}
