package dto

import (
     
     
     "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type IpGetPageReq struct {
	dto.Pagination     `search:"-"`
    Ip string `form:"ip"  search:"type:exact;column:ip;table:ip" comment:"ip地址"`
    Status string `form:"status"  search:"type:exact;column:status;table:ip" comment:"状态"`
    PublishAt time.Time `form:"publishAt"  search:"type:exact;column:publish_at;table:ip" comment:"发布时间"`
    IpOrder
}

type IpOrder struct {Id int `form:"idOrder"  search:"type:order;column:id;table:ip"`
    Ip string `form:"ipOrder"  search:"type:order;column:ip;table:ip"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:ip"`
    PublishAt time.Time `form:"publishAtOrder"  search:"type:order;column:publish_at;table:ip"`
    CreatedAt time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:ip"`
    UpdatedAt time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ip"`
    DeletedAt time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ip"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ip"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ip"`
    
}

func (m *IpGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type IpInsertReq struct {
    Id int `json:"-" comment:"编码"` // 编码
    Ip string `json:"ip" comment:"ip地址"`
    Status string `json:"status" comment:"状态"`
    PublishAt time.Time `json:"publishAt" comment:"发布时间"`
    common.ControlBy
}

func (s *IpInsertReq) Generate(model *models.Ip)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Ip = s.Ip
    model.Status = s.Status
    model.PublishAt = s.PublishAt
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *IpInsertReq) GetId() interface{} {
	return s.Id
}

type IpUpdateReq struct {
    Id int `uri:"id" comment:"编码"` // 编码
    Ip string `json:"ip" comment:"ip地址"`
    Status string `json:"status" comment:"状态"`
    PublishAt time.Time `json:"publishAt" comment:"发布时间"`
    common.ControlBy
}

func (s *IpUpdateReq) Generate(model *models.Ip)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Ip = s.Ip
    model.Status = s.Status
    model.PublishAt = s.PublishAt
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *IpUpdateReq) GetId() interface{} {
	return s.Id
}

// IpGetReq 功能获取请求参数
type IpGetReq struct {
     Id int `uri:"id"`
}
func (s *IpGetReq) GetId() interface{} {
	return s.Id
}

// IpDeleteReq 功能删除请求参数
type IpDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *IpDeleteReq) GetId() interface{} {
	return s.Ids
}