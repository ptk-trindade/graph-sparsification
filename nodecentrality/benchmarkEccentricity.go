package nodecentrality

import "math/rand"

type EccentricityNode struct {
	CloserBfs            int
	ExpectedEccentricity int
	RealEccentricity     int
	// ExpectedCloseness    float64
}

func ApproximateEccentricityCloseless(adjList [][]int, nSamples int) ([]EccentricityNode, float64) {
	eccentricityNodes := make([]EccentricityNode, len(adjList))
	for _, c := range eccentricityNodes {
		c.CloserBfs = len(adjList)
	}

	distanceSum := make([]int, len(adjList))

	cornerNode := rand.Intn(len(adjList)) // first node is randomly choosen
	for iteration := 0; iteration < nSamples; iteration++ {
		levels := bfs(adjList, cornerNode)
		eccentricityNodes[cornerNode].CloserBfs = 0

		var nextCornerNode, higherDistanceSum, equalyGoodOptions int
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
					equalyGoodOptions = 1
				} else { //distanceSum[i] == higherDistanceSum
					equalyGoodOptions++
					if rand.Intn(equalyGoodOptions) == 0 {
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
	var expectedSqrErr int
	for _, node := range eccentricityNodes {
		if node.CloserBfs == 0 { // was explored
			diff := (node.RealEccentricity - node.ExpectedEccentricity)
			expectedSqrErr += (diff * diff)
		}
	}
	expectedMse := float64(expectedSqrErr) / float64(nSamples)

	return eccentricityNodes, expectedMse
}
