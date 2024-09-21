package utils

import (
	"fmt"
	"sort"

	"github.com/dgryski/go-onlinestats"
)

func CompareSpearman(metric1 []float64, metric2 []float64) (rs float64, p float64) {
	return onlinestats.Spearman(metric1, metric2)
}

/* Get the top nodes from each metric and returns size(A ∩ B)/size(A ∪ B)
*/
func CompareJaccard(metric1 []float64, metric2 []float64, topK float64) float64 {
	if len(metric1) != len(metric2) {
		fmt.Println("Error CompareJaccard: metrics should have same length")
		return 0.0
	}

	threshold1 := findTopKThreshold(metric1, topK)
	threshold2 := findTopKThreshold(metric2, topK)
	
	// get the intersection
	var intersection, union int
	for i := range metric1 {
		if metric1[i] >= threshold1 || metric2[i] >= threshold2 { // at least one number is high
			union++
			if metric1[i] >= threshold1 && metric2[i] >= threshold2 { // both numbers are high
				intersection++
			}
		}
	}

	return float64(intersection) / float64(union)
}

func findTopKThreshold(slice []float64, topK float64) float64 {
	copiedSlice := make([]float64, len(slice))
	copy(copiedSlice, slice)
	
	sort.Float64s(copiedSlice)
	n := len(copiedSlice)
	topKIndex := int(float64(n) * (1 - topK))

	threshold := copiedSlice[topKIndex]
	return threshold
}