package services

import "gorm.io/gorm"

type CommentService struct {
	db *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{
		db: db,
	}
}
