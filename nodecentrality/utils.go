package nodecentrality

func bfs(adjList [][]int, start int) []int {

	levels := make([]int, len(adjList))
	fillSlice(levels, -1)

	levels[start] = 0
	queue := []int{start}
	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		for _, neighbor := range adjList[currNode] {
			if levels[neighbor] == -1 {
				levels[neighbor] = levels[currNode] + 1
				queue = append(queue, neighbor)
			}
		}
	}

	return levels
}

func fillSlice[T any](slice []T, val T) {
	for i := range slice {
		slice[i] = val
	}
}
