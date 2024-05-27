package entity

import (
	"time"
	"tracerstudy-post-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const (
	CommentTableName = "comments"
)

type Comment struct {
	Id        uint64         `json:"id"`
	PostId    uint64         `json:"post_id"`
	CommentId uint64         `json:"comment_id"`
	Name      string         `json:"name"`
	Content   string         `json:"content"`
	Level     uint32         `json:"level"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (c *Comment) TableName() string {
	return CommentTableName
}

func NewComment(id, postId, commentId uint64, name, content string, level uint32) *Comment {
	return &Comment{
		Id:        id,
		PostId:    postId,
		CommentId: commentId,
		Name:      name,
		Content:   content,
		Level:     level,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func ConvertEntityToProto(c *Comment) *pb.Comment {
	return &pb.Comment{
		Id:        c.Id,
		PostId:    c.PostId,
		CommentId: c.CommentId,
		Name:      c.Name,
		Content:   c.Content,
		Level:     c.Level,
		CreatedAt: timestamppb.New(c.CreatedAt),
		UpdatedAt: timestamppb.New(c.UpdatedAt),
	}
}
