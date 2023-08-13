package models

import "time"

type RoleModel struct {
	ID        uint   `gorm:"primaryKey"`
	RoleName  string `json:"role_name"`
	MenuIds   []uint `json:"menu_ids"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (r RoleModel) TableName() string {
	return "backend_roles"
}
