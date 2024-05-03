package file_handler

import (
	"os"
	"os/exec"
	"fmt"
	"path/filepath"
	"strings"
	"oasis/internal/graph"
)

func GetFiles(path string) []string {
	var files []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".ml") || strings.HasSuffix(path, ".mli") {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}

func GetRawDependencyMap(files []string) map[string][]string {
	dependencyMap := make(map[string][]string)

	for _, file := range files {
		cmd := exec.Command("ocamldep", "-modules", file)
		output, err := cmd.Output()
		if err != nil {
			panic(err)
		}

		parts := strings.Split(strings.TrimSpace(string(output)), " ")
		dependencyMap[file] = parts[1:]
	}

	return dependencyMap
}

func ResolveDependencies(dependencyMap map[string][]string, libraries []string) map[string][]string {
	// Remove libraries from dependency map
	for file, deps := range dependencyMap {
		for _, lib := range libraries {
			for i, dep := range deps {
				if dep == lib {
					deps = append(deps[:i], deps[i+1:]...)
				}
			}
		}

		dependencyMap[file] = deps
	}

	// NOTE: For now, assume that all files are in the same directory
	for file, deps := range dependencyMap {
		for index, dep := range deps {
			dep = strings.ToLower(dep)
			deps[index] = dep + ".ml"
		}

		dependencyMap[file] = deps
	}

	return dependencyMap
}

func BuildGraph(dependencyMap map[string][]string) *graph.Graph {
	graph := graph.NewGraph()

	for file, deps := range dependencyMap {
		if len(deps) == 0 { // If a file has no dependencies
			graph.AddNode(file)
		} else {
			for _, dep := range deps {
				graph.AddEdge(file, dep)
			}
		}
	}

	return graph
}

func BuildExecutable(files []string, libraries []string) {
	var libraryFlags []string
	for _, lib := range libraries {
		lib = strings.ToLower(lib)
		libraryFlags = append(libraryFlags, "-package", lib)
	}

	for index := len(files) - 1; index >= 0; index-- {
		file := files[index]

		compileFlags := append([]string{"ocamlopt"}, libraryFlags...)
		compileFlags = append(compileFlags, "-c", file)

		cmd := exec.Command("ocamlfind", compileFlags...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error compiling %s: %s\nOutput: %s\n", file, err, string(output))
			panic(err)
		}
	}

	var linkFlags []string
	linkFlags = append(linkFlags, "ocamlopt")
	linkFlags = append(linkFlags, libraryFlags...)
	linkFlags = append(linkFlags, "-linkpkg")
	for index := len(files) - 1; index >= 0; index-- {
		if filepath.Ext(files[index]) != ".mli" {
			file := strings.Replace(files[index], ".ml", ".cmx", 1)
			linkFlags = append(linkFlags, file)
		}
	}

	cmd := exec.Command("ocamlfind", linkFlags...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error linking files: %s\nOutput: %s\n", err, string(output))
		panic(err)
	}
}
