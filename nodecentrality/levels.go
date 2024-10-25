package nodecentrality

import (
	"github.com/ptk-trindade/graph-sparsification/utils"
	"math/rand"
)

type Levels struct {
	levels [][]int // [node][iteration]level
	sum    []int
	min    []int
	max    []int
}

func NewLevels(nodeQty int) *Levels {
	return &Levels{
		levels: make([][]int, nodeQty),
		sum:    make([]int, nodeQty),
		min:    make([]int, nodeQty),
		max:    make([]int, nodeQty),
	}
}

func (lvl *Levels) AddLevels(levels []int) {
	for node, level := range levels {
		lvl.levels[node] = append(lvl.levels[node], level)
	}
}

/*
Finds the Corner Node (the node with higher sum of levels)
output: id of CornerNode
*/
func (lvl Levels) GetCornerNode() (int, bool) {
	var cornerNode, maxSum int
	wasCornerBefore := false
	for node, nodeLevels := range lvl.levels {

		currSum := 0
		wasCB := false // this node was a corner before
		for _, level := range nodeLevels {
			currSum += level
			if level == 0 {
				wasCB = true
			}
		}

		if currSum > maxSum {
			maxSum = currSum
			cornerNode = node
			wasCornerBefore = wasCB
		}
	}

	return cornerNode, wasCornerBefore
}

/*
Finds the node with lowest Closeness (the node with higher avg of levels) based on the BFSs runned up to now
output: id of CornerNode, Closeless node was already picked as a corner (therefore wasn't picked again)
*/
func (lvl Levels) GetCloselessNode() (int, bool) {
	var pickingNode, equalyGoodOptions int
	var maxAvgDistance, pickingAvgDistance float64

	for node, nodeLevels := range lvl.levels {

		wasCB := includes(nodeLevels, 0)
		avgDistance := utils.Avg(nodeLevels...)

		if avgDistance > maxAvgDistance {
			maxAvgDistance = avgDistance
		}

		if !wasCB {
			if avgDistance > pickingAvgDistance {
				pickingAvgDistance = avgDistance
				pickingNode = node

			} else if avgDistance == pickingAvgDistance { // a tie happened, will run the tie breaker
				equalyGoodOptions++
				if rand.Intn(equalyGoodOptions) == 0 { // makes sure every node has the same probability of being chosen
					pickingAvgDistance = avgDistance
					pickingNode = node
				}

			}
		}
	}

	closelessWasCornerBefore := (maxAvgDistance > pickingAvgDistance)
	return pickingNode, closelessWasCornerBefore
}

/*
Finds the node with highest min distance from a 'BFSed' node
output: id of CornerNode, distance from the closer 'BFSed' node
*/
func (lvl Levels) GetFurtherBfsedNode() (int, int) {
	var cornerNode, closerBfsedNode, equalyCloseOptions int
	
	for node, nodeLevels := range lvl.levels {

		closerBN := utils.Min(nodeLevels...)

		if closerBN > closerBfsedNode {
			cornerNode = node
			closerBfsedNode = closerBN
			equalyCloseOptions = 1

		} else if closerBN == closerBfsedNode { // current node is as close as closerBfsedNode
			equalyCloseOptions++
			if rand.Intn(equalyCloseOptions) == 0 { // makes sure every node has the same probability of being chosen
				cornerNode = node
				closerBfsedNode = closerBN
			}
		}
	}

	return cornerNode, closerBfsedNode
}

func (lvl Levels) GetRandomNode() int {

	numbers := make([]int, len(lvl.levels))
	for i := range numbers {
		numbers[i] = i
	}
	utils.Scramble(numbers)

	for _, nodeId := range numbers {

		wasCB := includes(lvl.levels[nodeId], 0)

		if !wasCB {
			return nodeId
		}
	}

	return -1
}

/* Get levels avg, max and min for every node
 */
func (lvl Levels) GetMetrics() ([]float64, []int) {
	avgs := make([]float64, len(lvl.levels))
	maxs := make([]int, len(lvl.levels))

	for node, nodeLevels := range lvl.levels {
		avgs[node] = utils.Avg(nodeLevels...)
		maxs[node] = utils.Max(nodeLevels...)
	}

	return avgs, maxs
}

/*
Methods to determine corner:
- max sum of levels
- max min level
	after some iterations nodes on the center might become considered "corners"
	this might be a good way to find well spartified nodes

Methods do determine center:
- min sum of levels
- min max level (centers are close to every node)
- max min level (counterintuitive, but maybe this could be used to determine either corners or centers)
*/

func includes[E comparable](slice []E, findVal E) bool {
	for _, val := range slice {
		if val == findVal {
			return true
		}
	}
	return false
}
