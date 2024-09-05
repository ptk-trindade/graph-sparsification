package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/ptk-trindade/graph-sparsification/edgebetweenness"
	"github.com/ptk-trindade/graph-sparsification/effectiveresistance"
	"github.com/ptk-trindade/graph-sparsification/utils"
)

func min_(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// fmt.Println("Start!")
	// fmt.Println("avgDegree,stdDev,minDegree,maxDegree,components,smallest comp,biggest comp,EBS EBD rs,EBS EBD p-value,EBS EBD jcc,EBD ER rs,EBD ER p-value,EBD ER jcc,EBS ER rs,EBS ER p-value,EBS ER jcc")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	// divideValue, err := strconv.ParseBool(scanner.Text())
	// if err != nil {
	// 	fmt.Println("Error reading divide value")
	// 	return
	// }
	method := scanner.Text()
	if method != "effectiveresistance" && method != "edgebetweenness_divide" && method != "edgebetweenness_sum" && method != "compare" && method != "test" {
		fmt.Println("Invalid method")
		return
	}

	adjList, err := utils.ReadGraph(scanner)
	if err != nil {
		fmt.Println("Error reading graph")
		return
	}

	degrees := make([]int, len(adjList))
	avgDegree := 0.0
	for i := range adjList {
		degrees[i] = len(adjList[i])
		avgDegree += float64(len(adjList[i]))
	}

	avgDegree /= float64(len(adjList))

	stdDev := 0.0
	min := degrees[0]
	max := degrees[0]
	for _, d := range degrees {
		stdDev += (float64(d) - avgDegree) * (float64(d) - avgDegree)

		if d < min {
			min = d
		}

		if d > max {
			max = d
		}
	}
	stdDev = stdDev / float64(len(adjList))
	stdDev = math.Sqrt(stdDev)

	// fmt.Print(avgDegree, ",", stdDev, ",", min, ",", max, ",")

	if method == "effectiveresistance" {
		edgeWeight := effectiveresistance.EffectiveResistance(adjList)
		fmt.Println(edgeWeight)

	} else if method == "edgebetweenness_sum" || method == "edgebetweenness_divide" {
		var edgeWeight *utils.EdgeWeight
		if method == "edgebetweenness_divide" {
			edgeWeight = edgebetweenness.EdgeBetweennessDivide(adjList)
			edgeWeight.Show()
		} else {
			edgeWeight = edgebetweenness.EdgeBetweennessSum(adjList)
			edgeWeight.Show()
		}
		fmt.Println(edgeWeight)
	} else if method == "compare" {
		edgeWeightEBS, pathQty1 := edgebetweenness.EdgeBetweennessCountPaths(adjList, false)
		// edgeWeightEBS.Normalize()
		// fmt.Println("sum")
		// edgeWeight.Show()

		fmt.Print(pathQty1, ",")

		edgeWeightEBD := edgebetweenness.EdgeBetweennessDivide(adjList)
		// edgeWeightEBD.Normalize()
		// fmt.Println("divide")
		// edgeWeight.Show()

		edgeWeightER := effectiveresistance.EffectiveResistance(adjList)
		if edgeWeightER == nil {
			fmt.Println("Error computing effective resistance")
			return
		}
		// edgeWeightER.Normalize()

		rs, p := edgeWeightEBS.CompareSpearman(edgeWeightEBD)
		jcc := edgeWeightEBS.CompareJaccard(edgeWeightEBD, 0.01)
		fmt.Print(rs, ",", p, ",", jcc, ",")

		rs, p = edgeWeightEBD.CompareSpearman(edgeWeightER)
		jcc = edgeWeightEBD.CompareJaccard(edgeWeightER, 0.01)
		fmt.Print(rs, ",", p, ",", jcc, ",")

		rs, p = edgeWeightEBS.CompareSpearman(edgeWeightER)
		jcc = edgeWeightEBS.CompareJaccard(edgeWeightER, 0.01)
		fmt.Print(rs, ",", p, ",", jcc)
	} else if method == "test" {

		nodeQty := len(adjList)

		edgeWeightFull := edgebetweenness.EdgeBetweennessDivide(adjList)

		fmt.Printf("")
		// costs := []int{2 * nodeQty, 10 * nodeQty, int(math.Sqrt(float64(nodeQty))) * nodeQty}
		// label := []string{"2n", "10n", "n*sqrt(n)"}
		// for j := range costs {

		// 	fmt.Printf("\n\n%s\n", label[j])

		// start
		i := 0
		// topNodes := sortByNeighborQty(adjList)
		topNodes := scrambleNumbers(nodeQty)
		for i < 1000 {
			if i < 10 {
				i += 1
			} else {
				i += 10
			}

			bfsQty := nodeQty * i / 1000

			// CHANGE COST  HERE!
			cost := 10 * nodeQty

			depth := cost / bfsQty
			depth = min_(depth, nodeQty)

			edgeWeightPctg := edgebetweenness.EdgeBetweennessParcial(adjList, true, topNodes, bfsQty, depth, true)

			jcc01 := edgeWeightFull.CompareJaccard(edgeWeightPctg, 0.001)
			jcc05 := edgeWeightFull.CompareJaccard(edgeWeightPctg, 0.005)
			jcc1 := edgeWeightFull.CompareJaccard(edgeWeightPctg, 0.01)
			jcc2 := edgeWeightFull.CompareJaccard(edgeWeightPctg, 0.02)

			fmt.Printf("%d\t%d\t-\t%.4f\t%.4f\t%.4f\t%.4f\n", bfsQty, depth, jcc01, jcc05, jcc1, jcc2)

			// depth := costs[j] / bfsQty
			// depth = min_(depth, nodeQty)

			// edgeWeightPctg := edgebetweenness.EdgeBetweennessParcial(adjList, true, topNodes, bfsQty, depth, false)

			// jcc01 := edgeWeightFull.CompareJaccard(edgeWeightPctg, 0.001)
			// jcc05 := edgeWeightFull.CompareJaccard(edgeWeightPctg, 0.005)
			// jcc1 := edgeWeightFull.CompareJaccard(edgeWeightPctg, 0.01)
			// jcc2 := edgeWeightFull.CompareJaccard(edgeWeightPctg, 0.02)

			// fmt.Printf("\t%d\t%d\t-\t%.4f\t%.4f\t%.4f\t%.4f\n", bfsQty, depth, jcc01, jcc05, jcc1, jcc2)
		}
		// }
		// end
		return

		bfsQty := nodeQty
		bfsQtyStr := "n"
		topNodes = sortByNeighborQty(adjList)

		sqrtn := int(math.Sqrt(float64(nodeQty)))
		edgeWeightSqrtN := edgebetweenness.EdgeBetweennessParcial(adjList, true, topNodes, bfsQty, sqrtn, true)

		logn := int(math.Log2(float64(nodeQty)) * 2)
		edgeWeightLogN := edgebetweenness.EdgeBetweennessParcial(adjList, true, topNodes, bfsQty, logn, true)

		halfn := nodeQty / 2
		edgeWeightHalfN := edgebetweenness.EdgeBetweennessParcial(adjList, true, topNodes, bfsQty, halfn, true)

		fmt.Println("n:", nodeQty, "\nSqrt(n):", sqrtn, "\n2*Log(n):", logn, "\nn/2:", halfn)

		// Comparisons
		// fmt.Println("\n\nFull - Sqrt(n)")
		jcc01 := edgeWeightFull.CompareJaccard(edgeWeightSqrtN, 0.001)
		// fmt.Printf("0.1%%: %.4f\n", jcc01)

		jcc05 := edgeWeightFull.CompareJaccard(edgeWeightSqrtN, 0.005)
		// fmt.Printf("0.5%%: %.4f\n", jcc05)

		jcc1 := edgeWeightFull.CompareJaccard(edgeWeightSqrtN, 0.01)
		// fmt.Printf("1%%: %.4f\n", jcc1)

		jcc2 := edgeWeightFull.CompareJaccard(edgeWeightSqrtN, 0.02)
		// fmt.Printf("2%%: %.4f\n", jcc2)

		fmt.Printf("%s\t%d\tsqrt(n)\t%d\t-\t%.4f\t%.4f\t%.4f\t%.4f\n", bfsQtyStr, bfsQty, sqrtn, jcc01, jcc05, jcc1, jcc2)

		// ------------------------

		// fmt.Println("\n\nFull - Log(n)")
		jcc01 = edgeWeightFull.CompareJaccard(edgeWeightLogN, 0.001)
		// fmt.Printf("0.1%%: %.4f\n", jcc01)

		jcc05 = edgeWeightFull.CompareJaccard(edgeWeightLogN, 0.005)
		// fmt.Printf("0.5%%: %.4f\n", jcc05)

		jcc1 = edgeWeightFull.CompareJaccard(edgeWeightLogN, 0.01)
		// fmt.Printf("1%%: %.4f\n", jcc1)

		jcc2 = edgeWeightFull.CompareJaccard(edgeWeightLogN, 0.02)
		// fmt.Printf("2%%: %.4f\n", jcc2)

		fmt.Printf("%s\t%d\tlog(n)\t%d\t-\t%.4f\t%.4f\t%.4f\t%.4f\n", bfsQtyStr, bfsQty, logn, jcc01, jcc05, jcc1, jcc2)

		// ------------------------

		// fmt.Println("\n\nFull - n/2")
		jcc01 = edgeWeightFull.CompareJaccard(edgeWeightHalfN, 0.001)
		// fmt.Printf("0.1%%: %.4f\n", jcc01)

		jcc05 = edgeWeightFull.CompareJaccard(edgeWeightHalfN, 0.005)
		// fmt.Printf("0.5%%: %.4f\n", jcc05)

		jcc1 = edgeWeightFull.CompareJaccard(edgeWeightHalfN, 0.01)
		// fmt.Printf("1%%: %.4f\n", jcc1)

		jcc2 = edgeWeightFull.CompareJaccard(edgeWeightHalfN, 0.02)
		// fmt.Printf("2%%: %.4f\n", jcc2)

		fmt.Printf("%s\t%d\tn/2\t%d\t-\t%.4f\t%.4f\t%.4f\t%.4f\n", bfsQtyStr, bfsQty, halfn, jcc01, jcc05, jcc1, jcc2)

		// ------------------------

		// fmt.Println("\n\nSqrt(n) - Log(n)")
		// jcc01 = edgeWeightSqrtN.CompareJaccard(edgeWeightLogN, 0.001)
		// // fmt.Printf("0.1%%: %.4f\n", jcc01)

		// jcc05 = edgeWeightSqrtN.CompareJaccard(edgeWeightLogN, 0.005)
		// // fmt.Printf("0.5%%: %.4f\n", jcc05)

		// jcc1 = edgeWeightSqrtN.CompareJaccard(edgeWeightLogN, 0.01)
		// // fmt.Printf("1%%: %.4f\n", jcc1)

		// jcc2 = edgeWeightSqrtN.CompareJaccard(edgeWeightLogN, 0.02)
		// // fmt.Printf("2%%: %.4f\n", jcc2)

		// // ------------------------

		// fmt.Println("\n\nSqrt(n) - n/2")
		// jcc01 = edgeWeightSqrtN.CompareJaccard(edgeWeightHalfN, 0.001)
		// // fmt.Printf("0.1%%: %.4f\n", jcc01)

		// jcc05 = edgeWeightSqrtN.CompareJaccard(edgeWeightHalfN, 0.005)
		// // fmt.Printf("0.5%%: %.4f\n", jcc05)

		// jcc1 = edgeWeightSqrtN.CompareJaccard(edgeWeightHalfN, 0.01)
		// // fmt.Printf("1%%: %.4f\n", jcc1)

		// jcc2 = edgeWeightSqrtN.CompareJaccard(edgeWeightHalfN, 0.02)
		// // fmt.Printf("2%%: %.4f\n", jcc2)

		// // ------------------------

		// fmt.Println("\n\nLog(n) - n/2")
		// jcc01 = edgeWeightLogN.CompareJaccard(edgeWeightHalfN, 0.001)
		// // fmt.Printf("0.1%%: %.4f\n", jcc01)

		// jcc05 = edgeWeightLogN.CompareJaccard(edgeWeightHalfN, 0.005)
		// // fmt.Printf("0.5%%: %.4f\n", jcc05)

		// jcc1 = edgeWeightLogN.CompareJaccard(edgeWeightHalfN, 0.01)
		// // fmt.Printf("1%%: %.4f\n", jcc1)

		// jcc2 = edgeWeightLogN.CompareJaccard(edgeWeightHalfN, 0.02)
		// // fmt.Printf("2%%: %.4f\n", jcc2)

	}
}

// Top 10 nodes with higher degree
type Node struct {
	Index       int
	NeighborQty int
}

func sortByNeighborQty(adjList [][]int) []int {
	neighborQtys := make([]Node, len(adjList))
	for i, neighbors := range adjList {
		neighborQtys[i] = Node{Index: i, NeighborQty: len(neighbors)}
	}

	// Sort nodes by neighbor count in descending order
	sort.Slice(neighborQtys, func(i, j int) bool {
		return neighborQtys[i].NeighborQty > neighborQtys[j].NeighborQty
	})

	// Get the top n_value nodes (or less if there are fewer than n_value nodes)
	tops := make([]int, 0)
	for i := 0; i < len(neighborQtys); i++ {
		tops = append(tops, neighborQtys[i].Index)
	}

	return tops
}

// scrambleNumbers generates a slice [0, 1, ... n] scrambled
func scrambleNumbers(n int) []int {
	rand.Seed(time.Now().UnixNano())

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}

	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}

	return slice
}
