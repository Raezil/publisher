package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Check if there are enough arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: publisher <version> <module-path>")
		return
	}

	// Get the version and module path from the command-line arguments
	version := os.Args[1]
	modulePath := os.Args[2]

	// Step 1: Tag the current commit with the specified version
	err := gitTag(version)
	if err != nil {
		fmt.Printf("Error tagging version: %v\n", err)
		return
	}

	// Step 2: Push the tag to the remote repository
	err = gitPush(version)
	if err != nil {
		fmt.Printf("Error pushing tag: %v\n", err)
		return
	}

	// Step 3: Update Go module proxy by listing the module
	err = goListModule(modulePath)
	if err != nil {
		fmt.Printf("Error listing Go module: %v\n", err)
		return
	}

	fmt.Println("Library published successfully!")
}

// Function to run git tag
func gitTag(version string) error {
	cmd := exec.Command("git", "tag", version)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Function to run git push origin version
func gitPush(version string) error {
	cmd := exec.Command("git", "push", "origin", version)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Function to run go list to update Go proxy
func goListModule(modulePath string) error {
	cmd := exec.Command("GOPROXY=proxy.golang.org", "go", "list", "-m", modulePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
