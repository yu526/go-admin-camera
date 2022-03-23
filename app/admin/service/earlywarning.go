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

type Earlywarning struct {
	service.Service
}

// GetPage 获取Earlywarning列表
func (e *Earlywarning) GetPage(c *dto.EarlywarningGetPageReq, p *actions.DataPermission, list *[]models.Earlywarning, count *int64) error {
	var err error
	var data models.Earlywarning

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EarlywarningService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Earlywarning对象
func (e *Earlywarning) Get(d *dto.EarlywarningGetReq, p *actions.DataPermission, model *models.Earlywarning) error {
	var data models.Earlywarning

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEarlywarning error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Earlywarning对象
func (e *Earlywarning) Insert(c *dto.EarlywarningInsertReq) error {
    var err error
    var data models.Earlywarning
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EarlywarningService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Earlywarning对象
func (e *Earlywarning) Update(c *dto.EarlywarningUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Earlywarning{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if db.Error != nil {
        e.Log.Errorf("EarlywarningService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Earlywarning
func (e *Earlywarning) Remove(d *dto.EarlywarningDeleteReq, p *actions.DataPermission) error {
	var data models.Earlywarning

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveEarlywarning error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}