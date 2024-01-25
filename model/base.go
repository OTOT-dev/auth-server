package model

type BaseProps struct {
	ID        int64 `gorm:"primaryKey" json:"id"`
	CreatedAt int   `json:"created_at"`
	UpdatedAt int   `json:"updated_at"`
}
