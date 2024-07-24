package response

import "aiServer/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
