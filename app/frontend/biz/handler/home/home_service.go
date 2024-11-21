package home

import (
	"context"
	"net/http"

	"github.com/HCH1212/gomall/app/frontend/biz/service"
	"github.com/HCH1212/gomall/app/frontend/biz/utils"
	home "github.com/HCH1212/gomall/app/frontend/hertz_gen/frontend/home"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Home .
// @router / [GET]
func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req home.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(http.StatusOK, "home.html", resp)
	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
