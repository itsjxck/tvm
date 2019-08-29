package terraform

import "github.com/spf13/cobra"

// ProcessArgs =>
func (t *Terraform) ProcessArgs(cmd *cobra.Command, args []string) (err error) {
	semverIdent := "*"
	if len(args) > 0 {
		semverIdent = args[0]
	}

	if err := t.SetSemverConstraint(semverIdent); err != nil {
		return err
	}

	if t.UseLatestVersion, err = cmd.Parent().PersistentFlags().GetBool("latest"); err != nil {
		return err
	}

	return nil

}
