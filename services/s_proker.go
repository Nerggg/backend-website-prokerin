package services

import "gorm.io/gorm"

type ProkerService struct {
	db *gorm.DB
}

func NewProkerService(db *gorm.DB) *ProkerService {
	return &ProkerService{
		db: db,
	}
}
