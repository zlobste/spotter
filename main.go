package spotter

import (
	"github.com/zlobste/spotter/app"
	"github.com/zlobste/spotter/app/config"
	"os"
)

func main() {
	cfg := config.New(os.Getenv("CONFIG"))

	if err := app.New(cfg).Run(); err != nil {
		panic(err)
	}
}