package metric

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Mock functions

type CPUMock struct{}

func (mock CPUMock) Percent(duration time.Duration, perCPU bool) ([]float64, error) {
	return []float64{42.0}, nil
}

func TestGetUsedCPUPercent(t *testing.T) {
	assert := assert.New(t)
	mockCPU := CPUMock{}

	result := GetUsedCPUPercent(mockCPU)

	assert.Equal(42.0, result)
}
