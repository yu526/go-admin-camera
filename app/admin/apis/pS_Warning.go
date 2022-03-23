package apis

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"go-admin/Camera"
	"go-admin/common/apis"
	"os/signal"
	"runtime"
	"syscall"
)

type RealWarning struct {
	apis.Api
}

// 实时Gps信息
var real_Warning Camera.Result

// 订阅PS_DeviceStatus，返回实时结果
func (e RealGps)GetPS_Warning(c *gin.Context) {
	err := e.MakeContext(c).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}


	//fmt.Println("----",data,"----")
	e.OK(real_Warning,"success")
}

// 获取实时gps信息
func PS_Warning(){
	fmt.Println("PS_Warning 订阅服务已启动")
	nc, _ := nats.Connect(nats_ip)

	_, err := nc.Subscribe("PS_Warning", func(m *nats.Msg) {
		//fmt.Printf("Received a message: %s\n", string(m.Data))
		//fmt.Println(string(m.Data))
		json.Unmarshal(m.Data, &real_Warning)
	})
	if err != nil {
		//hand error
	}

	// 阻止进程结束而收不到消息
	signal.Ignore(syscall.SIGHUP)
	runtime.Goexit()

}