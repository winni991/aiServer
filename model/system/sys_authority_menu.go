package system

type SysMenu struct {
	SysBaseMenu
	MenuId      uint      `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId uint      `json:"-" gorm:"comment:角色ID"`
	Children    []SysMenu `json:"children" gorm:"-"`
}

type SysAuthorityMenu struct {
	MenuId      string `json:"menuId" gorm:"comment:菜单ID;column:sys_base_menu_id"`
	AuthorityId string `json:"-" gorm:"comment:角色ID;column:sys_authority_authority_id"`
}

func (s SysAuthorityMenu) TableName() string {
	return "sys_authority_menus"
}
