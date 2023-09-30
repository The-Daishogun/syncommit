package cmd

import (
	"fmt"
	"log"
	"syncommit/scripts"
	"syncommit/structs"
	"syncommit/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCommand)
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "setup the current repository to sync with github",
	Run: func(cmd *cobra.Command, args []string) {
		privateRepoFound, err := utils.SearchDir(structs.ConfigFolderPath, structs.RepoLocation)
		if err != nil {
			log.Fatal("failed to read directory")
		}
		privateRepoUrlFound, err := utils.SearchDir(structs.ConfigFolderPath, structs.RepoFileName)
		if err != nil {
			log.Fatal("failed to read directory")
		}
		if !privateRepoFound || !privateRepoUrlFound {
			utils.GetPrivateRepo()
		}
		gitFolderFound, err := utils.SearchDir(".", ".git")
		if err != nil {
			log.Fatal("failed to read current directory", err)
		} else if !gitFolderFound {
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
