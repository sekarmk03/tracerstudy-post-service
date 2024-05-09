package handler

import (
	"context"
	"log"
	"net/http"
	"tracerstudy-post-service/common/config"
	"tracerstudy-post-service/common/errors"
	"tracerstudy-post-service/modules/comment/entity"
	"tracerstudy-post-service/modules/comment/service"
	"tracerstudy-post-service/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type CommentHandler struct {
	pb.UnimplementedCommentServiceServer
	config     config.Config
	commentSvc service.CommentServiceUseCase
}

func NewCommentHandler(config config.Config, commentService service.CommentServiceUseCase) *CommentHandler {
	return &CommentHandler{
		config:     config,
		commentSvc: commentService,
	}
}

func (ch *CommentHandler) GetAllComments(ctx context.Context, req *emptypb.Empty) (*pb.GetAllCommentsResponse, error) {
	comments, err := ch.commentSvc.FindAll(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [CommentHandler - GetAllComments] Error while get all comments:", parseError.Message)
		return nil, err
	}

	var commentArr []*pb.Comment
	for _, c := range comments {
		commentProto := entity.ConvertEntityToProto(c)
		commentArr = append(commentArr, commentProto)
	}

	return &pb.GetAllCommentsResponse{
		Code:    uint32(http.StatusOK),
		Message: "get all comments success",
		Data:    commentArr,
	}, nil
}

func (ch *CommentHandler) GetCommentsByPostId(ctx context.Context, req *pb.GetCommentsByPostIdRequest) (*pb.GetAllCommentsResponse, error) {
	comments, err := ch.commentSvc.FindCommentsByPostId(ctx, req.GetPostId())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [CommentHandler - GetCommentsByPostId] Error while get comments by post id:", parseError.Message)
		return nil, err
	}

	var commentArr []*pb.Comment
	for _, c := range comments {
		commentProto := entity.ConvertEntityToProto(c)
		commentArr = append(commentArr, commentProto)
	}

	return &pb.GetAllCommentsResponse{
		Code:    uint32(http.StatusOK),
		Message: "get comments by post id success",
		Data:    commentArr,
	}, nil
}

func (ch *CommentHandler) GetCommentById(ctx context.Context, req *pb.GetCommentByIdRequest) (*pb.GetCommentResponse, error) {
	comment, err := ch.commentSvc.FindById(ctx, req.GetId())
	if err != nil {
		if comment == nil {
			log.Println("WARNING: [CommentHandler - GetCommentById] Resource comment not found for id:", req.GetId())
			return nil, err
		}
		parseError := errors.ParseError(err)
		log.Println("ERROR: [CommentHandler - GetCommentById] Internal server error:", parseError.Message)
		return nil, err
	}

	commentProto := entity.ConvertEntityToProto(comment)
	return &pb.GetCommentResponse{
		Code:    uint32(http.StatusOK),
		Message: "get comment success",
		Data:    commentProto,
	}, nil
}

func (ch *CommentHandler) CreateComment(ctx context.Context, req *pb.Comment) (*pb.GetCommentResponse, error) {
	comment, err := ch.commentSvc.Create(ctx, req.GetPostId(), 0, req.GetName(), req.GetContent(), 0)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [CommentHandler - CreateComment] Error while create comment:", parseError.Message)
		return nil, err
	}

	commentProto := entity.ConvertEntityToProto(comment)
	return &pb.GetCommentResponse{
		Code:    uint32(http.StatusCreated),
		Message: "create comment success",
		Data:    commentProto,
	}, nil
}

func (ch *CommentHandler) ReplyComment(ctx context.Context, req *pb.Comment) (*pb.GetCommentResponse, error) {
	comment, err := ch.commentSvc.Create(ctx, req.GetPostId(), req.GetCommentId(), req.GetName(), req.GetContent(), 1)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [CommentHandler - ReplyComment] Error while reply comment:", parseError.Message)
		return nil, err
	}

	commentProto := entity.ConvertEntityToProto(comment)
	return &pb.GetCommentResponse{
		Code:    uint32(http.StatusCreated),
		Message: "reply comment success",
		Data:    commentProto,
	}, nil
}

func (ch *CommentHandler) DeleteComment(ctx context.Context, req *pb.GetCommentByIdRequest) (*pb.DeleteCommentResponse, error) {
	err := ch.commentSvc.Delete(ctx, req.GetId())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [CommentHandler - DeleteComment] Error while delete comment:", parseError.Message)
		return nil, err
	}

	return &pb.DeleteCommentResponse{
		Code:    uint32(http.StatusOK),
		Message: "delete comment success",
	}, nil
}
