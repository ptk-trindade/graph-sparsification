#!/bin/bash

# Run for the first input file and save output
cat ./txtFiles/inputs/ABCommunityD/abcd_graph_10.txt | go run ./cmd/main.go > results_abcd.txt

# Run for the second input file and save output
cat ./txtFiles/inputs/erdosRenyi/erdosRenyi_4000.txt | go run ./cmd/main.go > results_erdos.txt

# Run for the third input file and save output
cat ./txtFiles/inputs/real_graphs/facebook.txt | go run ./cmd/main.go > results_fb.txt

# Run for the fourth input file and save output
cat ./txtFiles/inputs/real_graphs/CA-GrQc.txt | go run ./cmd/main.go > results_collab.txt
