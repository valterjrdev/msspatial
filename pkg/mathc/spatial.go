package mathc

type (
	Vector []int
)

// ManhattanDistance2D returns the distance between vectors
// V1: (x1, y1)
// V2: (x2, y2)
// Formula: |x1-x2| + |y1-y2|
func ManhattanDistance2D(v1 Vector, v2 Vector) int {
	return Abs(v1[0]-v2[0]) + Abs(v1[1]-v2[1])
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
