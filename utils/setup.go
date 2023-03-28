package utils

import (
	"log"
	"os"
	"path/filepath"
)

const RepoFileName = ".repo"
const RepoLocation = "repo"

var HomePath = os.Getenv("HOME")

var ConfigFolderPath = filepath.Join(HomePath, "/.syncommit")
var RepoPath = filepath.Join(ConfigFolderPath, RepoLocation)

const ConfigFolderPermission = 0755

func SetupConfigFolder() {
	err := os.MkdirAll(ConfigFolderPath, ConfigFolderPermission)
	if err != nil {
		log.Fatal("failed to setup", err)
	}
}
