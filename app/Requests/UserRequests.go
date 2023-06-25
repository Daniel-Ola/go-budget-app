package Requests

type CreateUserRequest struct {
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	UserName    string `json:"userName" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
