package cli

import (
	"github.com/alecthomas/kingpin"
	"github.com/zlobste/spotter/internal/config"
	"github.com/zlobste/spotter/internal/services/api"
	"os"
)

func Run(args []string) bool {
	cfg := config.New(os.Getenv("CONFIG"))
	log := cfg.Logging()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.Error("internal panicked", rvr)
		}
	}()

	app := kingpin.New("spotter", "")
	runCmd := app.Command("run", "run command")
	serviceCmd := runCmd.Command("service", "run service") // you can insert custom help

	migrateCmd := app.Command("migrate", "migrate command")
	migrateUpCmd := migrateCmd.Command("up", "migrate db up")
	migrateDownCmd := migrateCmd.Command("down", "migrate db down")

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	switch cmd {
	case serviceCmd.FullCommand():
		if err := api.New(cfg).Run(); err != nil {
			log.Error("failed to start api", err)
			return false
		}
	case migrateUpCmd.FullCommand():
		err = MigrateUp(cfg)
	case migrateDownCmd.FullCommand():
		err = MigrateDown(cfg)
	default:
		log.Errorf("unknown command %s", cmd)
		return false
	}

	if err != nil {
		log.WithError(err).Error("failed to exec cmd")
		return false
	}

	return true
}
