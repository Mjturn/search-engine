package routes

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func HandleRoutes(router *gin.Engine) {
    router.GET("/", func(context *gin.Context) {
        context.File("../frontend/static/html/index.html")
    })
    
    router.GET("/search", func(context *gin.Context) {
        searchQuery := context.Query("search-query-input")
        
        context.HTML(http.StatusOK, "search-results.html", gin.H {
            "searchQuery": searchQuery,
        })
    })
}
