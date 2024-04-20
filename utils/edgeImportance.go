package utils

import "strconv"

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
