package trim

import (
	"slices"

	"math"

	"golang.org/x/exp/constraints"
)

// Trim cuts a sorted slice of numbers by the proportions in q
// First argument of q is lower quantile, second is upper.  Beyond two get ignored.
// If zero quantiles are passed no trimming occurs.
// If one quantile is passed both ends are trimmed by that quantile.
func Trim[T number](slice []T, quants ...float64) []T {
	slices.Sort(slice)
	length := len(slice)

	q := quantiles{}
	switch n := len(quants); n {
	case 0: // skip; no need to update q
	case 1:
		q.Low = quants[0]
	default:
		q.Low = quants[0]
		q.High = quants[1]
	}

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

	lowTrim := int(q.Low * float64(length))
	lowIndex := lowTrim
	highTrim := int(q.High * float64(length))
	highIndex := length - highTrim
	return slice[lowIndex:highIndex]

}

func TrimmedMean[T number](slice []T, q ...float64) float64 {
	return mean(Trim(slice, q...))
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
