// HBM is an application to authorize and manage authorized docker command using Docker AuthZ plugin.
// Copyright (C) 2016-2017 Kassisol inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"github.com/juliengk/go-log"
	"github.com/juliengk/go-utils/user"
	"github.com/kassisol/hbm/cli/command/commands"
	"github.com/spf13/cobra"
)

func newHbmCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hbm",
		Short: "HBM is an application to authorize and manage authorized docker commands",
		Long:  "HBM is an application to authorize and manage authorized docker commands",
	}

	cmd.SetHelpTemplate(helpTemplate)
	cmd.SetUsageTemplate(usageTemplate)

	commands.AddCommands(cmd)

	return cmd
}

func main() {
	u := user.New()
	l, _ := log.NewDriver("standard", nil)

	if !u.IsRoot() {
		l.Error("You must be root to run that command")
	}

	cmd := newHbmCommand()
	if err := cmd.Execute(); err != nil {
		l.Error(err)
	}
}

var usageTemplate = `{{ .Short | trim }}

Usage:{{ if .Runnable }}
  {{ if .HasAvailableFlags }}{{ appendIfNotPresent .UseLine "[flags]" }}{{ else }}{{ .UseLine }}{{ end }}{{ end }}{{ if .HasAvailableSubCommands }}
  {{ .CommandPath }} [command]{{ end }}{{ if gt .Aliases 0 }}

Aliases:
  {{ .NameAndAliases }}{{ end }}{{ if .HasExample }}

Examples:
  {{ .Example }}{{ end }}{{ if .HasAvailableSubCommands}}

Available Commands:{{ range .Commands }}{{ if .IsAvailableCommand }}
  {{ rpad .Name .NamePadding }} {{ .Short }}{{ end }}{{ end }}{{ end }}{{ if .HasAvailableLocalFlags }}

Flags:
  {{ .LocalFlags.FlagUsages | trimRightSpace }}{{ end }}{{ if .HasAvailableInheritedFlags }}

Global Flags:
  {{ .InheritedFlags.FlagUsages | trimRightSpace }}{{ end }}{{ if .HasHelpSubCommands }}

Additional help topics: {{ range .Commands }}{{ if .IsHelpCommand }}
  {{ rpad .CommandPath .CommandPathPadding }} {{ .Short }}{{ end }}{{ end }}{{ end }}{{ if .HasAvailableSubCommands }}

Use "{{ .CommandPath }} [command] --help" for more information about a command.{{ end }}
`

var helpTemplate = `
{{ if or .Runnable .HasSubCommands }}{{ .UsageString }}{{ end }}`
