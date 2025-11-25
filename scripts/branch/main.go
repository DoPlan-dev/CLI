package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/DoPlan-dev/CLI/internal/git"
)

func main() {
	action := flag.String("action", "", "Action: create, checkout, push, or info")
	taskID := flag.String("task", "", "Task ID (e.g., 5.2)")
	projectPath := flag.String("project", ".", "Project path")
	branchName := flag.String("branch", "", "Branch name (optional, will be generated from task ID if not provided)")
	setUpstream := flag.Bool("upstream", true, "Set upstream when pushing")
	flag.Parse()

	if *action == "" {
		fmt.Fprintf(os.Stderr, "Error: --action is required (create, checkout, push, or info)\n")
		os.Exit(1)
	}

	absPath, err := filepath.Abs(*projectPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to resolve project path: %v\n", err)
		os.Exit(1)
	}

	// Check if it's a Git repository
	if !git.IsGitRepository(absPath) {
		fmt.Fprintf(os.Stderr, "Warning: %s is not a Git repository. Branch operations skipped.\n", absPath)
		os.Exit(0)
	}

	switch *action {
	case "create":
		if *taskID == "" && *branchName == "" {
			fmt.Fprintf(os.Stderr, "Error: --task or --branch is required for create action\n")
			os.Exit(1)
		}
		if *branchName == "" {
			*branchName = git.GetTaskBranchName(*taskID)
		}
		if err := git.CreateAndCheckoutBranch(absPath, *branchName); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Branch created/checked out: %s\n", *branchName)

	case "checkout":
		if *branchName == "" {
			if *taskID == "" {
				fmt.Fprintf(os.Stderr, "Error: --task or --branch is required for checkout action\n")
				os.Exit(1)
			}
			*branchName = git.GetTaskBranchName(*taskID)
		}
		if err := git.CreateAndCheckoutBranch(absPath, *branchName); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Checked out branch: %s\n", *branchName)

	case "push":
		if *branchName == "" {
			currentBranch, err := git.GetCurrentBranch(absPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to get current branch: %v\n", err)
				os.Exit(1)
			}
			*branchName = currentBranch
		}
		if err := git.PushBranch(absPath, *branchName, *setUpstream); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Pushed branch: %s\n", *branchName)

	case "info":
		currentBranch, err := git.GetCurrentBranch(absPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to get current branch: %v\n", err)
			os.Exit(1)
		}
		clean, err := git.IsCleanWorkingTree(absPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to check working tree: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Current branch: %s\n", currentBranch)
		if clean {
			fmt.Println("Working tree: clean")
		} else {
			fmt.Println("Working tree: has uncommitted changes")
		}

	default:
		fmt.Fprintf(os.Stderr, "Error: unknown action: %s\n", *action)
		fmt.Fprintf(os.Stderr, "Valid actions: create, checkout, push, info\n")
		os.Exit(1)
	}
}

