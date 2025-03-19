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

	// Must follow the same order as Graphs constant
	flattenArgs := []bool{args.ShowRAM, args.ShowCPU, args.ShowPing, args.ShowProcesses}

	defer ui.Close()

	width, height, error := GetTerminalSize()
	if error != nil {
		log.Fatalf("cannot retrieve terminal size : %v", error)
		return
	}

	/* We want to display 2 graphs by row */
	rows := arraylist.New(func(a, b bool) int {
		if a == b {
			return 0
		}

		return -1
	}, flattenArgs...).Filter(func(element bool) bool {
		return element
	}).Size()

	halfWidth := width / 2
	graphHeight := height / int((math.Ceil(float64(rows) / float64(2))))
	currentHeight := 0

	graphsDisplayed := 0

	for i, arg := range flattenArgs {
		x1 := (graphsDisplayed % 2) * halfWidth
		x2 := ((graphsDisplayed % 2) + 1) * halfWidth

		if arg {
			graphsDisplayed++
			if graphsDisplayed%2 == 1 && graphsDisplayed > 1 {
				currentHeight += graphHeight
			}

			y1 := currentHeight
			y2 := currentHeight + graphHeight

			graphInformation := Graphs[i]
			// display the associated graph
			go DisplayGraph(GraphPosition{
				x1: x1,
				x2: x2,
				y1: y1,
				y2: y2,
			}, graphInformation.callback, graphInformation.title)
		}
	}
	uiEvents := ui.PollEvents()

	// listen to user events
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		}
	}
}
