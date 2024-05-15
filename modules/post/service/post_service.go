package service

import (
	"context"
	"log"
	"time"
	"tracerstudy-post-service/common/config"
	"tracerstudy-post-service/common/errors"
	"tracerstudy-post-service/common/utils"
	"tracerstudy-post-service/modules/post/entity"
	"tracerstudy-post-service/modules/post/repository"
)

type PostService struct {
	cfg            config.Config
	postRepository repository.PostRepositoryUseCase
}

func NewPostService(cfg config.Config, postRepository repository.PostRepositoryUseCase) *PostService {
	return &PostService{
		cfg:            cfg,
		postRepository: postRepository,
	}
}

type PostServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Post, error)
	FindById(ctx context.Context, id uint64) (*entity.Post, error)
	Create(ctx context.Context, title, content, mainImagePath, mainImageCaption, tipe string, isFeatured uint32, createdBy, tags string) (*entity.Post, error)
	Update(ctx context.Context, id uint64, fields *entity.Post) (*entity.Post, error)
	Delete(ctx context.Context, id uint64) error
	IncrementVisitor(ctx context.Context, id uint64) (*entity.Post, error)
}

func (svc *PostService) FindAll(ctx context.Context, req any) ([]*entity.Post, error) {
	res, err := svc.postRepository.FindAll(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostService - FindAll] Error while find all post:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *PostService) FindById(ctx context.Context, id uint64) (*entity.Post, error) {
	res, err := svc.postRepository.FindById(ctx, id)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostService - FindById] Error while find post by id:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *PostService) Create(ctx context.Context, title, content, mainImagePath, mainImageCaption, tipe string, isFeatured uint32, createdBy, tags string) (*entity.Post, error) {
	post := &entity.Post{
		Title:        title,
		Slug:         utils.GenerateSlug(title),
		Content:      content,
		ImagePath:    mainImagePath,
		ImageCaption: mainImageCaption,
		Type:         tipe,
		IsFeatured:   isFeatured,
		Visitors:     0,
		CreatedBy:    createdBy,
		UpdatedBy:    createdBy,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Tags:         tags,
	}

	res, err := svc.postRepository.Create(ctx, post)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostService - Create] Error while create post:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *PostService) Update(ctx context.Context, id uint64, fields *entity.Post) (*entity.Post, error) {
	post, err := svc.postRepository.FindById(ctx, id)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostService - Update] Error while find post by id:", parseError.Message)
		return nil, err
	}

	updatedMap := make(map[string]interface{})

	utils.AddItemToMap(updatedMap, "title", fields.Title)
	if fields.Title != "" && fields.Title != post.Title {
		utils.AddItemToMap(updatedMap, "slug", utils.GenerateSlug(fields.Title))
	}
	utils.AddItemToMap(updatedMap, "content", fields.Content)
	utils.AddItemToMap(updatedMap, "image_path", fields.ImagePath)
	utils.AddItemToMap(updatedMap, "image_caption", fields.ImageCaption)
	utils.AddItemToMap(updatedMap, "type", fields.Type)
	utils.AddItemToMap(updatedMap, "is_featured", fields.IsFeatured)
	utils.AddItemToMap(updatedMap, "updated_by", fields.UpdatedBy)
	utils.AddItemToMap(updatedMap, "tags", fields.Tags)

	res, err := svc.postRepository.Update(ctx, post, updatedMap)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostService - Update] Error while update post:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *PostService) Delete(ctx context.Context, id uint64) error {
	err := svc.postRepository.Delete(ctx, id)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostService - Delete] Error while delete post:", parseError.Message)
		return err
	}

	return nil
}

func (svc *PostService) IncrementVisitor(ctx context.Context, id uint64) (*entity.Post, error) {
	post, err := svc.postRepository.FindById(ctx, id)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostService - IncrementVisitor] Error while find post by id:", parseError.Message)
		return nil, err
	}

	updatedMap := make(map[string]interface{})
	utils.AddItemToMap(updatedMap, "visitors", post.Visitors+1)

	post, err = svc.postRepository.Update(ctx, post, updatedMap)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostService - IncrementVisitor] Error while increment visitor:", parseError.Message)
		return nil, err
	}

	return post, nil
}
