package models

import "time"

type Permission struct {
	ID             uint   `gorm:"primaryKey"`
	PermissionName string `json:"permission_name"`
	Path           string `json:"path"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

func (p Permission) TableName() string {
	return "backend_permissions"
}
