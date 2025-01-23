package sliceskit_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/umefy/godash/sliceskit"
)

type FilterSuite struct {
	suite.Suite
}

// Filter should return nil when input slice is nil
func (s *FilterSuite) TestFilter_NilSlice() {
	result := sliceskit.Filter[[]int](nil, func(e int) bool { return e%2 == 0 })
	s.Nil(result)
}

// Filter should return filtered slice
func (s *FilterSuite) TestFilter_FilteredSlice() {
	slice := []int{1, 2, 3, 4, 5}
	result := sliceskit.Filter(slice, func(e int) bool { return e%2 == 0 })
	s.Equal([]int{2, 4}, result)
}

// FilterWithIndex should return nil when input slice is nil
func (s *FilterSuite) TestFilterWithIndex_NilSlice() {
	result := sliceskit.FilterWithIndex[[]int](nil, func(_ int, i int) bool { return i%2 == 0 })
	s.Nil(result)
}

// FilterWithIndex should return filtered slice
func (s *FilterSuite) TestFilterWithIndex_FilteredSlice() {
	slice := []int{0, 1, 2}
	result := sliceskit.FilterWithIndex(slice, func(e int, i int) bool { return (e*i)%2 == 0 })
	s.Equal([]int{0, 2}, result)
}

// FilterWithFuncErr should return error when filter function return error
func (s *FilterSuite) TestFilterWithFuncErr_FilterFuncErr() {
	slice := []int{1, 2, 3}
	result, err := sliceskit.FilterWithFuncErr(slice, func(e int) (bool, error) {
		if e == 2 {
			return false, errors.New("error")
		}
		return e%2 == 0, nil
	})
	s.NotNil(err)
	s.Nil(result)
}

// FilterWithIndexAndFuncErr should return error when filter function return error
func (s *FilterSuite) TestFilterWithIndexAndFuncErr_FilterFuncErr() {
	slice := []int{1, 2, 3}
	result, err := sliceskit.FilterWithIndexAndFuncErr(slice, func(e int, i int) (bool, error) {
		if i == 2 {
			return false, errors.New("error")
		}
		return e%2 == 0, nil
	})
	s.NotNil(err)
	s.Nil(result)
}

func TestFilterSuite(t *testing.T) {
	suite.Run(t, new(FilterSuite))
}
