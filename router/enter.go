package router

import (
	"aiServer/router/example"
	"aiServer/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}
