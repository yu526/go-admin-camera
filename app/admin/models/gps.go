package models

import (
     
     
     "time"

	"go-admin/common/models"

)

type Gps struct {
    models.Model
    
    Timestamp time.Time `json:"timestamp" gorm:"type:timestamp;comment:时间戳"` 
    Lat string `json:"lat" gorm:"type:varchar(128);comment:纬度"` 
    Lon string `json:"lon" gorm:"type:varchar(128);comment:经度"` 
    models.ModelTime
    models.ControlBy
}

func (Gps) TableName() string {
    return "gps"
}

func (e *Gps) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Gps) GetId() interface{} {
	return e.Id
}