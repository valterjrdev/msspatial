package repository

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"ms/spatial/pkg/common"
	"ms/spatial/pkg/persistence/entity"
	"ms/spatial/tests"
	"testing"
)

func TestPoint_FindAll(t *testing.T) {
	expected := []*entity.Point{
		{
			X: -8,
			Y: 10,
		},
		{
			X: 15,
			Y: -10,
		},
	}

	data, err := json.Marshal(expected)
	assert.Nil(t, err)

	payload := bytes.NewReader(data)
	pointRepository := NewPoint(nil, payload)
	points, err := pointRepository.FindAll()
	assert.Nil(t, err)
	assert.Len(t, points, 2)
	assert.Equal(t, expected, points)
}

func TestPoint_ReadAll_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := common.NewMockLogger(ctrl)
	logger.EXPECT().Errorf(gomock.Any(), gomock.Any())

	pointRepository := NewPoint(logger, tests.Reader(0))
	points, err := pointRepository.FindAll()
	assert.Nil(t, points)
	assert.EqualError(t, err, ErrPointLoadFile.Error())
}

func TestPoint_FindAll_Unmarshal_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := common.NewMockLogger(ctrl)
	logger.EXPECT().Errorf(gomock.Any(), gomock.Any())

	payload := bytes.NewReader([]byte(`{`))
	pointRepository := NewPoint(logger, payload)
	points, err := pointRepository.FindAll()
	assert.Nil(t, points)
	assert.EqualError(t, err, ErrPointUnmarshal.Error())
}
