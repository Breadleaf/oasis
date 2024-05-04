package project_handler

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"encoding/json"
)

// GetProjectInfo:
// Returns an error if there is one, else returns the project name
func GetProjectInfo() (error, string) {
	// Create a new scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Get project name
	fmt.Print("Enter project name: ")

	var name string
	if scanner.Scan() {
		name = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Error reading from stdin: %v", err), ""
	}

	// Validate project name
	if name == "" {
		return fmt.Errorf("Project name cannot be empty"), ""
	}

	// Format project name
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, " ", "_")

	return nil, name
}

func InitProject(projectName string) error {
	// Create project directory
	permission := os.FileMode(0755) // rwxr-xr-x

	// Check if project directory already exists; if not make it
	if err := os.Mkdir(projectName, permission); err != nil {
		return err
	}

	// Create project build json file and write to it
	buildFile, err := os.Create(projectName + "/oasis.json")
	if err != nil {
		return err
	}

	build := map[string]interface{}{
		"name": projectName,
		"libs": []interface{}{},
	}

	buildJson, err := json.MarshalIndent(build, "", "  ")
	if err != nil {
		return err
	}

	buildFile.Write(buildJson)

	return nil
}
