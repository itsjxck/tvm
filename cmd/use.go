package cmd

import (
	"github.com/itsjxck/tvm/terraform"

	"github.com/spf13/cobra"
)

var useLatest bool

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:     "use <tool> [semver]",
	Aliases: []string{"u"},
	Short:   "Switch between installed versions of a tool.",
	Long: `Switch between installed versions of a tool.

Examples:
tvm use terraform -l	# Switches to the latest installed version of Terraform
tvm use terraform 0.12	# Switches to the latest 0.12 version of Terraform
tvm use terraform 0.12.6	# Switches to version 0.12.6 of Terraform
`,
	Args: validateArgs,
	Run:  useHandler,
}

func init() {
	rootCmd.AddCommand(useCmd)
}

func useHandler(cmd *cobra.Command, args []string) {
	terraform := terraform.Terraform{}

	if err := terraform.ProcessArgs(cmd, args); err != nil {
		terraform.Quit("error processing args:", err)
	}

	if err := terraform.DiscoverInstalledVersions(); err != nil {
		terraform.Quit("error discovering available versions:", err)
	}

	if len(terraform.InstalledVersions) < 1 {
		terraform.Quit("No versions installed matching SemVer input", args[0])
	}

	if err := terraform.DetermineVersionToUse(); err != nil {
		terraform.Quit("error determining version to use:", err)
	}

	if err := terraform.Use(); err != nil {
		terraform.Quit("error switching to version:", err)
	}
}
