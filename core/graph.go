package core

import (
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type DataGenerator func() float64

type GraphPosition struct {
	x1 int
	x2 int
	y1 int
	y2 int
}

const MAX_SIZE = 20

/*
*
This function display a plot graph. Each second, it will call the
data generator to know the next point to calculate and display.
It should be called once the UI lib is initialized !
*/
func DisplayGraph(position GraphPosition, dataGenerator DataGenerator, title string) {
	plot := widgets.NewPlot()
	plot.Title = title
	plot.Data = [][]float64{{0}}
	plot.SetRect(position.x1, position.y1, position.x2, position.y2)
	plot.AxesColor = ui.ColorWhite
	plot.LineColors[0] = ui.ColorGreen

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			value := dataGenerator()
			plot.Data[0] = append(plot.Data[0], value)
			if len(plot.Data[0]) > 20 {
				plot.Data[0] = plot.Data[0][1:]
			}
			ui.Render(plot)
		}
	}
}
