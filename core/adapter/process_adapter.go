package adapter

import "github.com/shirou/gopsutil/process"

type GopsProcessAdapter struct{}

func (g GopsProcessAdapter) Count() int {
	processes, _ := process.Processes()
	return len(processes)
}
