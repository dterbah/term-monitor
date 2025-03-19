package metric

import (
	"time"
)

type CPUInterface interface {
	Percent(duration time.Duration, percpu bool) ([]float64, error)
}

func GetUsedCPUPercent(cpu CPUInterface) float64 {
	cpuPercent, _ := cpu.Percent(time.Second, false)
	return cpuPercent[0]
}
