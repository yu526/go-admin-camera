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

type Result struct {
	service.Service
}

// GetPage 获取Result列表
func (e *Result) GetPage(c *dto.ResultGetPageReq, p *actions.DataPermission, list *[]models.Result, count *int64) error {
	var err error
	var data models.Result

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ResultService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Result对象
func (e *Result) Get(d *dto.ResultGetReq, p *actions.DataPermission, model *models.Result) error {
	var data models.Result

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetResult error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Result对象
func (e *Result) Insert(c *dto.ResultInsertReq) error {
    var err error
    var data models.Result
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ResultService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Result对象
func (e *Result) Update(c *dto.ResultUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Result{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if db.Error != nil {
        e.Log.Errorf("ResultService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Result
func (e *Result) Remove(d *dto.ResultDeleteReq, p *actions.DataPermission) error {
	var data models.Result

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveResult error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}