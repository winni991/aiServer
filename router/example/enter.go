package example

import (
	api "aiServer/api/v1"
)

type RouterGroup struct {
	FileUploadAndDownloadRouter
}

var (
	exaFileUploadAndDownloadApi = api.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi
)
