package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

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

func ListFiles() map[string][]string {
	files, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	dependencies := make(map[string][]string)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) != ".ml" {
			continue
		}

		cmd := exec.Command("ocamldep", "-modules", file.Name())
		output, err := cmd.Output()
		if err != nil {
			panic(err)
		}

		parsed := strings.Split(strings.TrimSpace(string(output)), " ")
		parsed[0] = parsed[0][:len(parsed[0]) - 1]
		dependencies[file.Name()] = parsed[1:]
	}

	fmt.Println("Files and dependencies:")
	fmt.Println(dependencies)

	return dependencies
}

func ResolveDependencies(dependencies map[string][]string, libraries []string) map[string][]string {
	for file, deps := range dependencies {
		for _, lib := range libraries {
			for i, dep := range deps {
				if dep == lib {
					deps = append(deps[:i], deps[i+1:]...)
				}
			}
		}

		dependencies[file] = deps
	}

	for file, deps := range dependencies {
		for i, dep := range deps {
			dep = strings.ToLower(dep)
			deps[i] = dep + ".ml"
		}

		dependencies[file] = deps
	}

	fmt.Println("Resolved dependencies:")
	fmt.Println(dependencies)

	return dependencies
}

func BuildGraph(dependencies map[string][]string) *Graph {
	graph := NewGraph()
	for file, deps := range dependencies {
		if len(deps) == 0 {
			graph.AddNode(file)
		} else {
			for _, dep := range deps {
				graph.AddEdge(file, dep)
			}
		}
	}

	return graph
}

func InterfaceFiles() []string {
	files, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	var interfaces []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) != ".mli" {
			continue
		}

		interfaces = append(interfaces, file.Name())
	}

	return interfaces
}

func CompileFiles(files []string, libraries []string) {
	var library_flags []string
	for _, library := range libraries {
		library = strings.ToLower(library)
		library_flags = append(library_flags, "-package", library)
	}

	for index := len(files) - 1; index >= 0; index-- {
		file := files[index]
		compile_args := append([]string{"ocamlopt"}, library_flags...)
		compile_args = append(compile_args, "-c", file)

		fmt.Println(compile_args)

		cmd := exec.Command("ocamlfind", compile_args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error compiling %s: %s\nOutput: %s\n", file, err, string(output))
			panic(err)
		}

		fmt.Println("Compiled", file)
	}

	var link_args []string
	link_args = append(link_args, "ocamlopt")
	link_args = append(link_args, library_flags...)
	link_args = append(link_args, "-linkpkg")
	for index := len(files) - 1; index >= 0; index-- {
		if filepath.Ext(files[index]) != ".mli" {
			link_args = append(link_args, strings.Replace(files[index], ".ml", ".cmx", 1))
		}
	}

	cmd := exec.Command("ocamlfind", link_args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error linking: %s\nOutput: %s\n", err, string(output))
		panic(err)
	}

	fmt.Println("Linked files")
}

func main() {
	libraries := []string{
		"Re",
		"Random",
	}

	dependencies := ListFiles()
	dependencies = ResolveDependencies(dependencies, libraries)
	graph := BuildGraph(dependencies)
	sorted := graph.TopologicalSort()
	sorted = append(sorted, InterfaceFiles()...)
	CompileFiles(sorted, libraries)
}
