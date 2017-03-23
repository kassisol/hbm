package main

import (
	"fmt"

	"github.com/juliengk/go-utils/filedir"
	"github.com/kassisol/hbm/cli/command/commands"
	"github.com/spf13/cobra"
)

func main() {
	scPath := "/tmp/hbm/shellcompletion"
	bashTarget := fmt.Sprintf("%s/bash", scPath)

	if err := filedir.CreateDirIfNotExist(scPath, true, 0755); err != nil {
		fmt.Println(err)
	}

	cmd := &cobra.Command{Use: "hbm"}
	commands.AddCommands(cmd)
	cmd.DisableAutoGenTag = true

	if err := cmd.GenBashCompletionFile(bashTarget); err != nil {
		fmt.Println(err)
	}
}
