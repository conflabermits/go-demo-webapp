package server

import (
	"context"
	"net/http"

	"github.com/your-org/my-app/internal/logger"
)

type Server struct {
	addr   string
	http   *http.Server
	logger *logger.Logger
}

func New(addr string, httpHandler http.Handler) *Server {
	return &Server{
		addr: addr,
		http: &http.Server{
			Addr:    addr,
			Handler: httpHandler,
		},
		logger: logger.NewLogger(),
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("Starting server on", zap.String("addr", s.addr))

	go func() {
		<-ctx.Done()
		s.logger.Info("Shutting down server")
		s.http.Shutdown(ctx)
	}()

	return s.http.ListenAndServe()
}
