package service

import (
	"context"

	home "github.com/HCH1212/gomall/app/frontend/hertz_gen/frontend/home"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (resp map[string]any, err error) {
	items := []map[string]any{
		{"Name": "T-short-1", "Price": 100, "Picture": "../static/image/t-shirt-1.jpeg"},
		{"Name": "T-short-2", "Price": 110, "Picture": "../static/image/t-shirt-2.jpeg"},
		{"Name": "T-short-3", "Price": 100, "Picture": "../static/image/t-shirt-1.jpeg"},
		{"Name": "T-short-4", "Price": 120, "Picture": "../static/image/t-shirt-2.jpeg"},
		{"Name": "T-short-5", "Price": 100, "Picture": "../static/image/t-shirt-2.jpeg"},
		{"Name": "T-short-6", "Price": 130, "Picture": "../static/image/t-shirt-2.jpeg"},
		{"Name": "T-short-7", "Price": 100, "Picture": "../static/image/t-shirt-2.jpeg"},
		{"Name": "T-short-8", "Price": 140, "Picture": "../static/image/t-shirt.jpeg"},
	}
	resp = make(map[string]any)
	resp["Title"] = "Hot Sale"
	resp["Items"] = items
	return
}
