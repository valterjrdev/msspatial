package handler

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"ms/spatial/pkg/persistence/entity"
	"ms/spatial/pkg/persistence/repository"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHandlerPoints_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []struct {
		description string
		input       map[string]string
		points      []*entity.Point
		expected    string
	}{
		{
			description: "returns the points in order increasing",
			input: map[string]string{
				"x":        "2",
				"y":        "-3",
				"distance": "20",
			},
			points: []*entity.Point{
				{
					X: 2,
					Y: -8,
				},
				{
					X: 18,
					Y: -3,
				},
			},
			expected: `[{"x":2,"y":-8,"distance":5},{"x":18,"y":-3,"distance":16}]`,
		},
		{
			description: "returns the points collection empty",
			input: map[string]string{
				"x":        "2",
				"y":        "80",
				"distance": "20",
			},
			points: []*entity.Point{
				{
					X: 2,
					Y: -8,
				},
				{
					X: 18,
					Y: -3,
				},
			},
			expected: `[]`,
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.description, func(t *testing.T) {
			server := echo.New()
			queryString := make(url.Values)
			queryString.Set("x", tt.input["x"])
			queryString.Set("y", tt.input["y"])
			queryString.Set("distance", tt.input["distance"])
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/?%s", PointGet, queryString.Encode()), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			pointRepository := repository.NewMockPoints(ctrl)
			pointRepository.EXPECT().FindAll().Return(tt.points, nil)
			h := NewPoint(PointOpts{
				PointRepository: pointRepository,
			})
			if assert.NoError(t, h.Get(server.NewContext(req, rec))) {
				assert.Equal(t, http.StatusOK, rec.Code)
				assert.JSONEq(t, tt.expected, rec.Body.String())
			}
		})
	}
}

func TestHandlerPoints_Get_FindAllPoints_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	server := echo.New()
	queryString := make(url.Values)
	queryString.Set("x", "1")
	queryString.Set("y", "2")
	queryString.Set("distance", "20")
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/?%s", PointGet, queryString.Encode()), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	pointRepository := repository.NewMockPoints(ctrl)
	pointRepository.EXPECT().FindAll().Return(nil, repository.ErrPointLoadFile)
	h := NewPoint(PointOpts{
		PointRepository: pointRepository,
	})

	err := h.Get(server.NewContext(req, rec))
	assert.EqualError(t, err, "code=400, message=failed to load points files")
}

func TestHandlerPoints_Validate_Error(t *testing.T) {
	server := echo.New()
	queryString := make(url.Values)
	queryString.Set("x/", "")
	queryString.Set("y/", "")
	queryString.Set("distance", "")
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/?%s", PointGet, queryString.Encode()), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	h := NewPoint(PointOpts{PointRepository: nil})
	err := h.Get(server.NewContext(req, rec))
	assert.EqualError(t, err, "code=400, message=distance: cannot be blank; x: cannot be blank; y: cannot be blank.")
}
