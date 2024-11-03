package nodecentrality

import "math/rand"

type EccentricityNode struct {
	CloserBfs            int
	ExpectedEccentricity int
	RealEccentricity     int
	// ExpectedCloseness    float64
}

func calculateExpectedMseEccentricity(eccentricityNodes []EccentricityNode, nSamples int) float64 {
	var expectedSqrErr int
	for _, node := range eccentricityNodes {
		if node.CloserBfs == 0 { // was explored
			diff := (node.RealEccentricity - node.ExpectedEccentricity)
			expectedSqrErr += (diff * diff)
		}
	}
	expectedMse := float64(expectedSqrErr) / float64(nSamples)
	return expectedMse
}

func ApproximateEccentricityCloseless(adjList [][]int, nSamples int) ([]EccentricityNode, float64) {
	eccentricityNodes := make([]EccentricityNode, len(adjList))
	for i := range eccentricityNodes {
		eccentricityNodes[i].CloserBfs = len(adjList)
	}

	distanceSum := make([]int, len(adjList))

	cornerNode := rand.Intn(len(adjList)) // first node is randomly choosen
	for iteration := 0; iteration < nSamples; iteration++ {
		levels := bfs(adjList, cornerNode)
		eccentricityNodes[cornerNode].CloserBfs = 0

		var nextCornerNode, higherDistanceSum, equallyGoodOptions int
		for i, lvl := range levels {
			distanceSum[i] += lvl
			eccentricityNodes[i].ExpectedEccentricity = max(eccentricityNodes[i].ExpectedEccentricity, lvl)
			eccentricityNodes[cornerNode].RealEccentricity = max(eccentricityNodes[cornerNode].RealEccentricity, lvl)
			// eccentricityNodes[i].CloserBfs = min(eccentricityNodes[i].CloserBfs, lvl) // doesn't matter here, it only matters when 0, so I update it outside the loop

			// looks for next corner
			if distanceSum[i] >= higherDistanceSum && eccentricityNodes[i].CloserBfs > 0 {
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

		eccentricityNodes[cornerNode].ExpectedEccentricity = 0

		cornerNode = nextCornerNode
	}

	// Calculates expected MSE
	expectedMse := calculateExpectedMseEccentricity(eccentricityNodes, nSamples)

	return eccentricityNodes, expectedMse
}

func ApproximateEccentricityFurtherBfs(adjList [][]int, nSamples int) ([]EccentricityNode, float64) {
	eccentricityNodes := make([]EccentricityNode, len(adjList))
	for i := range eccentricityNodes {
		eccentricityNodes[i].CloserBfs = len(adjList)
	}

	distanceSum := make([]int, len(adjList))
	cornerNode := rand.Intn(len(adjList)) // first node is randomly chosen

	for iteration := 0; iteration < nSamples; iteration++ {
		levels := bfs(adjList, cornerNode)

		var nextCornerNode, bfsDistance, equallyGoodOptions int
		for i, lvl := range levels {
			distanceSum[i] += lvl
			eccentricityNodes[i].ExpectedEccentricity = max(eccentricityNodes[i].ExpectedEccentricity, lvl)
			eccentricityNodes[cornerNode].RealEccentricity = max(eccentricityNodes[cornerNode].RealEccentricity, lvl)
			eccentricityNodes[i].CloserBfs = min(eccentricityNodes[i].CloserBfs, lvl)

			// Selects the next corner node
			if eccentricityNodes[i].CloserBfs >= bfsDistance {
				if eccentricityNodes[i].CloserBfs > bfsDistance {
					bfsDistance = eccentricityNodes[i].CloserBfs
					nextCornerNode = i
					equallyGoodOptions = 1
				} else {
					equallyGoodOptions++
					if rand.Intn(equallyGoodOptions) == 0 {
						bfsDistance = eccentricityNodes[i].CloserBfs
						nextCornerNode = i
					}
				}
			}
		}

		cornerNode = nextCornerNode
	}

	// Calculates expected MSE
	expectedMse := calculateExpectedMseEccentricity(eccentricityNodes, nSamples)
	return eccentricityNodes, expectedMse
}

func ApproximateEccentricityRandom(adjList [][]int, nSamples int) ([]EccentricityNode, float64) {
	eccentricityNodes := make([]EccentricityNode, len(adjList))
	for i := range eccentricityNodes {
		eccentricityNodes[i].CloserBfs = len(adjList)
	}

	samplesIds := pickRandomNumbers(len(adjList), nSamples)
	distanceSum := make([]int, len(adjList))

	for _, sampleId := range samplesIds {
		levels := bfs(adjList, sampleId)

		for i, lvl := range levels {
			distanceSum[i] += lvl
			eccentricityNodes[i].ExpectedEccentricity = max(eccentricityNodes[i].ExpectedEccentricity, lvl)
			eccentricityNodes[sampleId].RealEccentricity = max(eccentricityNodes[sampleId].RealEccentricity, lvl)
		}
		eccentricityNodes[sampleId].CloserBfs = 0
	}

	// Calculates expected MSE
	expectedMse := calculateExpectedMseEccentricity(eccentricityNodes, nSamples)
	return eccentricityNodes, expectedMse
}
