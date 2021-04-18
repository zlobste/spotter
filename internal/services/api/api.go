package api

import (
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zlobste/spotter/internal/config"
	"github.com/zlobste/spotter/internal/context"
	"github.com/zlobste/spotter/internal/data/postgres"
	"github.com/zlobste/spotter/internal/services/api/handlers"
	"github.com/zlobste/spotter/internal/services/api/middlewares"
	"net/http"
)

type App interface {
	Run() error
}

type app struct {
	log    *logrus.Logger
	config config.Config
}

func New(cfg config.Config) App {
	return &app{
		config: cfg,
		log:    cfg.Logging(),
	}
}

func (a *app) Run() error {
	defer func() {
		if rvr := recover(); rvr != nil {
			a.log.Error("internal panicked\n", rvr)
		}
	}()

	a.log.WithField("port", a.config.Listener()).Info("Starting server")
	if err := http.ListenAndServe(a.config.Listener(), a.router()); err != nil {
		return errors.Wrap(err, "listener failed")
	}
	return nil
}

func (a *app) router() chi.Router {
	router := chi.NewRouter()

	router.Use(
		middlewares.CorsMiddleware(),
		middlewares.LoggingMiddleware(a.log),
		middlewares.CtxMiddleware(
			context.CtxLog(a.log),
			context.CtxConfig(a.config),
			context.CtxUsers(postgres.NewUsersStorage(a.config)),
		),
	)

	router.Route("/users", func(r chi.Router) {
		r.Post("/", handlers.CreateUser)
		r.Get("/{id}", handlers.GetUser)
	})

	return router
}
