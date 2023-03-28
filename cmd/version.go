package cmd

import (
	"fmt"
	"syncommit/utils"

	"github.com/spf13/cobra"
)

const VERSION = "0.0.1"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var branchName, commitHash string = utils.GetGitBranchName(), utils.GetGitHash()

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of syncommit",
	Long:  `All software has versions. This is syncommit's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("VERSION: %s\nBranchName: %s\nCommitHash: %s\n", VERSION, branchName, commitHash)
	},
}
