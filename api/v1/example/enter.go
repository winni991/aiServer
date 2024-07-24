package example

import "aiServer/service"

type ApiGroup struct {
	FileUploadAndDownloadApi
}

var (
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
