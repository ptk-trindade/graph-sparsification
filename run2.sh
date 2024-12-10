#!/bin/bash

# Run for the first input file and save output
go run ./cmd/main.go < ./txtFiles/inputs/ABCommunityD/abcd_graph.txt > abcd_output.txt

# Run for the second input file and save output
go run ./cmd/main.go < ./txtFiles/inputs/erdosRenyi/erdosRenyi.txt > erdos_output.txt

# Run for the third input file and save output
go run ./cmd/main.go < ./txtFiles/inputs/real_graphs/facebook.txt > fb_output.txt

# Run for the fourth input file and save output
go run ./cmd/main.go < ./txtFiles/inputs/real_graphs/CA-GrQc.txt > collab_output.txt
