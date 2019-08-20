package cmd

import (
	"github.com/itsjxck/hashiman/tool"
	"github.com/spf13/cobra"
)

// installCmd represents the get command
var installCmd = &cobra.Command{
	Use:     "install <tool> [semver]",
	Aliases: []string{"i"},
	Short:   "Install a new version of a Hashicorp tool.",
	Long: `Install a new version of a Hashicorp tool, optionally specifying
a semver identifier to apply version constraints.
	
Examples:
hashiman install terraform		# Installs the latest version of Terraform
hashiman install consul ~1.5	# Installs the latest 1.5 version of Consul
hashiman i consul 1.5.x	# Installs the latest 1.5 version of Consul using the command alias
hashiman i nomad 0.9.4	# Installs specifically the 0.9.4 version of Nomad using the command alias`,
	Args: validateArgs,
	Run:  installTool,
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func installTool(cmd *cobra.Command, args []string) {
	tool := tool.Tool{}

	if err := tool.ProcessArgs(cmd, args); err != nil {
		tool.Quit("error processing args:", err)
	}

	if err := tool.DiscoverAvailableVersions(); err != nil {
		tool.Quit("error discovering available versions:", err)
	}

	if err := tool.DetermineVersionToInstall(); err != nil {
		tool.Quit("error determining which version to install:", err)
	}

	if err := tool.Download(); err != nil {
		tool.Quit("error downloading version:", err)
	}

	if err := tool.Install(); err != nil {
		tool.Quit("error installing version:", err)
	}

	if err := tool.Use(); err != nil {
		tool.Quit("error switching to version:", err)
	}

	tool.Log("Done")
}
