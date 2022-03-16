package contract

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	PointRequest struct {
		X        int `json:"x" query:"x"`
		Y        int `json:"y" query:"y"`
		Distance int `json:"distance" query:"distance"`
	}

	PointResponse struct {
		X        int `json:"x"`
		Y        int `json:"y"`
		Distance int `json:"distance"`
	}

	PointCollectionResponse []PointResponse
)

func (a PointRequest) Validate() error {
	return validation.ValidateStruct(
		&a,
		validation.Field(&a.X, validation.Required),
		validation.Field(&a.Y, validation.Required),
		validation.Field(&a.Distance, validation.Required),
	)
}

func (p PointCollectionResponse) Len() int {
	return len(p)
}

func (p PointCollectionResponse) Less(i, j int) bool {
	return p[i].Distance < p[j].Distance
}

func (p PointCollectionResponse) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
