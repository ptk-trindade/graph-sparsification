package nodecentrality

import (
	"fmt"
	"math"

	"github.com/ptk-trindade/graph-sparsification/utils"
)

func NodeClosenessAndEccentricity(adjList [][]int) ([]float64, []int) {
	nodeCloseness := make([]float64, len(adjList))
	nodeEccentricity := make([]int, len(adjList))

	for vertex := range adjList {
		levels := bfs(adjList, vertex)

		nodeCloseness[vertex] = float64(len(adjList)-1) / utils.Sum(levels...)
		nodeEccentricity[vertex] = utils.Max(levels...)
	}

	return nodeCloseness, nodeEccentricity
}

func NodeEccentricity(adjList [][]int) []int {
	nodeEccentricity := make([]int, len(adjList))

	for vertex := range adjList {
		levels := bfs(adjList, vertex)

		nodeEccentricity[vertex] = utils.Max(levels...)
	}

	return nodeEccentricity
}

func NodeCloseness(adjList [][]int) []float64 {
	nodeCloseness := make([]float64, len(adjList))

	for vertex := range adjList {
		levels := bfs(adjList, vertex)

		fmt.Println("l:", levels[:10])

		nodeCloseness[vertex] = float64(len(adjList)-1) / utils.Sum(levels...)
	}

	return nodeCloseness
}

func NodeBetweenness(adjList [][]int) []float64 {
	nodeBetweenness := make([]float64, len(adjList))

	for vertex := range adjList {
		pathDAG, foundOrder, _ := findPaths(adjList, vertex)

		pathsBy := make([]float64, len(adjList)) // paths[node] = amount of paths that go through node

		for i := len(foundOrder) - 1; i >= 0; i-- {
			node := foundOrder[i]

			pathWeight := (pathsBy[node] + 1.0) / float64(len(pathDAG[node]))
			for _, parent := range pathDAG[node] {
				pathsBy[parent] += pathWeight
			}

		}

		if math.Abs(pathsBy[vertex]-float64(len(adjList)-1)) > 0.5 {
			fmt.Println("Error: int(pathsBy[vertex]) != len(adjList) - 1", pathsBy[vertex], len(adjList)-1)
		}

		pathsBy[vertex] = 0
		for i := range nodeBetweenness {
			nodeBetweenness[i] += pathsBy[i]
		}
	}

	return nodeBetweenness
}

/*
Returns a DAG with all the minimum paths to source node ('source' node has outdegree of 0)
*/
func findPaths(adjList [][]int, source int) ([][]int, []int, []float64) {
	pathDAG := make([][]int, len(adjList))        // same structure as adjList but it's a DAG
	foundOrder := make([]int, 1, len(adjList))    // order in which nodes were found while running the BFS
	minimalPaths := make([]float64, len(adjList)) // amount of minimal paths between [node] and source

	level := make([]int, len(adjList))
	fillSlice(level, -1)

	level[source] = 0
	minimalPaths[source] = 1

	foundOrder[0] = source
	queue := []int{source}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range adjList[node] {
			if level[neighbor] == -1 { // not visited
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
