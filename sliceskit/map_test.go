package sliceskit

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MapSuite struct {
	suite.Suite
}

// Map should return nil when input slice is nil
func (s *MapSuite) TestMapFunc_NilSlice() {
	result := Map[[]int](nil, func(e int) int { return e * 2 })
	s.Nil(result)
}

// Map should return nil when mapFunc is nil
func (s *MapSuite) TestMapFunc_NilMapFunc() {
	result := Map[[]int, int]([]int{1, 2, 3}, nil)
	s.Nil(result)
}

// Map should return mapped slice
func (s *MapSuite) TestMapFunc_MappedSlice() {
	slice := []int{1, 2, 3}
	result := Map(slice, func(e int) int { return e * 2 })
	s.Equal([]int{2, 4, 6}, result)
}

// MapWithIndex should return nil when input slice is nil
func (s *MapSuite) TestMapFuncIndex_NilSlice() {
	result := MapWithIndex[[]int](nil, func(e int, i int) int { return e * i })
	s.Nil(result)
}

// MapWithIndex should return nil when mapFunc is nil
func (s *MapSuite) TestMapFuncIndex_NilMapFunc() {
	result := MapWithIndex[[]int, int]([]int{1, 2, 3}, nil)
	s.Nil(result)
}

// MapWithIndex should return mapped slice
func (s *MapSuite) TestMapFuncIndex_MappedSlice() {
	slice := []int{1, 2, 3}
	result := MapWithIndex(slice, func(e int, i int) int { return e * i })
	s.Equal([]int{0, 2, 6}, result)
}

func TestMapSuite(t *testing.T) {
	suite.Run(t, new(MapSuite))
}
