package dto

// RegisterDTO is uded when client post form /register url
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:1"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" from:"password" binding:"required" validate:"min:6"`
}
