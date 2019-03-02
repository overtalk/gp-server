package gate

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
)

var closed = make(chan struct{})

// Service describes game gate service
type Service struct {
	addr     string
	certFile string
	keyFile  string
	gin      *gin.Engine
	srv      *http.Server
	routeMap map[string]module.Router
}

// NewService creates a game gate service
func NewService(addr string) *Service {
	s := &Service{
		addr:     addr,
		gin:      gin.New(),
		routeMap: make(map[string]module.Router),
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
	logger.Sugar.Debugf("all registered router : %v", s.routeMap)

	// register all routes that have been registered to routeMap to the gate
	s.registerToGate()

	var err error

	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.gin,
	}

	s.srv = srv

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
