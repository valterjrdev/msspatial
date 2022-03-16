package contract

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	Point struct {
		X        int `json:"x" query:"x"`
		Y        int `json:"y" query:"y"`
		Distance int `json:"distance" query:"distance"`
	}

	PointCollection []Point
)

func (a Point) Validate() error {
	return validation.ValidateStruct(
		&a,
		validation.Field(&a.X, validation.Required),
		validation.Field(&a.Y, validation.Required),
		validation.Field(&a.Distance, validation.Required),
	)
}

func (p PointCollection) Len() int {
	return len(p)
}

func (p PointCollection) Less(i, j int) bool {
	return p[i].Distance < p[j].Distance
}

func (p PointCollection) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
