package metric

type ProcessInterface interface {
	Count() int
}

func GetProcessesCount(process ProcessInterface) int {
	return process.Count()
}
