package models

type User struct {
	ID         uint64       `gorm:"primary_key:auto_increment" json:"id"`
	Full_Name  string       `gorm:"type:varchar(255)" json:"full_name"`
	Phone      string       `gorm:"type:varchar(255)" json:"phone"`
	Email      string       `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password   string       `gorm:"->;<-;not null" json:"-" validate:"required"`
	Token      string       `json:"token,omitempty"`
	Postingans *[]Postingan `json:"post,omitempty"`
}
