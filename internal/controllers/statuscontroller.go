package controllers

import (
	"github.com/dkadio/dysr/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"net"
	"time"
)

func Healthcheck(_ *gin.Context) (*models.Healthcheck, error) {
	host := "example.com"
	port := "80"
	timeout := time.Duration(1 * time.Second)
	_, healthy := net.DialTimeout("tcp", host+":"+port, timeout)

	if healthy != nil {
		return &models.Healthcheck{}, errors.Errorf("Healthcheck Failed!")
	}

	return &models.Healthcheck{
		Message: "The API is healthy.",
	}, nil
}
