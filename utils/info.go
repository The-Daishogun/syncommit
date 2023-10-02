package utils

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func runCommandAndGetOutput(cmd *exec.Cmd) string {
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(out.String())
}
