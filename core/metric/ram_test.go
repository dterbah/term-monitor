package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type RAMMock struct{}

func (mock RAMMock) Percent() float64 {
	return 42.0
}

func TestGetUsedRAM(t *testing.T) {
	assert := assert.New(t)
	mockRAM := RAMMock{}

	result := GetUsedRamPercent(mockRAM)

	assert.Equal(42.0, result)
}
