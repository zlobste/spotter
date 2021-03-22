package cmd

import (
	"context"
	"github.com/alecthomas/kingpin"
	"github.com/zlobste/spotter/internal/config"
	"github.com/zlobste/spotter/internal/services/spotter"
	"os"
)

func Run(args []string) bool {
	cfg := config.NewConfig(os.Getenv("CONFIG"))
	log := cfg.Logger()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.Error("app panicked", rvr)
		}
	}()

	app := kingpin.New("spotter-svc", "")

	runCmd := app.Command("run", "run command")
	spotterService := runCmd.Command("spotter", "run a spotter service")

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	switch cmd {
	case spotterService.FullCommand():
		svc := spotter.New(cfg)
		err := svc.Run(context.Background())
		if err != nil {
			log.WithError(err).Error("failed to run spotter")
			return false
		}
	default:
		log.WithField("command", cmd).Error("Unknown command")
		return false
	}

	return true
}
