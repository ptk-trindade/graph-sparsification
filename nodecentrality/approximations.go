package nodecentrality

import (
	"fmt"
	"math/rand"

	"github.com/ptk-trindade/graph-sparsification/utils"
)

type ResultsDTO struct {
	Aux                 []int // wcb or closestBFS
	Mse                 []float64
	SpearmanCorrelation []float64
	SpearmanP           []float64
	Jaccard1percent     []float64
	Jaccard5percent     []float64
}

func NewResultDTO(n int) ResultsDTO {
	return ResultsDTO{
		Aux:                 make([]int, n),
		Mse:                 make([]float64, n),
		SpearmanCorrelation: make([]float64, n),
		SpearmanP:           make([]float64, n),
		Jaccard1percent:     make([]float64, n),
		Jaccard5percent:     make([]float64, n),
	}
}

// func ApproximateNodeCentrality(adjList [][]int, minIterations int, maxIterations int, pickCriteria string) ([]float64, []int) {

// 	firstLevels := bfs(adjList, 0) // start at any vertex

// 	// find first corner
// 	var cornerNode, maxLevel int
// 	for node, level := range firstLevels {
// 		if level > maxLevel {
// 			maxLevel = level
// 			cornerNode = node
// 		}
// 	}

// 	levels := NewLevels(len(adjList))
// 	wcb := false // node was corner before
// 	iterations := 0
// 	for (!wcb || iterations < minIterations) && iterations < maxIterations {
// 		iterations++
// 		foundLevels := bfs(adjList, cornerNode)

// 		levels.AddLevels(foundLevels)

// 		cornerNode, wcb = levels.GetCloselessNode()
// 	}
// 	fmt.Println("iterations:", iterations)
// 	closeness, eccentricity := levels.GetMetrics()

// 	return closeness, eccentricity
// }

func ApproximateCompareNodeCentrality(adjList [][]int, pickCriteria string, realCloseness []float64, realEccentricity []int, graphName string) {

	src := rand.Intn(len(adjList))
	firstLevels := bfs(adjList, src) // start at any vertex

	// find first corner
	var cornerNode, maxLevel, equalyGoodOptions int
	for node, level := range firstLevels {
		if level > maxLevel {
			maxLevel = level
			cornerNode = node
		} else if level == maxLevel {
			equalyGoodOptions++
			if rand.Intn(equalyGoodOptions) == 0 { // makes sure every node has the same probability of being chosen
				maxLevel = level
				cornerNode = node
			}
		}
	}

	levels := NewLevels(len(adjList))
	for iterations := 0; iterations < len(adjList); iterations++ {
		foundLevels := bfs(adjList, cornerNode)

		levels.AddLevels(foundLevels)

		if pickCriteria == "closeless" {
			var wcb bool
			cornerNode, wcb = levels.GetCloselessNode()

			avgDistance, approxEccentricity := levels.GetMetrics()
			approxCloseness := make([]float64, len(adjList))
			for i := range approxCloseness {
				approxCloseness[i] = 1.0 / avgDistance[i]
			}

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

			fmt.Printf("%s;closeless;%d;%t;%e;%e;%e;%e;%e;%e;%e;%e;%e;%e\n", graphName, iterations, wcb, mseCloseness, spearmanCloseness, pC, mseEccentricity, spearmanEccentricity, pE, jaccard1Closeness, jaccard5Closeness, jaccard1Eccentricity, jaccard5Eccentricity)

		} else { // furtherBfsed
			var distanceFromBfsedNode int
			cornerNode, distanceFromBfsedNode = levels.GetFurtherBfsedNode()

			avgDistance, approxEccentricity := levels.GetMetrics()
			approxCloseness := make([]float64, len(adjList))
			for i := range approxCloseness {
				approxCloseness[i] = 1.0 / avgDistance[i]
			}

			mseCloseness := utils.CalculateMSE(realCloseness, approxCloseness)
			mseEccentricity := utils.CalculateMSE(realEccentricity, approxEccentricity)

			spearmanCloseness, pC := utils.CompareSpearman(realCloseness, approxCloseness)
			spearmanEccentricity, pE := utils.CompareSpearman(sliceIntToFloat(realEccentricity), sliceIntToFloat(approxEccentricity))

			jaccard1Closeness := utils.CompareJaccard(realCloseness, approxCloseness, 0.01)
			jaccard5Closeness := utils.CompareJaccard(realCloseness, approxCloseness, 0.05)

			jaccard1Eccentricity := utils.CompareJaccard(sliceIntToFloat(realEccentricity), sliceIntToFloat(approxEccentricity), 0.01)
			jaccard5Eccentricity := utils.CompareJaccard(sliceIntToFloat(realEccentricity), sliceIntToFloat(approxEccentricity), 0.05)

			fmt.Printf("%s;further_bfsed;%d;%d;%e;%e;%e;%e;%e;%e;%e;%e;%e;%e\n", graphName, iterations, distanceFromBfsedNode, mseCloseness, spearmanCloseness, pC, mseEccentricity, spearmanEccentricity, pE, jaccard1Closeness, jaccard5Closeness, jaccard1Eccentricity, jaccard5Eccentricity)
		}
	}
}

