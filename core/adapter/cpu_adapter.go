package adapter

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type GopsCPUAdapter struct{}

func (g GopsCPUAdapter) Percent(duration time.Duration, percpu bool) ([]float64, error) {
	return cpu.Percent(duration, percpu)
}
