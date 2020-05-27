package main

import (
	"testing"
)

func TestGetCombinations(t *testing.T) {
	result := GetCombinations(4)
	n := len(result)

	if n != 16 {
		t.Fail()
	}
}

func TestAggregateParameters(t *testing.T) {
	result := AggregateParameters([]float64{0.1, 0.11, 0.111, 0.111})
	n := len(result)

	if n != 16 {
		t.Fail()
	}
}

//func TestCase1(t *testing.T) {
//	t.Logf("Test case 1 \n\n\n")
//	args := []float64{0.1, 7, 0.1, 10}
//	expected := 0.322942087109504
//
//	actual := calculate(args)
//
//	if expected != actual {
//		t.Errorf("%f is not equal to %f", expected, actual)
//		t.Fail()
//	}
//}

func TestCase5(t *testing.T) {
	t.Logf("Test case 2 \n\n\n")
	args := []float64{1.2,	35,	1.3,	100}
	expected := 1.0935291036795

	actual := calculate(args)

	if expected != actual {
		t.Errorf("%f is not equal to %f", expected, actual)
		t.Fail()
	}
}

