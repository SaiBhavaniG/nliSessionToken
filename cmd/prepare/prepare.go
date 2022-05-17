package prepare

import (
	//"context"
	"nliSessionToken/internal/dao"
	"nliSessionToken/internal/handler"
	"nliSessionToken/internal/models"
	"nliSessionToken/internal/service"

	"github.com/go-playground/validator/v10"
	cu "github.hdfcbank.com/HDFCBANK/mb-microservices-utils/cache-util"
	"github.hdfcbank.com/HDFCBANK/mb-microservices-utils/logger"
	"github.hdfcbank.com/HDFCBANK/mb-microservices-utils/parser"
	//"go.uber.org/zap"
)

func Prepare() {
	var config models.Config
	if err := parser.ParseEnv(&config); err != nil {
		panic("failed to read/parse config, error: " + err.Error())
	}

	v := validator.New()
	if err := v.Struct(config); err != nil {
		panic("invalid configration provided, error: " + err.Error())
	}

	logger := logger.New(config.AppID)
	defer logger.Sync()
	//ctx := context.Background()

	// configurable and pluggable doa and client
	daoAeroClient := GetDaoAccessor(logger)
	srv := service.NewService(daoAeroClient, logger)

	router := handler.NewController(logger, srv)

	// serve the app
	serveApp(router, logger, config.HostAddress)
}

func GetDaoAccessor(logger *logger.Logger) dao.DaoAccessor {
	cacheAccessor, err := cu.GetAeroCacheAccessorService(logger)
	if err != nil {
		logger.Fatal("failed to initalise aero client, error: " + err.Error())
	}
	return dao.NewDao(logger, cacheAccessor)
}

/*func prepareAeroClient(logger *logger.Logger, config models.Config) dao.Dao {
	switch config.DaoClient {
	case "aerospike":
		aeroClient, err := aero.NewClient(config.AeroHostName, config.AeroPort)
		if err != nil {
			logger.Fatal("failed to initalise aero client, error: " + err.Error())
		}
		return dao.NewAeroClient(logger, aeroClient)
	default:
		logger.Fatal("unknown doa client defined")
	}

	return nil
}*/
