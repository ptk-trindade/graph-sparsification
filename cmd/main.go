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
	if method != "effectiveresistance" && method != "edgebetweenness_divide" && method != "edgebetweenness_sum" {
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
		return
	}

	divideValue := (method == "edgebetweenness_divide")

	edgeWeight := edgebetweenness.EdgeBetweenness(adjList, divideValue)
	fmt.Println(edgeWeight)
}
