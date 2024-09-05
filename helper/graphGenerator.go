package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// input: nodeQty edgeProb
func main() {
	// rand.Seed(time.Now().UnixNano())
	fmt.Println("compare")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	//scan 2 values (int, float) in same line (eg. 10 0.5)
	line := scanner.Text()

	lineSlice := strings.Fields(line)

	nodeQty, err := strconv.Atoi(lineSlice[0])
	if err != nil {
		fmt.Println("Error reading node quantity")
		panic(err)
	}

	edgeProb, err := strconv.ParseFloat(lineSlice[1], 64)
	if err != nil {
		fmt.Println("Error reading edge probability")
		panic(err)
	}

	if edgeProb < 0 || edgeProb > 1 {
		fmt.Println("Invalid edge probability")
		return
	}

	edges := make([]string, 0) //[]string{"0 1", "0 2", "2 3", ...}
	for i := 0; i < nodeQty; i++ {
		for j := i + 1; j < nodeQty; j++ {
			if rand.Float64() < edgeProb {
				edges = append(edges, fmt.Sprintf("%d %d", i, j))
			}
		}
	}

	// output
	fmt.Println(nodeQty)
	fmt.Println(len(edges))
	for _, edge := range edges {
		fmt.Println(edge)
	}

}
