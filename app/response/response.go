package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    data,
	})
	ctx.Abort()
}

func Fail(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    400,
		"message": err,
		"data":    nil,
	})
	ctx.Abort()
}

func ValidFail(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    10001,
		"message": err,
		"data":    nil,
	})
	ctx.Abort()
}

func TokenExpired(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"code":    10002,
		"message": "token expired",
		"data":    nil,
	})
	ctx.Abort()
}

func AuthFail(ctx *gin.Context) {
	ctx.JSON(http.StatusBadGateway, struct{}{})
	ctx.Abort()
}
