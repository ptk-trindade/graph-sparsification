package utils

type NodeWeight struct {
	weight map[[2]int]float64
}

// func NewEdgeWeight() *EdgeWeight {
// 	return &EdgeWeight{
// 		weight: make(map[[2]int]float64),
// 	}
// }

// func (ew *EdgeWeight) getKey(node1, node2 int) [2]int {
// 	if node1 < node2 {
// 		return [2]int{node1, node2}
// 	}
// 	return [2]int{node2, node1}
// }

// func (ew *EdgeWeight) AddWeight(node1, node2 int, importance float64) {
// 	ew.weight[ew.getKey(node1, node2)] += importance
// }

// // Makes the avg weight of the edges be 1
// func (ew *EdgeWeight) Normalize() {
// 	avg := 0.0
// 	for _, w := range ew.weight {
// 		avg += w
// 	}
// 	avg /= float64(len(ew.weight))

// 	for k := range ew.weight {
// 		ew.weight[k] /= avg
// 	}
// }

// func (ew *EdgeWeight) Show() {
// 	// sort edges by weight
// 	edges := make([][2]int, 0, len(ew.weight))
// 	for key := range ew.weight {
// 		edges = append(edges, key)
// 	}

// 	sort.Slice(edges, func(i, j int) bool {
// 		if ew.weight[edges[i]] == ew.weight[edges[j]] {
// 			return edges[i][0] < edges[j][0]
// 		}
// 		return ew.weight[edges[i]] > ew.weight[edges[j]]
// 	})

// 	for _, key := range edges {
// 		fmt.Printf("%d-%d: %.3f\n", key[0], key[1], ew.weight[key])
// 	}

// }

// func (ew1 *EdgeWeight) CompareSpearman(ew2 *EdgeWeight) (rs float64, p float64) {
// 	weight1 := make([]float64, 0, len(ew1.weight))
// 	weight2 := make([]float64, 0, len(ew2.weight))

// 	for edgeKey := range ew1.weight {
// 		weight1 = append(weight1, ew1.weight[edgeKey])
// 		weight2 = append(weight2, ew2.weight[edgeKey])
// 	}

// 	return onlinestats.Spearman(weight1, weight2)
// }

// func (ew1 *EdgeWeight) CompareJaccard(ew2 *EdgeWeight, top float64) float64 {
// 	edges1 := make([][2]int, 0, len(ew1.weight))
// 	edges2 := make([][2]int, 0, len(ew2.weight))

// 	for key := range ew1.weight {
// 		edges1 = append(edges1, key)
// 	}

// 	for key := range ew2.weight {
// 		edges2 = append(edges2, key)
// 	}

// 	sort.Slice(edges1, func(i, j int) bool {
// 		// if ew1.weight[edges1[i]] == ew1.weight[edges1[j]] { // if the weights are the same, sort by other metric
// 		// 	w_i, ok_i := ew2.weight[edges1[i]]
// 		// 	w_j, ok_j := ew2.weight[edges1[j]]
// 		// 	if ok_i && ok_j {
// 		// 		return w_i > w_j
// 		// 	}
// 		// 	return ok_i
// 		// }
// 		return ew1.weight[edges1[i]] > ew1.weight[edges1[j]]
// 	})

// 	sort.Slice(edges2, func(i, j int) bool {
// 		// if ew2.weight[edges2[i]] == ew2.weight[edges2[j]] {
// 		// 	w_i, ok_i := ew1.weight[edges2[i]]
// 		// 	w_j, ok_j := ew1.weight[edges2[j]]
// 		// 	if ok_i && ok_j {
// 		// 		return w_i > w_j
// 		// 	}
// 		// 	return ok_i
// 		// }
// 		return ew2.weight[edges2[i]] > ew2.weight[edges2[j]]
// 	})

// 	// get the top edges
// 	topEdges1 := make(map[[2]int]float64)
// 	for i := 0; i < int(top*float64(len(edges1))); i++ {
// 		topEdges1[edges1[i]] = ew1.weight[edges1[i]]
// 	}

// 	topEdges2 := make(map[[2]int]float64)
// 	for i := 0; i < int(top*float64(len(edges2))); i++ {
// 		topEdges2[edges2[i]] = ew2.weight[edges2[i]]
// 	}

// 	// get the intersection
// 	intersection := 0
// 	for key := range topEdges1 {
// 		if _, ok := topEdges2[key]; ok {
// 			intersection++
// 		}
// 	}

// 	return float64(intersection) / float64(len(topEdges1)+len(topEdges2)-intersection)
// }
