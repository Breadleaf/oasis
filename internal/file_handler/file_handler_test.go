package file_handler

import (
	"testing"
	"fmt"
)

func TestReadFile(t *testing.T) {
	files := GetFiles("test_files")

	if len(files) != 5 {
		t.Errorf("Expected 5 files, got %d", len(files))
	}

	dmap := GetRawDependencyMap(files)

	if len(dmap["test_files/main.ml"]) != 3 {
		t.Errorf("Expected 3 dependencies, got %d", len(dmap["test_files/main.ml"]))
	}

	resolved := ResolveDependencies(dmap, []string{"Random", "Re"})

	dependencyOrder := BuildGraph(resolved).TopologicalSort()

	fmt.Println(dependencyOrder)

	// TODO: test BuildExecutable after implementing a better resolution algorithm
}
