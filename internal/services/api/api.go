package api

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zlobste/spotter/internal/config"
	"github.com/zlobste/spotter/internal/context"
	"github.com/zlobste/spotter/internal/data/postgres"
	"github.com/zlobste/spotter/internal/services/api/handlers"
	"github.com/zlobste/spotter/internal/services/api/middlewares"
	"github.com/zlobste/spotter/internal/services/timer"
	"net/http"
)

type API interface {
	Run() error
}

type api struct {
	log    *logrus.Logger
	config config.Config
	db     *sql.DB
}

func New(cfg config.Config) API {
	return &api{
		config: cfg,
		log:    cfg.Logging(),
		db:     cfg.DB(),
	}
}

func (a *api) Run() error {
	defer func() {
		if rvr := recover(); rvr != nil {
			a.log.Error("internal panicked\n", rvr)
		}
	}()

	go func() {
		if err := timer.New(a.config).Run(); err != nil {
			panic(errors.Wrap(err, "failed to start timer"))
		}
	}()

	a.log.WithField("port", a.config.Listener()).Info("Starting server")
	if err := http.ListenAndServe(a.config.Listener(), a.router()); err != nil {
		return errors.Wrap(err, "listener failed")
	}
	return nil
}

func (a *api) router() chi.Router {
	router := chi.NewRouter()

	router.Use(
		middlewares.CorsMiddleware(),
		middlewares.LoggingMiddleware(a.log),
		middlewares.CtxMiddleware(
			context.CtxLog(a.log),
			context.CtxConfig(a.config),
			context.CtxUsers(postgres.NewUsersStorage(a.db)),
			context.CtxGroups(postgres.NewGroupsStorage(a.db)),
		),
	)

	router.Route("/users", func(r chi.Router) {
		r.Post("/create", handlers.CreateUserHandler)
		r.Get("/{user_id}", handlers.GetUserHandler)
	})
	router.Route("/groups", func(r chi.Router) {
		r.Post("/create", handlers.CreateGroupHandler)
		r.Get("/{group_id}", handlers.GetGroupHandler)
		r.Get("/{group_id}/users", handlers.GetGroupUsersHandler)
		r.Post("/add_user", handlers.AddUserToGroupHandler)
	})
	/*	router.Route("/timer", func(r chi.Router) {
		r.Post("/create", handlers.CreateGroupHandler)
		r.Get("/{id}", handlers.GetGroupHandler)
		r.Post("/{id}/users", handlers.AddUserToGroupHandler)
	})*/

	return router
}
