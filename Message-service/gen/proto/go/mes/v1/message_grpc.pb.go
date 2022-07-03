// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package message

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

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageClient interface {
	//测试
	Message(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	//向手机号发送验证码
	SentMessageCode(ctx context.Context, in *SentMessageCodeRequest, opts ...grpc.CallOption) (*SentMessageCodeResponse, error)
	//校验验证码
	CheckMessageCode(ctx context.Context, in *CheckMessageCodeRequest, opts ...grpc.CallOption) (*CheckMessageCodeResponse, error)
}

type messageClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageClient(cc grpc.ClientConnInterface) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) Message(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/mes.message/message", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) SentMessageCode(ctx context.Context, in *SentMessageCodeRequest, opts ...grpc.CallOption) (*SentMessageCodeResponse, error) {
	out := new(SentMessageCodeResponse)
	err := c.cc.Invoke(ctx, "/mes.message/SentMessageCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) CheckMessageCode(ctx context.Context, in *CheckMessageCodeRequest, opts ...grpc.CallOption) (*CheckMessageCodeResponse, error) {
	out := new(CheckMessageCodeResponse)
	err := c.cc.Invoke(ctx, "/mes.message/CheckMessageCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
// All implementations should embed UnimplementedMessageServer
// for forward compatibility
type MessageServer interface {
	//测试
	Message(context.Context, *MessageRequest) (*MessageResponse, error)
	//向手机号发送验证码
	SentMessageCode(context.Context, *SentMessageCodeRequest) (*SentMessageCodeResponse, error)
	//校验验证码
	CheckMessageCode(context.Context, *CheckMessageCodeRequest) (*CheckMessageCodeResponse, error)
}

// UnimplementedMessageServer should be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (UnimplementedMessageServer) Message(context.Context, *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Message not implemented")
}
func (UnimplementedMessageServer) SentMessageCode(context.Context, *SentMessageCodeRequest) (*SentMessageCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SentMessageCode not implemented")
}
func (UnimplementedMessageServer) CheckMessageCode(context.Context, *CheckMessageCodeRequest) (*CheckMessageCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckMessageCode not implemented")
}

// UnsafeMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServer will
// result in compilation errors.
type UnsafeMessageServer interface {
	mustEmbedUnimplementedMessageServer()
}

func RegisterMessageServer(s grpc.ServiceRegistrar, srv MessageServer) {
	s.RegisterService(&Message_ServiceDesc, srv)
}

func _Message_Message_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).Message(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mes.message/message",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).Message(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_SentMessageCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentMessageCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SentMessageCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mes.message/SentMessageCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SentMessageCode(ctx, req.(*SentMessageCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_CheckMessageCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckMessageCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).CheckMessageCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mes.message/CheckMessageCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).CheckMessageCode(ctx, req.(*CheckMessageCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Message_ServiceDesc is the grpc.ServiceDesc for Message service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Message_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mes.message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "message",
			Handler:    _Message_Message_Handler,
		},
		{
			MethodName: "SentMessageCode",
			Handler:    _Message_SentMessageCode_Handler,
		},
		{
			MethodName: "CheckMessageCode",
			Handler:    _Message_CheckMessageCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
