package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"syncommit/structs"
)

func ValidateGitUrl(gitUrl string) bool {
	re := regexp.MustCompile(`(?m)^git@github\.com:.*\.git$`)
	return re.Match([]byte(gitUrl))
}

func ClonePrivateRepo(repoUrl string) error {
	dirs, err := os.ReadDir(structs.ConfigFolderPath)
	if err != nil {
		log.Fatal("failed to read the contents of ConfigFolderPath. ", err)
	}
	for _, dir := range dirs {
		if dir.Name() == structs.RepoLocation {
			return nil
		}
	}
	cmd := exec.Command("git", "clone", "-q", repoUrl, filepath.Join(structs.ConfigFolderPath, structs.RepoLocation))
	return cmd.Run()

}

func GetPrivateRepo() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Create a new private repo on github (https://github.com/new) and paste the ssh url here:")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	validated := ValidateGitUrl(input)
	if !validated {
		log.Fatal("invalid git url. make sure it's the ssh url and the url is correct.")
	}
	file, err := os.Create(filepath.Join(structs.ConfigFolderPath, structs.RepoFileName))
	if err != nil {
		log.Fatal("failed to create .repo file. ", err)
	}
	defer file.Close()

	_, err = file.WriteString(input)
	if err != nil {
		os.Remove(filepath.Join(structs.ConfigFolderPath, structs.RepoFileName))
		log.Fatal("failed to write to .repo file. ", err)
	}
	fmt.Println("Starting to clone the repo.")
	err = ClonePrivateRepo(input)
	if err != nil {
		os.Remove(filepath.Join(structs.ConfigFolderPath, structs.RepoFileName))
		os.Remove(filepath.Join(structs.ConfigFolderPath, structs.RepoLocation))
		log.Fatal("failed to clone repo. make sure repo url is correct. ", err)
	}
	fmt.Println("Cloning successful.")
}

func createCommitMasterString(commits []structs.Commit) string {
	var masterString = ""
	for _, commit := range commits {
		masterString = masterString + commit.Hash + " "
	}
	return masterString
}

func FilterSyncedCommits(allCommits []structs.Commit, syncedCommits []structs.Commit) []structs.Commit {
	var commitsToSync []structs.Commit
	masterString := createCommitMasterString(syncedCommits)
	for _, commit := range allCommits {
		if strings.Contains(masterString, commit.Hash) {
			continue
		}
		commitsToSync = append(commitsToSync, commit)
	}
	return commitsToSync
}
