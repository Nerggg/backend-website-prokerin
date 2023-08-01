package models

import "time"

type Comment struct {
	ID          string     `gorm:"primarykey;size:36;default:gen_random_uuid()" json:"id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"update_at"`
	Status      int8       `gorm:"not null" json:"status"`
	IdCreator   string     `gorm:"not null" json:"id_creator"`
	IdProker    string     `gorm:"not null" json:"id_proker"`
	Description string     `gorm:"" json:"description"`
	Like        int64      `json:"like"`
}
