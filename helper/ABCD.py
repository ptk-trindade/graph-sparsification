import numpy as np
import networkx as nx
from itertools import combinations
from random import sample

def abcd_graph_generator(num_nodes, num_communities, intra_alpha, inter_alpha):
    """
    Generate a graph using the ABCD model.

    Parameters:
        num_nodes (int): Total number of nodes in the graph.
        num_communities (int): Number of communities.
        intra_alpha (float): Power-law exponent for intra-community degree distribution.
        inter_alpha (float): Power-law exponent for inter-community degree distribution.

    Returns:
        networkx.Graph: The generated graph.
    """
    # Step 1: Create communities
    community_sizes = [num_nodes // num_communities] * num_communities
    for i in range(num_nodes % num_communities):
        community_sizes[i] += 1  # Adjust sizes to sum to num_nodes
    
    # Assign nodes to communities
    communities = []
    node_id = 0
    for size in community_sizes:
        communities.append(list(range(node_id, node_id + size)))
        node_id += size
    
    # Step 2: Assign intra-community degrees
    intra_degrees = []
    for size in community_sizes:
        degrees = np.random.zipf(a=intra_alpha, size=size)
        
        # Ensure the sum of degrees is even
        if np.sum(degrees) % 2 != 0:
            degrees[np.random.randint(0, size)] += 1
        
        intra_degrees.append(degrees)
    
    # Step 3: Generate intra-community edges
    graph = nx.Graph()
    for comm_nodes, comm_degrees in zip(communities, intra_degrees):
        comm_graph = nx.configuration_model(comm_degrees)
        comm_graph = nx.Graph(comm_graph)  # Convert to simple graph
        comm_graph.remove_edges_from(nx.selfloop_edges(comm_graph))
        mapping = {old: new for old, new in zip(range(len(comm_graph)), comm_nodes)}
        comm_graph = nx.relabel_nodes(comm_graph, mapping)
        graph.update(comm_graph)
    
    # Step 4: Assign inter-community degrees
    inter_degrees = np.random.zipf(a=inter_alpha, size=num_nodes)
    
    # Step 5: Create inter-community edges
    for comm1, comm2 in combinations(range(num_communities), 2):
        comm1_nodes = communities[comm1]
        comm2_nodes = communities[comm2]
        potential_edges = [(u, v) for u in comm1_nodes for v in comm2_nodes]
        
        # Use inter-community degrees to probabilistically choose edges
        num_edges = min(len(potential_edges), inter_degrees[comm1] + inter_degrees[comm2])
        chosen_edges = sample(potential_edges, num_edges)
        graph.add_edges_from(chosen_edges)
    
    return graph


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



# Example: Generate an ABCD graph
num_nodes = 100
num_communities = 4
intra_alpha = 2.5
inter_alpha = 2.0

graph = abcd_graph_generator(num_nodes, num_communities, intra_alpha, inter_alpha)

# Visualize the graph
import matplotlib.pyplot as plt
pos = nx.spring_layout(graph)
nx.draw(graph, pos, with_labels=False, node_size=30)
plt.show()


output_file = "abcd_graph_10_v2.txt"
save_graph_to_file(graph, output_file)