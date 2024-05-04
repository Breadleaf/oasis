package project_handler

import (
	"testing"
	"os"
)

func TestInitProject(t *testing.T) {
	projectName := "testProject"

	// Create a project directory
	if err := InitProject(projectName); err != nil {
		t.Errorf("Error creating project directory: %s", err)
	}

	// Check if the project directory was created
	if _, err := os.Stat(projectName); os.IsNotExist(err) {
		t.Errorf("Project directory was not created")
	}

	// Check that oasis.json was created
	if _, err := os.Stat(projectName + "/oasis.json"); os.IsNotExist(err) {
		t.Errorf("oasis.json was not created")
	}

	// remove the project directory
	if err := os.RemoveAll(projectName); err != nil {
		t.Errorf("Error deleting project directory")
	}
}
