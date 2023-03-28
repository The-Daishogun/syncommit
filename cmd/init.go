package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"syncommit/scripts"
	"syncommit/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCommand)
}

func validateGitUrl(gitUrl string) bool {
	re := regexp.MustCompile(`(?m)^git@github\.com:.*\.git$`)
	return re.Match([]byte(gitUrl))
}

func clonePrivateRepo(repoUrl string) error {
	dirs, err := os.ReadDir(utils.ConfigFolderPath)
	if err != nil {
		log.Fatal("failed to read the contents of ConfigFolderPath. ", err)
	}
	for _, dir := range dirs {
		if dir.Name() == utils.RepoLocation {
			return nil
		}
	}
	cmd := exec.Command("git", "clone", "-q", repoUrl, filepath.Join(utils.ConfigFolderPath, utils.RepoLocation))
	return cmd.Run()

}

func getPrivateRepo() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Create a new private repo on github (https://github.com/new) and paste the ssh url here:")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	validated := validateGitUrl(input)
	if !validated {
		log.Fatal("invalid git url. make sure it's the ssh url and the url is correct.")
	}
	file, err := os.Create(filepath.Join(utils.ConfigFolderPath, utils.RepoFileName))
	if err != nil {
		log.Fatal("failed to create .repo file. ", err)
	}
	defer file.Close()

	_, err = file.WriteString(input)
	if err != nil {
		os.Remove(filepath.Join(utils.ConfigFolderPath, utils.RepoFileName))
		log.Fatal("failed to write to .repo file. ", err)
	}
	fmt.Println("Starting to clone the repo.")
	err = clonePrivateRepo(input)
	if err != nil {
		os.Remove(filepath.Join(utils.ConfigFolderPath, utils.RepoFileName))
		os.Remove(filepath.Join(utils.ConfigFolderPath, utils.RepoLocation))
		log.Fatal("failed to clone repo. make sure repo url is correct. ", err)
	}
	fmt.Println("Cloning successful.")
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "setup the current repository to sync with github",
	Run: func(cmd *cobra.Command, args []string) {
		getPrivateRepo()
		dirs, err := os.ReadDir(".")
		if err != nil {
			log.Println("failed to read current directory", err)
			os.Exit(1)
		}

		var found bool

		for _, dir := range dirs {
			if dir.Name() == ".git" {
				found = true
				break
			}
		}

		if !found {
			log.Fatal("You are not in a git initialized repo!\nnavigate to a folder containing a '.git' folder then run the command again...")
		}
		err = scripts.InjectPostCommitScript()
		if err != nil {
			log.Fatal("Failed to inject post-commit script. ", err)
		}
		err = scripts.InjectPrePushScript()
		if err != nil {
			log.Fatal("Failed to inject pre-push script. ", err)
		}
		fmt.Println("All Done. all your commits will be synced with the private repo.")
	},
}
