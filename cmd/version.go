package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

const VERSION = "0.0.1"

func init() {
	rootCmd.AddCommand(versionCmd)
}

func getGitInfo() (string, string) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	branchName := strings.TrimSpace(out.String())

	cmd = exec.Command("git", "rev-parse", "--short", "HEAD")
	out.Reset()
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	commitHash := strings.TrimSpace(out.String())
	return branchName, commitHash
}

var branchName, commitHash string = getGitInfo()

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of syncommit",
	Long:  `All software has versions. This is syncommit's`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("VERSION: %s\nBranchName: %s\nCommitHash: %s\n", VERSION, branchName, commitHash)
	},
}
