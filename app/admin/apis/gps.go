package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type Gps struct {
	api.Api
}

// GetPage 获取gps信息列表
// @Summary 获取gps信息列表
// @Description 获取gps信息列表
// @Tags gps信息
// @Param timestamp query time.Time false "时间戳"
// @Param lat query string false "纬度"
// @Param lon query string false "经度"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Gps}} "{"code": 200, "data": [...]}"
// @Router /api/v1/gps [get]
// @Security Bearer
func (e Gps) GetPage(c *gin.Context) {
    req := dto.GpsGetPageReq{}
    s := service.Gps{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.Gps, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取gps信息 失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取gps信息
// @Summary 获取gps信息
// @Description 获取gps信息
// @Tags gps信息
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.Gps} "{"code": 200, "data": [...]}"
// @Router /api/v1/gps/{id} [get]
// @Security Bearer
func (e Gps) Get(c *gin.Context) {
	req := dto.GpsGetReq{}
	s := service.Gps{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Gps

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取gps信息失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建gps信息
// @Summary 创建gps信息
// @Description 创建gps信息
// @Tags gps信息
// @Accept application/json
// @Product application/json
// @Param data body dto.GpsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/gps [post]
// @Security Bearer
func (e Gps) Insert(c *gin.Context) {
    req := dto.GpsInsertReq{}
    s := service.Gps{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建gps信息  失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改gps信息
// @Summary 修改gps信息
// @Description 修改gps信息
// @Tags gps信息
// @Accept application/json
// @Product application/json
// @Param data body dto.GpsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/gps/{id} [put]
// @Security Bearer
func (e Gps) Update(c *gin.Context) {
    req := dto.GpsUpdateReq{}
    s := service.Gps{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改gps信息 失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除gps信息
// @Summary 删除gps信息
// @Description 删除gps信息
// @Tags gps信息
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/gps [delete]
// @Security Bearer
func (e Gps) Delete(c *gin.Context) {
    s := service.Gps{}
    req := dto.GpsDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除gps信息失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}