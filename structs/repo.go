package structs

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const RepoFileName = ".repo"
const RepoLocation = "repo"

var HomePath = os.Getenv("HOME")

var ConfigFolderPath = filepath.Join(HomePath, "/.syncommit")
var RepoPath = filepath.Join(ConfigFolderPath, RepoLocation)

const hashPrefix = "hash: "

type Repo struct {
	Name               string
	Path               string
	BranchName         string
	CurrentAuthorName  string
	CurrentAuthorEmail string
}

func (r *Repo) GetRepoCommitsForCurrentAuthor() (commits []Commit) {
	cmd := exec.Command("git", "--no-pager", "log", fmt.Sprintf("--author=%s", r.CurrentAuthorEmail), "--pretty=format:%h$//%s$//%ad", "--date=unix", "--no-merges")
	commitsBytes, err := cmd.Output()
	if err != nil && err.Error() == "exit status 128" {
		// User has no commits
		return
	}
	if err != nil {
		log.Fatal("Failed to get commits for current author.\nError: ", err.Error())
	}
	commitsString := strings.Split(string(commitsBytes), "\n")
	for _, commitString := range commitsString {
		if len(commitString) < 1 {
			continue
		}

		var commit Commit
		if r.Path != RepoPath {
			commit = parseGeneralRepoCommitString(commitString)
		} else {
			if !strings.Contains(commitString, hashPrefix) {
				continue
			}
			commit = parseSyncRepoCommitString(commitString)
		}
		commits = append(commits, commit)
	}
	return
}

func parseSyncRepoCommitString(commitString string) Commit {
	hashWithMessageAndTime := strings.Split(commitString, "$//")
	commitTimeStr, _ := strconv.Atoi(hashWithMessageAndTime[2])
	originalCommitHash := strings.Replace(hashWithMessageAndTime[0], hashPrefix, "", 1)[0:7]
	return Commit{Hash: originalCommitHash, Message: hashWithMessageAndTime[1], Time: time.Unix(int64(commitTimeStr), 0)}
}

func parseGeneralRepoCommitString(commitString string) Commit {
	hashWithMessageAndTime := strings.Split(commitString, "$//")
	commitTimeStr, _ := strconv.Atoi(hashWithMessageAndTime[2])
	return Commit{Hash: hashWithMessageAndTime[0], Message: hashWithMessageAndTime[1], Time: time.Unix(int64(commitTimeStr), 0)}
}

func (r *Repo) Push() error {
	cmd := exec.Command("git", "push", "-fu")
	cmd.Dir = r.Path
	return cmd.Run()
}

func GetRepoAtPath(path string) Repo {
	err := os.Chdir(path)
	if err != nil {
		log.Fatal("failed to change directory,\nError: ", err.Error())
	}
	branchNameBytes, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		log.Fatal("failed to get repo information.\nError: ", err.Error())
	}
	branchName := strings.ReplaceAll(string(branchNameBytes), "\n", "")
	repoPathBytes, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		log.Fatal("failed to get repo information.\nError: ", err.Error())
	}
	repoPath := strings.ReplaceAll(string(repoPathBytes), "\n", "")
	repoNameBytes, err := exec.Command("basename", string(repoPath)).Output()
	if err != nil {
		log.Fatal("failed to get repo information.\nError: ", err.Error())
	}
	repoName := strings.ReplaceAll(string(repoNameBytes), "\n", "")
	authorNameBytes, err := exec.Command("git", "config", "user.name").Output()
	if err != nil {
		log.Fatal("failed to get repo information.\nError: ", err.Error())
	}
	authorName := strings.TrimSpace(string(authorNameBytes))
	authorEmailBytes, err := exec.Command("git", "config", "user.email").Output()
	if err != nil {
		log.Fatal("failed to get repo information.\nError:", err.Error())
	}
	authorEmail := strings.TrimSpace(string(string(authorEmailBytes)))
	return Repo{
		Name:               repoName,
		Path:               repoPath,
		BranchName:         branchName,
		CurrentAuthorName:  authorName,
		CurrentAuthorEmail: authorEmail,
	}
}
