package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint64      `gorm:"primary key;autoIncrement" json:"id"`
	Full_Name  string      `gorm:"type:varchar(255)" json:"full_name" validate:"required,max=30"`
	Phone      uint64      `json:"phone" validate:"required"`
	Email      string      `gorm:"uniqueIndex;type:varchar(255)" json:"email" validate:"email,required"`
	Password   string      `gorm:"->;<-;not null" json:"-" validate:"required"`
	Token      string      `json:"token,omitempty"`
	Created_at time.Time   `json:"-"`
	Updated_at time.Time   `json:"-"`
	Deleted_at time.Time   `json:"-"`
	Postingans []Postingan `gorm:"foreignKey:User_ID" json:"post"`
	Post_Types *PostType
}
