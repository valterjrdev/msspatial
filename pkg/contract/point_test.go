package contract

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestPointCollection_Len(t *testing.T) {
	input := PointCollectionResponse([]PointResponse{
		{
			X:        10,
			Y:        -3,
			Distance: 30,
		},
		{
			X:        10,
			Y:        -3,
			Distance: 20,
		},
	})

	expected := PointCollectionResponse([]PointResponse{
		{
			X:        10,
			Y:        -3,
			Distance: 20,
		},
		{
			X:        10,
			Y:        -3,
			Distance: 30,
		},
	})

	sort.Sort(input)
	assert.Equal(t, expected, input)
	assert.Equal(t, 2, input.Len())
}

func TestPoint_Validate(t *testing.T) {
	point := PointRequest{}
	assert.EqualError(t, point.Validate(), "distance: cannot be blank; x: cannot be blank; y: cannot be blank.")
}
