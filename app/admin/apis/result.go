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

type Result struct {
	api.Api
}

// GetPage 获取检测结果列表
// @Summary 获取检测结果列表
// @Description 获取检测结果列表
// @Tags 检测结果
// @Param timestamp query time.Time false "时间戳"
// @Param imagepath query string false "图片地址"
// @Param result query string false "识别结果"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Result}} "{"code": 200, "data": [...]}"
// @Router /api/v1/result [get]
// @Security Bearer
func (e Result) GetPage(c *gin.Context) {
    req := dto.ResultGetPageReq{}
    s := service.Result{}
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
	list := make([]models.Result, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取检测结果 失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取检测结果
// @Summary 获取检测结果
// @Description 获取检测结果
// @Tags 检测结果
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.Result} "{"code": 200, "data": [...]}"
// @Router /api/v1/result/{id} [get]
// @Security Bearer
func (e Result) Get(c *gin.Context) {
	req := dto.ResultGetReq{}
	s := service.Result{}
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
	var object models.Result

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取检测结果失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建检测结果
// @Summary 创建检测结果
// @Description 创建检测结果
// @Tags 检测结果
// @Accept application/json
// @Product application/json
// @Param data body dto.ResultInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/result [post]
// @Security Bearer
func (e Result) Insert(c *gin.Context) {
    req := dto.ResultInsertReq{}
    s := service.Result{}
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
		e.Error(500, err, fmt.Sprintf("创建检测结果  失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改检测结果
// @Summary 修改检测结果
// @Description 修改检测结果
// @Tags 检测结果
// @Accept application/json
// @Product application/json
// @Param data body dto.ResultUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/result/{id} [put]
// @Security Bearer
func (e Result) Update(c *gin.Context) {
    req := dto.ResultUpdateReq{}
    s := service.Result{}
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
		e.Error(500, err, fmt.Sprintf("修改检测结果 失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除检测结果
// @Summary 删除检测结果
// @Description 删除检测结果
// @Tags 检测结果
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/result [delete]
// @Security Bearer
func (e Result) Delete(c *gin.Context) {
    s := service.Result{}
    req := dto.ResultDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除检测结果失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}