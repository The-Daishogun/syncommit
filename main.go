package main

import (
	"log"
	"syncommit/cmd"
	"syncommit/utils"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	utils.RunChecks()
	utils.SetupConfigFolder()
	cmd.Execute()
}
