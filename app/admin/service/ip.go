package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type Ip struct {
	service.Service
}

// GetPage 获取Ip列表
func (e *Ip) GetPage(c *dto.IpGetPageReq, p *actions.DataPermission, list *[]models.Ip, count *int64) error {
	var err error
	var data models.Ip

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("IpService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Ip对象
func (e *Ip) Get(d *dto.IpGetReq, p *actions.DataPermission, model *models.Ip) error {
	var data models.Ip

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetIp error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Ip对象
func (e *Ip) Insert(c *dto.IpInsertReq) error {
    var err error
    var data models.Ip
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("IpService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Ip对象
func (e *Ip) Update(c *dto.IpUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Ip{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if db.Error != nil {
        e.Log.Errorf("IpService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Ip
func (e *Ip) Remove(d *dto.IpDeleteReq, p *actions.DataPermission) error {
	var data models.Ip

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveIp error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}