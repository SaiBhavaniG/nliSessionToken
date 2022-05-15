package prepare

import (
	"context"
	"nliSessionToken/internal/service"

	"github.com/go-playground/validator/v10"
	"github.hdfcbank.com/HDFCBANK/mb-microservices-utils/logger"
	"github.hdfcbank.com/HDFCBANK/mb-microservices-utils/parser"
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
	ctx := context.Background()

	// configurable and pluggable doa and client
	//daoClient := prepareAeroClient(logger, config)
	srv := service.NewService(daoClient, TokenClient, logger)
	router := controller.NewController(logger, srv)

	// serve the app
	serveApp(router, logger, config.HostAddress)
}

/*func prepareAeroClient(logger *logger.Logger, config invalid type) {
	panic("unimplemented")
}*/
