package httpserver

import (
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	middlewareOpenAPI "github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

// Server represents our http server.
type Server struct {
	addr        string
	logger      *log.Logger
	ibanService ibanService
	ln          net.Listener
}

// NewServer creates a new instance of the http server that will listen on addr.
func NewServer(addr string, ibanService ibanService, logger *log.Logger) *Server {
	return &Server{
		addr:        addr,
		ibanService: ibanService,
		logger:      logger,
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

	// Attach middlewares.
	r.Use(
		redocMiddleware,
	)

	r.Route("/api", func(r chi.Router) {
		r.Route("/validate/iban", newIBANHandler(s.ibanService, s.logger).Routes)
	})

	fs := http.FileServer(http.Dir("./docs"))
	r.Mount("/api-spec/", http.StripPrefix("/api-spec/", fs))

	return r
}

// redocMiddleware will host a ReDoc from the generated OpenAPI specification.
func redocMiddleware(next http.Handler) http.Handler {
	return middlewareOpenAPI.Redoc(middlewareOpenAPI.RedocOpts{
		BasePath: "/",
		Path:     "docs",
		SpecURL:  "./api-spec/swagger.yaml",
		Title:    "IBAN validation Service API Documentation",
	}, next)
}
