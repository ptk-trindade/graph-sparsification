package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"

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

	nRuns := 20
	
	// benchmark
	sampleSizes := []int{20, 100, 500, 1000}
	benchmark(adjList, sampleSizes, nRuns)
	
	// compute
	// sampleSizes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 25, 26, 27, 28, 30, 31, 33, 34, 36, 38, 39, 41, 43, 45, 47, 50, 52, 54, 57, 60, 63, 66, 69, 72, 75, 79, 83, 87, 91, 95, 100, 104, 109, 114, 120, 125, 131, 138, 144, 151, 158, 165, 173, 181, 190, 199, 208, 218, 229, 239, 251, 263, 275, 288, 301, 316, 331, 346, 363, 380, 398, 416, 436, 457, 478, 501, 524, 549, 575, 602, 630, 660, 691, 724, 758, 794, 831, 870, 912, 954, 1000, 1047, 1096, 1148, 1202, 1258, 1318, 1380, 1445, 1513, 1584, 1659, 1737, 1819, 1905, 1995, 2089, 2187, 2290, 2398, 2511, 2630, 2754, 2884, 3019, 3162, 3311, 3467, 3630, 3801, 3981, len(adjList)}
	// fmt.Println("Closeness")
	// compareCloseness(adjList, sampleSizes, nRuns)
	// fmt.Println("\nEccentricity")
	// compareEccentricity(adjList, sampleSizes, nRuns)
}

// Helper function to calculate the average of a slice of floats
func average(times []float64) float64 {
	var sum float64
	for _, time := range times {
		sum += time
	}
	return sum / float64(len(times))
}

// Helper function to calculate the standard deviation of a slice of floats
func stdDev(times []float64, avg float64) float64 {
	var sum float64
	for _, time := range times {
		sum += math.Pow(time-avg, 2)
	}
	return math.Sqrt(sum / float64(len(times)))
}

func benchmark(adjList [][]int, sampleSizes []int, nRuns int) {
	// CLOSENESS
	fmt.Println("Closeness")
	methodsCloseness := []func([][]int, int) ([]nodecentrality.ClosenessNode, float64){
		nodecentrality.ApproximateClosenessCloseless,
		nodecentrality.ApproximateClosenessFurtherBfs,
		nodecentrality.ApproximateClosenessRandom,
	}

	for methodI, method := range methodsCloseness {
		for _, nSamples := range sampleSizes {
			var times []float64
			for i := 0; i < nRuns; i++ {
				start := time.Now()
				method(adjList, nSamples)
				elapsed := time.Since(start).Seconds() * 1000 // Convert to miliseconds for precision
				times = append(times, elapsed)
			}

			avg := average(times)
			stddev := stdDev(times, avg)
			fmt.Printf("Method %d > Samples: %d, Average Time: %.2f ± %.2f ms\n", methodI, nSamples, avg, stddev)
		}
	}

	// ECCENTRICITY
	fmt.Println("Eccentricity")
	methodsEccentricity := []func([][]int, int) ([]nodecentrality.EccentricityNode, float64){
		nodecentrality.ApproximateEccentricityCloseless,
		nodecentrality.ApproximateEccentricityFurtherBfs,
		nodecentrality.ApproximateEccentricityRandom,
	}

	for methodI, method := range methodsEccentricity {
		for _, nSamples := range sampleSizes {
			var times []float64
			for i := 0; i < nRuns; i++ {
				start := time.Now()
				method(adjList, nSamples)
				elapsed := time.Since(start).Seconds() * 1e9 // Convert to nanoseconds
				times = append(times, elapsed)
			}

			avg := average(times)
			stddev := stdDev(times, avg)
			fmt.Printf("Method %d > Samples: %d, Average Time: %.2f ± %.2f ns\n", methodI, nSamples, avg, stddev)
		}
	}
}

type functionResults struct {
	expectedMse []float64
	mse         []float64
	spCor       []float64
	Jcc1        []float64
	Jcc5        []float64
}

