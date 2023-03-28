package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syncommit/utils"

	"github.com/spf13/cobra"
)

var Message string

func init() {
	commitCommand.Flags().StringVarP(&Message, "message", "m", "", "message for the sync commit")
	commitCommand.MarkFlagRequired("message")
	rootCmd.AddCommand(commitCommand)
}

var commitCommand = &cobra.Command{
	Use:   "commit",
	Short: "Add a commit to the sync repo",
	Long:  `Add a commit to the sync repo`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.Chdir(utils.RepoPath)
		if err != nil {
			log.Fatal("failed to cd into repo directory. ", err.Error())
		}
		commitCmd := exec.Command("git", "commit", "-m", Message, "--allow-empty")
		err = commitCmd.Run()
		if err != nil {
			log.Fatal("failed to commit. ", err.Error())
		}
		fmt.Println("sync commit committed successfully.")
	},
}
