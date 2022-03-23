package Camera
// 与nats的通信

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// 订阅RR_CameraParam并回复摄像头参数
func RR_CameraParam() {
	// 等待后端启动时间
	time.Sleep(time.Second * 6)
	fmt.Println("RR_CameraParam 订阅服务已启动")
	nc, _ := nats.Connect(nats_ip)

	_, _ = nc.Subscribe("RR_CameraParam", func(m *nats.Msg){
		//fmt.Println("receive a message: %s\n", string(m.Data))
		if string(m.Data) != " " {
			fmt.Println("RR_CameraParam 请求相机参数，正在获取.....")
			// 获取相机信息
			s := getStatusCamera()
			// 获取成功，返回相机参数
			fmt.Println("获取成功，返回参数")
			//fmt.Println(s)
			_ = nc.Publish(m.Reply, s)
		}
	})
	//阻止进程结束而收不到消息
	signal.Ignore(syscall.SIGHUP)
	runtime.Goexit()
}
// 发布摄像头参数， 发布条件为ip地址的表增加 删除 修改
func RR_CameraOpera() {
	nc, err := nats.Connect(nats_ip)
	if err != nil {
		// handle error
	}
	s := getStatusCamera
	msg, err := nc.Request("RR_CameraSetting", s(), 2 * time.Second)
	if err != nil {
		return
	}
	fmt.Print(string(msg.Data))

	// 关闭连接
	nc.Close()
}

// 订阅PS_CameraResult 检测结果 ，并存入数据库
func PS_CameraResult() {
	// 等待后端启动时间
	time.Sleep(time.Second * 6)
	fmt.Println("PS_CameraResult 订阅服务已启动")
	nc, _ := nats.Connect(nats_ip)

	_, _ = nc.Subscribe("PS_CameraResult", func(m *nats.Msg){
		//fmt.Println("receive a message: %s\n", string(m.Data))

		fmt.Println("PS_CameraResult 获取到检测结果，正在存入数据库")
		// 解析 msg
		var data Result
		json.Unmarshal(m.Data, &data)
		// 存入数据库
		table := "result"
		err := insertResult(data, table)
		if err != nil {
			fmt.Println("存储失败")
		} else {
			fmt.Println("存储成功")
		}
	})
	//阻止进程结束而收不到消息
	signal.Ignore(syscall.SIGHUP)
	runtime.Goexit()
}

// 订阅PQ_Warning 预警结果 ，并存入数据库
func PQ_Warning() {
	// 等待后端启动时间
	time.Sleep(time.Second * 6)
	fmt.Println("PQ_Warning 订阅服务已启动")
	nc, _ := nats.Connect(nats_ip)

	_, _ = nc.QueueSubscribe("PQ_Warning", "queue", func(m *nats.Msg) {
		//fmt.Println("receive a message: %s\n", string(m.Data))

		fmt.Println("PQ_Warning 获取到预警结果，正在存入数据库")
		// 解析 msg
		var data Result
		json.Unmarshal(m.Data, &data)
		// 存入数据库
		table := "earlywarning"
		err := insertResult(data, table)
		if err != nil {
			fmt.Println("存储失败")
		} else {
			fmt.Println("存储成功")
		}
	})
	//阻止进程结束而收不到消息
	signal.Ignore(syscall.SIGHUP)
	runtime.Goexit()
}

// 订阅PQ_DeviceStatus gps结果 ，并存入数据库
func PQ_DeviceStatus() {
	// 等待后端启动时间
	time.Sleep(time.Second * 6)
	fmt.Println("PQ_DeviceStatus 订阅服务已启动")
	nc, _ := nats.Connect(nats_ip)

	_, _ = nc.QueueSubscribe("PQ_DeviceStatus", "queue", func(m *nats.Msg) {
		//fmt.Println("receive a message: %s\n", string(m.Data))
		fmt.Println("PQ_DeviceStatus 获取到gps结果，正在存入数据库")
		// 解析 msg
		var data Gps
		json.Unmarshal(m.Data, &data)
		// 存入数据库
		table := "gps"
		err := insertGps(data, table)
		if err != nil {
			fmt.Println("存储失败")
		} else {
			fmt.Println("存储成功")
		}
	})
	//阻止进程结束而收不到消息
	signal.Ignore(syscall.SIGHUP)
	runtime.Goexit()
}




