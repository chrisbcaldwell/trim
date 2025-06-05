package trim

import (
	"slices"

	"golang.org/x/exp/constraints"
)

// Trim cuts a slice of numbers by the proportions in q
func Trim[T number](slice []T, q quantiles) []T {
	slices.Sort(slice)
	length := len(slice)
	if length == 0 {
		return []T{}
	}

	// set default quantile behavior
	if q.High == 0 {
		q.High = q.Low
	}

	lowI := int(q.Low * float64(length))
	highI := length - int(q.High*float64(length)) - 1
	return slice[lowI:highI]

}

func TrimmedMean[T number](slice []T, q quantiles) T {
	return mean(Trim(slice, q))
}

type quantiles struct {
	Low  float64 // default will be 0 automatically
	High float64 // default will be Low first, then 0
}

type number interface {
	constraints.Integer | constraints.Float
}

func mean[T number](slice []T) T {
	sum := slice[0]
	if len(slice) == 1 {
		return sum
	}
	for i := 1; i < len(slice); i++ {
		sum += slice[i]
	}
	return sum / T(len(slice))
}
