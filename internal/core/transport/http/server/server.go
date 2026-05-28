package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/0ScPro0/go-todolist/internal/core/config"
	"github.com/0ScPro0/go-todolist/internal/core/logger"
	core_http_middleware "github.com/0ScPro0/go-todolist/internal/core/transport/http/middleware"
)

type HTTPServer struct {
	mux    *http.ServeMux
	config *config.ServerConfig
	log    *logger.Logger

	middleware []core_http_middleware.Middleware
}

func NewHTTPServer(
	cfg *config.Config,
	log *logger.Logger,
	middleware ...core_http_middleware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		mux: http.NewServeMux(),
		config: &cfg.Server,
		log: log,
		middleware: middleware,
	}
}

func (h *HTTPServer) RegisterAPIVersionRouter(routers ...APIVersionRouter){
	for _, router := range routers{
		prefix := "/api/" + string(router.apiVersion)

		h.mux.Handle(
			prefix + "/", 
			http.StripPrefix(prefix, router),
		)
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddleware(h.mux, h.middleware...)

	addr := net.JoinHostPort(h.config.Host, strconv.Itoa(h.config.Port))
	server := &http.Server{
		Addr: addr,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		h.log.Warn("Start HTTP server", zap.String("addr", addr))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("Listen and server HTTP: %w", err)
		}
	case <-ctx.Done():
		h.log.Warn("Shutdown HTTP server...")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			h.config.ShutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("Shutdown HTTP server: %w", err)
		}
	}

	h.log.Warn("HTTP server stopped")
	return nil

}