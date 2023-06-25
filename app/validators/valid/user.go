package valid

import (
	"github.com/gin-gonic/gin"
	"silent-cxl.top/app/handlers"
	"silent-cxl.top/app/validators"
	"silent-cxl.top/app/validators/data"
)

func VUserQuery(ctx *gin.Context) {
	var params data.VUserQueryData
	validators.BindData(ctx, &params, (handlers.User{}).UserQuery)
}
