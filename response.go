package golib

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response gin.H

const (
	OKCode           int    = 200
	OKMsg            string = "SUCCESS"
	InternalErrCode  int    = 500
	InternalErrMsg   string = "服务内部错误,请联系管理员!"
	AccessDeniedCode int    = 403
	AccessDeniedMsg  string = "该用户无权访问!"
	NotOpenMsg       string = "该API暂不开放!"
	NoUriCode        int    = 404
	NoUriMsg         string = "该API不存在!"
)

func Result(code int, msg string, data interface{}, c *gin.Context) {
	if c == nil {
		return
	}
	c.JSON(http.StatusOK, Response{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func OkWithMsgAndData(msg string, data interface{}, c *gin.Context) {
	Result(OKCode, msg, data, c)
}

func OkWithoutMsgAndData(c *gin.Context) {
	Result(OKCode, OKMsg, nil, c)
}

func FailWithMsg(code int, msg string, c *gin.Context) {
	Result(code, msg, nil, c)
}

func FailWithInternal(c *gin.Context) {
	Result(InternalErrCode, InternalErrMsg, nil, c)
}
