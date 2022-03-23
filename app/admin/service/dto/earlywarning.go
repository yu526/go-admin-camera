package dto

import (
     
     
     "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EarlywarningGetPageReq struct {
	dto.Pagination     `search:"-"`
    Timestamp time.Time `form:"timestamp"  search:"type:exact;column:timestamp;table:earlywarning" comment:"时间戳"`
    Imagepath string `form:"imagepath"  search:"type:exact;column:imagepath;table:earlywarning" comment:"图片地址"`
    Result string `form:"result"  search:"type:exact;column:result;table:earlywarning" comment:"预警类型"`
    EarlywarningOrder
}

type EarlywarningOrder struct {Id int `form:"idOrder"  search:"type:order;column:id;table:earlywarning"`
    Timestamp time.Time `form:"timestampOrder"  search:"type:order;column:timestamp;table:earlywarning"`
    Imagepath string `form:"imagepathOrder"  search:"type:order;column:imagepath;table:earlywarning"`
    Result string `form:"resultOrder"  search:"type:order;column:result;table:earlywarning"`
    CreatedAt time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:earlywarning"`
    UpdatedAt time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:earlywarning"`
    DeletedAt time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:earlywarning"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:earlywarning"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:earlywarning"`
    
}

func (m *EarlywarningGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EarlywarningInsertReq struct {
    Id int `json:"-" comment:"编码"` // 编码
    Timestamp time.Time `json:"timestamp" comment:"时间戳"`
    Imagepath string `json:"imagepath" comment:"图片地址"`
    Result string `json:"result" comment:"预警类型"`
    common.ControlBy
}

func (s *EarlywarningInsertReq) Generate(model *models.Earlywarning)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Timestamp = s.Timestamp
    model.Imagepath = s.Imagepath
    model.Result = s.Result
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *EarlywarningInsertReq) GetId() interface{} {
	return s.Id
}

type EarlywarningUpdateReq struct {
    Id int `uri:"id" comment:"编码"` // 编码
    Timestamp time.Time `json:"timestamp" comment:"时间戳"`
    Imagepath string `json:"imagepath" comment:"图片地址"`
    Result string `json:"result" comment:"预警类型"`
    common.ControlBy
}

func (s *EarlywarningUpdateReq) Generate(model *models.Earlywarning)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Timestamp = s.Timestamp
    model.Imagepath = s.Imagepath
    model.Result = s.Result
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *EarlywarningUpdateReq) GetId() interface{} {
	return s.Id
}

// EarlywarningGetReq 功能获取请求参数
type EarlywarningGetReq struct {
     Id int `uri:"id"`
}
func (s *EarlywarningGetReq) GetId() interface{} {
	return s.Id
}

// EarlywarningDeleteReq 功能删除请求参数
type EarlywarningDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EarlywarningDeleteReq) GetId() interface{} {
	return s.Ids
}