package main

import (
	"fmt"

	"github.com/juliengk/go-utils/filedir"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/cli/command/commands"
	"github.com/spf13/cobra/doc"
)

func main() {
	manPath := "/tmp/hbm/man"
	man8 := fmt.Sprintf("%s/man8", manPath)

	if err := filedir.CreateDirIfNotExist(man8, true, 0755); err != nil {
		fmt.Println(err)
	}

	header := &doc.GenManHeader{
		Title:   "HBM",
		Section: "8",
		Source:  "Harbormaster",
	}
	opts := doc.GenManTreeOptions{
		Header:           header,
		Path:             man8,
		CommandSeparator: "-",
	}

	cmd := command.NewHBMCommand()
	commands.AddCommands(cmd)
	cmd.DisableAutoGenTag = true

	if err := doc.GenManTreeFromOpts(cmd, opts); err != nil {
		fmt.Println(err)
	}
}
