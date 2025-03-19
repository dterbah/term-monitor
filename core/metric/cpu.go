package metric

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func GetUsedCPUPercent() float64 {
	cpuPercent, _ := cpu.Percent(time.Second, false)
	return cpuPercent[0]
}
