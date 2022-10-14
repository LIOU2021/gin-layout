package controller

import (
	"net/http"

	"github.com/LIOU2021/gin-layout/env"
	"github.com/LIOU2021/gin-layout/helpers"
	"github.com/gin-gonic/gin"
)

type csrfTokenController struct{}

var CsrfTokenController = &csrfTokenController{}

// response Csrf-Token in header
func (controller *csrfTokenController) Generate(ctx *gin.Context) {
	key := env.AppSetting.Key
	hash := helpers.HashMake(key)
	ctx.Header("Csrf-Token", hash)
	ctx.String(http.StatusOK, "success")
}

// verify csrf token
func (controller *csrfTokenController) Verify(ctx *gin.Context) {
	ctx.String(http.StatusOK, "csrf token verify success")
}
