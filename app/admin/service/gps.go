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

type Gps struct {
	service.Service
}

// GetPage 获取Gps列表
func (e *Gps) GetPage(c *dto.GpsGetPageReq, p *actions.DataPermission, list *[]models.Gps, count *int64) error {
	var err error
	var data models.Gps

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("GpsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Gps对象
func (e *Gps) Get(d *dto.GpsGetReq, p *actions.DataPermission, model *models.Gps) error {
	var data models.Gps

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetGps error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Gps对象
func (e *Gps) Insert(c *dto.GpsInsertReq) error {
    var err error
    var data models.Gps
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("GpsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Gps对象
func (e *Gps) Update(c *dto.GpsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Gps{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if db.Error != nil {
        e.Log.Errorf("GpsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Gps
func (e *Gps) Remove(d *dto.GpsDeleteReq, p *actions.DataPermission) error {
	var data models.Gps

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveGps error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}