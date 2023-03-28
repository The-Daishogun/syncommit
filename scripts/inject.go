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

func makeFileExecutable(filepath string) error {
	return os.Chmod(filepath, 0755)
}

func InjectPostCommitScript() error {
	file, err := os.Create(postCommitScriptPath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(postCommitScript)
	if err != nil {
		os.Remove(postCommitScriptPath)
		return err
	}
	err = makeFileExecutable((postCommitScriptPath))
	if err != nil {
		os.Remove(postCommitScriptPath)
		return err
	}
	return nil
}

func InjectPrePushScript() error {
	file, err := os.Create(prePushScriptPath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(prePushScript)
	if err != nil {
		os.Remove(prePushScriptPath)
		return err
	}
	err = makeFileExecutable((prePushScriptPath))
	if err != nil {
		os.Remove(prePushScriptPath)
		return err
	}
	return nil
}
