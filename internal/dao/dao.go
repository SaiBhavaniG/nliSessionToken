package dao

import (
	cachemodels "github.hdfcbank.com/HDFCBANK/mb-microservices-utils/cache-util/models"
	"github.hdfcbank.com/HDFCBANK/mb-microservices-utils/cache-util/types"
	"github.hdfcbank.com/HDFCBANK/mb-microservices-utils/logger"
	"go.uber.org/zap"
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

//user_device_mapping

func (d *daoClient) CheckdeviceID(SetName, device_ID string) (bool, error) {
	keyvalue := map[string]interface{}{
		"deviceID": device_ID,
	}

	data := cachemodels.GetCacheQueryConfig{
		SetName: SetName,
		KeyVals: keyvalue,
	}
	check, err := d.client.FilterCache(data)
	if err != nil {
		//need to include multilingual error util in handler
		d.Error("failed to get key from db", zap.Error(err))

		return false, err
	}
	if check == nil {
		d.Error("no data found", zap.Error(err))
		return false, err
	}
	ret := true
	return ret, nil
}

/*type GetCacheQueryConfig struct {
	SetName string                 `validate:"required"` //set or table name
	Fields  []string               `validate:"required"`
	KeyVals map[string]interface{} `validate:"required"` // Key value pair for quering
}*/
