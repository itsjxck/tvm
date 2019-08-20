package tool

import "github.com/spf13/cobra"

// ProcessArgs =>
func (t *Tool) ProcessArgs(cmd *cobra.Command, args []string) (err error) {
	tool, semverIdent := args[0], "*"
	if len(args) > 1 {
		semverIdent = args[1]
	}

	t.Name = tool
	if err := t.SetSemverConstraint(semverIdent); err != nil {
		return err
	}

	if t.UseLatestVersion, err = cmd.Parent().PersistentFlags().GetBool("latest"); err != nil {
		return err
	}

	return nil

}
