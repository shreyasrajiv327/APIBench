package rest

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	router.POST("/messages", handler.Publish)
	router.GET("/messages/next", handler.Poll)
    router.POST("/messages/:id/ack", handler.Ack)
    router.POST("/messages/:id/nack", handler.Nack)
}