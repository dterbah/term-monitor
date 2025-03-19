package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ProcessMock struct{}

func (mock ProcessMock) Count() int {
	return 123
}

func TestGetProcessesCount(t *testing.T) {
	assert := assert.New(t)
	mockProcess := ProcessMock{}

	result := GetProcessesCount(mockProcess)

	assert.Equal(123, result)
}
