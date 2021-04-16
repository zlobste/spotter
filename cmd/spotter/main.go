package spotter

import (
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/zlobste/spotter/internal/assets"
	"github.com/zlobste/spotter/internal/config"
	"github.com/zlobste/spotter/internal/services/api"
	"os"
)

func main() {
	defer func() {
		if rvr := recover(); rvr != nil {
			logrus.New().Error("internal panicked\n", rvr)
		}
	}()

	app := initApp()
	if err := app.Run(os.Args); err != nil {
		logrus.New().WithError(err).Fatal("internal failed")
	}
}

func run(*cli.Context) error {
	cfg := config.New(os.Getenv("CONFIG"))

	srv := api.New(cfg)
	if err := srv.Run(); err != nil {
		return errors.Wrap(err, "failed to start api")
	}

	return nil
}

func migrateCmd(dir migrate.MigrationDirection) func(ctx *cli.Context) error {
	return func(c *cli.Context) error {
		cfg := config.New(os.Getenv("CONFIG"))

		source := migrate.PackrMigrationSource{
			Box: assets.Migrations,
		}

		count, err := migrate.Exec(cfg.DB(), "postgres", source, dir)
		if err != nil {
			return errors.Wrap(err, "failed to run migrations")
		}

		cfg.Logging().WithField("count", count).
			Info("applied migrations")

		return nil
	}
}

func initApp() *cli.App {
	return &cli.App{
		Usage:  "run spotter service",
		Action: run,
		Commands: []*cli.Command{
			{
				Name:  "migrate",
				Usage: "run db migration",
				Subcommands: []*cli.Command{
					{
						Name:   "up",
						Usage:  "apply new migrations",
						Action: migrateCmd(migrate.Up),
					},
					{
						Name:   "down",
						Usage:  "rollback migrations",
						Action: migrateCmd(migrate.Down),
					},
				},
			},
		},
	}
}
