import numpy as np

def read_graph(file):
    with open(file, 'r') as infile:
        num_nodes = int(infile.readline().strip())
        num_edges = int(infile.readline().strip())

        edges = []
        for line in infile:
            node1, node2 = map(int, line.strip().split())
            edges.append((node1, node2))

    return num_nodes, edges

def find_components(num_nodes, edges):
    # Initialize adjacency list
    adjacency_list = {i: [] for i in range(num_nodes)}
    for node1, node2 in edges:
        adjacency_list[node1].append(node2)
        adjacency_list[node2].append(node1)

    visited = [False] * num_nodes
    components = []

    def dfs(node):
        stack = [node]
        size = 0
        while stack:
            current = stack.pop()
            if not visited[current]:
                visited[current] = True
                size += 1
                stack.extend(adjacency_list[current])
        return size

    for node in range(num_nodes):
        if not visited[node]:
            component_size = dfs(node)
            components.append(component_size)

    return components


# Function to compute the diameter of the graph
def graph_diameter(num_nodes, edges):
    from collections import deque

    def bfs(graph, start_node, num_nodes):
        distances = [-1] * num_nodes  # Initialize distances as -1 (unreachable)
        distances[start_node] = 0  # Distance to itself is 0
        queue = deque([start_node])
        
        while queue:
            node = queue.popleft()
            
            for neighbor in graph[node]:
                if distances[neighbor] == -1:  # If the neighbor hasn't been visited
                    distances[neighbor] = distances[node] + 1
                    queue.append(neighbor)
        
        return distances
    
    # Step 1: Convert edge list to adjacency list
    graph = {i: set() for i in range(num_nodes)}
    for u, v in edges:
        graph[u].add(v)
        graph[v].add(u)
    
    # Step 2: Use BFS to compute the shortest paths from each node and find the longest shortest path
    max_distance = 0
    
    for node in range(num_nodes):
        distances = bfs(graph, node, num_nodes)
        # Get the maximum distance from the current node
        # (excluding -1 values as they represent unreachable nodes)
        max_distance = max(max_distance, max(d for d in distances if d != -1))
    
    return max_distance
    

def count_triangles(num_nodes, edges):
    # Step 1: Convert the edge list to an adjacency list
    graph = {i: set() for i in range(num_nodes)}  # Create an empty graph with `num_nodes` nodes
    for u, v in edges:
        graph[u].add(v)
        graph[v].add(u)
    
    # Step 2: Count triangles
    triangle_count = 0
    # Iterate through all edges (u, v), with u < v to avoid double counting
    for u, v in edges:
        if u < v:  # Ensure each triangle is counted once
            # Find the intersection of neighbors of u and v
            common_neighbors = graph[u].intersection(graph[v])
            # The number of common neighbors is the number of triangles
            triangle_count += len(common_neighbors)
    
    return triangle_count


def compute_degrees(num_nodes, edges):
    degrees = [0] * num_nodes
    for node1, node2 in edges:
        degrees[node1] += 1
        degrees[node2] += 1

    return degrees


def print_degree_statistics(degrees):
    avg_degree = np.mean(degrees)
    std_dev_degree = np.std(degrees)
    min_degree = np.min(degrees)
    max_degree = np.max(degrees)

    print(f"Degree of each node: {degrees}")
    print(f"Average degree: {avg_degree:.2f}")
    print(f"Standard deviation: {std_dev_degree:.2f}")
    print(f"Minimum degree: {min_degree}")
    print(f"Maximum degree: {max_degree}")

def main():
    input_file = r'txtFiles/inputs/ABCommunityD/abcd_graph.txt'

    num_nodes, edges = read_graph(input_file)

    components = find_components(num_nodes, edges)
    print(f"Number of components: {len(components)}")
    print("Size of each component:", components)

    degrees = compute_degrees(num_nodes, edges)
    print_degree_statistics(degrees)

    triangle_count = count_triangles(num_nodes, edges)
    print("triangle_count:", triangle_count)

    diameter = graph_diameter(num_nodes, edges)
    print("diameter:", diameter)



if __name__ == "__main__":
    main()
