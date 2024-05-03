package graph

import "testing"

func (graph *Graph) VerifySort(order []string) bool {
	position := make(map[string]int)
	for index, node := range order {
		position[node] = index
	}

	for node, neighbors := range graph.nodes {
		for neighbor := range neighbors {
			if position[node] >= position[neighbor] {
				return false
			}
		}
	}

	return true
}

func TestGraph(t *testing.T) {
	graph := NewGraph()

	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("A", "D")
	graph.AddEdge("C", "B")
	graph.AddEdge("C", "D")
	graph.AddEdge("B", "E")
	graph.AddEdge("D", "E")

	graph.AddNode("F")

	sorted := graph.TopologicalSort()

	if !graph.VerifySort(sorted) {
		t.Error("Topological sort failed")
	} else {
		t.Log("Topological sort verified successfully")
	}
}
