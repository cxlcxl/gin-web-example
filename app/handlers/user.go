package handlers

import (
	"github.com/gin-gonic/gin"
	"silent-cxl.top/app/model"
	"silent-cxl.top/app/response"
	"silent-cxl.top/app/utils"
	"silent-cxl.top/app/validators/data"
	"silent-cxl.top/app/vars"
)

type User struct{}

func (h User) UserQuery(ctx *gin.Context, p interface{}) {
	params := p.(*data.VUserQueryData)
	offset := utils.GetPages(params.Page, params.Limit)
	users, total, err := model.DbUser(vars.Mysql).List("", -1, offset, params.Limit)
	if err != nil {
		response.Fail(ctx, "请求失败")
		return
	}

	response.Success(ctx, gin.H{"total": total, "list": users})
}
