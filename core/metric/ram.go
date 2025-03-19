package metric

import "github.com/shirou/gopsutil/v3/mem"

func GetUsedRamPercent() float64 {
	v, _ := mem.VirtualMemory()
	return v.UsedPercent
}
