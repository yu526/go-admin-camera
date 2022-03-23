package models

import (
     
     
     "time"

	"go-admin/common/models"

)

type Ip struct {
    models.Model
    
    Ip string `json:"ip" gorm:"type:varchar(128);comment:ip地址"` 
    Status string `json:"status" gorm:"type:int(1);comment:状态"` 
    PublishAt time.Time `json:"publishAt" gorm:"type:timestamp;comment:发布时间"` 
    models.ModelTime
    models.ControlBy
}

func (Ip) TableName() string {
    return "ip"
}

func (e *Ip) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Ip) GetId() interface{} {
	return e.Id
}