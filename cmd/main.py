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


n = graph.numberOfNodes()
intervals = [(1, 100, 1), (100, 1000, 10), (1000, n, 100), (n, n+1, 1)]
sampleSizes = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 25, 26, 27, 28, 30, 31, 33, 34, 36, 38, 39, 41, 43, 45, 47, 50, 52, 54, 57, 60, 63, 66, 69, 72, 75, 79, 83, 87, 91, 95, 100, 104, 109, 114, 120, 125, 131, 138, 144, 151, 158, 165, 173, 181, 190, 199, 208, 218, 229, 239, 251, 263, 275, 288, 301, 316, 331, 346, 363, 380, 398, 416, 436, 457, 478, 501, 524, 549, 575, 602, 630, 660, 691, 724, 758, 794, 831, 870, 912, 954, 1000, 1047, 1096, 1148, 1202, 1258, 1318, 1380, 1445, 1513, 1584, 1659, 1737, 1819, 1905, 1995, 2089, 2187, 2290, 2398, 2511, 2630, 2754, 2884, 3019, 3162, 3311, 3467, 3630, 3801, 3981, n]
nRuns = 20

mse_values = [0] * len(sampleSizes)  # Store MSE for each sample size
spearman_values = [0] * len(sampleSizes)  # Store Spearman correlation for each sample size
jaccard1_values = [0] * len(sampleSizes)
jaccard5_values = [0] * len(sampleSizes)

for ss_i, nSamples in enumerate(sampleSizes):
    start = time.time()
    for _ in range(nRuns):
        # Run ApproxCloseness with nSamples
        approx_closeness = nk.centrality.ApproxCloseness(graph, nSamples, epsilon, normalized, nk.centrality.ClosenessType.OUTBOUND)
        approx_closeness.run()
        approx_closeness_scores = approx_closeness.scores()
        
        # Calculate the MSE between approximated and exact closeness
        mse = calculate_mse(approx_closeness_scores, exact_closeness_scores)
        mse_values[ss_i] += (mse/nRuns)
        
        # Calculate Spearman correlation between approximated and exact closeness
        spearman_corr, spearman_p = spearmanr(approx_closeness_scores, exact_closeness_scores)
        spearman_values[ss_i] += (spearman_corr/nRuns)

        jaccard1 = compare_jaccard(approx_closeness_scores, exact_closeness_scores, 0.01)
        jaccard1_values[ss_i] += (jaccard1/nRuns)

        jaccard5 = compare_jaccard(approx_closeness_scores, exact_closeness_scores, 0.05)
        jaccard5_values[ss_i] += (jaccard5/nRuns)

    # Print MSE and Spearman correlation for this number of samples
    took = time.time() - start
    print(f"nSamples = {nSamples}, took: {took}")


write_lists_to_file([sampleSizes,mse_values,spearman_values,jaccard1_values,jaccard5_values], "networkit_results.csv")

