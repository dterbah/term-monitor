package main

import (
	"flag"
	"log"
	"os"

	"github.com/dterbah/term-monitor/core"
)

func main() {
	// flags for the CLI
	showRAM := flag.Bool("ram", false, "Display usage of RAM")
	showCPU := flag.Bool("cpu", false, "Display usage of CPU")
	showPing := flag.Bool("ping", false, "Display current ping")
	showProcess := flag.Bool("proc", false, "Display number of processes")

	flag.Parse()

	if !(*showRAM || *showCPU || *showPing || *showProcess) {
		log.Fatal("Error: you must specify at least one option (--ram, --cpu, --ping, --proc)")
		flag.Usage()
		os.Exit(1)
	}

	core.RunCLI(core.CLIArgs{
		ShowRAM:       *showRAM,
		ShowCPU:       *showCPU,
		ShowPing:      *showPing,
		ShowProcesses: *showProcess,
	})
}
