package dto

import (
     
     
     "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type GpsGetPageReq struct {
	dto.Pagination     `search:"-"`
    Timestamp time.Time `form:"timestamp"  search:"type:exact;column:timestamp;table:gps" comment:"时间戳"`
    Lat string `form:"lat"  search:"type:exact;column:lat;table:gps" comment:"纬度"`
    Lon string `form:"lon"  search:"type:exact;column:lon;table:gps" comment:"经度"`
    GpsOrder
}

type GpsOrder struct {Id int `form:"idOrder"  search:"type:order;column:id;table:gps"`
    Timestamp time.Time `form:"timestampOrder"  search:"type:order;column:timestamp;table:gps"`
    Lat string `form:"latOrder"  search:"type:order;column:lat;table:gps"`
    Lon string `form:"lonOrder"  search:"type:order;column:lon;table:gps"`
    CreatedAt time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:gps"`
    UpdatedAt time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:gps"`
    DeletedAt time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:gps"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:gps"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:gps"`
    
}

func (m *GpsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type GpsInsertReq struct {
    Id int `json:"-" comment:"编码"` // 编码
    Timestamp time.Time `json:"timestamp" comment:"时间戳"`
    Lat string `json:"lat" comment:"纬度"`
    Lon string `json:"lon" comment:"经度"`
    common.ControlBy
}

func (s *GpsInsertReq) Generate(model *models.Gps)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Timestamp = s.Timestamp
    model.Lat = s.Lat
    model.Lon = s.Lon
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *GpsInsertReq) GetId() interface{} {
	return s.Id
}

type GpsUpdateReq struct {
    Id int `uri:"id" comment:"编码"` // 编码
    Timestamp time.Time `json:"timestamp" comment:"时间戳"`
    Lat string `json:"lat" comment:"纬度"`
    Lon string `json:"lon" comment:"经度"`
    common.ControlBy
}

func (s *GpsUpdateReq) Generate(model *models.Gps)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Timestamp = s.Timestamp
    model.Lat = s.Lat
    model.Lon = s.Lon
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *GpsUpdateReq) GetId() interface{} {
	return s.Id
}

// GpsGetReq 功能获取请求参数
type GpsGetReq struct {
     Id int `uri:"id"`
}
func (s *GpsGetReq) GetId() interface{} {
	return s.Id
}

// GpsDeleteReq 功能删除请求参数
type GpsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *GpsDeleteReq) GetId() interface{} {
	return s.Ids
}