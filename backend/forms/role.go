package forms

type RoleCreateForm struct {
	RoleName string `form:"role_name" binding:"required" json:"role_name"`
	MenuIds  []uint `form:"menu_ids" binding:"required" json:"menu_ids"`
}

type RoleUpdateForm struct {
	ID       uint   `form:"id" binding:"required" json:"id"`
	RoleName string `form:"role_name" binding:"required" json:"role_name"`
	MenuIds  []uint `form:"menu_ids" binding:"required" json:"menu_ids"`
}
