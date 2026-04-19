package rest

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/tuanta7/chasingpaper/internal/transport/rest/handler"
	"github.com/tuanta7/chasingpaper/internal/transport/rest/middleware"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

type Server struct {
	server      *http.Server
	router      chi.Router
	planHandler *handler.PlanHandler
	meter       metric.Meter
}

func NewServer(
	addr string,
	planHandler *handler.PlanHandler,
) *Server {
	router := chi.NewRouter()

	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
		router:      router,
		planHandler: planHandler,
		meter:       otel.Meter("rest_server_meter"),
	}
}

func (s *Server) Run() error {
	if err := middleware.InitMetricsMiddleware(s.meter); err != nil {
		return err
	}

	s.registerRoutes()

	log.Printf("Server is running on %s\n", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Timeout() time.Duration {
	return 20 * time.Second
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) registerRoutes() {
	s.router.Route("/api/v1/plans", func(r chi.Router) {
		r.Use(middleware.WithMetric)

		r.With(middleware.Pagination).Get("/", s.planHandler.ListPlans)
		r.Post("/", s.planHandler.CreatePlan)
		r.Get("/{id}", s.planHandler.GetPlanByID)
		r.Put("/{id}", s.planHandler.UpdatePlan)
		r.Delete("/{id}", s.planHandler.DeletePlan)
	})
}
