package delivery

import (
	"github.com/KrishKJ/targeting-engine/db"
	"github.com/KrishKJ/targeting-engine/delivery/models"
	"github.com/KrishKJ/targeting-engine/delivery/utils"
	"github.com/gin-gonic/gin"
)

// LoadServices attaches delivery-related routes to the router group
func LoadServices(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		v1.GET("/delivery", HandleDelivery)
		v1.GET("/refresh-cache", RefreshCache)
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

// RefreshCache triggers a refresh of the Redis cache for campaigns
func RefreshCache(c *gin.Context) {
	db.LoadCampaignsToCache()
	c.JSON(200, gin.H{"status": "Cache refreshed successfully"})
}

