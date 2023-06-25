package validators

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
	"silent-cxl.top/app/response"
	"strings"
)

type Handler func(ctx *gin.Context, v interface{})

// ValidatorMessages 验证器自定义错误信息字典
type ValidatorMessages map[string]string

// GetErrorMsg 获取自定义错误信息
func GetErrorMsg(p interface{}, err error) string {
	for _, v := range err.(validator.ValidationErrors) {
		if message, exist := GetMessage(p)[v.Field()+"."+v.Tag()]; exist {
			return message
		}
		return v.Error()
	}
	return "Parameter error"
}

func GetMessage(p interface{}) map[string]string {
	s := reflect.TypeOf(p).Elem()
	messages := make(map[string]string)
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		name, tag, msg := field.Name, field.Tag.Get("binding"), field.Tag.Get("msg")
		if tag != "" && msg != "" {
			tags := strings.Split(tag, ",")
			for _, t := range tags {
				messages[fmt.Sprintf("%s.%s", name, t)] = msg
			}
		}
	}
	return messages
}

func BindData(ctx *gin.Context, v interface{}, h Handler, bindFs ...func(*gin.Context, interface{}) error) {
	if err := ctx.ShouldBind(v); err != nil {
		response.ValidFail(ctx, GetErrorMsg(v, err))
		return
	}
	for _, f := range bindFs {
		if err := f(ctx, v); err != nil {
			response.ValidFail(ctx, err.Error())
			return
		}
	}

	h(ctx, v)
}
