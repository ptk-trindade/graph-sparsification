# This code do many things:
# Remove duplicated edges
# Make sure the ids are sequential
def process_graph(input_file, output_file):
    with open(input_file, 'r') as infile:
        lines = infile.readlines()

    num_nodes = int(lines[0].strip())
    num_edges = int(lines[1].strip())
    edges = set()

    node_mapping = {}
    current_id = 0

    for line in lines[2:]:
        node1, node2 = map(int, line.strip().split())
        if node1 not in node_mapping:
            node_mapping[node1] = current_id
            current_id += 1
        if node2 not in node_mapping:
            node_mapping[node2] = current_id
            current_id += 1
        
        mapped_node1 = node_mapping[node1]
        mapped_node2 = node_mapping[node2]

        if mapped_node1 < mapped_node2:
            edges.add((mapped_node1, mapped_node2))
        else:
            edges.add((mapped_node2, mapped_node1))

    with open(output_file, 'w') as outfile:
        outfile.write(f"{num_nodes}\n")
        outfile.write(f"{len(edges)}\n")
        for edge in sorted(edges):
            outfile.write(f"{edge[0]} {edge[1]}\n")

# Usage
input_file = 'CA-GrQc.txt'
output_file = 'cleaned_graph.txt'
process_graph(input_file, output_file)

