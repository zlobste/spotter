package context

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/zlobste/spotter/app/config"
	"github.com/zlobste/spotter/app/data/postgres"
	"net/http"
)

const (
	ctxLog    = "ctxLog"
	ctxConfig = "ctxConfig"
	ctxUsers  = "ctxUsers"
)

func CtxConfig(cfg config.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxConfig, cfg)
	}
}

func Config(r *http.Request) config.Config {
	return r.Context().Value(ctxConfig).(config.Config)
}

func CtxLog(log *logrus.Logger) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxLog, log)
	}
}

func Log(r *http.Request) *logrus.Logger {
	return r.Context().Value(ctxLog).(*logrus.Logger)
}

func CtxUsers(users postgres.UsersStorage) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxUsers, users)
	}
}

func Users(r *http.Request) postgres.UsersStorage {
	return r.Context().Value(ctxUsers).(postgres.UsersStorage).New()
}