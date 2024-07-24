package system

import api "aiServer/api/v1"

type RouterGroup struct {
	JwtRouter
	BaseRouter
	InitRouter
	MenuRouter
	UserRouter
	CasbinRouter
	AuthorityRouter
	OperationRecordRouter
}

var (
	dbApi              = api.ApiGroupApp.SystemApiGroup.DBApi
	jwtApi             = api.ApiGroupApp.SystemApiGroup.JwtApi
	baseApi            = api.ApiGroupApp.SystemApiGroup.BaseApi
	casbinApi          = api.ApiGroupApp.SystemApiGroup.CasbinApi
	authorityApi       = api.ApiGroupApp.SystemApiGroup.AuthorityApi
	authorityMenuApi   = api.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	operationRecordApi = api.ApiGroupApp.SystemApiGroup.OperationRecordApi
)
