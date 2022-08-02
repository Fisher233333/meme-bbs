package admin

import (
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/web"
	"github.com/mlogclub/simple/web/params"

	"bbs-go/services"
)

type SysConfigController struct {
	Ctx iris.Context
}

func (c *SysConfigController) GetBy(id int64) *web.JsonResult {
	t := services.SysConfigService.Get(id)
	if t == nil {
		return web.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return web.JsonData(t)
}

func (c *SysConfigController) AnyList() *web.JsonResult {
	list, paging := services.SysConfigService.FindPageByParams(params.NewQueryParams(c.Ctx).PageByReq().Desc("id"))
	return web.JsonData(&web.PageResult{Results: list, Page: paging})
}

func (c *SysConfigController) GetAll() *web.JsonResult {
	config := services.SysConfigService.GetConfig()
	return web.JsonData(config)
}

func (c *SysConfigController) PostSave() *web.JsonResult {
	config := c.Ctx.FormValue("config")
	if err := services.SysConfigService.SetAll(config); err != nil {
		return web.JsonErrorMsg(err.Error())
	}
	return web.JsonSuccess()
}
