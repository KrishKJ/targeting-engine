package delivery

import (
	"github.com/gin-gonic/gin"
)

// LoadServices attaches delivery-related routes to the router group
func LoadServices(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		v1.GET("/delivery", HandleDelivery)
	}
}

// HandleDelivery will be implemented later
func HandleDelivery(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delivery endpoint hit!",
	})
}
