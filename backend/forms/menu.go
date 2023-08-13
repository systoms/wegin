package forms

type MenuCreateForm struct {
	MenuName      string `form:"menu_name" binding:"required" json:"menu_name"`
	Pid           uint   `form:"pid" binding:"required" json:"pid"`
	Path          string `form:"path" binding:"required" json:"path"`
	Identifying   string `form:"identifying" binding:"required" json:"identifying"`
	PermissionIds string `form:"permission_ids" binding:"required" json:"permission_ids"`
}

type MenuUpdateForm struct {
	ID            uint   `form:"id" binding:"required" json:"id"`
	MenuName      string `form:"menu_name" binding:"required" json:"menu_name"`
	Pid           uint   `form:"pid" binding:"required" json:"pid"`
	Path          string `form:"path" binding:"required" json:"path"`
	Identifying   string `form:"identifying" binding:"required" json:"identifying"`
	PermissionIds string `form:"permission_ids" binding:"required" json:"permission_ids"`
}
