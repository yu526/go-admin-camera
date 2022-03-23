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

type Earlywarning struct {
	api.Api
}

// GetPage 获取预警类型列表
// @Summary 获取预警类型列表
// @Description 获取预警类型列表
// @Tags 预警类型
// @Param timestamp query time.Time false "时间戳"
// @Param imagepath query string false "图片地址"
// @Param result query string false "预警类型"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Earlywarning}} "{"code": 200, "data": [...]}"
// @Router /api/v1/earlywarning [get]
// @Security Bearer
func (e Earlywarning) GetPage(c *gin.Context) {
    req := dto.EarlywarningGetPageReq{}
    s := service.Earlywarning{}
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
	list := make([]models.Earlywarning, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取预警类型 失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取预警类型
// @Summary 获取预警类型
// @Description 获取预警类型
// @Tags 预警类型
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.Earlywarning} "{"code": 200, "data": [...]}"
// @Router /api/v1/earlywarning/{id} [get]
// @Security Bearer
func (e Earlywarning) Get(c *gin.Context) {
	req := dto.EarlywarningGetReq{}
	s := service.Earlywarning{}
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
	var object models.Earlywarning

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取预警类型失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建预警类型
// @Summary 创建预警类型
// @Description 创建预警类型
// @Tags 预警类型
// @Accept application/json
// @Product application/json
// @Param data body dto.EarlywarningInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/earlywarning [post]
// @Security Bearer
func (e Earlywarning) Insert(c *gin.Context) {
    req := dto.EarlywarningInsertReq{}
    s := service.Earlywarning{}
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
		e.Error(500, err, fmt.Sprintf("创建预警类型  失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改预警类型
// @Summary 修改预警类型
// @Description 修改预警类型
// @Tags 预警类型
// @Accept application/json
// @Product application/json
// @Param data body dto.EarlywarningUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/earlywarning/{id} [put]
// @Security Bearer
func (e Earlywarning) Update(c *gin.Context) {
    req := dto.EarlywarningUpdateReq{}
    s := service.Earlywarning{}
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
		e.Error(500, err, fmt.Sprintf("修改预警类型 失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除预警类型
// @Summary 删除预警类型
// @Description 删除预警类型
// @Tags 预警类型
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/earlywarning [delete]
// @Security Bearer
func (e Earlywarning) Delete(c *gin.Context) {
    s := service.Earlywarning{}
    req := dto.EarlywarningDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除预警类型失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}