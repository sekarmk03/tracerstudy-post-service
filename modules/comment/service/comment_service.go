package service

import (
	"context"
	"log"
	"time"
	"tracerstudy-post-service/common/config"
	"tracerstudy-post-service/common/errors"
	"tracerstudy-post-service/modules/comment/entity"
	"tracerstudy-post-service/modules/comment/repository"
)

type CommentService struct {
	cfg               config.Config
	commentRepository repository.CommentRepositoryUseCase
}

func NewCommentService(cfg config.Config, commentRepository repository.CommentRepositoryUseCase) *CommentService {
	return &CommentService{
		cfg:               cfg,
		commentRepository: commentRepository,
	}
}

type CommentServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Comment, error)
	FindCommentsByPostId(ctx context.Context, postId uint64) ([]*entity.Comment, error)
	FindById(ctx context.Context, id uint64) (*entity.Comment, error)
	Create(ctx context.Context, postId, commentId uint64, name, content string, level uint32) (*entity.Comment, error)
	Delete(ctx context.Context, id uint64) error
}

func (svc *CommentService) FindAll(ctx context.Context, req any) ([]*entity.Comment, error) {
	res, err := svc.commentRepository.FindAll(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [CommentService - FindAll] Error while find all comment:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *CommentService) FindCommentsByPostId(ctx context.Context, postId uint64) ([]*entity.Comment, error) {
	res, err := svc.commentRepository.FindCommentsByPostId(ctx, postId)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [CommentService - FindCommentsByPostId] Error while find comments by post id:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *CommentService) FindById(ctx context.Context, id uint64) (*entity.Comment, error) {
	res, err := svc.commentRepository.FindById(ctx, id)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [CommentService - FindById] Error while find comment by id:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *CommentService) Create(ctx context.Context, postId, commentId uint64, name, content string, level uint32) (*entity.Comment, error) {
	comment := &entity.Comment{
		PostId:    postId,
		CommentId: commentId,
		Name:      name,
		Content:   content,
		Level:     level,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res, err := svc.commentRepository.Create(ctx, comment)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [CommentService - Create] Error while create comment:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *CommentService) Delete(ctx context.Context, id uint64) error {
	err := svc.commentRepository.Delete(ctx, id)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [CommentService - Delete] Error while delete comment:", parseError.Message)
		return err
	}

	return nil
}