func compareCloseness(adjList [][]int, sampleSizes []int, nRuns int) {

	methods := []func([][]int, int) ([]nodecentrality.ClosenessNode, float64){
		nodecentrality.ApproximateClosenessCloseless,
		nodecentrality.ApproximateClosenessFurtherBfs,
		nodecentrality.ApproximateClosenessRandom,
	}

	realCloseness := nodecentrality.NodeCloseness(adjList)
	tableVals := make([]functionResults, len(methods))

	for methodI, method := range methods {
		expectedMse := make([]float64, len(sampleSizes))
		mse := make([]float64, len(sampleSizes))
		spCor := make([]float64, len(sampleSizes))
		Jcc1 := make([]float64, len(sampleSizes))
		Jcc5 := make([]float64, len(sampleSizes))

		for ssI, nSamples := range sampleSizes {
			// start := time.Now()
			for i := 0; i < nRuns; i++ {
				approxClosenessNodes, expMse := method(adjList, nSamples)

				approxCloseness := make([]float64, len(approxClosenessNodes))
				for i, node := range approxClosenessNodes {
					if node.CloserBfs == 0 {
						approxCloseness[i] = node.RealCloseness
					} else {
						approxCloseness[i] = node.ExpectedCloseness
					}
				}

				expectedMse[ssI] += expMse / float64(nRuns)

				realMse := utils.CalculateMSE(approxCloseness, realCloseness)
				mse[ssI] += realMse / float64(nRuns)

				currSpCor, _ := utils.CompareSpearman(approxCloseness, realCloseness)
				spCor[ssI] += currSpCor / float64(nRuns)

				currJc1 := utils.CompareJaccard(approxCloseness, realCloseness, 0.01)
				Jcc1[ssI] += currJc1 / float64(nRuns)

				currJc5 := utils.CompareJaccard(approxCloseness, realCloseness, 0.05)
				Jcc5[ssI] += currJc5 / float64(nRuns)

			}
			// elapsed := time.Since(start)
			// fmt.Printf("%d took: %s\n", nSamples, elapsed)
		}

		tableVals[methodI] = functionResults{
			expectedMse: expectedMse,
			mse:         mse,
			spCor:       spCor,
			Jcc1:        Jcc1,
			Jcc5:        Jcc5,
		}

	}

	fmt.Println(";Closeless;Further BFS;Random")

	fmt.Print("nSamples")
	fmt.Print(";Expected MSE;MSE;Spearman Correlation;Jaccard 1%;Jaccard 5%")
	fmt.Print(";Expected MSE;MSE;Spearman Correlation;Jaccard 1%;Jaccard 5%")
	fmt.Print(";Expected MSE;MSE;Spearman Correlation;Jaccard 1%;Jaccard 5%")
	fmt.Println()

	for ssI, ss := range sampleSizes {
		fmt.Print(ss, ";")
		for methodI := 0; methodI < len(methods); methodI++ {
			fmt.Print(tableVals[methodI].expectedMse[ssI], ";")
			fmt.Print(tableVals[methodI].mse[ssI], ";")
			fmt.Print(tableVals[methodI].spCor[ssI], ";")
			fmt.Print(tableVals[methodI].Jcc1[ssI], ";")
			fmt.Print(tableVals[methodI].Jcc5[ssI], ";")
		}
		fmt.Println()
	}
}

