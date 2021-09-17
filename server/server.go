package server

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Server interface {
	Start(address string)
	Register(path string, handlerFunc http.HandlerFunc)
	Shutdown(ctx context.Context) error
	ShutdownWithCallBack(fn func())
	Stop() error
}

type server struct {
	*http.Server
	*http.ServeMux
	log *log.Logger
}

func New(log *log.Logger) Server {
	return &server{
		Server:   &http.Server{},
		ServeMux: http.NewServeMux(),
		log:      log,
	}
}

func (s *server) Start(address string) {
	s.Addr = address
	s.Server.Handler = s.ServeMux
	s.log.Debug(fmt.Sprintf("%s %s", "Server Started at", address))
	if err := s.ListenAndServe(); err != nil {
		s.log.Error(err.Error())
	}
}

func (s *server) Register(path string, handlerFunc http.HandlerFunc) {
	s.log.Info(fmt.Sprintf("Register with path=%s is registered", path))
	s.HandleFunc(path, handlerFunc)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.log.Debug(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
	s.ServeMux.ServeHTTP(w, r)
}

func (s *server) Shutdown(ctx context.Context) error {
	s.log.Info("http server Shutdown with context is registered")
	return s.Server.Shutdown(ctx)
}

func (s *server) ShutdownWithCallBack(fn func()) {
	s.log.Info("Shutdown With CallBack is registered")
	s.Server.RegisterOnShutdown(fn)
}

func (s *server) Stop() error {
	s.log.Info("Server is stopped")
	return s.Server.Close()
}
