package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/juliengk/go-utils"
	"github.com/juliengk/go-utils/json"
	"github.com/juliengk/go-utils/validation"
	u "github.com/kassisol/hbm/pkg/utils"
	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/storage/driver"
	"github.com/spf13/cobra"
)

var (
	resourceListFilter []string

	resourceAddType            string
	resourceAddValue           string
	resourceAddVolumeRecursive bool
	resourceAddVolumeNoSuid    bool

	resourceMemberAdd    bool
	resourceMemberRemove bool
)

var resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Manage whitelisted resources",
	Long:  "Manage whitelisted resources",
}

var resourceListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List whitelisted resources",
	Long:  "List whitelisted resources",
}

var resourceAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add resource to the whitelist",
	Long:  "Add resource to the whitelist",
}

var resourceRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove resource from the whitelist",
	Long:  "Remove resource from the whitelist",
}

var resourceExistsCmd = &cobra.Command{
	Use:   "find",
	Short: "Verify if resource exists in the whitelist",
	Long:  "Verify if resource exists in the whitelist",
}

var resourceMemberCmd = &cobra.Command{
	Use:   "member",
	Short: "Verify if resource exists in the whitelist",
	Long:  "Verify if resource exists in the whitelist",
}

func init() {
	RootCmd.AddCommand(resourceCmd)
	resourceCmd.AddCommand(resourceListCmd)
	resourceCmd.AddCommand(resourceAddCmd)
	resourceCmd.AddCommand(resourceRemoveCmd)
	resourceCmd.AddCommand(resourceExistsCmd)
	resourceCmd.AddCommand(resourceMemberCmd)

	resourceListCmd.Flags().StringSliceVarP(&resourceListFilter, "filter", "f", []string{}, "Filter output based on conditions provided")
	resourceAddCmd.Flags().StringVarP(&resourceAddType, "type", "t", "action", "Add resource to group")
	resourceAddCmd.Flags().StringVarP(&resourceAddValue, "value", "v", "", "Add resource to group")
	resourceAddCmd.Flags().BoolVarP(&resourceAddVolumeRecursive, "recursive", "", false, "Add resource to group")
	resourceAddCmd.Flags().BoolVarP(&resourceAddVolumeNoSuid, "no-suid", "", false, "Allow volume mounted with nosuid option")

	resourceMemberCmd.Flags().BoolVarP(&resourceMemberAdd, "add", "a", false, "Add resource to group")
	resourceMemberCmd.Flags().BoolVarP(&resourceMemberRemove, "remove", "r", false, "Remove resource to group")

	resourceCmd.Run = resourceUsage
	resourceListCmd.Run = resourceList
	resourceAddCmd.Run = resourceAdd
	resourceRemoveCmd.Run = resourceRemove
	resourceExistsCmd.Run = resourceExists
	resourceMemberCmd.Run = resourceMember
}

func resourceUsage(cmd *cobra.Command, args []string) {
	resourceCmd.Usage()
	os.Exit(-1)
}

func resourceList(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	filters := utils.ConvertSliceToMap("=", resourceListFilter)

	resources := s.ListResources(filters)

	if len(resources) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 2, ' ', 0)

		fmt.Fprintln(w, "NAME\tTYPE\tVALUE\tOPTION\tCOLLECTIONS")

		for resource, collections := range resources {
			if len(collections) > 0 {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", resource.Name, resource.Type, resource.Value, resource.Option, strings.Join(collections, ", "))
			} else {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", resource.Name, resource.Type, resource.Value, u.RemoveLastChar(resource.Option))
			}
		}

		w.Flush()
	}
}

func resourceAdd(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	if resourceAddType != "volume" && resourceAddVolumeRecursive {
		fmt.Printf("Conflicting options --type %s and --recursive\n", resourceAddType)
	}

	if resourceAddType != "volume" && resourceAddVolumeNoSuid {
		fmt.Printf("Conflicting options --type %s and --no-suid\n", resourceAddType)
	}

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if err = validation.IsValidName(args[0]); err != nil {
		utils.Exit(err)
	}

	if s.FindResource(args[0]) {
		utils.Exit(fmt.Errorf("%s already exists", args[0]))
	}

	rt := NewResourceTypes()
	if err = rt.IsValidResourceType(resourceAddType); err != nil {
		utils.Exit(err)
	}

	options := ""
	if resourceAddType == "volume" && resourceAddVolumeRecursive || resourceAddVolumeNoSuid {
		vo := driver.VolumeOptions{Recursive: resourceAddVolumeRecursive, NoSuid: resourceAddVolumeNoSuid}
		jsonR := json.Encode(vo)
		options = jsonR.String()
	}

	s.AddResource(args[0], resourceAddType, resourceAddValue, options)
}

func resourceRemove(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if !s.FindResource(args[0]) {
		utils.Exit(fmt.Errorf("%s does not exist", args[0]))
	}

	if err = s.RemoveResource(args[0]); err != nil {
		utils.Exit(err)
	}
}

func resourceExists(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	result := s.FindResource(args[0])

	fmt.Println(result)
}

func resourceMember(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	s, err := storage.NewDriver("sqlite", appPath)
	if err != nil {
		utils.Exit(err)
	}
	defer s.End()

	if !s.FindCollection(args[0]) {
		utils.Exit(fmt.Errorf("%s does not exist", args[0]))
	}

	if !s.FindResource(args[1]) {
		utils.Exit(fmt.Errorf("%s does not exist", args[1]))
	}

	if resourceMemberAdd {
		s.AddResourceToCollection(args[0], args[1])
	}
	if resourceMemberRemove {
		s.RemoveResourceFromCollection(args[0], args[1])
	}
}
