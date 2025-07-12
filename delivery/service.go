package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/KrishKJ/targeting-engine/delivery/models"
	"github.com/KrishKJ/targeting-engine/delivery/utils"
)

// LoadServices attaches delivery-related routes to the router group
func LoadServices(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		v1.GET("/delivery", HandleDelivery)
	}
}

// HandleDelivery processes the delivery request and returns matching campaigns
// It validates the request, processes it, and returns the response
func HandleDelivery(c *gin.Context) {
	var (
		req  models.DeliveryRequest
		data []models.CampaignResponse
		err  error
	)

	if req, err = utils.ValidateDeliveryRequest(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if data, err = utils.ProcessDelivery(req); err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	if len(data) == 0 {
		c.Status(204)
		return
	}

	c.JSON(200, data)
}

