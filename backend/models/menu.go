package models

import "time"

type MenuModel struct {
	ID            uint   `gorm:"primaryKey"`
	MenuName      string `json:"menu_name"`
	Pid           uint   `json:"pid"`
	Path          string `json:"path"`
	Identifying   string `json:"identifying"`
	PermissionIds string `json:"permission_ids"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

func (m MenuModel) TableName() string {
	return "backend_menus"
}
