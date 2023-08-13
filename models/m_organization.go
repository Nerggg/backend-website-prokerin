package models

import (
	"time"
)

type Organization struct {
	ID          string     `gorm:"primarykey;size:36;default:gen_random_uuid()" json:"id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"update_at"`
	Status      int8       `gorm:"not null" json:"status"`
	Name        string     `gorm:"not null" json:"name"`
	Logo        string     `gorm:"not null" json:"logo"`
	Description string     `gorm:"" json:"description"`
}
