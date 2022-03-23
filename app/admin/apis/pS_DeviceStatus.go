package apis

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"go-admin/common/apis"
	"os/signal"
	"runtime"
	"syscall"
	"go-admin/Camera"
)

type RealGps struct {
	apis.Api
}

const (
	// nats的ip以及端口号
	nats_ip = "114.213.213.217:4222"
)

// 实时Gps信息
var real_Gps Camera.Gps

// 订阅PS_DeviceStatus，返回实时结果
func (e RealGps)GetPS_DeviceStatus(c *gin.Context) {
	err := e.MakeContext(c).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}


	//fmt.Println("----",data,"----")
	e.OK(real_Gps,"success")
}

// 获取实时gps信息
func PS_DeviceStatus(){
	fmt.Println("PS_DeviceStatus 订阅服务已启动")
	nc, _ := nats.Connect(nats_ip)

	_, err := nc.Subscribe("PS_DeviceStatus", func(m *nats.Msg) {
		//fmt.Printf("Received a message: %s\n", string(m.Data))
		//fmt.Println(string(m.Data))
		json.Unmarshal(m.Data, &real_Gps)
	})
	if err != nil {
		//hand error
	}

	// 阻止进程结束而收不到消息
	signal.Ignore(syscall.SIGHUP)
	runtime.Goexit()

}