func compareEccentricity(adjList [][]int, sampleSizes []int, nRuns int) {

	methods := []func([][]int, int) ([]nodecentrality.EccentricityNode, float64){
		nodecentrality.ApproximateEccentricityCloseless,
		nodecentrality.ApproximateEccentricityFurtherBfs,
		nodecentrality.ApproximateEccentricityRandom,
	}

	realEccentricityInt := nodecentrality.NodeEccentricity(adjList)
	realEccentricity := make([]float64, len(realEccentricityInt))
	for i, val := range realEccentricityInt {
		realEccentricity[i] = float64(val)
	}

	tableVals := make([]functionResults, len(methods))

	for methodI, method := range methods {
		expectedMse := make([]float64, len(sampleSizes))
		mse := make([]float64, len(sampleSizes))
		spCor := make([]float64, len(sampleSizes))
		Jcc1 := make([]float64, len(sampleSizes))
		Jcc5 := make([]float64, len(sampleSizes))

		for ssI, nSamples := range sampleSizes {
			for i := 0; i < nRuns; i++ {
				approxEccentricityNodes, expMse := method(adjList, nSamples)

				approxEccentricity := make([]float64, len(approxEccentricityNodes))
				for i, node := range approxEccentricityNodes {
					if node.CloserBfs == 0 {
						approxEccentricity[i] = float64(node.RealEccentricity)
					} else {
						approxEccentricity[i] = float64(node.ExpectedEccentricity)
					}
				}

				expectedMse[ssI] += expMse / float64(nRuns)

				realMse := utils.CalculateMSE(approxEccentricity, realEccentricity)
				mse[ssI] += float64(realMse) / float64(nRuns)

				currSpCor, _ := utils.CompareSpearman(approxEccentricity, realEccentricity)
				spCor[ssI] += currSpCor / float64(nRuns)

				currJc1 := utils.CompareJaccardBottom(approxEccentricity, realEccentricity, 0.01)
				Jcc1[ssI] += currJc1 / float64(nRuns)

				currJc5 := utils.CompareJaccardBottom(approxEccentricity, realEccentricity, 0.05)
				Jcc5[ssI] += currJc5 / float64(nRuns)

			}
		}

		tableVals[methodI] = functionResults{
			expectedMse: expectedMse,
			mse:         mse,
			spCor:       spCor,
			Jcc1:        Jcc1,
			Jcc5:        Jcc5,
		}

	}

	fmt.Println("Closeless;Further BFS;Random")

	fmt.Print("nSamples")
	fmt.Print(";Expected MSE;MSE;Spearman Correlation;Jaccard 1%;Jaccard 5%")
	fmt.Print(";Expected MSE;MSE;Spearman Correlation;Jaccard 1%;Jaccard 5%")
	fmt.Print(";Expected MSE;MSE;Spearman Correlation;Jaccard 1%;Jaccard 5%")
	fmt.Println()

	for ssI, ss := range sampleSizes {
		fmt.Print(ss, ";")
		for methodI := 0; methodI < len(methods); methodI++ {
			fmt.Print(tableVals[methodI].expectedMse[ssI], ";")
			fmt.Print(tableVals[methodI].mse[ssI], ";")
			fmt.Print(tableVals[methodI].spCor[ssI], ";")
			fmt.Print(tableVals[methodI].Jcc1[ssI], ";")
			fmt.Print(tableVals[methodI].Jcc5[ssI], ";")
		}
		fmt.Println()
	}
}

// func test1(adjList [][]int) {

// 	realCloseness, realEccentricity := nodecentrality.NodeClosenessAndEccentricity(adjList)

// 	graphName := "graph_name"

// 	// randomCloseness := nodecentrality.NewResultDTO()
// 	// randomEccentricity := nodecentrality.NewResultDTO()

// 	resultsCloseness := make([]nodecentrality.ResultsDTO, 20)
// 	resultsEccentricity := make([]nodecentrality.ResultsDTO, 20)

// 	// random
// 	for i := 0; i < 20; i++ {
// 		closeness, eccentricity := nodecentrality.ApproximateCompareNodeCentralityRandom(adjList, realCloseness, realEccentricity, graphName)

// 		resultsCloseness[i] = closeness
// 		resultsEccentricity[i] = eccentricity
// 	}
// 	randomAvgResultCloseness := ResultsDTOAvg(resultsCloseness)
// 	randomAvgResultEccentricity := ResultsDTOAvg(resultsEccentricity)

// 	// closeless
// 	for i := 0; i < 20; i++ {
// 		closeness, eccentricity := nodecentrality.ApproximateCompareNodeCentrality(adjList, "closeless", realCloseness, realEccentricity, graphName)

// 		resultsCloseness[i] = closeness
// 		resultsEccentricity[i] = eccentricity
// 	}
// 	closelessAvgResultCloseness := ResultsDTOAvg(resultsCloseness)
// 	closelessAvgResultEccentricity := ResultsDTOAvg(resultsEccentricity)

// 	//furtherBfs
// 	for i := 0; i < 20; i++ {
// 		closeness, eccentricity := nodecentrality.ApproximateCompareNodeCentrality(adjList, "furtherBfsed", realCloseness, realEccentricity, graphName)

// 		resultsCloseness[i] = closeness
// 		resultsEccentricity[i] = eccentricity
// 	}
// 	furtherBfsAvgResultCloseness := ResultsDTOAvg(resultsCloseness)
// 	furtherBfsAvgResultEccentricity := ResultsDTOAvg(resultsEccentricity)

// 	//output

// }

// func ResultsDTOAvg(results []nodecentrality.ResultsDTO) nodecentrality.ResultsDTO {
// 	n := len(results)
// 	avg := NewResultDTO(len(results.Mse[0]))
// 	for _, r := range results {
// 		avg.Aux += float64(r.Aux) / n
// 		avg.Mse += r.Mse / n
// 		avg.SpearmanCorrelation += r.SpearmanCorrelation / n
// 		avg.SpearmanP += r.SpearmanP / n
// 		avg.Jaccard1percent += r.Jaccard1percent / n
// 		avg.Jaccard5percent += r.Jaccard5percent / n
// 	}

// 	return avg
// }

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
