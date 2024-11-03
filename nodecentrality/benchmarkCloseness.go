package nodecentrality

import (
	"math/rand"
)

type ClosenessNode struct {
	CloserBfs         int
	ExpectedCloseness float64
	RealCloseness     float64
}

// Fills ExpectedCloseness and Calculates expected MSE
func calculateExpectedMseCloseness(closenessNodes []ClosenessNode, nSamples int, distanceSum []int) float64 {
	var expectedMse float64
	for i := range closenessNodes {
		if closenessNodes[i].CloserBfs == 0 { // was explored
			closenessNodes[i].ExpectedCloseness = float64(nSamples-1) / float64(distanceSum[i])

			diff := (closenessNodes[i].ExpectedCloseness - closenessNodes[i].RealCloseness)
			expectedMse += (diff * diff)
		} else {
			closenessNodes[i].ExpectedCloseness = float64(nSamples) / float64(distanceSum[i])
		}
	}
	expectedMse /= float64(nSamples)
	return expectedMse
}

func ApproximateClosenessCloseless(adjList [][]int, nSamples int) ([]ClosenessNode, float64) {
	closenessNodes := make([]ClosenessNode, len(adjList))
	for i := range closenessNodes {
		closenessNodes[i].CloserBfs = len(adjList)
	}

	distanceSum := make([]int, len(adjList))

	cornerNode := rand.Intn(len(adjList)) // first node is randomly choosen
	for iteration := 0; iteration < nSamples; iteration++ {
		levels := bfs(adjList, cornerNode)
		closenessNodes[cornerNode].CloserBfs = 0

		var lvlsSum, nextCornerNode, higherDistanceSum, equallyGoodOptions int
		for i, lvl := range levels {
			distanceSum[i] += lvl
			lvlsSum += lvl
			// closenessNodes[i].CloserBfs = min(closenessNodes[i].CloserBfs, lvl) // doesn't matter here, it only matters when 0, so I update it outside the loop

			// looks for next corner
			if distanceSum[i] >= higherDistanceSum && closenessNodes[i].CloserBfs > 0 {
				if distanceSum[i] > higherDistanceSum {
					higherDistanceSum = distanceSum[i]
					nextCornerNode = i
					equallyGoodOptions = 1
				} else { //distanceSum[i] == higherDistanceSum
					equallyGoodOptions++
					if rand.Intn(equallyGoodOptions) == 0 {
						higherDistanceSum = distanceSum[i]
						nextCornerNode = i
					}
				}
			}
		}

		closenessNodes[cornerNode].RealCloseness = float64(len(adjList)-1) / float64(lvlsSum)

		cornerNode = nextCornerNode
	}

	// Fills ExpectedCloseness and Calculates expected MSE
	expectedMse := calculateExpectedMseCloseness(closenessNodes, nSamples, distanceSum)

	return closenessNodes, expectedMse
}

func ApproximateClosenessFurtherBfs(adjList [][]int, nSamples int) ([]ClosenessNode, float64) {
	closenessNodes := make([]ClosenessNode, len(adjList))
	for i := range closenessNodes {
		closenessNodes[i].CloserBfs = len(adjList)
	}

	distanceSum := make([]int, len(adjList))

	cornerNode := rand.Intn(len(adjList)) // first node is randomly choosen
	for iteration := 0; iteration < nSamples; iteration++ {
		levels := bfs(adjList, cornerNode)

		var lvlsSum, nextCornerNode, bfsDistance, equallyGoodOptions int
		for i, lvl := range levels {
			distanceSum[i] += lvl
			lvlsSum += lvl
			closenessNodes[i].CloserBfs = min(closenessNodes[i].CloserBfs, lvl)

			// looks for next corner
			if closenessNodes[i].CloserBfs >= bfsDistance {
				if closenessNodes[i].CloserBfs > bfsDistance {
					bfsDistance = closenessNodes[i].CloserBfs
					nextCornerNode = i
					equallyGoodOptions = 1
				} else { // closenessNodes[i].CloserBfs == bfsDistance
					equallyGoodOptions++
					if rand.Intn(equallyGoodOptions) == 0 {
						bfsDistance = closenessNodes[i].CloserBfs
						nextCornerNode = i
					}
				}
			}
		}

		closenessNodes[cornerNode].RealCloseness = float64(len(adjList)-1) / float64(lvlsSum)

		cornerNode = nextCornerNode
	}

	// Fills ExpectedCloseness and Calculates expected MSE
	expectedMse := calculateExpectedMseCloseness(closenessNodes, nSamples, distanceSum)
	return closenessNodes, expectedMse
}

func pickRandomNumbers(maxNumber int, samples int) []int {
	numbers := make([]int, maxNumber)

	for i := range numbers {
		numbers[i] = i
	}

	// Shuffle the slice using the Fisher-Yates algorithm
	for i := 0; i < min(samples, maxNumber-1); i++ {
		j := i + rand.Intn(maxNumber-1-i)               // Generate a random index from 0 to i
		numbers[i], numbers[j] = numbers[j], numbers[i] // Swap elements
	}

	return numbers[:samples]
}

func ApproximateClosenessRandom(adjList [][]int, nSamples int) ([]ClosenessNode, float64) {
	closenessNodes := make([]ClosenessNode, len(adjList))
	for i := range closenessNodes {
		closenessNodes[i].CloserBfs = len(adjList)
	}

	samplesIds := pickRandomNumbers(len(adjList), nSamples)
	distanceSum := make([]int, len(adjList))

	for _, sampleId := range samplesIds {
		levels := bfs(adjList, sampleId)

		var lvlsSum int
		for i, lvl := range levels {
			distanceSum[i] += lvl
			lvlsSum += lvl
			// closenessNodes[i].CloserBfs = min(closenessNodes[i].CloserBfs, lvl) // doesn't matter here, it only matters when 0, so I update it outside the loop
		}
		closenessNodes[sampleId].CloserBfs = 0

		closenessNodes[sampleId].RealCloseness = float64(len(adjList)-1) / float64(lvlsSum)
	}

	// Fills ExpectedCloseness and Calculates expected MSE
	expectedMse := calculateExpectedMseCloseness(closenessNodes, nSamples, distanceSum)

	return closenessNodes, expectedMse
}
