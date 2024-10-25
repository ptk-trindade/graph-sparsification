package utils

import (
	"fmt"
	"math/rand"
)

func Scramble[T any](slice []T) {
	// Shuffle the slice using the Fisher-Yates algorithm
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)                   // Generate a random index from 0 to i
		slice[i], slice[j] = slice[j], slice[i] // Swap elements
	}
}

type Number interface {
	int | float64
}

// CalculateMSE computes the Mean Squared Error between two slices.
func CalculateMSE[E Number](slice1, slice2 []E) float64 {
	if len(slice1) != len(slice2) {
		fmt.Println("Error: Slices with different lengths", len(slice1), len(slice2))
		return 0
	}

	var sumSquaredError E
	for i := 0; i < len(slice1); i++ {
		diff := slice1[i] - slice2[i]
		sumSquaredError += diff * diff
	}

	mse := float64(sumSquaredError) / float64(len(slice1))
	return mse
}

func Max[E Number](vals ...E) E {
	maxVal := vals[0]
	for _, val := range vals {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func Min[E Number](vals ...E) E {
	minVal := vals[0]
	for _, val := range vals {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func Avg[E Number](vals ...E) float64 {
	var sum float64
	for _, val := range vals {
		sum += float64(val)
	}

	sum /= float64(len(vals))

	return sum
}

func Sum[E Number](vals ...E) float64 {
	var sum float64
	for _, val := range vals {
		sum += float64(val)
	}

	return sum
}
