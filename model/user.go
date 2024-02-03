package model

type User struct {
	BaseInfo
	Username string `json:"username" gorm:"unique_index" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" gorm:"unique_index" binding:"required" validate:"email"`
	Avatar   string `json:"avatar"`
	Salt     string `json:"salt"`
}

type UpdateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

type GetUserParams struct {
	Id string `uri:"id" validate:"required"`
}
