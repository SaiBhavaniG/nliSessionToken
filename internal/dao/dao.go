package dao

import (
	cu "github.hdfcbank.com/HDFCBANK/mb-microservices-utils/cache-util"
	cachemodels "github.hdfcbank.com/HDFCBANK/mb-microservices-utils/cache-util/models"
	"github.hdfcbank.com/HDFCBANK/mb-microservices-utils/cache-util/types"
	"github.hdfcbank.com/HDFCBANK/mb-microservices-utils/logger"
)

type dao struct {
	client types.CacheAccessorService //cacheaccessor
	logger *logger.Logger
}

type DaoAccessor interface {
	CheckdeviceID(SetName, deviceID string) (bool, error)
}

func GetDaoAccessor(logger *logger.Logger) (DaoAccessor, error) {
	cacheAccessor, err := cu.GetAeroCacheAccessorService(logger)
	if err != nil {
		return &dao{}, nil
	}
	obj := &dao{
		client: cacheAccessor,
		logger: logger,
	}
	return obj, nil
}

//user_device_mapping to be Changed(naming Convention)

func (d *dao) CheckdeviceID(SetName, deviceID string) (bool, error) {

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
