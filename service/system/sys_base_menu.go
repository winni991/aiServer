package system

import (
	"errors"

	"aiServer/global"
	"aiServer/model/system"
	"gorm.io/gorm"
)

type BaseMenuService struct{}

//@author: [xt]
//@function: DeleteBaseMenu
//@description: 删除基础路由
//@param: id float64
//@return: err error

var BaseMenuServiceApp = new(BaseMenuService)

func (baseMenuService *BaseMenuService) DeleteBaseMenu(id int) (err error) {
	err = global.AI_DB.First(&system.SysBaseMenu{}, "parent_id = ?", id).Error
	if err != nil {
		return global.AI_DB.Transaction(func(tx *gorm.DB) error {

			err = tx.Delete(&system.SysBaseMenu{}, "id = ?", id).Error
			if err != nil {
				return err
			}

			err = tx.Delete(&system.SysAuthorityMenu{}, "sys_base_menu_id = ?", id).Error
			if err != nil {
				return err
			}
			return nil
		})
	}
	return errors.New("此菜单存在子菜单不可删除")
}

//@author: [xt]
//@function: UpdateBaseMenu
//@description: 更新路由
//@param: menu model.SysBaseMenu
//@return: err error

func (baseMenuService *BaseMenuService) UpdateBaseMenu(menu system.SysBaseMenu) (err error) {
	var oldMenu system.SysBaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
	upDateMap["close_tab"] = menu.CloseTab
	upDateMap["default_menu"] = menu.DefaultMenu
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["active_name"] = menu.ActiveName
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort

	err = global.AI_DB.Transaction(func(tx *gorm.DB) error {
		tx.Where("id = ?", menu.ID).Find(&oldMenu)
		if oldMenu.Name != menu.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
				global.AI_LOG.Debug("存在相同name修改失败")
				return errors.New("存在相同name修改失败")
			}
		}
		txErr := tx.Model(&oldMenu).Updates(upDateMap).Error
		if txErr != nil {
			global.AI_LOG.Debug(txErr.Error())
			return txErr
		}
		return nil
	})
	return err
}

//@author: [xt]
//@function: GetBaseMenuById
//@description: 返回当前选中menu
//@param: id float64
//@return: menu system.SysBaseMenu, err error

func (baseMenuService *BaseMenuService) GetBaseMenuById(id int) (menu system.SysBaseMenu, err error) {
	err = global.AI_DB.Where("id = ?", id).First(&menu).Error
	return
}
