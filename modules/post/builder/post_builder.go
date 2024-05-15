package builder

import (
	"tracerstudy-post-service/common/config"
	"tracerstudy-post-service/modules/post/client"
	"tracerstudy-post-service/modules/post/handler"
	"tracerstudy-post-service/modules/post/repository"
	"tracerstudy-post-service/modules/post/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildPostHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.PostHandler {
	postRepo := repository.NewPostRepository(db)
	imageSvc := service.NewImageService(cfg)
	postSvc := service.NewPostService(cfg, postRepo)
	authSvc := client.BuildAuthServiceClient(cfg.ClientURL.Auth)

	return handler.NewPostHandler(cfg, postSvc, imageSvc, authSvc)
}
