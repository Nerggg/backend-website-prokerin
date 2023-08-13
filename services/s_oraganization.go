package services

import "gorm.io/gorm"

type OrganizationService struct {
	db *gorm.DB
}

func NewOrganizationService(db *gorm.DB) *OrganizationService {
	return &OrganizationService{
		db: db,
	}
}
