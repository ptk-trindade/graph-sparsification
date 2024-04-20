package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ptk-trindade/graph-sparsification/edgebetweenness"
	"github.com/ptk-trindade/graph-sparsification/utils"
)

func main() {
	fmt.Println("Hello, cmd!")

	scanner := bufio.NewScanner(os.Stdin)
	adjList, err := utils.ReadGraph(scanner)
	if err != nil {
		fmt.Println("Error reading graph")
		return
	}

	edgeWeight := edgebetweenness.EdgeBetweenness(adjList)
	fmt.Println(edgeWeight)
}
