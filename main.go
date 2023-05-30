package main

import (
	"lion/cmd"
	"lion/internal/interfaces/container"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Start(container.Setup())
}
