package test_test

import (
	"fmt"
	calc "golang/testing/calculator"
	"testing"
)

type testData struct {
	arg1, arg2, expected int
}

type testPowData struct {
	arg, expected int
}

func TestAdd(t *testing.T) {
	cases := []testData{
		{2, 3, 5},
		{10, 5, 15},
		{-8, -3, -11},
	}
	for _, it := range cases {
		t.Run(
			fmt.Sprintf("%d+%d=%d", it.arg1, it.arg2, it.expected), // имя подтеста
			func(t *testing.T) { // подтест
				result := calc.Add(it.arg1, it.arg2)
				if result != it.expected {
					t.Errorf("result %d, expected %d", result, it.expected)
				}
			},
		)
	}
}

func TestSub(t *testing.T) {
	cases := []testData{
		{2, 3, -1},
		{10, 5, 5},
		{-8, -3, -5},
	}
	for _, it := range cases {
		result := calc.Sub(it.arg1, it.arg2)
		if result != it.expected {
			t.Errorf("result %d, expected %d", result, it.expected)
		}
	}
}

func TestMul(t *testing.T) {
	cases := []testData{
		{2, 3, 6},
		{10, 5, 50},
		{-8, -3, 24},
	}
	for _, it := range cases {
		result := calc.Mul(it.arg1, it.arg2)
		if result != it.expected {
			t.Errorf("result %d, expected %d", result, it.expected)
		}
	}
}

func TestPow2(t *testing.T) {
	cases := []testPowData{
		{2, 4},
		{-2, 4},
		{3, 9},
	}
	for _, it := range cases {
		result := calc.Pow2(it.arg)
		if result != it.expected {
			t.Errorf("result %d, expected %d", result, it.expected)
		}
	}
}
