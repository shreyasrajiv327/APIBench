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

	c.JSON(http.StatusCreated, ToMessageResponse(msg))
}

func (h *Handler) Poll(c *gin.Context) {
	msg, err := h.queue.Poll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ToMessageResponse(msg))
}

func (h *Handler) Ack(c *gin.Context) {
	id := c.Param("id")

	if err := h.queue.Ack(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "acknowledged",
	})
}

func (h *Handler) Nack(c *gin.Context) {
	id := c.Param("id")

	if err := h.queue.Nack(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "nacked",
	})
}