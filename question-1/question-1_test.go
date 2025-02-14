package question_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuestion1(t *testing.T) {
	input := [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}
	expectedOutput := 237

	output := maxPathSum(input)
	assert.Equal(t, expectedOutput, output)
}
