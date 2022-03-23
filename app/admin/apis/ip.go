package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/Camera"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type Ip struct {
	api.Api
}

// GetPage 获取ip地址列表
// @Summary 获取ip地址列表
// @Description 获取ip地址列表
// @Tags ip地址
// @Param ip query string false "ip地址"
// @Param status query string false "状态"
// @Param publishAt query time.Time false "发布时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Ip}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ip [get]
// @Security Bearer
func (e Ip) GetPage(c *gin.Context) {
    req := dto.IpGetPageReq{}
    s := service.Ip{}
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
	list := make([]models.Ip, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取ip地址 失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取ip地址
// @Summary 获取ip地址
// @Description 获取ip地址
// @Tags ip地址
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.Ip} "{"code": 200, "data": [...]}"
// @Router /api/v1/ip/{id} [get]
// @Security Bearer
func (e Ip) Get(c *gin.Context) {
	req := dto.IpGetReq{}
	s := service.Ip{}
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
	var object models.Ip

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取ip地址失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建ip地址
// @Summary 创建ip地址
// @Description 创建ip地址
// @Tags ip地址
// @Accept application/json
// @Product application/json
// @Param data body dto.IpInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ip [post]
// @Security Bearer
func (e Ip) Insert(c *gin.Context) {
    req := dto.IpInsertReq{}
    s := service.Ip{}
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
		e.Error(500, err, fmt.Sprintf("创建ip地址  失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
	// 发送修改后信息给nats
	Camera.RR_CameraOpera()
}

// Update 修改ip地址
// @Summary 修改ip地址
// @Description 修改ip地址
// @Tags ip地址
// @Accept application/json
// @Product application/json
// @Param data body dto.IpUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ip/{id} [put]
// @Security Bearer
func (e Ip) Update(c *gin.Context) {
    req := dto.IpUpdateReq{}
    s := service.Ip{}
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
		e.Error(500, err, fmt.Sprintf("修改ip地址 失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
	// 发送修改后信息给nats
	Camera.RR_CameraOpera()
}

// Delete 删除ip地址
// @Summary 删除ip地址
// @Description 删除ip地址
// @Tags ip地址
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ip [delete]
// @Security Bearer
func (e Ip) Delete(c *gin.Context) {
    s := service.Ip{}
    req := dto.IpDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除ip地址失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")

	// 发送修改后信息给nats
	Camera.RR_CameraOpera()
}