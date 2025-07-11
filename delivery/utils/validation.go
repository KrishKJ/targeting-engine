package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/KrishKJ/targeting-engine/delivery/models"
)

func ValidateDeliveryRequest(c *gin.Context) (models.DeliveryRequest, error) {
	app := c.Query("app")
	country := c.Query("country")
	os := c.Query("os")

	if app == "" {
		return models.DeliveryRequest{}, errors.New("missing app param")
	}
	if country == "" {
		return models.DeliveryRequest{}, errors.New("missing country param")
	}
	if os == "" {
		return models.DeliveryRequest{}, errors.New("missing os param")
	}

	return models.DeliveryRequest{
		App:     app,
		Country: country,
		OS:      os,
	}, nil
}
