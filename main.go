package main

import (
	"github.com/demonoid81/dsp/cmd"
	"os"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}