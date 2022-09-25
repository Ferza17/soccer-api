package main

import (
	"runtime"

	"github.com/Ferza17/soccer-api/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Run()
}
