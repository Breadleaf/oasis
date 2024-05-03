package graph

type Graph struct {
	nodes map[string]map[string]bool
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]map[string]bool),
	}
}

func (graph *Graph) AddNode(node string) {
	if _, exists := graph.nodes[node]; !exists {
		graph.nodes[node] = make(map[string]bool)
	}
}

func (graph *Graph) AddEdge(from, to string) {
	graph.AddNode(from)
	graph.AddNode(to)
	graph.nodes[from][to] = true
}

func (graph *Graph) TopologicalSort() []string {
	visited := make(map[string]bool)
	var stack []string

	var dfs func(node string)
	dfs = func(node string) {
		visited[node] = true

		for neighbor := range graph.nodes[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}

		stack = append([]string{node}, stack...)
	}

	for node := range graph.nodes {
		if !visited[node] {
			dfs(node)
		}
	}

	return stack
}
