package main

import (
	"fmt"
	"log/slog"
)

func Lol() {
	fmt.Println("walking mmcmcmcmmc")
}

func main() {
	slog.SetLogLoggerLevel(10)
	GoRebuildUrself()

	RunCmd("go", "build", "urmom")
}
