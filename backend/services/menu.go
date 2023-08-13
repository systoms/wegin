package services

import (
	"errors"
	"github.com/systoms/wegin/backend/forms"
	"github.com/systoms/wegin/backend/models"
	"github.com/systoms/wegin/utils"
	"gorm.io/gorm"
)

type MenuService struct {
	db *gorm.DB
}

func NewMenuService(db *gorm.DB) *MenuService {
	return &MenuService{db: db}
}

func (ctx MenuService) GetList(page int) (utils.ResultData, error) {
	var resultData utils.ResultData
	resultData.Page = page
	resultData.PageSize = 20

	var menuCount int64
	result := ctx.db.Count(&menuCount)

	var menuModels []models.MenuModel

	result = ctx.db.Offset((resultData.Page - 1) * resultData.PageSize).Limit(resultData.PageSize).Find(&menuModels)
	if result.Error != nil {
		return resultData, errors.New("查询失败")
	}

	resultData.Total = int(menuCount)
	resultData.List = result
	resultData.CalcPageCount()
	return resultData, nil
}

func (ctx MenuService) Create(menuCreateForm forms.MenuCreateForm) (models.MenuModel, error) {
	var menuModel models.MenuModel
	menuModel.MenuName = menuCreateForm.MenuName
	menuModel.Pid = menuCreateForm.Pid
	menuModel.Path = menuCreateForm.Path
	menuModel.Identifying = menuCreateForm.Identifying
	menuModel.PermissionIds = menuCreateForm.PermissionIds

	result := ctx.db.Create(&menuModel)
	if result.Error != nil {
		return menuModel, errors.New("添加失败")
	}

	return menuModel, nil
}

func (ctx MenuService) Update(menuUpdateForm forms.MenuUpdateForm) (models.MenuModel, error) {
	var menuModel models.MenuModel
	menuModel.ID = menuUpdateForm.ID
	menuModel.MenuName = menuUpdateForm.MenuName
	menuModel.Pid = menuUpdateForm.Pid
	menuModel.Path = menuUpdateForm.Path
	menuModel.Identifying = menuUpdateForm.Identifying
	menuModel.PermissionIds = menuUpdateForm.PermissionIds

	result := ctx.db.Save(&menuModel)
	if result.Error != nil {
		return menuModel, errors.New("更新失败")
	}

	return menuModel, nil
}

func (ctx MenuService) GetOneById(id int) (models.MenuModel, error) {
	var menuModel models.MenuModel

	result := ctx.db.Where("id = ?", id).First(&menuModel)
	if result.Error != nil {
		return menuModel, errors.New("菜单不存在")
	}

	return menuModel, nil
}
func (ctx MenuService) Delete(menuModel models.MenuModel) error {

	result := ctx.db.Where("id = ?", menuModel.ID).Delete(&menuModel)
	if result.Error != nil {
		return errors.New("删除失败")
	}

	return nil
}

type MenuNode struct {
	ID            uint        `gorm:"primaryKey"`
	MenuName      string      `json:"menu_name"`
	Pid           uint        `json:"pid"`
	Path          string      `json:"path"`
	Identifying   string      `json:"identifying"`
	PermissionIds string      `json:"permission_ids"`
	Children      []*MenuNode `json:"children"`
}

func (ctx MenuService) GenerateMenuTree(menuItems []models.MenuModel) []*MenuNode {
	items := make(map[uint]*MenuNode)
	var tree []*MenuNode

	for _, item := range menuItems {
		node := &MenuNode{
			ID:            item.ID,
			MenuName:      item.MenuName,
			Pid:           item.Pid,
			Path:          item.Path,
			Identifying:   item.Identifying,
			PermissionIds: item.PermissionIds,
		}
		items[node.ID] = node
		if _, exists := items[node.Pid]; exists {
			items[node.Pid].Children = append(items[node.Pid].Children, node)
		} else {
			tree = append(tree, node)
		}
	}

	return tree
}

func (ctx MenuService) Tree() ([]*MenuNode, error) {
	var menuModels []models.MenuModel

	result := ctx.db.Select("id,menu_name,path,pid,identifying,permission_ids").Find(&menuModels)
	if result.Error != nil {
		return nil, errors.New("查询失败")
	}

	menuNodes := ctx.GenerateMenuTree(menuModels)
	return menuNodes, nil
}
