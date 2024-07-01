package service

import (
	"context"
	"log"
	"os"
	"tracerstudy-post-service/common/config"
)

type ImageService struct {
	cfg config.Config
}

func NewImageService(cfg config.Config) *ImageService {
	return &ImageService{
		cfg: cfg,
	}
}

type ImageServiceUseCase interface {
	UploadImage(ctx context.Context, fileName string, image []byte) (string, error)
	DeleteImage(ctx context.Context, image string) error
}

func (svc *ImageService) UploadImage(ctx context.Context, fileName string, image []byte) (string, error) {
	err := os.WriteFile(svc.cfg.StoragePath+fileName, image, 0644)
	if err != nil {
		log.Println("ERROR: [ImageService - UploadImage] Error while upload image:", err)
		return "", err
	}

	return svc.cfg.PublicStoragePath + fileName, nil
}

func (svc *ImageService) DeleteImage(ctx context.Context, image string) error {
	err := os.Remove(image)
	if err != nil {
		log.Println("ERROR: [ImageService - DeleteImage] Error while delete image:", err)
		return err
	}

	return nil
}
