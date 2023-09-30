package cmd

import (
	"fmt"
	"log"
	"syncommit/structs"

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
		syncRepo := structs.GetRepoAtPath(structs.RepoPath)
		err := syncRepo.Push()
		if err != nil {
			log.Fatal("failed to push. ", err.Error())
		}
		fmt.Println("sync commits pushed successfully.")
	},
}
