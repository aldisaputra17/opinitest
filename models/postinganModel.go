package models

type Postingan struct {
	Postingan_ID uint64 `gorm:"primaryKey" json:"postingan_id"`
	Description  string `gorm:"type:text" json:"description"`
	User_ID      uint64 `json:"user_id"`
	// Post_Types   *[]PostType
}
