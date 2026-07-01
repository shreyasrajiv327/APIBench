package rest

import(
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shreyasrajiv327/APIBench/internal/service"
)

type Handler struct{
	queue service.QueueService
}


func NewHandler(queue service.QueueService) *Handler {
	return &Handler{
		queue: queue,
	}
}

type PublishRequest struct {
	Payload string `json:"payload"`
}

func (h *Handler) Publish(c *gin.Context) {
	var req PublishRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	msg, err := h.queue.Publish([]byte(req.Payload))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, msg)
}