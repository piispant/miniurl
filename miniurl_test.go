package miniurl_test

import (
	"testing"

	"github.com/piispant/miniurl"
	"github.com/stretchr/testify/assert"
)

func TestHashLength(t *testing.T) {
	const (
		input          = "https://github.com/piispant/miniurl"
		expectedLength = 32
	)

	output := miniurl.Hash(input)

	assert.Len(t, output, expectedLength)
}

func TestHashIsDeterministic(t *testing.T) {
	const input = "https://github.com/piispant/miniurl"

	output1 := miniurl.Hash(input)
	output2 := miniurl.Hash(input)

	assert.Equal(t, output1, output2)
}
