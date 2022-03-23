package Camera
//	与数据库的操作
import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	// nats的ip以及端口号
	nats_ip = "114.213.213.217:4222"
)

// 从数据库获取摄像头参数 ip 和 开启状态（on|off）
func getStatusCamera() []byte{
	resp, err := http.Get("http://localhost:8000/api/v1/ip")
	if err != nil {
		// handle error
	}
	//fmt.Print(resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	//var camera Camera
	//json.Unmarshal(body, &camera)
	//fmt.Println(camera)
	return body
}

// 将检测结果 或者 预警结果 插入到数据库的 相应 表中, data是预警和检测结果结构体，table是表名
func insertResult(data Result, table string) error{
	da, err := json.Marshal(data)
	if err != nil {
		return err
		// handle error
	}
	resp, err := http.Post("http://localhost:8000/api/v1/" + table, "application/json", bytes.NewBuffer(da))
	if err != nil {
		return err
		// handle error
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	var response Response
	json.Unmarshal(result, &response)
	if response.Msg == "创建成功" {
		return nil
	}
	return errors.New("创建失败")
}

// 将 GPS到数据库的 gps 表中,table是表名
func insertGps(data Gps, table string) error{
	da, err := json.Marshal(data)
	if err != nil {
		return err
		// handle error
	}
	resp, err := http.Post("http://localhost:8000/api/v1/" + table, "application/json", bytes.NewBuffer(da))
	if err != nil {
		return err
		// handle error
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	var response Response
	json.Unmarshal(result, &response)
	if response.Msg == "创建成功" {
		return nil
	}
	return errors.New("创建失败")
}