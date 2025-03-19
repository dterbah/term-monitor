package metric

type RAMInterface interface {
	Percent() float64
}

func GetUsedRamPercent(ram RAMInterface) float64 {
	return ram.Percent()
}
