package cmd

import (
	"github.com/itsjxck/hashiman/tool"

	"github.com/spf13/cobra"
)

var useLatest bool

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use <tool> [semver]",
	Short: "Switch between installed versions of a tool.",
	Long: `Switch between installed versions of a tool.

Examples:
hashiman use terraform -l	# Switches to the latest installed version of Terraform
hashiman use terraform 0.12	# Switches to the latest 0.12 version of Terraform
hashiman use terraform 0.12.6	# Switches to version 0.12.6 of Terraform
`,
	Args: validateArgs,
	Run:  useHandler,
}

func init() {
	rootCmd.AddCommand(useCmd)
}

func useHandler(cmd *cobra.Command, args []string) {
	tool := tool.Tool{}

	if err := tool.ProcessArgs(cmd, args); err != nil {
		tool.Quit("error processing args:", err)
	}

	if err := tool.DiscoverInstalledVersions(); err != nil {
		tool.Quit("error discovering available versions:", err)
	}

	if err := tool.DetermineVersionToUse(); err != nil {
		tool.Quit("error determining version to use:", err)
	}

	if err := tool.Use(); err != nil {
		tool.Quit("error switching to version:", err)
	}
}
