package trim

import (
	"slices"

	"math"

	"golang.org/x/exp/constraints"
)

// Trim cuts a slice of numbers by the proportions in q
func Trim[T number](slice []T, q quantiles) []T {
	slices.Sort(slice)
	length := len(slice)

	// some conditions that return empty slice
	// empty slice will force float64 NaN returns for TrimmedMean
	if length == 0 {
		return []T{}
	}

	for _, n := range q.values() {
		if n < 0 || n > 1 {
			return []T{}
		}
	}

	if q.Low+q.High > 1 {
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

func TrimmedMean[T number](slice []T, q quantiles) float64 {
	return mean(Trim(slice, q))
}

type quantiles struct {
	Low  float64 // default will be 0 automatically
	High float64 // default will be Low first, then 0
}

func (q quantiles) values() []float64 {
	return []float64{q.Low, q.High}
}

type number interface {
	constraints.Integer | constraints.Float
}

func mean[T number](slice []T) float64 {
	if len(slice) == 0 {
		return math.NaN()
	}
	var sum float64

	for _, n := range slice {
		sum += float64(n)
	}
	return sum / float64(len(slice))
}
