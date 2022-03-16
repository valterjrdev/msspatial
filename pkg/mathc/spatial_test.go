package mathc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManhattanDistance(t *testing.T) {
	cases := []struct {
		input    []Vector
		expected int
	}{
		{
			input: []Vector{
				{2, -3},
				{18, -3},
			},
			expected: 16,
		},
		{
			input: []Vector{
				{10, -5},
				{-30, -38},
			},
			expected: 73,
		},
		{
			input: []Vector{
				{-2, -8},
				{-20, -10},
			},
			expected: 20,
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			sum := ManhattanDistance2D(tt.input[0], tt.input[1])
			assert.Equal(t, tt.expected, sum)
		})
	}
}

func TestAbs(t *testing.T) {
	cases := []struct {
		input    int
		expected int
	}{
		{
			input:    1,
			expected: 1,
		},
		{
			input:    -1,
			expected: 1,
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tt.expected, Abs(tt.input))
		})
	}
}
