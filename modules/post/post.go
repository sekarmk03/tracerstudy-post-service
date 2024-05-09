package post

import (
	"tracerstudy-post-service/common/config"
	"tracerstudy-post-service/modules/post/builder"
	"tracerstudy-post-service/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	post := builder.BuildPostHandler(cfg, db, grpcConn)
	pb.RegisterPostServiceServer(server, post)
}
