package utils

import (
	"fmt"
	"sort"
)

func FindComponents(adjList [][]int) [][]int {
	visited := make([]bool, len(adjList))
	components := make([][]int, 0)
	for node_i := 0; node_i < len(adjList); node_i++ { // for each vertex
		if !visited[node_i] { // if not visited, belongs to a new component
			component := dfs(adjList, node_i, visited)
			components = append(components, component)
		}
	}

	sort.Slice(components, func(i, j int) bool { return len(components[i]) < len(components[j]) }) // ascending order

	singleComponents := make([]int, 0)
	fmt.Print(len(components), ",", len(components[0]), ",", len(components[len(components)-1]), ",")
	for _, component := range components {
		// fmt.Print(len(component), " ")
		if len(component) == 1 {
			singleComponents = append(singleComponents, component[0])
		}
	}

	n := len(adjList)
	for _, c := range components {
		n -= len(c)
	}

	return components
}

// searches a component, changing the visited vector
func dfs(adjList [][]int, start int, visited []bool) []int {
	// ----- DFS -----
	component := make([]int, 0)
	stack := []int{start}

	visited[start] = true
	for len(stack) > 0 {
		// pop(-1)
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		component = append(component, current)

		// add neighbors to queue
		for _, neighbor := range adjList[current] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
				visited[neighbor] = true
			}
		}
	}

	return component
}
