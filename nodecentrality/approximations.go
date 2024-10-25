package nodecentrality

import (
	"fmt"

	"github.com/ptk-trindade/graph-sparsification/utils"
)

func ApproximateNodeCentrality(adjList [][]int, minIterations int, maxIterations int, pickCriteria string) ([]float64, []int) {

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
	iterations := 0
	for (!wcb || iterations < minIterations) && iterations < maxIterations {
		iterations++
		foundLevels := bfs(adjList, cornerNode)

		levels.AddLevels(foundLevels)

		cornerNode, wcb = levels.GetCloselessNode()
	}
	fmt.Println("iterations:", iterations)
	closeness, eccentricity := levels.GetMetrics()

	return closeness, eccentricity
}

func ApproximateCompareNodeCentrality(adjList [][]int, pickCriteria string, realCloseness []float64, realEccentricity []int, graphName string) {

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
	for iterations := 0; iterations < len(adjList); iterations++ {
		foundLevels := bfs(adjList, cornerNode)

		levels.AddLevels(foundLevels)

		if pickCriteria == "closeless" {
			var wcb bool
			cornerNode, wcb = levels.GetCloselessNode()

			approxCloseness, approxEccentricity := levels.GetMetrics()

			// fmt.Println("C:", realCloseness[:10], approxCloseness[:10])
			// fmt.Println("E:", realEccentricity[:10], approxEccentricity[:10])

			mseCloseness := utils.CalculateMSE(realCloseness, approxCloseness)
			mseEccentricity := utils.CalculateMSE(realEccentricity, approxEccentricity)

			spearmanCloseness, pC := utils.CompareSpearman(realCloseness, approxCloseness)
			spearmanEccentricity, pE := utils.CompareSpearman(sliceIntToFloat(realEccentricity), sliceIntToFloat(approxEccentricity))

			jaccard1Closeness := utils.CompareJaccard(realCloseness, approxCloseness, 0.01)
			jaccard5Closeness := utils.CompareJaccard(realCloseness, approxCloseness, 0.05)

			jaccard1Eccentricity := utils.CompareJaccard(sliceIntToFloat(realEccentricity), sliceIntToFloat(approxEccentricity), 0.01)
			jaccard5Eccentricity := utils.CompareJaccard(sliceIntToFloat(realEccentricity), sliceIntToFloat(approxEccentricity), 0.05)

			fmt.Printf("%s;closeless;%d;%t;%.3f;%.3f;%.3f;%.3f;%.3f;%.3f;%.3f;%.3f;%.3f;%.3f\n", graphName, iterations, wcb, mseCloseness, spearmanCloseness, pC, mseEccentricity, spearmanEccentricity, pE, jaccard1Closeness, jaccard5Closeness, jaccard1Eccentricity, jaccard5Eccentricity)

		} else { // furtherBfsed
			var distanceFromBfsedNode int
			cornerNode, distanceFromBfsedNode = levels.GetFurtherBfsedNode()

			approxCloseness, approxEccentricity := levels.GetMetrics()

			mseCloseness := utils.CalculateMSE(realCloseness, approxCloseness)
			mseEccentricity := utils.CalculateMSE(realEccentricity, approxEccentricity)

			spearmanCloseness, pC := utils.CompareSpearman(realCloseness, approxCloseness)
			spearmanEccentricity, pE := utils.CompareSpearman(sliceIntToFloat(realEccentricity), sliceIntToFloat(approxEccentricity))

			jaccard1Closeness := utils.CompareJaccard(realCloseness, approxCloseness, 0.01)
			jaccard5Closeness := utils.CompareJaccard(realCloseness, approxCloseness, 0.05)

			jaccard1Eccentricity := utils.CompareJaccard(sliceIntToFloat(realEccentricity), sliceIntToFloat(approxEccentricity), 0.01)
			jaccard5Eccentricity := utils.CompareJaccard(sliceIntToFloat(realEccentricity), sliceIntToFloat(approxEccentricity), 0.05)

			fmt.Printf("%s;further_bfsed;%d;%d;%.3f;%.3f;%.3f;%.3f;%.3f;%.3f;%.3f;%.3f;%.3f;%.3f\n", graphName, iterations, distanceFromBfsedNode, mseCloseness, spearmanCloseness, pC, mseEccentricity, spearmanEccentricity, pE, jaccard1Closeness, jaccard5Closeness, jaccard1Eccentricity, jaccard5Eccentricity)
		}
	}
	// graph;bfs_qty;MSE_closeness;spearman_closeness;spearman_p_closeness;MSE_eccentricity;spearman_eccentricity,spearman_p_eccentricity

}

func ApproximateCompareNodeCentralityRandom(adjList [][]int, realCloseness []float64, realEccentricity []int, graphName string) {

	nodeIdxes := make([]int, len(adjList))
	for i := range nodeIdxes {
		nodeIdxes[i] = i
	}
	utils.Scramble(nodeIdxes)

	levels := NewLevels(len(adjList))
	for iterations, nodeId := range nodeIdxes {
		foundLevels := bfs(adjList, nodeId)

		levels.AddLevels(foundLevels)

		approxCloseness, approxEccentricity := levels.GetMetrics()

		mseCloseness := utils.CalculateMSE(realCloseness, approxCloseness)
		mseEccentricity := utils.CalculateMSE(realEccentricity, approxEccentricity)

		spearmanCloseness, pC := utils.CompareSpearman(realCloseness, approxCloseness)
		spearmanEccentricity, pE := utils.CompareSpearman(sliceIntToFloat(realEccentricity), sliceIntToFloat(approxEccentricity))

		jaccard1Closeness := utils.CompareJaccard(realCloseness, approxCloseness, 0.01)
		jaccard5Closeness := utils.CompareJaccard(realCloseness, approxCloseness, 0.05)

		jaccard1Eccentricity := utils.CompareJaccard(sliceIntToFloat(realEccentricity), sliceIntToFloat(approxEccentricity), 0.01)
		jaccard5Eccentricity := utils.CompareJaccard(sliceIntToFloat(realEccentricity), sliceIntToFloat(approxEccentricity), 0.05)

		fmt.Printf("%s;random;%d;-;%.3f;%.3f;%.3f;%.3f;%.3f;%.3f;%.3f;%.3f;%.3f;%.3f\n", graphName, iterations, mseCloseness, spearmanCloseness, pC, mseEccentricity, spearmanEccentricity, pE, jaccard1Closeness, jaccard5Closeness, jaccard1Eccentricity, jaccard5Eccentricity)

	}
	// graph;bfs_qty;MSE_closeness;spearman_closeness;spearman_p_closeness;MSE_eccentricity;spearman_eccentricity,spearman_p_eccentricity

}

func sliceIntToFloat(slice []int) []float64 {
	sliceFloat := make([]float64, len(slice))
	for i, v := range slice {
		sliceFloat[i] = float64(v)
	}
	return sliceFloat
}
