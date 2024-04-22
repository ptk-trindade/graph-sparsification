package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/ptk-trindade/graph-sparsification/edgebetweenness"
	"github.com/ptk-trindade/graph-sparsification/utils"
)

func main() {
	fmt.Println("Hello, cmd!")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	divideValue, err := strconv.ParseBool(scanner.Text())
	if err != nil {
		fmt.Println("Error reading divide value")
		return
	}

	adjList, err := utils.ReadGraph(scanner)
	if err != nil {
		fmt.Println("Error reading graph")
		return
	}

	edgeWeight := edgebetweenness.EdgeBetweenness(adjList, divideValue)
	fmt.Println(edgeWeight)
}
