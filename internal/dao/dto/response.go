package dto

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ParseValidateErrorsResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode int, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  GetMsg(errCode),
		Data: data,
	})
	return
}

// Response setting gin.JSON
func (g *Gin) DirectResponse(httpCode int, msg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  msg,
		Data: data,
	})
	return
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := EnMsg[code]
	if ok {
		return msg
	}

	return EnMsg[NETWORK_ERROR]
}
