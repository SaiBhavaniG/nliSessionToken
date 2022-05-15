package service

import (
	"nliSessionToken/internal/client"
	"nliSessionToken/internal/dao"

	randtok "github.com/mazen160/go-random"
	"github.hdfcbank.com/HDFCBANK/mb-microservices-utils/logger"
	"github.hdfcbank.com/HDFCBANK/mb-microservices-utils/response"
	"go.uber.org/zap"
)

type TokenService interface {
	GenerateToken(SetName, deviceID string) (string, error)
}

type serviceImpl struct {
	doa    dao.DaoAccessor
	client client.TokenClient
	*logger.Logger
	*response.Response
}

func NewService(doa dao.DaoAccessor, client client.TokenClient, logger *logger.Logger) TokenService {
	return &serviceImpl{
		doa:      doa,
		client:   client,
		Logger:   logger,
		Response: response.NewResponse(logger),
	}
}

func (s serviceImpl) GenerateToken(SetName, deviceID string) (string, error) {
	deviceStatus, err := s.doa.CheckdeviceID(SetName, deviceID)
	if !deviceStatus {
		s.Logger.Error("device is not registered", zap.Error(err))
		return "", err
	}

	//Generate Token Randomly, TODO: Generate JWT Token
	token, err := randtok.String(4)
	if err != nil {
		return "", err
	}
	return token, nil

	//call update details function by Saravanan to nli table function
}
