package dataobject

type RegisterObject struct {
	Full_Name string `json:"full_name" form:"name" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required,email" `
	Phone     string `json:"phone" binding:"required" `
	Password  string `json:"password" form:"password" binding:"required"`
}
