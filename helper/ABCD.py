import random
import numpy as np
import networkx as nx
import matplotlib.pyplot as plt

def generate_abcd_graph(n, avg_degree, mu, num_communities):
    """
    Generate a random graph using the ABCD method.
    
    Parameters:
    - n: Number of nodes
    - avg_degree: Average degree of nodes
    - mu: Mixing parameter (controls intra- and inter-community edges)
    - num_communities: Number of communities
    
    Returns:
    - G: Generated graph with community structure
    """
    # Step 1: Assign nodes to communities
    community_sizes = [n // num_communities] * num_communities
    for i in range(n % num_communities):
        community_sizes[i] += 1
    
    communities = []
    node_idx = 0
    for size in community_sizes:
        communities.append(list(range(node_idx, node_idx + size)))
        node_idx += size
    
    # Step 2: Assign degrees to nodes (power-law distribution)
    degrees = np.random.zipf(a=2.0, size=n)  # Zipf distribution (approx. power-law)
    degrees = np.clip(degrees, 1, 2 * avg_degree)  # Clip to avoid overly large degrees
    degree_sum = sum(degrees)
    
    # Normalize degrees to match the expected total degree sum
    scaling_factor = (n * avg_degree) / degree_sum
    degrees = np.round(degrees * scaling_factor).astype(int)
    
    # Step 3: Create edges
    G = nx.Graph()
    G.add_nodes_from(range(n))
    
    for node in range(n):
        community = next(c for c in communities if node in c)
        intra_community_edges = int((1 - mu) * degrees[node])  # Intra-community edges
        inter_community_edges = degrees[node] - intra_community_edges  # Inter-community edges
        
        # Step 3a: Add intra-community edges
        possible_intra_edges = [v for v in community if v != node]
        if possible_intra_edges:
            intra_edges = random.sample(possible_intra_edges, min(intra_community_edges, len(possible_intra_edges)))
            G.add_edges_from([(node, v) for v in intra_edges])
        
        # Step 3b: Add inter-community edges
        possible_inter_edges = [v for v in range(n) if v not in community and v != node]
        if possible_inter_edges:
            inter_edges = random.sample(possible_inter_edges, min(inter_community_edges, len(possible_inter_edges)))
            G.add_edges_from([(node, v) for v in inter_edges])
    
    return G

def save_graph_to_file(G, file_path):
    """
    Save the graph to a file in the format:
    <quantity of nodes>
    <quantity of edges>
    <edge1>
    <edge2>
    ...
    
    Parameters:
    - G: The graph (networkx.Graph)
    - file_path: Output file path
    """
    with open(file_path, 'w') as f:
        # Write number of nodes
        f.write(f"{G.number_of_nodes()}\n")
        
        # Write number of edges
        f.write(f"{G.number_of_edges()}\n")
        
        # Write edges
        for u, v in G.edges():
            f.write(f"{u} {v}\n")

# Parameters for the ABCD graph
n = 4000  # Number of nodes
avg_degree = 40  # Average degree
mu = 0.3  # Mixing parameter (closer to 0 = strong community structure)
num_communities = 50  # Number of communities

# Generate the graph
G = generate_abcd_graph(n, avg_degree, mu, num_communities)

# Save the graph to a file
output_file = "abcd_graph.txt"
save_graph_to_file(G, output_file)

# # Plot the graph
# pos = nx.spring_layout(G)  # Layout for visualization
# nx.draw(G, pos, with_labels=True, node_color='lightblue', node_size=500, edge_color='gray')
# plt.show()
