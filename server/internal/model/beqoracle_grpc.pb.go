// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/beqoracle.proto

package model

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BeqOracleClient is the client API for BeqOracle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BeqOracleClient interface {
	CreateAnswer(ctx context.Context, in *CreateAnswerRequest, opts ...grpc.CallOption) (*Answer, error)
	UpdateAnswer(ctx context.Context, in *UpdateAnswerRequest, opts ...grpc.CallOption) (*Answer, error)
	GetAnswer(ctx context.Context, in *GetAnswerRequest, opts ...grpc.CallOption) (*Answer, error)
	DeleteAnswer(ctx context.Context, in *DeleteAnswerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetAnswerHistory(ctx context.Context, in *GetAnswerHistoryRequest, opts ...grpc.CallOption) (*EventList, error)
}

type beqOracleClient struct {
	cc grpc.ClientConnInterface
}

func NewBeqOracleClient(cc grpc.ClientConnInterface) BeqOracleClient {
	return &beqOracleClient{cc}
}

func (c *beqOracleClient) CreateAnswer(ctx context.Context, in *CreateAnswerRequest, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/model.BeqOracle/CreateAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beqOracleClient) UpdateAnswer(ctx context.Context, in *UpdateAnswerRequest, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/model.BeqOracle/UpdateAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beqOracleClient) GetAnswer(ctx context.Context, in *GetAnswerRequest, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/model.BeqOracle/GetAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beqOracleClient) DeleteAnswer(ctx context.Context, in *DeleteAnswerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/model.BeqOracle/DeleteAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beqOracleClient) GetAnswerHistory(ctx context.Context, in *GetAnswerHistoryRequest, opts ...grpc.CallOption) (*EventList, error) {
	out := new(EventList)
	err := c.cc.Invoke(ctx, "/model.BeqOracle/GetAnswerHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BeqOracleServer is the server API for BeqOracle service.
// All implementations must embed UnimplementedBeqOracleServer
// for forward compatibility
type BeqOracleServer interface {
	CreateAnswer(context.Context, *CreateAnswerRequest) (*Answer, error)
	UpdateAnswer(context.Context, *UpdateAnswerRequest) (*Answer, error)
	GetAnswer(context.Context, *GetAnswerRequest) (*Answer, error)
	DeleteAnswer(context.Context, *DeleteAnswerRequest) (*empty.Empty, error)
	GetAnswerHistory(context.Context, *GetAnswerHistoryRequest) (*EventList, error)
	mustEmbedUnimplementedBeqOracleServer()
}

// UnimplementedBeqOracleServer must be embedded to have forward compatible implementations.
type UnimplementedBeqOracleServer struct {
}

func (UnimplementedBeqOracleServer) CreateAnswer(context.Context, *CreateAnswerRequest) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnswer not implemented")
}
func (UnimplementedBeqOracleServer) UpdateAnswer(context.Context, *UpdateAnswerRequest) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAnswer not implemented")
}
func (UnimplementedBeqOracleServer) GetAnswer(context.Context, *GetAnswerRequest) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnswer not implemented")
}
func (UnimplementedBeqOracleServer) DeleteAnswer(context.Context, *DeleteAnswerRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAnswer not implemented")
}
func (UnimplementedBeqOracleServer) GetAnswerHistory(context.Context, *GetAnswerHistoryRequest) (*EventList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnswerHistory not implemented")
}
func (UnimplementedBeqOracleServer) mustEmbedUnimplementedBeqOracleServer() {}

// UnsafeBeqOracleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BeqOracleServer will
// result in compilation errors.
type UnsafeBeqOracleServer interface {
	mustEmbedUnimplementedBeqOracleServer()
}

func RegisterBeqOracleServer(s grpc.ServiceRegistrar, srv BeqOracleServer) {
	s.RegisterService(&BeqOracle_ServiceDesc, srv)
}

func _BeqOracle_CreateAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeqOracleServer).CreateAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.BeqOracle/CreateAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeqOracleServer).CreateAnswer(ctx, req.(*CreateAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeqOracle_UpdateAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeqOracleServer).UpdateAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.BeqOracle/UpdateAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeqOracleServer).UpdateAnswer(ctx, req.(*UpdateAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeqOracle_GetAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeqOracleServer).GetAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.BeqOracle/GetAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeqOracleServer).GetAnswer(ctx, req.(*GetAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeqOracle_DeleteAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeqOracleServer).DeleteAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.BeqOracle/DeleteAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeqOracleServer).DeleteAnswer(ctx, req.(*DeleteAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeqOracle_GetAnswerHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnswerHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeqOracleServer).GetAnswerHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.BeqOracle/GetAnswerHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeqOracleServer).GetAnswerHistory(ctx, req.(*GetAnswerHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BeqOracle_ServiceDesc is the grpc.ServiceDesc for BeqOracle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BeqOracle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.BeqOracle",
	HandlerType: (*BeqOracleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAnswer",
			Handler:    _BeqOracle_CreateAnswer_Handler,
		},
		{
			MethodName: "UpdateAnswer",
			Handler:    _BeqOracle_UpdateAnswer_Handler,
		},
		{
			MethodName: "GetAnswer",
			Handler:    _BeqOracle_GetAnswer_Handler,
		},
		{
			MethodName: "DeleteAnswer",
			Handler:    _BeqOracle_DeleteAnswer_Handler,
		},
		{
			MethodName: "GetAnswerHistory",
			Handler:    _BeqOracle_GetAnswerHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/beqoracle.proto",
}