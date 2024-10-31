package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	// "github.com/ptk-trindade/graph-sparsification/edgebetweenness"
	// "github.com/ptk-trindade/graph-sparsification/effectiveresistance"

	"github.com/ptk-trindade/graph-sparsification/nodecentrality"
	"github.com/ptk-trindade/graph-sparsification/utils"
)

func start(calculateMetrics bool) [][]int {

	scanner := bufio.NewScanner(os.Stdin)

	// scanner.Scan()
	// _ = scanner.Text() //method

	adjList, err := utils.ReadGraph(scanner)
	if err != nil {
		// fmt.Println("Error reading graph")
		panic("Error reading graph")
	}

	if calculateMetrics {
		degrees := make([]int, len(adjList))
		avgDegree := 0.0
		for i := range adjList {
			degrees[i] = len(adjList[i])
			avgDegree += float64(len(adjList[i]))
		}

		avgDegree /= float64(len(adjList))

		stdDev := 0.0
		minDegree := degrees[0]
		maxDegree := degrees[0]
		for _, d := range degrees {
			tmp := (float64(d) - avgDegree)
			stdDev += tmp * tmp

			minDegree = min(minDegree, d)
			maxDegree = max(maxDegree, d)
		}

		stdDev = stdDev / float64(len(adjList))
		stdDev = math.Sqrt(stdDev)

		fmt.Printf("min: %d\navg: %.3f~%.3f\nmax: %d", minDegree, avgDegree, stdDev, maxDegree)
	}

	return adjList
}

func main() {
	adjList := start(false)

	test1(adjList)
}

func test1(adjList [][]int) {

	realCloseness, realEccentricity := nodecentrality.NodeClosenessAndEccentricity(adjList)

	graphName := "graph_name"

	// randomCloseness := nodecentrality.NewResultDTO()
	// randomEccentricity := nodecentrality.NewResultDTO()

	resultsCloseness := make([]nodecentrality.ResultsDTO, 20)
	resultsEccentricity := make([]nodecentrality.ResultsDTO, 20)

	// random
	for i := 0; i < 20; i++ {
		closeness, eccentricity := nodecentrality.ApproximateCompareNodeCentralityRandom(adjList, realCloseness, realEccentricity, graphName)

		resultsCloseness[i] = closeness
		resultsEccentricity[i] = eccentricity
	}
	randomAvgResultCloseness := ResultsDTOAvg(resultsCloseness)
	randomAvgResultEccentricity := ResultsDTOAvg(resultsEccentricity)

	// closeless
	for i := 0; i < 20; i++ {
		closeness, eccentricity := nodecentrality.ApproximateCompareNodeCentrality(adjList, "closeless", realCloseness, realEccentricity, graphName)

		resultsCloseness[i] = closeness
		resultsEccentricity[i] = eccentricity
	}
	closelessAvgResultCloseness := ResultsDTOAvg(resultsCloseness)
	closelessAvgResultEccentricity := ResultsDTOAvg(resultsEccentricity)

	//furtherBfs
	for i := 0; i < 20; i++ {
		closeness, eccentricity := nodecentrality.ApproximateCompareNodeCentrality(adjList, "furtherBfsed", realCloseness, realEccentricity, graphName)

		resultsCloseness[i] = closeness
		resultsEccentricity[i] = eccentricity
	}
	furtherBfsAvgResultCloseness := ResultsDTOAvg(resultsCloseness)
	furtherBfsAvgResultEccentricity := ResultsDTOAvg(resultsEccentricity)

	//output

}

func ResultsDTOAvg(results []nodecentrality.ResultsDTO) nodecentrality.ResultsDTO {
	n := len(results)
	avg := NewResultDTO(len(results.Mse[0]))
	for r := range results {
		avg.Aux += float64(r.Aux) / n
		avg.Mse += r.Mse / n
		avg.SpearmanCorrelation += r.SpearmanCorrelation / n
		avg.SpearmanP += r.SpearmanP / n
		avg.Jaccard1percent += r.Jaccard1percent / n
		avg.Jaccard5percent += r.Jaccard5percent / n
	}

	return avg
}

/*
	randFloats := make([]float64, len(adjList))
	for i := range randFloats {
		randFloats[i] = rand.Float64()
	}

	// since the Jaccard function will consider the bigger values
	// and I'm interested in the lower ones, I must invert the values
	for i := range avgs {
		avgs[i] = 1 / avgs[i]
		maxs[i] = 1 / maxs[i]
	}

	percentages := []float64{0.01, 0.05, 0.1}
	for _, p := range percentages {
		realMin := utils.CompareJaccard(realBetweenness, mins, p)
		realAvg := utils.CompareJaccard(realBetweenness, avgs, p)
		realMax := utils.CompareJaccard(realBetweenness, maxs, p)

		minAvg := utils.CompareJaccard(mins, avgs, p)
		minMax := utils.CompareJaccard(mins, maxs, p)

		avgMax := utils.CompareJaccard(avgs, maxs, p)

		fmt.Printf("%.4f\t%.4f\t%.4f\t%.4f\t%.4f\t%.4f\t%.4f\n", p, realMin, realAvg, realMax, minAvg, minMax, avgMax)

		randReal := utils.CompareJaccard(randFloats, realBetweenness, p)
		randMin := utils.CompareJaccard(randFloats, mins, p)
		randAvg := utils.CompareJaccard(randFloats, avgs, p)
		randMax := utils.CompareJaccard(randFloats, maxs, p)
		fmt.Printf("%.4f\t%.4f\t%.4f\t%.4f\n", randReal, randMin, randAvg, randMax)
	}

	realMinRS, realMinP := utils.CompareSpearman(realBetweenness, mins)
	realAvgRS, realAvgP := utils.CompareSpearman(realBetweenness, avgs)
	realMaxRS, realMaxP := utils.CompareSpearman(realBetweenness, maxs)

	minAvgRS, minAvgP := utils.CompareSpearman(mins, avgs)
	minMaxRS, minMaxP := utils.CompareSpearman(mins, maxs)

	avgMaxRS, avgMaxP := utils.CompareSpearman(avgs, maxs)

	fmt.Printf("%.4f %.4f\t%.4f %.4f\t%.4f %.4f\t%.4f %.4f\t%.4f %.4f\t%.4f %.4f\n", realMinRS, realMinP, realAvgRS, realAvgP, realMaxRS, realMaxP, minAvgRS, minAvgP, minMaxRS, minMaxP, avgMaxRS, avgMaxP)

	randRealRS, randRealP := utils.CompareSpearman(realBetweenness, randFloats)
	randMinRS, randMinP := utils.CompareSpearman(randFloats, mins)
	randAvgRS, randAvgP := utils.CompareSpearman(randFloats, avgs)
	randMaxRS, randMaxP := utils.CompareSpearman(randFloats, maxs)

	fmt.Printf("%.4f %.4f\t%.4f %.4f\t%.4f %.4f\t%.4f %.4f\n", randRealRS, randRealP, randMinRS, randMinP, randAvgRS, randAvgP, randMaxRS, randMaxP)
*/
