package models

import (
	"time"
)

type Proker struct {
	ID               string       `gorm:"primarykey;size:36;default:gen_random_uuid()" json:"id"`
	CreatedAt        *time.Time   `json:"created_at"`
	UpdatedAt        *time.Time   `json:"update_at"`
	Status           int8         `gorm:"not null" json:"status"`
	IdCreator        string       `gorm:"not null" json:"id_creator"`
	Name             string       `gorm:"not null" json:"name"`
	Image            string       `gorm:"not null" json:"image"` //in static url
	Description      string       `gorm:"" json:"description"`
	ShortDescription string       `gorm:"" json:"short_description"`
	TimeLineImage    string       `json:"time_line_image"`
	Like             int64        `json:"like"`
	OrganizationId   string       `json:"organization_id"`
	Organization     Organization `json:"organization" gorm:"foreignKey:organization_id"`
}
