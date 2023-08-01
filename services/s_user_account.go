package services

import "gorm.io/gorm"

type UserAccountService struct {
	db *gorm.DB
}

func NewUserAccountService(db *gorm.DB) *UserAccountService {
	return &UserAccountService{
		db: db,
	}
}
