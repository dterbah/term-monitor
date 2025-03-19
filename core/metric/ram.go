package metric

import "github.com/shirou/gopsutil/v3/mem"

type RAMInterface interface {
	Percent() float64
}

func GetUsedRamPercent() float64 {
	v, _ := mem.VirtualMemory()
	return v.UsedPercent
}
