package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syncommit/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pushCommand)
}

var pushCommand = &cobra.Command{
	Use:   "push",
	Short: "Pushes all the sync commits to github",
	Long:  `Pushes all the sync commits to github`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.Chdir(utils.RepoPath)
		if err != nil {
			log.Fatal("failed to cd into repo directory. ", err.Error())
		}
		commitCmd := exec.Command("git", "push", "-fu")
		err = commitCmd.Run()
		if err != nil {
			log.Fatal("failed to push. ", err.Error())
		}
		fmt.Println("sync commits pushed successfully.")
	},
}
