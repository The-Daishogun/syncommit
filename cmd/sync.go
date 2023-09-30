package cmd

import (
	"fmt"
	"log"
	"syncommit/structs"
	"syncommit/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(syncCommand)
}

var syncCommand = &cobra.Command{
	Use:   "sync",
	Short: "sync all commits to private repo",
	Long:  `commits and pushes all your commits in the current project to github.`,
	Run: func(cmd *cobra.Command, args []string) {
		repo := structs.GetRepoAtPath(".")
		allCommits := repo.GetRepoCommitsForCurrentAuthor()
		syncRepo := structs.GetRepoAtPath(structs.RepoPath)
		syncedCommits := syncRepo.GetRepoCommitsForCurrentAuthor()
		commitsToSync := utils.FilterSyncedCommits(allCommits, syncedCommits)
		for _, commit := range commitsToSync {
			err := commit.Commit(branchName, repo.Name)
			if err != nil {
				log.Fatal("failed to sync repo.\nError: ", err.Error())
			}
		}
		err := syncRepo.Push()
		if err != nil {
			log.Fatal("failed to push. ", err.Error())
		}
		fmt.Println("sync commits pushed successfully.")
	},
}
