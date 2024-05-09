package repository

import (
	"context"
	"errors"
	"log"
	"time"
	"tracerstudy-post-service/modules/post/entity"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

type PostRepositoryUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Post, error)
	FindById(ctx context.Context, id uint64) (*entity.Post, error)
	Create(ctx context.Context, req *entity.Post) (*entity.Post, error)
	Update(ctx context.Context, post *entity.Post, updatedFields map[string]interface{}) (*entity.Post, error)
	Delete(ctx context.Context, id uint64) error
}

func (p *PostRepository) FindAll(ctx context.Context, req any) ([]*entity.Post, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PostRepository - FindAll")
	defer span.End()

	var post []*entity.Post
	if err := p.db.Debug().WithContext(ctxSpan).Order("created_at desc").Find(&post).Error; err != nil {
		log.Println("ERROR: [PostRepository - FindAll] Internal server error:", err)
		return nil, err
	}

	return post, nil
}

func (p *PostRepository) FindById(ctx context.Context, id uint64) (*entity.Post, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PostRepository - FindById")
	defer span.End()

	var post entity.Post
	if err := p.db.Debug().WithContext(ctxSpan).Where("id = ?", id).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [PostRepository - FindById] Record not found for id", id)
			return nil, status.Errorf(codes.NotFound, "record not found for id %d", id)
		}
		log.Println("ERROR: [PostRepository - FindById] Internal server error:", err)
		return nil, err
	}

	return &post, nil
}

func (p *PostRepository) Create(ctx context.Context, req *entity.Post) (*entity.Post, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PostRepository - Create")
	defer span.End()

	if err := p.db.Debug().WithContext(ctxSpan).Create(req).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Println("WARNING: [PostRepository - Create] Record already exists")
			return nil, status.Errorf(codes.AlreadyExists, "record already exists")
		}
		log.Println("ERROR: [PostRepository - Create] Internal server error:", err)
		return nil, err
	}

	return req, nil
}

func (p *PostRepository) Update(ctx context.Context, post *entity.Post, updatedFields map[string]interface{}) (*entity.Post, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PostRepository - Update")
	defer span.End()

	updatedFields["updated_at"] = time.Now()
	if err := p.db.Debug().WithContext(ctxSpan).Model(&post).Updates(updatedFields).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Println("WARNING: [PostRepository - Update] Record already exists")
			return nil, status.Errorf(codes.AlreadyExists, "record already exists")
		}
		log.Println("ERROR: [PostRepository - Update] Internal server error:", err)
		return nil, err
	}

	return post, nil
}

func (p *PostRepository) Delete(ctx context.Context, id uint64) error {
	ctxSpan, span := trace.StartSpan(ctx, "PostRepository - Delete")
	defer span.End()

	if err := p.db.Debug().WithContext(ctxSpan).Where("id = ?", id).Delete(&entity.Post{}).Error; err != nil {
		log.Println("ERROR: [PostRepository - Delete] Internal server error:", err)
		return err
	}

	return nil
}
