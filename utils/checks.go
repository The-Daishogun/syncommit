package utils

import (
	"log"
	"os/exec"
)

func checkGitInstallation() error {
	cmd := exec.Command("git", "--version")
	_, err := cmd.Output()
	return err
}

func RunChecks() {
	gitErr := checkGitInstallation()
	if gitErr != nil {
		log.Fatal("Git is not installed, please install git using your package manager and try again...")
	}
}
