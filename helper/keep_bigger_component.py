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
        component = []
        while stack:
            current = stack.pop()
            if not visited[current]:
                visited[current] = True
                component.append(current)
                stack.extend(adjacency_list[current])
        return component

    for node in range(num_nodes):
        if not visited[node]:
            component = dfs(node)
            components.append(component)

    return components

def save_largest_component(input_file, output_file):
    num_nodes, edges = read_graph(input_file)
    components = find_components(num_nodes, edges)

    # Find the largest component
    largest_component = max(components, key=len)
    largest_component_set = set(largest_component)

    # Filter edges to keep only those in the largest component
    largest_component_edges = [
        (node1, node2) for node1, node2 in edges
        if node1 in largest_component_set and node2 in largest_component_set
    ]

    with open(output_file, 'w') as outfile:
        outfile.write(f"{len(largest_component)}\n")
        outfile.write(f"{len(largest_component_edges)}\n")
        for edge in sorted(largest_component_edges):
            outfile.write(f"{edge[0]} {edge[1]}\n")

# Usage
input_file = 'cleaned_graph.txt'
output_file = 'largest_component_graph.txt'
save_largest_component(input_file, output_file)
