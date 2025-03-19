package adapter

import "github.com/shirou/gopsutil/mem"

type GopsRAMAdapter struct{}

func (g GopsRAMAdapter) Percent() float64 {
	v, _ := mem.VirtualMemory()
	return v.UsedPercent
}
