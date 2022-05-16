package handler

import (
	"nliSessionToken/internal/service"

	"github.hdfcbank.com/HDFCBANK/mb-microservices-utils/logger"
	"github.hdfcbank.com/HDFCBANK/mb-microservices-utils/response"
)

type Controller struct {
	response.Response
	service service.TokenService
	*logger.Logger
}

func NewController(logger *logger.Logger, service service.TokenService) *Controller {
	return &Controller{
		Response: *response.NewResponse(logger),
		service:  service,
		Logger:   logger,
	}
}
