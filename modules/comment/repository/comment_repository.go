package repository

import (
	"context"
	"errors"
	"log"
	"tracerstudy-post-service/modules/comment/entity"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

type CommentRepositoryUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Comment, error)
	FindCommentsByPostId(ctx context.Context, postId uint64) ([]*entity.Comment, error)
	FindById(ctx context.Context, id uint64) (*entity.Comment, error)
	Create(ctx context.Context, req *entity.Comment) (*entity.Comment, error)
	Delete(ctx context.Context, id uint64) error
}

func (c *CommentRepository) FindAll(ctx context.Context, req any) ([]*entity.Comment, error) {
	ctxSpan, span := trace.StartSpan(ctx, "CommentRepository - FindAll")
	defer span.End()

	var comment []*entity.Comment
	if err := c.db.Debug().WithContext(ctxSpan).Order("created_at desc").Find(&comment).Error; err != nil {
		log.Println("ERROR: [CommentRepository - FindAll] Internal server error:", err)
		return nil, err
	}

	return comment, nil
}

func (c *CommentRepository) FindCommentsByPostId(ctx context.Context, postId uint64) ([]*entity.Comment, error) {
	ctxSpan, span := trace.StartSpan(ctx, "CommentRepository - FindCommentsByPostId")
	defer span.End()

	var comment []*entity.Comment
	if err := c.db.Debug().WithContext(ctxSpan).Where("post_id = ?", postId).Order("created_at desc").Find(&comment).Error; err != nil {
		log.Println("ERROR: [CommentRepository - FindCommentsByPostId] Internal server error:", err)
		return nil, err
	}

	return comment, nil
}

func (c *CommentRepository) FindById(ctx context.Context, id uint64) (*entity.Comment, error) {
	ctxSpan, span := trace.StartSpan(ctx, "CommentRepository - FindById")
	defer span.End()

	var comment entity.Comment
	if err := c.db.Debug().WithContext(ctxSpan).Where("id = ?", id).First(&comment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [CommentRepository - FindById] Record not found for id", id)
			return nil, status.Errorf(codes.NotFound, "record not found for id %d", id)
		}
		log.Println("ERROR: [CommentRepository - FindById] Internal server error:", err)
		return nil, err
	}

	return &comment, nil
}

func (c *CommentRepository) Create(ctx context.Context, req *entity.Comment) (*entity.Comment, error) {
	ctxSpan, span := trace.StartSpan(ctx, "CommentRepository - Create")
	defer span.End()

	if err := c.db.Debug().WithContext(ctxSpan).Create(req).Error; err != nil {
		log.Println("ERROR: [CommentRepository - Create] Internal server error:", err)
		return nil, err
	}

	return req, nil
}

func (c *CommentRepository) Delete(ctx context.Context, id uint64) error {
	ctxSpan, span := trace.StartSpan(ctx, "CommentRepository - Delete")
	defer span.End()

	if err := c.db.Debug().WithContext(ctxSpan).Where("id = ?", id).Delete(&entity.Comment{}).Error; err != nil {
		log.Println("ERROR: [CommentRepository - Delete] Internal server error:", err)
		return err
	}

	return nil
}
