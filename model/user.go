package model

type UserProps struct {
	BaseProps
	Username string `json:"username" gorm:"unique_index" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" gorm:"unique_index" binding:"required"`
	Avatar   string `json:"avatar"`
	Salt     string `json:"salt"`
}

type UserUpdate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}
