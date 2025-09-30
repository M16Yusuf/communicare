package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m16yusuf/communicare/internal/models"
)

// NoRouteHandler
// @Tags NoRoute
// @Router 			/{any} [get]
// @Summary 		testing display for no route
// @Description if route not found, send 404 statusNotfound as response
// @Produce 		json
// @failure 		404		{object} 	models.NotFoundResponse "Not found"
func NoRouteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, models.NotFoundResponse{
		IsSuccess: false,
		Code:      http.StatusNotFound,
		Err:       "not found ...",
	})
}
