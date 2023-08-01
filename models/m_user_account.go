package models

import (
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserAccount struct {
	ID           string         `gorm:"primarykey;size:36;default:gen_random_uuid()" json:"id"`
	CreatedAt    *time.Time     `json:"created_at"`
	UpdatedAt    *time.Time     `json:"update_at"`
	Status       int8           `gorm:"not null" json:"status"`
	Email        string         `gorm:"not null,unique" json:"email"`
	Password     string         `gorm:"not null" json:"password"`
	NickName     string         `gorm:"not null,unique" json:"nick_name"`
	Instansi     string         `gorm:"" json:"instasnsi"`
	IsAdmin      int8           `gorm:"" json:"is_admin"` // 0 super, 1 user, 2 admin
	LikedProker  datatypes.JSON `gorm:"" json:"liked_proker"`
	LikedComment datatypes.JSON `gorm:"" json:"liked_comment"`
}

func (data *UserAccount) BeforeSave(*gorm.DB) error {

	bytes, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if err != nil {
		print(err)
	}
	data.Password = string(bytes)

	temp_email := html.EscapeString(strings.TrimSpace(data.Email))
	data.Email = temp_email
	return nil
}
