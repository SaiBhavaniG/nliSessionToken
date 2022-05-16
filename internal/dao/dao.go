package dao

import (
	cu "github.hdfcbank.com/HDFCBANK/mb-microservices-utils/cache-util"
	cachemodels "github.hdfcbank.com/HDFCBANK/mb-microservices-utils/cache-util/models"
	"github.hdfcbank.com/HDFCBANK/mb-microservices-utils/cache-util/types"
	"github.hdfcbank.com/HDFCBANK/mb-microservices-utils/logger"
)

type daoClient struct {
	client types.CacheAccessorService //cacheaccessor
	*logger.Logger
}

type DaoAccessor interface {
	CheckdeviceID(SetName, deviceID string) (bool, error)
}

func NewDao(logger *logger.Logger, client types.CacheAccessorService) DaoAccessor {
	return &daoClient{
		client: client,
		Logger: logger,
	}
}

func GetDaoAccessor(logger *logger.Logger) (DaoAccessor, error) {
	cacheAccessor, err := cu.GetAeroCacheAccessorService(logger)
	if err != nil {
		return &daoClient{}, nil
	}
	obj := &daoClient{
		client: cacheAccessor,
		Logger: logger,
	}
	return obj, nil
}

//user_device_mapping

func (d *daoClient) CheckdeviceID(SetName, deviceID string) (bool, error) {

	data := cachemodels.GetCacheConfig{
		Key:     deviceID,
		SetName: SetName,
	}
	check, err := d.client.CacheExists(data)
	if err != nil {
		//need to include multilingual error util
		//d.logger.Error("failed", zap.Error(err))
		return false, err
	}

	return check, err
}
