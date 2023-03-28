package scripts

import (
	_ "embed"
	"os"
	"path/filepath"
)

//go:embed post-commit.sh
var postCommitScript string

//go:embed pre-push.sh
var prePushScript string

const gitHooksLocation = ".git/hooks"

var postCommitScriptPath = filepath.Join(gitHooksLocation, "post-commit")

var prePushScriptPath = filepath.Join(gitHooksLocation, "pre-push")

const executableFilePermissions = 0755

func injectHook(filePath string, script string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(script)
	if err != nil {
		os.Remove(filePath)
		return err
	}
	err = os.Chmod(filePath, executableFilePermissions)
	if err != nil {
		os.Remove(filePath)
		return err
	}
	return nil
}

func InjectPostCommitScript() error {
	return injectHook(postCommitScriptPath, postCommitScript)
}

func InjectPrePushScript() error {
	return injectHook(prePushScriptPath, prePushScript)
}
