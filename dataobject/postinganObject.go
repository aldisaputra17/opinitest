package dataobject

type PostCreateObj struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
	TypeID      uint64 `json:"type_id,omitempty" form:"type_id,omitempty"`
	// PostType    `json:"post_type" form:"post_type"`
}

type PostUpdatedObj struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id" form:"user_id,omitempty"`
	// PostType    `json:"post_type" form:"post_type"`
}

type PostDeletedObj struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty" binding:"required"`
	// PostType    `json:"post_type" form:"post_type"`
}
