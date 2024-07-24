package edgebetweenness

import (
	"github.com/ptk-trindade/graph-sparsification/utils"
)

func EdgeBetweennessParcial(adjList [][]int, divideValue bool, edges []int, nodesPerBfs int) *utils.EdgeWeight {
	edgeWeight := utils.NewEdgeWeight()

	for _, vertex := range edges {
		pathDAG, foundOrder, minimalPaths := findPathsParcial(adjList, vertex, nodesPerBfs)

		pathsBy := make([]float64, len(minimalPaths))
		if divideValue {
			for i := range pathsBy {
				pathsBy[i] = 1.0
			}
		} else {
			copy(pathsBy, minimalPaths)
		}

		for i := len(foundOrder) - 1; i >= 0; i-- {
			// get last element of emptyIndegree
			node := foundOrder[i]

			var sumOutdegreePaths float64
			for _, parent := range pathDAG[node] {
				sumOutdegreePaths += minimalPaths[parent]
			}

			for _, parent := range pathDAG[node] {
				pathWeight := pathsBy[node] * (minimalPaths[parent] / sumOutdegreePaths)
				pathsBy[parent] += pathWeight
				edgeWeight.AddWeight(parent, node, pathWeight)

			}
		}

	}
	return edgeWeight
}

/*
Returns a DAG with all the minimum paths to source node ('source' node has outdegree of 0)
*/
func findPathsParcial(adjList [][]int, source int, nodesToExplore int) ([][]int, []int, []float64) {
	pathDAG := make([][]int, len(adjList))
	foundOrder := make([]int, 1, len(adjList))
	minimalPaths := make([]float64, len(adjList))

	level := make([]int, len(adjList))
	level[source] = 1
	minimalPaths[source] = 1

	foundOrder[0] = source
	queue := []int{source}
	for len(queue) > 0 && len(foundOrder) < nodesToExplore {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range adjList[node] {
			if level[neighbor] == 0 {
				level[neighbor] = level[node] + 1
				queue = append(queue, neighbor)
				foundOrder = append(foundOrder, neighbor)
			}
			if level[neighbor] == level[node]+1 {
				pathDAG[neighbor] = append(pathDAG[neighbor], node)
				minimalPaths[neighbor] += minimalPaths[node]
			}
		}
	}

	return pathDAG, foundOrder, minimalPaths
}
