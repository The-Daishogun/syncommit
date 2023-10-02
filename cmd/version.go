package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VERSION, BRANCH_NAME, HASH_NAME string

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of syncommit",
	Long:  `All software has versions. This is syncommit's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\nBranchName: %s\nCommitHash: %s\n", VERSION, BRANCH_NAME, HASH_NAME)
	},
}
