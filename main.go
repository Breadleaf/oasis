package main

import (
	// fh "oasis/internal/file_handler"
	ph "oasis/internal/project_handler"
	"flag"
	"fmt"
	"os"
)

func init() {
	initFlag := flag.Bool("i", false, "Initialize a new project")
	compileFlag := flag.Bool("c", false, "Compile the project")

	flag.Parse()

	if *initFlag && *compileFlag {
		fmt.Println("Cannot use both -i and -c flags at the same time")
		os.Exit(1)
	} else if *initFlag {
		// Get the project info
		err, name := ph.GetProjectInfo();
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Create the project
		if err := ph.InitProject(name); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if *compileFlag {
		fmt.Println("Compiling the project")
	} else {
		fmt.Printf("Usage: %s [-i] [-c]\n", os.Args[0])
		os.Exit(1)
	}
}

func main() {
	/*
	libraries := []string{"Re", "Random"}
	fh.BuildExecutable(fh.BuildGraph(fh.ResolveDependencies(fh.GetRawDependencyMap(fh.GetFiles(".")), libraries)).TopologicalSort(), libraries)
	*/
}
