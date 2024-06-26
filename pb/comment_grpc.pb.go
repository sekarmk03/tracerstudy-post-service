// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: comment.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CommentService_GetAllComments_FullMethodName      = "/tracer_study_grpc.CommentService/GetAllComments"
	CommentService_GetCommentsByPostId_FullMethodName = "/tracer_study_grpc.CommentService/GetCommentsByPostId"
	CommentService_GetCommentById_FullMethodName      = "/tracer_study_grpc.CommentService/GetCommentById"
	CommentService_CreateComment_FullMethodName       = "/tracer_study_grpc.CommentService/CreateComment"
	CommentService_ReplyComment_FullMethodName        = "/tracer_study_grpc.CommentService/ReplyComment"
	CommentService_DeleteComment_FullMethodName       = "/tracer_study_grpc.CommentService/DeleteComment"
)

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentServiceClient interface {
	GetAllComments(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllCommentsResponse, error)
	GetCommentsByPostId(ctx context.Context, in *GetCommentsByPostIdRequest, opts ...grpc.CallOption) (*GetAllCommentsResponse, error)
	GetCommentById(ctx context.Context, in *GetCommentByIdRequest, opts ...grpc.CallOption) (*GetCommentResponse, error)
	CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*GetCommentResponse, error)
	ReplyComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*GetCommentResponse, error)
	DeleteComment(ctx context.Context, in *GetCommentByIdRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
}

type commentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentServiceClient(cc grpc.ClientConnInterface) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) GetAllComments(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllCommentsResponse, error) {
	out := new(GetAllCommentsResponse)
	err := c.cc.Invoke(ctx, CommentService_GetAllComments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetCommentsByPostId(ctx context.Context, in *GetCommentsByPostIdRequest, opts ...grpc.CallOption) (*GetAllCommentsResponse, error) {
	out := new(GetAllCommentsResponse)
	err := c.cc.Invoke(ctx, CommentService_GetCommentsByPostId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetCommentById(ctx context.Context, in *GetCommentByIdRequest, opts ...grpc.CallOption) (*GetCommentResponse, error) {
	out := new(GetCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_GetCommentById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*GetCommentResponse, error) {
	out := new(GetCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_CreateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) ReplyComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*GetCommentResponse, error) {
	out := new(GetCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_ReplyComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) DeleteComment(ctx context.Context, in *GetCommentByIdRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	out := new(DeleteCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_DeleteComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
// All implementations must embed UnimplementedCommentServiceServer
// for forward compatibility
type CommentServiceServer interface {
	GetAllComments(context.Context, *emptypb.Empty) (*GetAllCommentsResponse, error)
	GetCommentsByPostId(context.Context, *GetCommentsByPostIdRequest) (*GetAllCommentsResponse, error)
	GetCommentById(context.Context, *GetCommentByIdRequest) (*GetCommentResponse, error)
	CreateComment(context.Context, *Comment) (*GetCommentResponse, error)
	ReplyComment(context.Context, *Comment) (*GetCommentResponse, error)
	DeleteComment(context.Context, *GetCommentByIdRequest) (*DeleteCommentResponse, error)
	mustEmbedUnimplementedCommentServiceServer()
}

// UnimplementedCommentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServiceServer struct {
}

func (UnimplementedCommentServiceServer) GetAllComments(context.Context, *emptypb.Empty) (*GetAllCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllComments not implemented")
}
func (UnimplementedCommentServiceServer) GetCommentsByPostId(context.Context, *GetCommentsByPostIdRequest) (*GetAllCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsByPostId not implemented")
}
func (UnimplementedCommentServiceServer) GetCommentById(context.Context, *GetCommentByIdRequest) (*GetCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentById not implemented")
}
func (UnimplementedCommentServiceServer) CreateComment(context.Context, *Comment) (*GetCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCommentServiceServer) ReplyComment(context.Context, *Comment) (*GetCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplyComment not implemented")
}
func (UnimplementedCommentServiceServer) DeleteComment(context.Context, *GetCommentByIdRequest) (*DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedCommentServiceServer) mustEmbedUnimplementedCommentServiceServer() {}

// UnsafeCommentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServiceServer will
// result in compilation errors.
type UnsafeCommentServiceServer interface {
	mustEmbedUnimplementedCommentServiceServer()
}

func RegisterCommentServiceServer(s grpc.ServiceRegistrar, srv CommentServiceServer) {
	s.RegisterService(&CommentService_ServiceDesc, srv)
}

func _CommentService_GetAllComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetAllComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetAllComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetAllComments(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetCommentsByPostId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsByPostIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetCommentsByPostId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetCommentsByPostId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetCommentsByPostId(ctx, req.(*GetCommentsByPostIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetCommentById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetCommentById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetCommentById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetCommentById(ctx, req.(*GetCommentByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).CreateComment(ctx, req.(*Comment))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_ReplyComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).ReplyComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_ReplyComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).ReplyComment(ctx, req.(*Comment))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).DeleteComment(ctx, req.(*GetCommentByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentService_ServiceDesc is the grpc.ServiceDesc for CommentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tracer_study_grpc.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllComments",
			Handler:    _CommentService_GetAllComments_Handler,
		},
		{
			MethodName: "GetCommentsByPostId",
			Handler:    _CommentService_GetCommentsByPostId_Handler,
		},
		{
			MethodName: "GetCommentById",
			Handler:    _CommentService_GetCommentById_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _CommentService_CreateComment_Handler,
		},
		{
			MethodName: "ReplyComment",
			Handler:    _CommentService_ReplyComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _CommentService_DeleteComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
