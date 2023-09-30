package utils

import (
	"log"
	"os"
	"syncommit/structs"
)

const ConfigFolderPermission = 0755

func SetupConfigFolder() {
	err := os.MkdirAll(structs.ConfigFolderPath, ConfigFolderPermission)
	if err != nil {
		log.Fatal("failed to setup", err)
	}
}
