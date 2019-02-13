package gate

import (
	"context"
	"net/http"
	"time"

	"github.com/QHasaki/Server/logger"
	"github.com/QHasaki/Server/model/v1"
)

var closed = make(chan struct{})

// Service describes game gate service
type Service struct {
	certFile string
	keyFile  string
	server   http.Server
}

// NewService creates a game gate service
func NewService(addr, httpDir string, dataStorage *model.DataStorage) *Service {
	s := new(Service)

	mux := http.NewServeMux()

	mux.Handle("/config/reload", http.StripPrefix("/config", s.route()))

	s.server = http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return s
}

// AddTLSConfig binds cert file and key file to game gate service
func (s *Service) AddTLSConfig(certFile, keyFile string) {
	s.certFile = certFile
	s.keyFile = keyFile
}

// Start game gate service
func (s *Service) Start() {
	var err error
	if s.certFile != "" && s.keyFile != "" {
		err = s.server.ListenAndServeTLS(s.certFile, s.keyFile)
	} else {
		err = s.server.ListenAndServe()
	}
	if err != http.ErrServerClosed {
		logger.Sugar.Fatalf("gate service ListenAndServe error: %v", err)
	}
}

// Stop game gate service
func (s *Service) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		logger.Sugar.Errorf("gate service Shutdown error: %v", err)
	}
	close(closed)
}
