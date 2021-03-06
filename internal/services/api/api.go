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
			context.CtxTimers(postgres.NewTimersStorage(a.db)),
			context.CtxProofs(postgres.NewProofsStorage(a.db)),
		),
	)

	router.Group(func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", handlers.RegisterHandler)
			r.Post("/login", handlers.LoginHandler)
		})
	})

	router.Group(func(r chi.Router) {
		r.Use(middlewares.JWTMiddleware)

		r.Route("/users", func(r chi.Router) {
			r.Get("/{user_id}", handlers.GetUserHandler)
			r.Get("/drivers", handlers.GetAllDriversHandler)
			r.Get("/managers", handlers.GetAllManagersHandler)

			r.Post("/set_manager/{user_id}", handlers.SetManagerHandler)
			r.Post("/block/{user_id}", handlers.BlockUserHandler)
			r.Post("/unblock/{user_id}", handlers.UnblockUserHandler)
		})

		r.Route("/timers", func(r chi.Router) {
			r.Get("/drivers/{user_id}", handlers.GetTimersByDriverHandler)
			r.Get("/{timer_id}", handlers.GetTimerHandler)
			r.Get("/pending/{user_id}", handlers.GetPendingTimerHandler)
			r.Post("/start/{user_id}", handlers.StartTimerHandler)
			r.Post("/stop/{user_id}", handlers.StopTimerHandler)
		})

		r.Route("/proofs", func(r chi.Router) {
			r.Get("/{timer_id}", handlers.GetProofsByTimerHandler)
			r.Post("/", handlers.MakeProofHandler)
		})
	})

	return router
}