func ApproximateCompareNodeCentralityRandom(adjList [][]int, realCloseness []float64, realEccentricity []int, graphName string) (ResultsDTO, ResultsDTO) {

	resultCloseness := NewResultDTO(len(adjList))
	resultEccentricity := NewResultDTO(len(adjList))

	nodeIdxes := make([]int, len(adjList))
	for i := range nodeIdxes {
		nodeIdxes[i] = i
	}
	utils.Scramble(nodeIdxes)

	levels := NewLevels(len(adjList))
	for iterations, nodeId := range nodeIdxes {
		foundLevels := bfs(adjList, nodeId)

		levels.AddLevels(foundLevels)

		avgDistance, approxEccentricity := levels.GetMetrics()
		approxCloseness := make([]float64, len(adjList))
		for i := range approxCloseness {
			approxCloseness[i] = 1.0 / avgDistance[i]
		}

		mseCloseness := utils.CalculateMSE(realCloseness, approxCloseness)
		mseEccentricity := utils.CalculateMSE(realEccentricity, approxEccentricity)

		spearmanCloseness, pC := utils.CompareSpearman(realCloseness, approxCloseness)
		spearmanEccentricity, pE := utils.CompareSpearman(sliceIntToFloat(realEccentricity), sliceIntToFloat(approxEccentricity))

		jaccard1Closeness := utils.CompareJaccard(realCloseness, approxCloseness, 0.01)
		jaccard5Closeness := utils.CompareJaccard(realCloseness, approxCloseness, 0.05)

		jaccard1Eccentricity := utils.CompareJaccard(sliceIntToFloat(realEccentricity), sliceIntToFloat(approxEccentricity), 0.01)
		jaccard5Eccentricity := utils.CompareJaccard(sliceIntToFloat(realEccentricity), sliceIntToFloat(approxEccentricity), 0.05)

		resultCloseness.Mse[iterations] = mseCloseness
		resultCloseness.SpearmanCorrelation[iterations] = spearmanCloseness
		resultCloseness.SpearmanP[iterations] = pC
		resultCloseness.Jaccard1percent[iterations] = jaccard1Closeness
		resultCloseness.Jaccard5percent[iterations] = jaccard5Closeness

		resultEccentricity.Mse[iterations] = mseEccentricity
		resultEccentricity.SpearmanCorrelation[iterations] = spearmanEccentricity
		resultEccentricity.SpearmanP[iterations] = pE
		resultEccentricity.Jaccard1percent[iterations] = jaccard1Eccentricity
		resultEccentricity.Jaccard5percent[iterations] = jaccard5Eccentricity

	}
	// graph;bfs_qty;MSE_closeness;spearman_closeness;spearman_p_closeness;MSE_eccentricity;spearman_eccentricity,spearman_p_eccentricity

	return resultCloseness, resultEccentricity
}

func sliceIntToFloat(slice []int) []float64 {
	sliceFloat := make([]float64, len(slice))
	for i, v := range slice {
		sliceFloat[i] = float64(v)
	}
	return sliceFloat
}
