package routes

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func HandleRoutes(router *gin.Engine) {
    router.GET("/", func(context *gin.Context) {
        context.HTML(http.StatusOK, "index.html", nil)
    })
}
