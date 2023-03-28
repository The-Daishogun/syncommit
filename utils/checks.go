package utils

import (
	"log"
	"os"
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

func SearchDir(path string, target string) (found bool, err error) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return false, err
	}
	for _, dir := range dirs {
		if dir.Name() == target {
			found = true
			break
		}
	}
	return found, err
}
