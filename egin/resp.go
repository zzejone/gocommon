package egin

import (
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type M map[string]interface{}

type JSONRET struct {
	ErrCode int                    `json:"errcode"`
	ErrMsg  string                 `json:"errmsg"`
	Data    map[string]interface{} `json:"data"`
}

// JSON 微信小程序返回错误信息：请联系客服
func JSON(c *gin.Context, err error, data map[string]interface{}) {
	var resp JSONRET
	if err == nil {
		resp = JSONRET{
			ErrCode: 0,
			ErrMsg:  "",
			Data:    data,
		}
	} else {
		_, file, line, _ := runtime.Caller(1)
		Logger.Warnf("%s:%d %s", file, line, err)
		// ErrMsg:  strings.Join(errContents, " ; "),
		resp = JSONRET{
			ErrCode: line,
			ErrMsg:  "请联系客服",
			Data:    nil,
		}
	}
	c.JSON(200, resp)
}

func Resp(c *gin.Context, err error, data map[string]interface{}) {
	var resp JSONRET
	if err == nil {
		resp = JSONRET{
			ErrCode: 0,
			ErrMsg:  "",
			Data:    data,
		}
	} else {
		_, file, line, _ := runtime.Caller(1)
		Logger.Warnf("%s:%d %s", file, line, err)
		resp = JSONRET{
			ErrCode: line,
			ErrMsg:  err.Error(),
			Data:    nil,
		}
	}
	c.JSON(200, resp)
}

func BindErr(c *gin.Context, err error) {
	var resp JSONRET
	_, file, line, _ := runtime.Caller(1)
	Logger.Warnf("%s:%d %s", file, line, err)
	// ErrMsg:  strings.Join(errContents, " ; "),
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		var errContents []string
		for _, v := range errs.Translate(trans) {
			errContents = append(errContents, v)
		}
		resp = JSONRET{
			ErrCode: line,
			ErrMsg:  strings.Join(errContents, " ; "),
		}
	} else {
		resp = JSONRET{
			ErrCode: line,
			ErrMsg:  err.Error(),
		}
	}
	c.JSON(200, resp)
}

func SpecialJson(c *gin.Context, errcode int, data map[string]any) {
	resp := JSONRET{
		ErrCode: errcode,
		ErrMsg:  "",
		Data:    data,
	}
	c.JSON(200, resp)
}
