package entity

import (
	"time"
	"tracerstudy-post-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const (
	PostTableName = "posts"
)

type Post struct {
	Id           uint64         `json:"id"`
	Title        string         `json:"title"`
	Slug         string         `json:"slug"`
	Content      string         `json:"content"`
	ImagePath    string         `json:"image_path"`
	ImageCaption string         `json:"image_caption"`
	Type         string         `json:"type"`
	IsFeatured   uint32         `json:"is_featured"`
	Visitors     uint64         `json:"visitors"`
	CreatedBy    string         `json:"created_by"`
	UpdatedBy    string         `json:"updated_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Tags         string         `json:"tags"`
}

func (p *Post) TableName() string {
	return PostTableName
}

func ConvertEntityToProto(p *Post) *pb.Post {
	return &pb.Post{
		Id:           p.Id,
		Title:        p.Title,
		Slug:         p.Slug,
		Content:      p.Content,
		ImagePath:    p.ImagePath,
		ImageCaption: p.ImageCaption,
		Type:         p.Type,
		IsFeatured:   p.IsFeatured,
		Visitors:     p.Visitors,
		CreatedBy:    p.CreatedBy,
		UpdatedBy:    p.UpdatedBy,
		CreatedAt:    timestamppb.New(p.CreatedAt),
		UpdatedAt:    timestamppb.New(p.UpdatedAt),
		Tags:         p.Tags,
	}
}

func ConvertProtoToEntity(p *pb.Post) *Post {
	return &Post{
		Id:           p.Id,
		Title:        p.Title,
		Slug:         p.Slug,
		Content:      p.Content,
		ImagePath:    p.ImagePath,
		ImageCaption: p.ImageCaption,
		Type:         p.Type,
		IsFeatured:   p.IsFeatured,
		Visitors:     p.Visitors,
		CreatedBy:    p.CreatedBy,
		UpdatedBy:    p.UpdatedBy,
		CreatedAt:    p.CreatedAt.AsTime(),
		UpdatedAt:    p.UpdatedAt.AsTime(),
		Tags:         p.Tags,
	}
}
