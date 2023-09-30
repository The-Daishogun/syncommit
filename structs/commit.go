package structs

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Commit struct {
	Hash    string
	Message string
	Time    time.Time
}

func (c *Commit) generateCommitMessage(repoName, branchName string) string {
	return fmt.Sprintf("hash: %s %s on branch: %s on repo: %s", c.Hash, c.Message, strings.TrimSpace(repoName), strings.TrimSpace(branchName))
}

func (c *Commit) Commit(repoName, branchName string) error {
	err := os.Chdir(RepoPath)
	if err != nil {
		log.Fatal("failed to cd into repo directory. ", err.Error())
	}
	return exec.Command("git", "commit", "-m", fmt.Sprintf("%q", c.generateCommitMessage(repoName, branchName)), "--allow-empty", fmt.Sprintf("--date='%s'", c.Time.Format(time.RFC1123Z))).Run()
}
