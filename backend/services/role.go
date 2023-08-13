package services

import (
	"errors"
	"github.com/systoms/wegin/backend/forms"
	"github.com/systoms/wegin/backend/models"
	"github.com/systoms/wegin/utils"
	"gorm.io/gorm"
)

type RoleService struct {
	db *gorm.DB
}

func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{db: db}
}

func (ctx RoleService) GetList(page int) (utils.ResultData, error) {
	var resultData utils.ResultData
	resultData.Page = page
	resultData.PageSize = 20

	var menuCount int64
	result := ctx.db.Count(&menuCount)

	var roleModels []models.RoleModel

	result = ctx.db.Offset((resultData.Page - 1) * resultData.PageSize).Limit(resultData.PageSize).Find(&roleModels)
	if result.Error != nil {
		return resultData, errors.New("查询失败")
	}

	resultData.Total = int(menuCount)
	resultData.List = result
	resultData.CalcPageCount()
	return resultData, nil
}

func (ctx RoleService) Create(roleCreateForm forms.RoleCreateForm) (models.RoleModel, error) {
	var roleModel models.RoleModel
	roleModel.RoleName = roleCreateForm.RoleName
	roleModel.MenuIds = roleCreateForm.MenuIds

	result := ctx.db.Create(&roleModel)
	if result.Error != nil {
		return roleModel, errors.New("添加失败")
	}

	return roleModel, nil
}

func (ctx RoleService) Update(roleUpdateForm forms.RoleUpdateForm) (models.RoleModel, error) {
	var roleModel models.RoleModel
	roleModel.ID = roleUpdateForm.ID
	roleModel.RoleName = roleUpdateForm.RoleName
	roleModel.MenuIds = roleUpdateForm.MenuIds

	result := ctx.db.Save(&roleModel)
	if result.Error != nil {
		return roleModel, errors.New("更新失败")
	}

	return roleModel, nil
}

func (ctx RoleService) GetOneById(id int) (models.RoleModel, error) {
	var roleModel models.RoleModel

	result := ctx.db.Where("id = ?", id).First(&roleModel)
	if result.Error != nil {
		return roleModel, errors.New("角色不存在")
	}

	return roleModel, nil
}
func (ctx RoleService) Delete(roleModel models.RoleModel) error {

	result := ctx.db.Where("id = ?", roleModel.ID).Delete(&roleModel)
	if result.Error != nil {
		return errors.New("删除失败")
	}

	return nil
}
