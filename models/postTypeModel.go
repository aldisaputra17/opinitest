package models

type postType string

const (
	Artikel postType = "Artikel"
	Idea    postType = "Idea"
)

type PostType struct {
	Type_ID    uint64   `gorm:"primaryKey" json:"type_id"`
	Post_Types postType `sql:"post_type"`
	Post_ID    uint64
	User_ID    uint64
}
