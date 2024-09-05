package edgebetweenness

/*
receive adjList
for each vertex:
	run modified BFS
	add weight to edges
*/

import (
	"github.com/ptk-trindade/graph-sparsification/utils"
)

func EdgeBetweennessSum(adjList [][]int) *utils.EdgeWeight {
	return edgeBetweenness(adjList, false)
}

func EdgeBetweennessDivide(adjList [][]int) *utils.EdgeWeight {
	return edgeBetweenness(adjList, true)
}

func edgeBetweenness(adjList [][]int, divideValue bool) *utils.EdgeWeight {
	edgeWeight := utils.NewEdgeWeight()
	for vertex := range adjList {
		pathDAG, foundOrder, minimalPaths := findPaths(adjList, vertex)

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

			// fmt.Println(node, ">", pathDAG[node])
			// nodeOutdegree := len(pathDAG[node])
			var sumOutdegreePaths float64
			for _, parent := range pathDAG[node] {
				sumOutdegreePaths += minimalPaths[parent]
			}

			for _, parent := range pathDAG[node] {
				pathWeight := pathsBy[node] * (minimalPaths[parent] / sumOutdegreePaths)
				pathsBy[parent] += pathWeight
				// fmt.Println("pathsBy", parent, pathsBy[parent])
				edgeWeight.AddWeight(parent, node, pathWeight)

				// indegree[parent]--
				// if indegree[parent] == 0 {
				// 	emptyIndegree = append(emptyIndegree, parent)
				// }
			}
			// fmt.Println(edgeWeight)
		}
		// fmt.Println("vertex:", vertex)
		// fmt.Println("pathsBy", pathsBy)
		// fmt.Println(edgeWeight)

	}
	return edgeWeight
}

/*
Returns a DAG with all the minimum paths to source node ('source' node has outdegree of 0)
*/
func findPaths(adjList [][]int, source int) ([][]int, []int, []float64) {
	pathDAG := make([][]int, len(adjList))
	foundOrder := make([]int, 1, len(adjList))
	minimalPaths := make([]float64, len(adjList))

	level := make([]int, len(adjList))
	level[source] = 1
	minimalPaths[source] = 1

	foundOrder[0] = source
	queue := []int{source}
	for len(queue) > 0 {
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

// count paths
func EdgeBetweennessCountPaths(adjList [][]int, divideValue bool) (*utils.EdgeWeight, int) {
	edgeWeight := utils.NewEdgeWeight()
	totalPaths := 0
	for vertex := range adjList {
		pathDAG, foundOrder, minimalPaths := findPaths(adjList, vertex)

		pathsBy := make([]float64, len(minimalPaths))
		if divideValue {
			for i := range pathsBy {
				pathsBy[i] = 1.0
			}
		} else {
			copy(pathsBy, minimalPaths)
		}

		for i := len(foundOrder) - 1; i >= 0; i-- {
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
		for _, paths := range minimalPaths {
			totalPaths += int(paths)
		}
	}
	return edgeWeight, totalPaths
}
