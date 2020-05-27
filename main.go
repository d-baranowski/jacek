package main

import (
	"fmt"
	"math"
)

/*
	Function equation_1(gamma As Double, tau As Double, beta As Double, theta As Double) As Double
    equation_1 = gamma * tau ^ (1.1) * (1.11 - 3 * (beta - 0.52) ^ 2) * Sin(theta) ^ (1.6)
	End Function
*/

/*
	trzeba przeliczyc funkcje 2 razy: function(alpha, beta, gamma)
	oraz function(alpha, betaMax, gamma)
	czyli raz dla beta = 1.1
	oraz raz dla beta = 1.0
*/

// beta gamma thau theta
// Borders (0.2, 1) (8, 32) (0.2, 1) (20, 90)
var (
	min = []float64{0.2, 8.0, 0.2, 20.0}
	max = []float64{1.0, 32.0, 1.0, 90.0}
)

func calculate(args []float64) float64 {
	computations := AggregateParameters(args)
	results := make([]float64, len(computations))

	for i, computation := range computations {
		results[i] = debug(computation[0], computation[1], computation[2], computation[3])
	}

	return FindMax(results)
}

func FindMax(in []float64) float64 {
	max := in[0]
	for _, f := range in {
		if f > max {
			max = f
		}
	}

	return max
}

func function(beta, gamma, thau, theta float64) float64 {
	return gamma * math.Pow(thau, 1.1) * (1.11 - 3 * math.Pow(beta - 0.52, 2)) * math.Pow(math.Sin(theta), 1.6)
}

func debug(beta, gamma, thau, theta float64) float64 {
	fmt.Printf("beta: %f gamma: %f thau: %f theta: %f\n", beta, gamma, thau, theta)
	return function(beta, gamma, thau, theta)
}

func GetCombinations(amountOfVariables int) [][]bool {
	result := make([][]bool, 0)

	for i := 0; i < (1 << amountOfVariables); i++ {
		boolArr := make([]bool, 0)

		for j := amountOfVariables - 1; j >= 0; j-- {
			boolArr = append(boolArr, i&(1<<j) > 0)
		}

		result = append(result, boolArr)
	}

	return result
}

func AggregateParameters(args []float64) [][]float64 {
	result := make([][]float64, 0)
	mods := []int{0, 0, 0, 0}
	outOfRange := make([]int, 0)

	for i, arg := range args {
		if arg < min[i] {
			mods[i] = -1
			outOfRange = append(outOfRange, i)
			continue
		}

		if arg > max[i] {
			mods[i] = 1
			outOfRange = append(outOfRange, i)
			continue
		}
	}

	combinations := GetCombinations(len(outOfRange))

	for _, combination := range combinations {
		argModifiers := []int{
			0, 0, 0, 0,
		}

		for i := range combination {
			if combination[i] {
				argModifiers[outOfRange[i]] = mods[i]
			}
		}

		argCombo := []float64{0, 0, 0, 0}

		for i, modifier := range argModifiers {
			if modifier == 0 {
				argCombo[i] = args[i]
			}

			if modifier == 1 {
				argCombo[i] = max[i]
			}

			if modifier == -1 {
				argCombo[i] = min[i]
			}
		}

		result = append(result, argCombo)
	}

	return result
}
