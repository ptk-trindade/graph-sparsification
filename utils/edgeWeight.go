package utils

import (
	"fmt"
	"sort"
	"strconv"
)

type EdgeWeight struct {
	weight map[string]float64
}

func NewEdgeWeight() *EdgeWeight {
	return &EdgeWeight{
		weight: make(map[string]float64),
	}
}

func (ew *EdgeWeight) getKey(node1, node2 int) string {
	if node1 < node2 {
		return strconv.Itoa(node1) + "-" + strconv.Itoa(node2)
	}
	return strconv.Itoa(node2) + "-" + strconv.Itoa(node1)
}

func (ew *EdgeWeight) AddWeight(node1, node2 int, importance float64) {
	ew.weight[ew.getKey(node1, node2)] += importance
}

// Makes the avg weight of the edges be 1
func (ew *EdgeWeight) Normalize() {
	avg := 0.0
	for _, w := range ew.weight {
		avg += w
	}
	avg /= float64(len(ew.weight))

	for k := range ew.weight {
		ew.weight[k] /= avg
	}
}

func (ew *EdgeWeight) Show() {
	// sort edges by weight
	edges := make([]string, 0, len(ew.weight))
	for k := range ew.weight {
		edges = append(edges, k)
		ew.weight[k] = RoundFloat(ew.weight[k], 3)
	}

	sort.Slice(edges, func(i, j int) bool {
		if ew.weight[edges[i]] == ew.weight[edges[j]] {
			return edges[i] < edges[j]
		}
		return ew.weight[edges[i]] > ew.weight[edges[j]]
	})

	for _, k := range edges {
		fmt.Println(k, ew.weight[k])
	}

}
