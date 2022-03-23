package main

import (
	"go-admin/Camera"
	"go-admin/cmd"
	"go-admin/app/admin/apis"
	//"os/signal"
	//"runtime"
	//"syscall"
)

//go:generate swag init --parseDependency --parseDepth=6

// @title go-admin API
// @version 2.0.0
// @description 基于Gin + Vue + Element UI的前后端分离权限管理系统的接口文档
// @license.name MIT
// @license.url https://github.com/go-admin-team/go-admin/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	go Camera.RR_CameraParam()
	go Camera.PS_CameraResult()
	go Camera.PQ_Warning()
	go Camera.PQ_DeviceStatus()
	go apis.PS_DeviceStatus()
	go apis.PS_Warning()
	cmd.Execute()
}



