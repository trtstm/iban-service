package http

import (
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

// Server represents our http server.
type Server struct {
	addr   string
	logger *log.Logger
	ln     net.Listener
}

// NewServer creates a new instance of the http server that will listen on addr.
func NewServer(addr string, logger *log.Logger) *Server {
	return &Server{
		addr:   addr,
		logger: logger,
	}
}

// Open will setup a tcp listener and start handling requests.
func (s *Server) Open() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return errors.Wrap(err, "could not start listening")
	}

	// Store listener so we can close it later.
	s.ln = ln

	// Start HTTP server.
	server := http.Server{
		Handler: s.Handler(),
	}

	return server.Serve(s.ln)
}

// Close will close the http server.
func (s *Server) Close() {
	if s.ln != nil {
		s.ln.Close()
	}
}

// Handler will set up all handlers/endpoints.
func (s *Server) Handler() http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Route("/validate/iban", newIBANHandler(s.logger).Routes)
	})

	return r
}
