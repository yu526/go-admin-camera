package models

import (
     
     
     "time"

	"go-admin/common/models"

)

type Result struct {
    models.Model
    
    Timestamp time.Time `json:"timestamp" gorm:"type:timestamp;comment:时间戳"` 
    Imagepath string `json:"imagepath" gorm:"type:varchar(128);comment:图片地址"` 
    Result string `json:"result" gorm:"type:json;comment:识别结果"` 
    models.ModelTime
    models.ControlBy
}

func (Result) TableName() string {
    return "result"
}

func (e *Result) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Result) GetId() interface{} {
	return e.Id
}