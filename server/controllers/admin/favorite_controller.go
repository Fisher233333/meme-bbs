package admin

import (
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/web"
	"github.com/mlogclub/simple/web/params"

	"bbs-go/model"
	"bbs-go/services"
)

type FavoriteController struct {
	Ctx iris.Context
}

func (c *FavoriteController) GetBy(id int64) *web.JsonResult {
	t := services.FavoriteService.Get(id)
	if t == nil {
		return web.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return web.JsonData(t)
}

func (c *FavoriteController) AnyList() *web.JsonResult {
	list, paging := services.FavoriteService.FindPageByParams(params.NewQueryParams(c.Ctx).PageByReq().Desc("id"))
	return web.JsonData(&web.PageResult{Results: list, Page: paging})
}

func (c *FavoriteController) PostCreate() *web.JsonResult {
	t := &model.Favorite{}
	params.ReadForm(c.Ctx, t)

	err := services.FavoriteService.Create(t)
	if err != nil {
		return web.JsonErrorMsg(err.Error())
	}
	return web.JsonData(t)
}

func (c *FavoriteController) PostUpdate() *web.JsonResult {
	id, err := params.FormValueInt64(c.Ctx, "id")
	if err != nil {
		return web.JsonErrorMsg(err.Error())
	}
	t := services.FavoriteService.Get(id)
	if t == nil {
		return web.JsonErrorMsg("entity not found")
	}

	params.ReadForm(c.Ctx, t)

	err = services.FavoriteService.Update(t)
	if err != nil {
		return web.JsonErrorMsg(err.Error())
	}
	return web.JsonData(t)
}
