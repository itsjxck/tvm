package cmd

import (
	"github.com/itsjxck/tvm/terraform"
	"github.com/spf13/cobra"
)

// listCmd represents the installed command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "Displays a list of all installed Terraform versions",
	Long: `Displays a list of all installed Terraform versions.
	
Examples:
tvm list		# Lists all installed versions of Terraform
tvm ls	# Lists all installed versions of Terraform using the command alias
`,
	Run: listInstalled,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listInstalled(cmd *cobra.Command, args []string) {
	terraform := terraform.Terraform{}

	if err := terraform.DiscoverInstalledVersions(); err != nil {
		terraform.Quit("unable to discover installed versions:", err)
	}

	if err := terraform.DetectCurrentVersion(); err != nil {
		terraform.Quit("unable to detect current version:", err)
	}

	terraform.Log("Installed Versions:")

	for _, v := range terraform.InstalledVersions {
		str := "	" + v.String()
		if v.String() == terraform.CurrentVersion.String() {
			str = str + "	(currently used)"
		}
		terraform.Log(str)
	}
}
