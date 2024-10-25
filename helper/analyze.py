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
    input_file = r'C:\Users\Ptk\graph-sparsification\txtFiles\inputs\erdosRenyi\redosErni_4000.txt'
    num_nodes, edges = read_graph(input_file)

    components = find_components(num_nodes, edges)
    print(f"Number of components: {len(components)}")
    print("Size of each component:", components)

    degrees = compute_degrees(num_nodes, edges)
    print_degree_statistics(degrees)

if __name__ == "__main__":
    main()
