package utils

func LaplaceMatrix(adjList [][]int) [][]int {
	laplaceMatrix := make([][]int, len(adjList))
	for i := 0; i < len(adjList); i++ {
		laplaceMatrix[i] = make([]int, len(adjList))
	}

	for i := 0; i < len(adjList); i++ {
		for j := 0; j < len(adjList[i]); j++ {
			laplaceMatrix[i][adjList[i][j]] = -1
		}
		laplaceMatrix[i][i] = len(adjList[i])
	}

	return laplaceMatrix
}
