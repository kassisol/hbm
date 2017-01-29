package resource

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/juliengk/go-utils"
	"github.com/kassisol/hbm/cli/command"
	"github.com/kassisol/hbm/storage"
	"github.com/spf13/cobra"
)

var resourceListFilter []string

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List whitelisted resources",
		Long:    listDescription,
		Run:     runList,
	}

	flags := cmd.Flags()
	flags.StringSliceVarP(&resourceListFilter, "filter", "f", []string{}, "Filter output based on conditions provided")

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", command.AppPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	filters := utils.ConvertSliceToMap("=", resourceListFilter)

	resources := s.ListResources(filters)

	if len(resources) > 0 {
		/*w := map[string]int{
			"name": 20,
			"type": 20,
			"value": 20,
			"options": 20,
			"collections": 20,
		}

		for resource, collections := range resources {
			nameLen := len(resource.Name)
			typeLen := len(resource.Type)
			valueLen := len(resource.Value)
			optionsLen := len(resource.Option)
			collectionsLen := len(strings.Join(collections, ", "))

			if nameLen > w["name"] {
				w["name"] = nameLen
			}
			if typeLen > w["type"] {
				w["type"] = typeLen
			}
			if valueLen > w["value"] {
				w["value"] = valueLen
			}
			if optionsLen > w["options"] {
				w["options"] = optionsLen
			}
			if collectionsLen > w["collections"] {
				w["collections"] = collectionsLen
			}
		}

		fmt.Printf("%s %s %s %s %s\n", stringpad("NAME", w["name"]), stringpad("TYPE", w["type"]), stringpad("VALUE", w["value"]), stringpad("OPTIONS", w["options"]), stringpad("COLLECTIONS", w["collections"]))
		for resource, collections := range resources {
			if len(collections) > 0 {
				fmt.Printf("%s %s %s %s %s\n", stringpad(resource.Name, w["name"]), stringpad(resource.Type, w["type"]), stringpad(resource.Value, w["value"]), stringpad(removeLastChar(resource.Option), w["options"]), stringpad(strings.Join(collections, ", "), w["collections"]))
			} else {
				fmt.Printf("%s %s %s %s %s\n", stringpad(resource.Name, w["name"]), stringpad(resource.Type, w["type"]), stringpad(resource.Value, w["value"]), stringpad(removeLastChar(resource.Option), w["options"]))
			}
		}*/

		tw := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)
		fmt.Fprintln(tw, "NAME\tTYPE\tVALUE\tOPTIONS\tCOLLECTIONS")

		for resource, collections := range resources {
			if len(collections) > 0 {
				fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\n", resource.Name, resource.Type, resource.Value, removeLastChar(resource.Option), strings.Join(collections, ", "))
			} else {
				fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", resource.Name, resource.Type, resource.Value, removeLastChar(resource.Option))
			}
		}

		tw.Flush()
	}
}

func removeLastChar(s string) string {
	strLen := len(s) - 1
	newStr := s
	if strLen > 0 {
		newStr = s[0:strLen]
	}

	return newStr
}

/*
func stringpad(s string, width int) string {
	l := len(s)
	g := width - l

	return s + strings.Repeat(" ", g)
}
*/

var listDescription = `
List whitelisted resources

`
