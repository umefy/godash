package sliceskit_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/umefy/godash/sliceskit"
)

type MapSuite struct {
	suite.Suite
}

// Map should return nil when input slice is nil
func (s *MapSuite) TestMapFunc_NilSlice() {
	result := sliceskit.Map[[]int](nil, func(e int) int { return e * 2 })
	s.Nil(result)
}

// Map should return mapped slice
func (s *MapSuite) TestMapFunc_MappedSlice() {
	slice := []int{1, 2, 3}
	result := sliceskit.Map(slice, func(e int) int { return e * 2 })
	s.Equal([]int{2, 4, 6}, result)
}

// MapWithIndex should return nil when input slice is nil
func (s *MapSuite) TestMapFuncIndex_NilSlice() {
	result := sliceskit.MapWithIndex[[]int](nil, func(e int, i int) int { return e * i })
	s.Nil(result)
}

// MapWithIndex should return mapped slice
func (s *MapSuite) TestMapFuncIndex_MappedSlice() {
	slice := []int{1, 2, 3}
	result := sliceskit.MapWithIndex(slice, func(e int, i int) int { return e * i })
	s.Equal([]int{0, 2, 6}, result)
}

// MapWithFuncErr should return nil when input slice is nil
func (s *MapSuite) TestMapFuncErr_NilSlice() {
	result, err := sliceskit.MapWithFuncErr[[]int](nil, func(e int) (int, error) { return e * 2, nil })
	s.Nil(result)
	s.Nil(err)
}

// MapWithFuncErr should return error when Map func return error
func (s *MapSuite) TestMapFuncErr_MapFuncErr() {
	slice := []int{1, 2, 3}
	result, err := sliceskit.MapWithFuncErr(slice, func(e int) (int, error) {
		if e == 2 {
			return 0, errors.New("error")
		}
		return e * 2, nil
	})
	s.NotNil(err)
	s.Nil(result)
}

// MapWithIndexAndFuncErr should return nil when input slice is nil
func (s *MapSuite) TestMapFuncIndexErr_NilSlice() {
	result, err := sliceskit.MapWithIndexAndFuncErr[[]int](nil, func(e int, i int) (int, error) { return e * i, nil })
	s.Nil(result)
	s.Nil(err)
}

// MapWithIndexAndFuncErr should return error when Map func return error
func (s *MapSuite) TestMapFuncIndexErr_MapFuncIndexErr() {
	slice := []int{1, 2, 3}
	result, err := sliceskit.MapWithIndexAndFuncErr(slice, func(e int, i int) (int, error) {
		if i == 1 {
			return 0, errors.New("error")
		}
		return e * i, nil
	})
	s.NotNil(err)
	s.Nil(result)
}

func TestMapSuite(t *testing.T) {
	suite.Run(t, new(MapSuite))
}
