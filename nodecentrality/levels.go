package nodecentrality


type Levels struct {
	levels [][]int // [node][iteration]level
	sum []int
	min []int
	max []int
}


func (lvl &Levels) NewLevels(nodeQty int) *Levels {
	return &Levels{
		levels: make([][]int, nodeQty),
		average: make([]int, nodeQty),
		min: make([]int, nodeQty),
		max: make([]int, nodeQty),
	}
}

func (lvl &Levels) AddLevels(levels []int) {
	for node, level := range levels {
		lvl.levels[node] = append(lvl.levels[node], level)
	}
}

/*Finds the Corner Node (the node with higher sum of levels)
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

/* Get levels avg, max and min for every node
*/
func (lvl Levels) GetMetrics() (maxs []int, avgs []float64, mins []int) {

	mins = make([]int, len(lvl.levels))
	avgs = make([]float64, len(lvl.levels))
	maxs = make([]int, len(lvl.levels))
	
	for node, nodeLevels := range lvl.levels {
	
		currSum, currMax, currMin := 0, 0, 
		for _, level := range nodeLevels {
			currSum += level
			currMax = max(currMax, level)
			currMin = min(currMin, level)
		}

		avgs[node] = float64(currSum)/float64(len(lvl.levels))
		maxs[node] = currMax
		mins[node] = currMin
	}

	return mins, avgs, maxs
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