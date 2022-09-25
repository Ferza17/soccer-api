package graphql

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	chim "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	l "github.com/treastech/logger"
	"go.uber.org/zap"

	"github.com/Ferza17/soccer-api/middleware"
	"github.com/Ferza17/soccer-api/model/graph/generated"
	"github.com/Ferza17/soccer-api/model/graph/resolver"
	"github.com/Ferza17/soccer-api/repository"
)

type (
	Server struct {
		codename      string
		host          string
		port          string
		logger        *zap.Logger
		playersDB     repository.Players
		teamsDB       repository.Teams
		router        *chi.Mux
		httpServer    *http.Server
		graphQLServer *handler.Server
	}
	Option func(s *Server)
)

func NewServer(codename, host, address string, option ...Option) *Server {
	s := &Server{
		codename: codename,
		host:     host,
		port:     address,
	}
	for _, o := range option {
		o(s)
	}
	s.setup()
	return s
}

func (srv *Server) Serve() {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		srv.logger.Info(fmt.Sprintf("%s %s", method, route))
		return nil
	}
	if err := chi.Walk(srv.router, walkFunc); err != nil {
		log.Panicln(errors.Cause(err))
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", srv.host, srv.port), srv.router))
}

func (srv *Server) setup() {
	c := generated.Config{
		Resolvers: &resolver.Resolver{},
	}
	gqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	srv.graphQLServer = gqlServer

	// HTTP Server
	srv.router = srv.routes()
	srv.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", srv.host, srv.port),
		Handler: srv.router,
	}
}

func (srv *Server) routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "CONNECT", "TRACE", "HEAD", "PATCH"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}),
		chim.RequestID,
		chim.RealIP,
		chim.Recoverer,
		chim.NoCache,
		render.SetContentType(render.ContentTypeJSON),
		l.Logger(srv.logger),
		middleware.Host(srv.codename),
		middleware.RegisterHeaderHTTPContext,
		middleware.RegisterTeamsDBHTTPContext(srv.teamsDB),
		middleware.RegisterPlayersDBHTTPContext(srv.playersDB),
		chim.Heartbeat("/ping"),
	)
	routes(r, srv.graphQLServer)
	return r
}
