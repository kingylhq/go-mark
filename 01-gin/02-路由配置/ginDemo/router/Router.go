package router

import (
	"ginDemo/common"
	v12 "ginDemo/controller/v1"
	v22 "ginDemo/controller/v2"
	"github.com/gin-gonic/gin"
	"net/url"
	"strconv"
)

func InitRouter(r *gin.Engine) {

	r.GET("/sn", SignDemo)

	// v1 版本
	GroupV1 := r.Group("/v1")
	{
		GroupV1.Any("/product/add", v12.AddProduct)
		GroupV1.Any("/member/add", v12.AddMember)
	}

	// v2 版本
	GroupV2 := r.Group("/v2", common.VerifySign)
	{
		GroupV2.Any("/product/add", v22.AddProduct)
		GroupV2.Any("/member/add", v22.AddMember)
	}
}

func SignDemo(c *gin.Context) {
	ts := strconv.FormatInt(common.GetTimeUnix(), 10)
	res := map[string]interface{}{}
	params := url.Values{
		"name":  []string{"a"},
		"price": []string{"10"},
		"ts":    []string{ts},
	}
	res["sn"] = common.CreateSign(params)
	res["ts"] = ts
	common.RetJson("200", "", res, c)
}
