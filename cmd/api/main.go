package main

import (
	"log/slog"

	"github.com/mottajunior/glauber-io/internal/config"
	"github.com/mottajunior/glauber-io/internal/handler"
	"github.com/mottajunior/glauber-io/internal/router"
)

func main() {

	//Initialize  Configs [Database]
	err := config.Init()
	if err != nil {
		slog.Error("error when initialize configurations:", err)
		return
	}

	//Initialize  Handlers Dependencies
	handler.Init()

	//Initialize Server And Routes
	router.Init()
}
