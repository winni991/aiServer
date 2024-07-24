package initialize

import (
	_ "aiServer/source/example"
	_ "aiServer/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
