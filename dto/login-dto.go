package dto

// LoginDTO is a model that used by client when post from /login url
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" from:"password" binding:"required" validate:"min:6"`
}
