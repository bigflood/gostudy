package gocalc

import (
	"testing"
)

func BenchmarkCalcExpr(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		CalcExpr("123/4534+56")
	}
}

func TestCalcExpr(t *testing.T) {
	testCases := []struct {
		src      string
		expected float64
	}{
		{"123*2", 246},
		{"123/2", 61.5},
		{"123+2", 125},
	}

	for i, tc := range testCases {
		actual, err := CalcExpr(tc.src)
		if err != nil {
			t.Errorf("%d: err should be nil. but %v", i, err)
			continue
		}

		if actual != tc.expected {
			t.Errorf("%d: %v expected. but %v", i, tc.expected, actual)
		}
	}
}
