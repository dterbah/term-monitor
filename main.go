package main

import (
	"flag"

	"github.com/dterbah/term-monitor/core"
)

func main() {
	// flags for the CLI
	showRAM := flag.Bool("ram", false, "Display usage of RAM")
	showCPU := flag.Bool("cpu", false, "Display usage of CPU")
	flag.Parse()

	core.RunCLI(core.CLIArgs{
		ShowRAM: *showRAM,
		ShowCPU: *showCPU,
	})
}
