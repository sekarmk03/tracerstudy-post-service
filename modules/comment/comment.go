package comment

import (
	"tracerstudy-post-service/common/config"
	"tracerstudy-post-service/modules/comment/builder"
	"tracerstudy-post-service/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	comment := builder.BuildCommentHandler(cfg, db, grpcConn)
	pb.RegisterCommentServiceServer(server, comment)
}
