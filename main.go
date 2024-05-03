package main

import (
	fh "oasis/internal/file_handler"
)

func main() {
	libraries := []string{"Re", "Random"}

	fh.BuildExecutable(fh.BuildGraph(fh.ResolveDependencies(fh.GetRawDependencyMap(fh.GetFiles(".")), libraries)).TopologicalSort(), libraries)
}
