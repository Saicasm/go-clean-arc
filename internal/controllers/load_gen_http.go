package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoadHttpController struct {
}

func NewLoadHttpController() *LoadHttpController {
	return &LoadHttpController{}
}

// CRUD controllers
func (logController *LoadHttpController) GetServerHealth(ctx *gin.Context) {
	res := make(map[string]interface{}, 0)
	res["status"] = 200
	res["healthy"] = "OK"
	ctx.JSON(http.StatusOK, res)
}
