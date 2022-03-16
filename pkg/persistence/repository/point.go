package repository

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"ms/spatial/pkg/common"
	"ms/spatial/pkg/persistence/entity"
)

var (
	ErrPointLoadFile  = errors.New("failed to load points files")
	ErrPointUnmarshal = errors.New("failed to load structure")
)

type (
	Points interface {
		FindAll() ([]*entity.Point, error)
	}

	Point struct {
		logger common.Logger
		file   io.Reader
	}
)

func NewPoint(logger common.Logger, file io.Reader) *Point {
	return &Point{
		logger: logger,
		file:   file,
	}
}

func (p *Point) FindAll() ([]*entity.Point, error) {
	payload, err := ioutil.ReadAll(p.file)
	if err != nil {
		p.logger.Errorf("ioutil.ReadAll failed with %s\n", err)
		return nil, ErrPointLoadFile
	}

	var points []*entity.Point
	if err = json.Unmarshal(payload, &points); err != nil {
		p.logger.Errorf("json.Unmarshal failed with %s\n", err)
		return nil, ErrPointUnmarshal
	}

	return points, nil
}
