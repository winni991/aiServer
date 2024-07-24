package request

import (
	"aiServer/model/common/request"
	"aiServer/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
