package cmd

import (
	"fmt"

	"github.com/itsjxck/hashiman/config"
	"github.com/itsjxck/hashiman/tool"
	"github.com/itsjxck/hashiman/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the installed command
var listCmd = &cobra.Command{
	Use:     "list [tool]",
	Aliases: []string{"ls", "l"},
	Short:   "Displays a list of all installed tool versions",
	Long: `Displays a list of all installed tool versions.
	
Examples:
hashiman list		# Lists all installed versions of all tools
hashiman list terraform	# Lists all installed versions of Terraform
hashiman ls terraform	# Lists all installed versions of Terraform using the command alias
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return nil
		} else if len(args) == 1 {
			if err := utils.ToolIsValid(args[0]); err != nil {
				return err
			}
			return nil
		} else {
			return fmt.Errorf("too many arguments")
		}
	},
	Run: listInstalled,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listInstalled(cmd *cobra.Command, args []string) {
	tools := []tool.Tool{}
	if len(args) > 1 {
		tools = append(tools, tool.Tool{Name: args[0]})
	} else {
		ts := config.SupportedTools
		for _, t := range ts {
			tools = append(tools, tool.Tool{Name: t})
		}
	}

	for _, tool := range tools {
		if err := tool.DiscoverInstalledVersions(); err != nil {
			tool.Quit("unable to discover installed versions:", err)
		}

		if err := tool.DetectCurrentVersion(); err != nil {
			tool.Quit("unable to detect current version:", err)
		}

		tool.Log("Installed Versions:")

		for _, v := range tool.InstalledVersions {
			str := "	" + v.String()
			if v.String() == tool.CurrentVersion.String() {
				str = str + "	(currently used)"
			}
			tool.Log(str)
		}
	}
}
