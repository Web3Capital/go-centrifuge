package api

import (
	"net/http"
	_ "net/http/pprof" // we need this side effect that loads the pprof endpoints to defaultServerMux
	"sync"
	"time"

	"github.com/centrifuge/go-centrifuge/httpapi"
	"github.com/centrifuge/go-centrifuge/utils/httputils"
	"github.com/go-chi/render"
	logging "github.com/ipfs/go-log"
	"golang.org/x/net/context"
)

var log = logging.Logger("api-server")

// Config defines methods required for the package api
type Config interface {
	GetServerAddress() string
	GetServerPort() int
	GetNetworkString() string
	IsPProfEnabled() bool
}

// apiServer is an implementation of node.Server interface for serving HTTP based Centrifuge API
type apiServer struct {
	config Config
}

func (apiServer) Name() string {
	return "APIServer"
}

// Serve exposes the client APIs for interacting with a centrifuge node
func (c apiServer) Start(ctx context.Context, wg *sync.WaitGroup, startupErr chan<- error) {
	defer wg.Done()

	apiAddr := c.config.GetServerAddress()
	mux, err := httpapi.Router(ctx)
	if err != nil {
		startupErr <- err
		return
	}

	if c.config.IsPProfEnabled() {
		log.Info("added pprof endpoints to the server")
		mux.Handle("/debug/", http.DefaultServeMux)
	}

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, httputils.HTTPError{
			Message: http.StatusText(http.StatusNotFound),
		})
	})
	srv := &http.Server{
		Addr:    apiAddr,
		Handler: mux,
	}

	startUpErrOut := make(chan error)
	go func(startUpErrInner chan<- error) {
		log.Infof("HTTP API running at: %s\n", c.config.GetServerAddress())
		log.Infof("Connecting to Network: %s\n", c.config.GetNetworkString())
		err = srv.ListenAndServe()
		if err != nil {
			startUpErrInner <- err
		}
	}(startUpErrOut)

	// listen to context events as well as http server startup errors
	select {
	case err := <-startUpErrOut:
		// this could create an issue if the listeners are blocking.
		// We need to only propagate the error if its an error other than a server closed
		if err != nil && err.Error() != http.ErrServerClosed.Error() {
			startupErr <- err
			return
		}
		// most probably a graceful shutdown
		log.Info(err)
		return
	case <-ctx.Done():
		ctxn, _ := context.WithTimeout(context.Background(), 1*time.Second)
		// gracefully shutdown the server
		// we can only do this because srv is thread safe
		log.Info("Shutting down API server")
		err := srv.Shutdown(ctxn)
		if err != nil {
			panic(err)
		}
		log.Info("API server stopped")
		return
	}
}
