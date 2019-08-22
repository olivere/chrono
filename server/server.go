package server

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server implements an HTTP server. It is set up as a http.Handler.
// Example:
//
//   srv := server.New(...)
//   httpSrv := &http.Server{
//     Addr: ":8080",
//     Handler: srv,
//   }
//   httpSrv.ListenAndServe()
//
type Server struct {
	logger log.Logger
	h      http.Handler
}

// ServerOption specifies a signature for configuring the server.
type ServerOption func(*Server)

// New creates and initializes a new server using the given options.
func New(options ...ServerOption) *Server {
	s := &Server{
		logger: log.NewNopLogger(),
	}
	for _, o := range options {
		o(s)
	}
	s.initRoutes()
	return s
}

// WithLogger allows to configure the logger being used.
// By default, the server will not log anything.
func WithLogger(logger log.Logger) ServerOption {
	return func(s *Server) {
		s.logger = logger
	}
}

func (s *Server) initRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler(s.logger))
	s.h = handlers.ProxyHeaders(r)
}

// ServeHTTP handles a HTTP request.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.h.ServeHTTP(w, r)
}
