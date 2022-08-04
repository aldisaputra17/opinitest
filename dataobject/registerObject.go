package dataobject

type RegisterObject struct {
	Full_Name string `json:"full_name" form:"name" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required,email" `
	Phone     uint64 `json:"phone" form:"phone" binding:"required,phone" `
	Password  string `json:"password" form:"password" binding:"required"`
}
