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

	// METHOD 1: Always get the exact top K percent, in case of ties, some are randomly chosen to be excluded
	// indexes1 := topKIndexes(metric1, topK)
	// indexes2 := topKIndexes(metric2, topK)

	// // get the intersection
	// intersection := 0
	// union := len(indexes1) + len(indexes2)
	// for idx := range indexes1 {
	// 	if indexes2[idx] {
	// 		intersection++
	// 		union--
	// 	}
	// }

	// METHOD 2: In case of a tie, get a few more elements
	threshold1 := findTopKThreshold(metric1, topK)
	threshold2 := findTopKThreshold(metric2, topK)

	// get the intersection
	var intersection, union int
	for i := range metric1 {
		isInTop1 := metric1[i] >= threshold1
		isInTop2 := metric2[i] >= threshold2

		if isInTop1 || isInTop2 { // at least one number is high
			union++
			if isInTop1 && isInTop2 { // both numbers are high
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
	topKIndex := int(float64(n) * (1.0 - topK))

	threshold := copiedSlice[topKIndex]
	return threshold
}

func topKIndexes(slice []float64, topK float64) map[int]bool {
	type item struct {
		index int
		value float64
	}

	itemsSlice := make([]item, len(slice))
	for i := range itemsSlice {
		itemsSlice[i] = item{i, slice[i]}
	}

	Scramble(itemsSlice) // guaranties that ties won't tend to be together

	sort.Slice(itemsSlice, func(i, j int) bool {
		return itemsSlice[i].value > itemsSlice[j].value
	})

	topIndexes := make(map[int]bool)

	indexesLen := int(float64(len(itemsSlice)) * topK)
	for i := 0; i < indexesLen; i++ {
		topIndexes[itemsSlice[i].index] = true
	}

	return topIndexes
}
