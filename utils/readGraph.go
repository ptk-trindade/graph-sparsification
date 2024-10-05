package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ReadGraph(scanner *bufio.Scanner) ([][]int, error) {
	scanner.Scan()
	nodeQtyStr := scanner.Text()
	nodeQty, err := strconv.Atoi(nodeQtyStr)
	if err != nil {
		fmt.Println("Error reading node quantity", nodeQtyStr)
		return nil, err
	}

	scanner.Scan()
	edgeQty, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Error reading edge quantity")
		return nil, err
	}

	adjList := make([][]int, nodeQty)
	for i := 0; i < edgeQty; i++ {
		scanner.Scan()
		edge := scanner.Text()
		nodes := strings.Fields(edge)

		node1, err := strconv.Atoi(nodes[0])
		if err != nil {
			fmt.Println("Error reading node 1")
			return nil, err
		}

		node2, err := strconv.Atoi(nodes[1])
		if err != nil {
			fmt.Println("Error reading node 2")
			return nil, err
		}

		adjList[node1] = append(adjList[node1], node2)
		adjList[node2] = append(adjList[node2], node1)
	}

	return adjList, nil
}
