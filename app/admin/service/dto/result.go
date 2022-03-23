package dto

import (
     
     
     "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ResultGetPageReq struct {
	dto.Pagination     `search:"-"`
    Timestamp time.Time `form:"timestamp"  search:"type:exact;column:timestamp;table:result" comment:"时间戳"`
    Imagepath string `form:"imagepath"  search:"type:exact;column:imagepath;table:result" comment:"图片地址"`
    Result string `form:"result"  search:"type:exact;column:result;table:result" comment:"识别结果"`
    ResultOrder
}

type ResultOrder struct {Id int `form:"idOrder"  search:"type:order;column:id;table:result"`
    Timestamp time.Time `form:"timestampOrder"  search:"type:order;column:timestamp;table:result"`
    Imagepath string `form:"imagepathOrder"  search:"type:order;column:imagepath;table:result"`
    Result string `form:"resultOrder"  search:"type:order;column:result;table:result"`
    CreatedAt time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:result"`
    UpdatedAt time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:result"`
    DeletedAt time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:result"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:result"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:result"`
    
}

func (m *ResultGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ResultInsertReq struct {
    Id int `json:"-" comment:"编码"` // 编码
    Timestamp time.Time `json:"timestamp" comment:"时间戳"`
    Imagepath string `json:"imagepath" comment:"图片地址"`
    Result string `json:"result" comment:"识别结果"`
    common.ControlBy
}

func (s *ResultInsertReq) Generate(model *models.Result)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Timestamp = s.Timestamp
    model.Imagepath = s.Imagepath
    model.Result = s.Result
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *ResultInsertReq) GetId() interface{} {
	return s.Id
}

type ResultUpdateReq struct {
    Id int `uri:"id" comment:"编码"` // 编码
    Timestamp time.Time `json:"timestamp" comment:"时间戳"`
    Imagepath string `json:"imagepath" comment:"图片地址"`
    Result string `json:"result" comment:"识别结果"`
    common.ControlBy
}

func (s *ResultUpdateReq) Generate(model *models.Result)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Timestamp = s.Timestamp
    model.Imagepath = s.Imagepath
    model.Result = s.Result
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *ResultUpdateReq) GetId() interface{} {
	return s.Id
}

// ResultGetReq 功能获取请求参数
type ResultGetReq struct {
     Id int `uri:"id"`
}
func (s *ResultGetReq) GetId() interface{} {
	return s.Id
}

// ResultDeleteReq 功能删除请求参数
type ResultDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ResultDeleteReq) GetId() interface{} {
	return s.Ids
}