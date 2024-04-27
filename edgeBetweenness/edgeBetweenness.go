package edgebetweenness

/*
receive adjList
for each vertex:
	run modified BFS
	add weight to edges
*/

import (
	"fmt"
	"math"
	"os"

	"github.com/ptk-trindade/graph-sparsification/utils"
)

/*
every node has:
current indegree
amount of paths
*/
func EdgeBetweenness(adjList [][]int, divideValue bool) *utils.EdgeWeight {
	edgeWeight := utils.NewEdgeWeight()
	for vertex := range adjList {
		pathDAG, indegree, initialPathsBy := findPaths(adjList, vertex)

		// get edges with indegree of 0
		emptyIndegree := make([]int, 0)
		for i, indegree := range indegree {
			if indegree == 0 {
				emptyIndegree = append(emptyIndegree, i)
			}
		}

		pathsBy := make([]float64, len(initialPathsBy))
		copy(pathsBy, initialPathsBy)
		for i := range pathsBy {
			var value float64
			if divideValue {
				value = 1.0
			} else {
				value = initialPathsBy[i]
			}

			pathsBy[i] = value
		}

		for len(emptyIndegree) > 0 {
			// get last element of emptyIndegree
			node := emptyIndegree[len(emptyIndegree)-1]
			emptyIndegree = emptyIndegree[:len(emptyIndegree)-1]

			// fmt.Println(node, ">", pathDAG[node])
			// nodeOutdegree := len(pathDAG[node])
			var sumOutdegreePaths float64
			for _, parent := range pathDAG[node] {
				sumOutdegreePaths += initialPathsBy[parent]
			}

			for _, parent := range pathDAG[node] {
				pathWeight := pathsBy[node] * (initialPathsBy[parent] / sumOutdegreePaths)
				pathsBy[parent] += pathWeight
				// fmt.Println("pathsBy", parent, pathsBy[parent])
				edgeWeight.AddWeight(parent, node, pathWeight)
				if !divideValue && math.Mod(pathWeight, 1.0) > 0.1 {
					fmt.Println("whaaaat:")
					fmt.Println(vertex, "<-", parent, "<-", node, ":", pathWeight, "=", pathsBy[node], "*", pathsBy[parent], "/", sumOutdegreePaths)
					fmt.Println("pathsBy", pathsBy)
					os.Exit(4)
				}
				indegree[parent]--
				if indegree[parent] == 0 {
					emptyIndegree = append(emptyIndegree, parent)
				}
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
Returns a DAG with all the minimum paths to start node ('start' node has outdegree of 0)
*/
func findPaths(adjList [][]int, source int) ([][]int, []int, []float64) {
	pathDAG := make([][]int, len(adjList))
	indegree := make([]int, len(adjList))
	pathsBy := make([]float64, len(adjList))

	level := make([]int, len(adjList))
	level[source] = 1
	pathsBy[source] = 1

	queue := []int{source}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range adjList[node] {
			if level[neighbor] == 0 {
				level[neighbor] = level[node] + 1
				queue = append(queue, neighbor)
			}
			if level[neighbor] == level[node]+1 {
				pathDAG[neighbor] = append(pathDAG[neighbor], node)
				indegree[node]++
				pathsBy[neighbor] += pathsBy[node]
			}
		}
	}
	return pathDAG, indegree, pathsBy
}
