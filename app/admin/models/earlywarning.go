package models

import (
     
     
     "time"

	"go-admin/common/models"

)

type Earlywarning struct {
    models.Model
    
    Timestamp time.Time `json:"timestamp" gorm:"type:timestamp;comment:时间戳"` 
    Imagepath string `json:"imagepath" gorm:"type:varchar(128);comment:图片地址"` 
    Result string `json:"result" gorm:"type:json;comment:预警类型"` 
    models.ModelTime
    models.ControlBy
}

func (Earlywarning) TableName() string {
    return "earlywarning"
}

func (e *Earlywarning) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Earlywarning) GetId() interface{} {
	return e.Id
}