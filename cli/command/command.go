package command

import (
	"github.com/kassisol/hbm/cli/command/commands"
	"github.com/spf13/cobra"
)

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

func NewHBMCommand() *cobra.Command {
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
