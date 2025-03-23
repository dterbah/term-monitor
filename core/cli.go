package core

import (
	"log"
	"math"

	"github.com/dterbah/gods/list/arraylist"
	"github.com/dterbah/term-monitor/core/adapter"
	"github.com/dterbah/term-monitor/core/metric"
	ui "github.com/gizak/termui/v3"
)

type GraphInformation struct {
	title    string
	callback DataGenerator
}

var Graphs = []GraphInformation{
	{title: "Used RAM (in %)", callback: func() float64 {
		ram := adapter.GopsRAMAdapter{}
		return metric.GetUsedRamPercent(ram)
	}},
	{title: "Used CPU (in %)", callback: func() float64 {
		cpu := adapter.GopsCPUAdapter{}
		return metric.GetUsedCPUPercent(cpu)
	}},
	{title: "Ping (in ms)", callback: func() float64 {
		ping, _ := metric.GetPing()
		return ping
	}},
	{title: "Proccesse count", callback: func() float64 {
		proccess := adapter.GopsProcessAdapter{}
		return float64(metric.GetProcessesCount(proccess))
	}},
}

// Args passed to the program to show different graph
type CLIArgs struct {
	ShowRAM       bool
	ShowCPU       bool
	ShowPing      bool
	ShowProcesses bool
}

func RunCLI(args CLIArgs) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui : %s", err)
	}
	defer ui.Close()

	width, height, err := GetTerminalSize()
	if err != nil {
		log.Fatalf("cannot retrieve terminal size : %v", err)
		return
	}

	// Liste des graphiques activés
	flattenArgs := []bool{args.ShowRAM, args.ShowCPU, args.ShowPing, args.ShowProcesses}
	activeGraphs := arraylist.New(func(a, b bool) int { return 0 }, flattenArgs...).
		Filter(func(e bool) bool { return e }).Size()

	if activeGraphs == 0 {
		log.Println("No metrics selected. Exiting.")
		return
	}

	cols := 2
	rows := int(math.Ceil(float64(activeGraphs) / float64(cols)))

	halfWidth := width / cols
	graphHeight := height / rows
	currentRow := 0
	currentCol := 0

	for i, showGraph := range flattenArgs {
		if showGraph {
			x1 := currentCol * halfWidth
			x2 := x1 + halfWidth
			y1 := currentRow * graphHeight
			y2 := y1 + graphHeight

			graphInfo := Graphs[i]
			go DisplayGraph(GraphPosition{x1, x2, y1, y2}, graphInfo.callback, graphInfo.title)

			currentCol++
			if currentCol >= cols {
				currentCol = 0
				currentRow++
			}
		}
	}

	// Gestion des événements
	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			if e.ID == "q" || e.ID == "<C-c>" {
				return
			}
		}
	}
}
