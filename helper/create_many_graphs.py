import sys
import os
import subprocess

def create_graph(N, avg_degree):
    # write to temp_input.txt
    with open("temp_input.txt", "w") as f:
        f.write(f"{str(N)} {str(avg_degree/N)[:5]}")
    
    printed = subprocess.call(["Get-Content", "temp_input.txt", "|", "./helper/graphGenerator.go" ">" "temp_graph.txt"])

def main():
    N = 1000
    avg_degree = 2

    # create graphs
    create_graph(N, avg_degree)

    # run the program
    printed = subprocess.call(["./main", "temp_graph.txt", "temp_output.txt"])