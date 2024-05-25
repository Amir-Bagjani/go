package main

import (
	"fmt"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(10, 20)
	if ans != 10 {
		t.Errorf("IntMin(10,20) = %d but want 20", ans)
	}
}

type IntMinTest struct {
	n1   int
	n2   int
	want int
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []IntMinTest{
		{n1: 10, n2: 20, want: 10},
		{n1: 15, n2: 8, want: 8},
		{n1: -15, n2: 0, want: -15},
		{n1: -30, n2: -33, want: -33},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("IntMin(%d,%d)", tt.n1, tt.n2)

		t.Run(testName, func(t *testing.T) {
			ans := IntMin(tt.n1, tt.n2)

			if ans != tt.want {
				t.Errorf("got %d but want %d", ans, tt.want)
			}
		})
	}
}
