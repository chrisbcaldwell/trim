package trim

import (
	"math"
	"slices"
	"testing"
)

// Testing data

// len = 11 so the breaks are weird, out of order, some repeated
var testInt = []int{-25, 56, 55, 56, 0, -1, 10, 1, 1, 5, 3}
var testFloat = []float64{-25.1, 56.9, 55.3, 56.9, 0, -1, 10.43, 1.112, 1.112, 5, 3}

func TestTrim(t *testing.T) {
	// initialize full, full unequal, half, and empty quantiles structs
	quants := make([]quantiles, 4)
	quants[0] = quantiles{0.1, 0.1}
	quants[1] = quantiles{0.25, 0.1}
	quants[2].Low = 0.1
	// q[3] remains empty (zero values)

	resInt := make([][]int, 4)
	resFloat := make([][]float64, 4)
	for i, q := range quants {
		resInt[i] = Trim(testInt, q)
		resFloat[i] = Trim(testFloat, q)
		slices.Sort(resInt[i])
		slices.Sort(resFloat[i])
	}

	exInt := make([][]int, 4)
	exInt[0] = []int{55, 56, 0, -1, 10, 1, 1, 5, 3}
	exInt[1] = []int{55, 56, 0, 10, 1, 1, 5, 3}
	exInt[2] = exInt[0]
	exInt[3] = testInt

	exFloat := make([][]float64, 4)
	exFloat[0] = []float64{55.3, 56.9, 0, -1, 10.43, 1.112, 1.112, 5, 3}
	exFloat[1] = []float64{55.3, 56.9, 0, 10.43, 1.112, 1.112, 5, 3}
	exFloat[2] = exFloat[0]
	exFloat[3] = testFloat

	for i, s := range exInt {
		slices.Sort(s)
		slices.Sort(exFloat[i])
	}

	for i, r := range resInt {
		if !slices.Equal(exInt[i], r) {
			t.Errorf("Integer slice: Trim(%v) = %v; expected: %v", testInt[i], r, exInt[i])
		}
	}

	for i, r := range resFloat {
		if !slices.Equal(exFloat[i], r) {
			t.Errorf("Float slice: Trim(%v) = %v; expected: %v", testFloat[i], r, exFloat[i])
		}
	}

}

func TestTrimmedMean(t *testing.T) {
	precision := 0.00000001

	// initialize full, full unequal, half, and empty quantiles structs
	quants := make([]quantiles, 4)
	quants[0] = quantiles{0.1, 0.1}
	quants[1] = quantiles{0.25, 0.1}
	quants[2].Low = 0.1
	// q[3] remains empty (zero values)

	resInt := make([]float64, 4)
	resFloat := make([]float64, 4)
	for i, q := range quants {
		resInt[i] = TrimmedMean(testInt, q)
		resFloat[i] = TrimmedMean(testFloat, q)
	}

	exInt := []float64{14.4444444444, 16.375, 14.4444444444, 14.6363636364}
	exFloat := []float64{14.6504444444, 16.60675, 14.6504444444, 14.8776363636}

	for i, r := range resInt {
		diff := r - exInt[i]
		if math.Abs(diff) > precision {
			t.Errorf("Integer slice: TrimmedMean(%v) = %v; expected: %v", testInt[i], r, exInt[i])
		}
	}

	for i, r := range resFloat {
		diff := r - exFloat[i]
		if math.Abs(diff) > precision {
			t.Errorf("Float slice: TrimmedMean(%v) = %v; expected: %v", testFloat[i], r, exFloat[i])
		}
	}

}
