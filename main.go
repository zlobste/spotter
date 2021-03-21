package main

import (
	"github.com/zlobste/spotter/internal/cli"
	"os"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
