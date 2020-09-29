package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.startLogger(); err != nil {
		return err
	}
	s.logger.Info("server is online")
	s.setupRouter()
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) startLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLev)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil

}

func (s *APIServer) setupRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "hello")
	}
}
