package main

import (
	"context"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/zlobste/spotter/internal/config"
	"github.com/zlobste/spotter/internal/services/spotter"
	"os"
)

func main() {
	defer func() {
		if rvr := recover(); rvr != nil {
			logrus.Error("app panicked\n", rvr)
		}
	}()

	app := initApp()
	if err := app.Run(os.Args); err != nil {
		logrus.New().WithError(err).Fatal("app failed")
	}
}

func initApp() *cli.App {
	return &cli.App{
		Usage:  "run spotter service",
		Action: run,
		Commands: []cli.Command{
			{
				Name:  "migrate",
				Usage: "run db migration",
				Subcommands: []cli.Command{
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

func run(*cli.Context) error {
	cfg := config.New(os.Getenv("CONFIG"))

	go func() {
		spotterSvc := spotter.New(cfg)
		if err := spotterSvc.Run(context.TODO()); err != nil {
			panic(errors.Wrapf(err, "failed to run spotter service"))
		}
	}()

	return nil
}

func migrateCmd(dir migrate.MigrationDirection) func(ctx *cli.Context) error {
/*	return func(c *cli.Context) error {
		cfg := config.New(os.Getenv("CONFIG"))

		source := migrate.PackrMigrationSource{
			Box: assets.Migrations,
		}

		count, err := migrate.Exec(cfg.RawDB(), "postgres", source, dir)
		if err != nil {
			return errors.Wrap(err, "failed to run migrations")
		}

		cfg.Logger().WithField("count", count).
			Info("applied migrations")

		return nil
	}*/
	return nil
}
