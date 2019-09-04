package cmd

import (
	"github.com/itsjxck/tvm/terraform"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Removes a version of Terraform",
	Long: `Removes a version of Terraform.

tvm remove 0.12	# Displays the 0.12 versions of Terraform and prompts to pick one to remove.
tvm rm 0.12.4	# Removes the 0.12.4 version of Terraform using the command alias.
tvm rm 0.11 -a	# Removes all the 0.11 versions of Terraform using the command alias.`,
	Args: validateArgs,
	Run:  removeTerraform,
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().BoolP("all", "a", false, "Remove all that match the SemVer constraint")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func removeTerraform(cmd *cobra.Command, args []string) {
	terraform := terraform.Terraform{}

	if err := terraform.ProcessArgs(cmd, args); err != nil {
		terraform.Quit("unable to process args:", err)
	}

	if err := terraform.DiscoverInstalledVersions(); err != nil {
		terraform.Quit("unable to discover installed versions:", err)
	}

	if err := terraform.DetectCurrentVersion(); err != nil {
		terraform.Quit("unable to detect current version:", err)
	}

	if !terraform.RemoveAll {
		if err := terraform.DetermineVersionToUse(); err != nil {
			terraform.Quit("error determining version to use:", err)
		}
	}

	if err := terraform.Remove(); err != nil {
		terraform.Quit("unable to remove:", err)
	}
}
