package handler

import (
	"github.com/labstack/echo/v4"
	"ms/spatial/pkg/contract"
	"ms/spatial/pkg/mathc"
	"ms/spatial/pkg/persistence/repository"
	"net/http"
	"sort"
)

const (
	PointGet = "/api/points"
)

type (
	PointOpts struct {
		PointRepository repository.Points
	}
	Point struct {
		PointOpts
	}
)

func NewPoint(opts PointOpts) *Point {
	return &Point{opts}
}

func (a *Point) Get(c echo.Context) error {
	request := &contract.PointRequest{}
	_ = c.Bind(request)
	if err := request.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := a.PointRepository.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	pointResponse := make([]contract.PointResponse, 0)
	for _, point := range data {
		v1 := mathc.Vector{request.X, request.Y}
		v2 := mathc.Vector{point.X, point.Y}

		distance := mathc.ManhattanDistance2D(v1, v2)
		if request.Distance >= distance {
			pointResponse = append(pointResponse, contract.PointResponse{
				X:        point.X,
				Y:        point.Y,
				Distance: distance,
			})
		}
	}

	collection := contract.PointCollectionResponse(pointResponse)
	sort.Sort(collection)
	return c.JSON(http.StatusOK, collection)
}
