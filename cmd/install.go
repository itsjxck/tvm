package cmd

import (
	"github.com/itsjxck/tvm/terraform"
	"github.com/spf13/cobra"
)

// installCmd represents the get command
var installCmd = &cobra.Command{
	Use:     "install [semver]",
	Aliases: []string{"i"},
	Short:   "Install a new version of Terraform.",
	Long: `Install a new version of Terraform, optionally specifying
a semver identifier to apply version constraints.
	
Examples:
tvm install		# Retrieves a list of available Terraform versions and prompts you for a selection
tvm install 0.12 -l	# Installs the latest 1.5 version of Consul
tvm i 0.12.6	# Installs specifically the 0.12.6 version of Terraform using the command alias`,
	Args: validateArgs,
	Run:  installTerraform,
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func installTerraform(cmd *cobra.Command, args []string) {
	terraform := terraform.Terraform{}

	if err := terraform.ProcessArgs(cmd, args); err != nil {
		terraform.Quit("error processing args:", err)
	}

	if err := terraform.DiscoverAvailableVersions(); err != nil {
		terraform.Quit("error discovering available versions:", err)
	}

	if err := terraform.DetermineVersionToInstall(); err != nil {
		terraform.Quit("error determining which version to install:", err)
	}

	if err := terraform.Download(); err != nil {
		terraform.Quit("error downloading version:", err)
	}

	if err := terraform.Install(); err != nil {
		terraform.Quit("error installing version:", err)
	}

	if err := terraform.Use(); err != nil {
		terraform.Quit("error switching to version:", err)
	}

	terraform.Log("Done")
}
