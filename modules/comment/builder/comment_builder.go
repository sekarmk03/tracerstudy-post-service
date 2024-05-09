package builder

import (
	"tracerstudy-post-service/common/config"
	"tracerstudy-post-service/modules/comment/handler"
	"tracerstudy-post-service/modules/comment/repository"
	"tracerstudy-post-service/modules/comment/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildCommentHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.CommentHandler {
	commentRepo := repository.NewCommentRepository(db)
	commentSvc := service.NewCommentService(cfg, commentRepo)

	return handler.NewCommentHandler(cfg, commentSvc)
}
