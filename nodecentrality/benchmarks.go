package nodecentrality


import (
	"fmt"
	"math/rand"

	"github.com/ptk-trindade/graph-sparsification/utils"
)

type ClosenessNode struct {
	explored bool
	expectedCloseness float64
	realCloseness float64
}

func findCorner(nodes []ClosenessNode) int {
	var cornerNode, equalyGoodOptions int
	var cornerCloseness float64 = 1.0
	for i, node := range nodes {
		if node.expectedCloseness =< cornerCloseness && !node.explored {
			if node.expectedCloseness < cornerCloseness {
				cornerCloseness = node.expectedCloseness
				cornerNode = node
				equalyGoodOptions = 1
			} else if node.expectedCloseness == cornerCloseness {
				equalyGoodOptions++
				if rand.Intn(equalyGoodOptions) == 0 {
					cornerCloseness = node.expectedCloseness
					cornerNode = node
				}
			}
		}
	}

	return cornerNode
}

func ApproximateClosenessCloseless(adjList [][]int, nSamples int) []ClosenessNode, float64 {
	
	closenessNodes := make([]ClosenessNode, len(adjList))
	samples := make([]*ClosenessNode, nSamples)

	cornerNode := rand.Intn(len(adjList)) // first node is randomly choosen
	for iteration := 0; iteration < nSamples; iteration++ {
		levels := bfs(adjList, cornerNode)
		
		currCloseness := 0
		for i, lvl:= range levels {
			closenessNodes[i].expectedCloseness += float64(nSamples)/float64(lvl) // ((1/lvl)/nSamples)
			currCloseness += float64(lvl)
		}
		closenessNodes[i].expectedCloseness
		currCloseness /= float64(len(adjList))

		samples[iteration] = &closenessNodes[cornerNode]
		closenessNodes[cornerNode].explored = true
		closenessNodes[cornerNode].realCloseness = currCloseness

		cornerNode = findCorner(closenessNodes)
	}

	// calculate expected mse
	var mse float64
	for _, node := range samples {
		diff = (node.expectedCloseness - node.realCloseness)
		mse += (diff*diff)
	}
	mse /= nSamples
}