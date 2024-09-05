package edgebetweenness

import (
	"github.com/ptk-trindade/graph-sparsification/utils"
)

/*
Run
*/
func runPartialForOneNode(vertex int, adjList [][]int, divideValue bool, nodesPerBfs int, edgeWeight *utils.EdgeWeight) map[int]bool {
	findPathsParcialDTO := findPathsParcial(adjList, vertex, nodesPerBfs)
	pathDAG := findPathsParcialDTO.pathDAG
	foundOrder := findPathsParcialDTO.foundOrder
	minimalPaths := findPathsParcialDTO.minimalPaths

	bannedNodes := make(map[int]bool)
	for idx, parents := range pathDAG {
		if len(parents) > 0 && parents[0] == vertex {
			bannedNodes[idx] = true
		}
	}

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
	return bannedNodes
}

func EdgeBetweennessParcial(adjList [][]int, divideValue bool, vertexes []int, bfsQty int, nodesPerBfs int, ban bool) *utils.EdgeWeight {
	edgeWeight := utils.NewEdgeWeight()

	var bannedNodes map[int]bool
	processed := make(map[int]bool)
	for qty := 0; qty < bfsQty; qty++ {
		var vertex int
		for _, vtx := range vertexes {
			if !ban || (!processed[vtx] && !bannedNodes[vtx]) {
				vertex = vtx
				break
			}
		}
		processed[vertex] = true
		bannedNodes = runPartialForOneNode(vertex, adjList, divideValue, nodesPerBfs, edgeWeight)
	}

	return edgeWeight
}

type findPathsParcialDTO struct {
	pathDAG      [][]int
	foundOrder   []int
	minimalPaths []float64
	bannedNodes  []int
}

/*
Returns a DAG with all the minimum paths to source node ('source' node has outdegree of 0)
*/
func findPathsParcial(adjList [][]int, source int, nodesToExplore int) findPathsParcialDTO {
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
	findPathsParcialDTO := findPathsParcialDTO{
		pathDAG:      pathDAG,
		foundOrder:   foundOrder,
		minimalPaths: minimalPaths,
	}

	return findPathsParcialDTO
}

/*
Returns a DAG with all the minimum paths to source node ('source' node has outdegree of 0)
*/
func findPathsParcialBan(adjList [][]int, source int, nodesToExplore int, levelsToBan int) findPathsParcialDTO {
	pathDAG := make([][]int, len(adjList))
	foundOrder := make([]int, 1, len(adjList))
	minimalPaths := make([]float64, len(adjList))

	bannedNodes := make([]int, 1, len(adjList))

	level := make([]int, len(adjList))
	level[source] = 1
	minimalPaths[source] = 1

	foundOrder[0] = source
	bannedNodes[0] = source
	queue := []int{source}
	for len(queue) > 0 && len(foundOrder) < nodesToExplore {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range adjList[node] {
			if level[neighbor] == 0 {
				level[neighbor] = level[node] + 1
				queue = append(queue, neighbor)
				foundOrder = append(foundOrder, neighbor)

				if level[node] <= levelsToBan { // ban neighbor
					bannedNodes = append(bannedNodes, neighbor)
				}
			}
			if level[neighbor] == level[node]+1 {
				pathDAG[neighbor] = append(pathDAG[neighbor], node)
				minimalPaths[neighbor] += minimalPaths[node]
			}
		}
	}

	findPathsParcialDTO := findPathsParcialDTO{
		pathDAG:      pathDAG,
		foundOrder:   foundOrder,
		minimalPaths: minimalPaths,
		bannedNodes:  bannedNodes,
	}

	return findPathsParcialDTO
}
