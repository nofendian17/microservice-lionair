package cmd

import (
	"lion/internal/interfaces/container"
	"lion/internal/interfaces/delivery/http"
)

func Start(container *container.Container) {
	http.Start(container)
}
