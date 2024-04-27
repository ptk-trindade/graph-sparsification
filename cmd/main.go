package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ptk-trindade/graph-sparsification/edgebetweenness"
	effectiveresisitance "github.com/ptk-trindade/graph-sparsification/effectiveresistance"
	"github.com/ptk-trindade/graph-sparsification/utils"
)

func main() {
	fmt.Println("Start!")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	// divideValue, err := strconv.ParseBool(scanner.Text())
	// if err != nil {
	// 	fmt.Println("Error reading divide value")
	// 	return
	// }
	method := scanner.Text()
	if method != "effectiveresistance" && method != "edgebetweenness_divide" && method != "edgebetweenness_sum" && method != "compare" {
		fmt.Println("Invalid method")
		return
	}

	adjList, err := utils.ReadGraph(scanner)
	if err != nil {
		fmt.Println("Error reading graph")
		return
	}

	if method == "effectiveresistance" {
		edgeWeight := effectiveresisitance.EffectiveResistance(adjList)
		fmt.Println(edgeWeight)

	} else if method == "edgebetweenness_sum" || method == "edgebetweenness_divide" {
		divideValue := (method == "edgebetweenness_divide")

		edgeWeight := edgebetweenness.EdgeBetweenness(adjList, divideValue)
		fmt.Println(edgeWeight)
	} else if method == "compare" {
		edgeWeight := edgebetweenness.EdgeBetweenness(adjList, false)
		// edgeWeight.Normalize()
		fmt.Println("sum")
		edgeWeight.Show()

		edgeWeight = edgebetweenness.EdgeBetweenness(adjList, true)
		// edgeWeight.Normalize()
		fmt.Println("divide")
		edgeWeight.Show()

		edgeWeight = effectiveresisitance.EffectiveResistance(adjList)
		if edgeWeight == nil {
			fmt.Println("Error computing effective resistance")
			return
		}
		// edgeWeight.Normalize()
		fmt.Println("resistance")
		edgeWeight.Show()
	}
}
