package main

import (
	"syncommit/cmd"
	"syncommit/utils"
)

func main() {
	utils.RunChecks()
	utils.SetupConfigFolder()
	cmd.Execute()
}
