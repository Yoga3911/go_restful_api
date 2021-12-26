package dto

// RegisterDTO is used when client post form /register url
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" from:"password" binding:"required"`
}
