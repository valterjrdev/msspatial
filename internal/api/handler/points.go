package handler

import (
	"github.com/labstack/echo/v4"
	"ms/spatial/pkg/contract"
	"ms/spatial/pkg/mathc"
	"net/http"
	"sort"
)

const (
	PointGet = "/api/points"
)

type (
	PointOpts struct {
		Points []contract.Point
	}
	Point struct {
		PointOpts
	}
)

func NewPoint(opts PointOpts) *Point {
	return &Point{opts}
}

func (a *Point) Get(c echo.Context) error {
	request := &contract.Point{}
	_ = c.Bind(request)
	if err := request.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	points := make([]contract.Point, 0)
	for _, point := range a.Points {
		v1 := mathc.Vector{request.X, request.Y}
		v2 := mathc.Vector{point.X, point.Y}

		distance := mathc.ManhattanDistance2D(v1, v2)
		if request.Distance >= distance {
			points = append(points, contract.Point{
				X:        point.X,
				Y:        point.Y,
				Distance: distance,
			})
		}
	}

	collection := contract.PointCollection(points)
	sort.Sort(collection)
	return c.JSON(http.StatusOK, collection)
}
