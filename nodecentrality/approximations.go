package nodecentrality

func ApproximateNodeCentrality(adjList [][]int) ([]float64, []float64, []float64) {
	
	firstLevels := bfs(adjList, 0) // start at any vertex
	
	// find first corner
	var cornerNode, maxLevel int
	for node, level := range firstLevels {
		if level > maxLevel {
			maxLevel = level
			cornerNode = node
		}
	}
	
	levels := NewLevels(len(adjList)) 
	wcb := false // node was corner before
	for !wcb {
		levels := bfs(adjList, cornerNode)

		levels.AddLevels(levels)
		cornerNode, wcb = levels.GetCornerNode()
	}

	minsInt, avgs, maxsInt := levels.GetMetrics()

	minsFloat := make([]float64, len(minsFloat))
	maxsFloat := make([]float64, len(maxsInts))
	for i := range minsInt {
		minsFloat[i] = float64(minsInt[i])
		maxsFloat[i] = float64(maxsInts[i])
	}

	return minsFloat, avgs, maxsFloat
}

func bfs(adjList [][]int, start int) []int {

	levels := make([]int, len(adjList))
	fillSlice(levels, -1)

	levels[start] = 0
	queue := []int{start}
	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		for neighbor := range adjList[currNode] {
			if levels[neighbor] == -1 {
				levels[neighbor] = levels[currNode] + 1
				queue = append(queue, currNode)
			}
		}
	}

	return levels
}