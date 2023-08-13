package models

import "time"

type UserModel struct {
	ID          uint   `gorm:"primaryKey"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
	Nickname    string `json:"nickname"`
	RoleIds     string `json:"role_ids"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func (u UserModel) TableName() string {
	return "backend_users"
}
