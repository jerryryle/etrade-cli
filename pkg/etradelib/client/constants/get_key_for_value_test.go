package constants

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetKeyForValue(t *testing.T) {
	testMap := map[int]string{
		1: "A",
		2: "B",
	}

	result, err := getKeyForValue(testMap, "A")
	assert.Nil(t, err)
	assert.Equal(t, 1, result)

	result, err = getKeyForValue(testMap, "C")
	assert.Error(t, err)
}
