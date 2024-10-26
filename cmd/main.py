from scipy.stats import spearmanr

import csv
import matplotlib.pyplot as plt
import networkit as nk
import numpy as np


def write_lists_to_file(lists, filename):
    # Transpose the list of lists to get columns
    columns = list(zip(*lists))
    
    with open(filename, 'w', newline='') as file:
        writer = csv.writer(file)
        writer.writerows(columns)


def load_graph_from_file(filename):
    # Open the file and read the data
    with open(filename, 'r') as file:
        # Read the number of nodes and edges
        num_nodes = int(file.readline().strip())
        num_edges = int(file.readline().strip())

        # Initialize an undirected graph with the given number of nodes
        graph = nk.graph.Graph(num_nodes, weighted=False, directed=False)

        # Read the edges and add them to the graph
        for _ in range(num_edges):
            node1, node2 = map(int, file.readline().strip().split())
            graph.addEdge(node1, node2)

    return graph

def calculate_mse(approx_scores, exact_scores):
    """Compute Mean Squared Error between two sets of scores."""
    for i in range(len(exact_scores)):
        diff = abs(approx_scores[i] - exact_scores[i]) / exact_scores[i]
        if diff > 1:
            print(i, approx_scores[i], exact_scores[i])

    return np.mean((np.array(approx_scores) - np.array(exact_scores)) ** 2)

def compare_jaccard(metric1, metric2, topK):
    def find_top_k_threshold(slice, topK):
        copied_slice = sorted(slice)
        n = len(copied_slice)
        top_k_index = int(n * (1.0 - topK))

        threshold = copied_slice[top_k_index]
        return threshold

    if len(metric1) != len(metric2):
        print("Error CompareJaccard: metrics should have same length")
        return 0.0

    threshold1 = find_top_k_threshold(metric1, topK)
    threshold2 = find_top_k_threshold(metric2, topK)

    # get the intersection
    intersection = 0
    union = 0
    for m1, m2 in zip(metric1, metric2):
        if m1 >= threshold1 or m2 >= threshold2:  # at least one number is high
            union += 1
            if m1 >= threshold1 and m2 >= threshold2:  # both numbers are high
                intersection += 1

    return intersection / union if union != 0 else 0.0

# Load the graph from the file
graph = load_graph_from_file(r"txtFiles\inputs\erdosRenyi\erdosRenyi_4000.txt")

# 1. Calculate the exact closeness centrality (this will be our reference)
normalized = True # Closeness normalization

exact_closeness = nk.centrality.Closeness(graph, normalized, nk.centrality.ClosenessVariant.Standard)
exact_closeness.run()
exact_closeness_scores = exact_closeness.scores()

# 2. Run ApproxCloseness with increasing sample sizes and calculate MSE and Spearman correlation
epsilon = 0.1     # Precision

n_samples = ["n_samples"]
mse_values = ["mse_values"]   # Store MSE for each sample size
spearman_values = ["spearman_values"]  # Store Spearman correlation for each sample size
spearman_ps = ["spearman_ps"]
jaccard1_values = ["jaccard1_values"]
jaccard5_values = ["jaccard5_values"]

# (min, max, step)
n = graph.numberOfNodes()
intervals = [(1, 100, 1), (100, 1000, 10), (1000, n, 100), (n, n+1, 1)]
for interval in intervals:
    start, end, step = interval
    for nSamples in range(start, end, 1):
        n_samples.append(nSamples)
        
        if (nSamples - start) % step != 0:
            mse_values.append('')
            spearman_values.append('')
            spearman_ps.append('')
            jaccard1_values.append('')
            jaccard5_values.append('')
            continue

        # Run ApproxCloseness with nSamples
        approx_closeness = nk.centrality.ApproxCloseness(graph, nSamples, epsilon, normalized, nk.centrality.ClosenessType.OUTBOUND)
        approx_closeness.run()
        approx_closeness_scores = approx_closeness.scores()
        
        # Calculate the MSE between approximated and exact closeness
        mse = calculate_mse(approx_closeness_scores, exact_closeness_scores)
        mse_values.append(mse)
        
        # Calculate Spearman correlation between approximated and exact closeness
        spearman_corr, spearman_p = spearmanr(approx_closeness_scores, exact_closeness_scores)
        spearman_values.append(spearman_corr)
        spearman_ps.append(spearman_p)

        jaccard1 = compare_jaccard(approx_closeness_scores, exact_closeness_scores, 0.01)
        jaccard1_values.append(jaccard1)

        jaccard5 = compare_jaccard(approx_closeness_scores, exact_closeness_scores, 0.05)
        jaccard5_values.append(jaccard5)

        # Print MSE and Spearman correlation for this number of samples
        print(f"nSamples = {nSamples}, MSE = {mse}, Spearman Correlation = {spearman_corr}")


write_lists_to_file([n_samples,mse_values,spearman_values,spearman_ps,jaccard1_values,jaccard5_values], "networkit_results.csv")

