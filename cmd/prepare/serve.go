package prepare

import (
	"net/http"
	"nliSessionToken/internal/handler"
	"sync"

	"github.hdfcbank.com/HDFCBANK/mb-microservices-utils/logger"
)

func serveApp(endpoint *handler.Controller, logger *logger.Logger, host string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		logger.Info("serving token-server in " + host)
		err := http.ListenAndServe(host, endpoint.Router())
		if err != nil {
			logger.Fatal("failed to serve the app, error: " + err.Error())
		}
		wg.Done()
	}()

	wg.Wait()
}
