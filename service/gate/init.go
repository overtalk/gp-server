package gate

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/rs/cors"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
)

var closed = make(chan struct{})

// Service describes game gate service
type Service struct {
	addr     string
	certFile string
	keyFile  string
	srv      *http.Server
	routeMap sync.Map
}

// NewService creates a game gate service
func NewService(addr string) *Service {
	s := &Service{
		addr: addr,
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
	logger.Sugar.Debugf("all router:")
	s.routeMap.Range(func(k, v interface{}) bool {
		handler := v.(module.Router)
		logger.Sugar.Debugf("[%s]----[%v]", handler.Method, k)
		return true
	})

	// register all routes that have been registered to routeMap to the gate
	mux := http.NewServeMux()
	s.registerToGate(mux)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	srv := &http.Server{
		Addr:    s.addr,
		Handler: c.Handler(mux),
	}

	s.srv = srv
	var err error
	if s.certFile != "" && s.keyFile != "" {
		err = srv.ListenAndServeTLS(s.certFile, s.keyFile)
	} else {
		err = srv.ListenAndServe()
	}
	if err != http.ErrServerClosed {
		logger.Sugar.Fatalf("gate service ListenAndServe error: %v", err)
	}
}

// Stop game gate service
func (s *Service) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		logger.Sugar.Errorf("gate service Shutdown error: %v", err)
	}
	close(closed)
}
