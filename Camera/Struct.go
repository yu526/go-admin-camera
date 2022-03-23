package Camera
// 结构体

import (
	"time"
)

// camera结构体
type Camera struct {
	Data struct {
		Count int `json:"count"`
		PageIndex int `json:"pageIndex"`
		PageSize int `json:"pageSize"`
		List []struct {
			IP string `json:"ip"`
			Status string `json:"status"`
			//PublishAt time.Time `json:"publishAt"`
		} `json:"list"`
	} `json:"data"`
}

// 检测结果结构体 and 预警类型结构体
type Result struct {
	Timestamp time.Time `json:"timestamp"`
	Imagepath string    `json:"imagepath"`
	Result    string    `json:"result"`
}

// result_api返回形式
type Response struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Data      int    `json:"data"`
}

// gps结构体
type Gps struct {
	Timestamp time.Time `json:"timestamp"`
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}