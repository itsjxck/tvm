package cmd

import (
	"errors"

	"github.com/itsjxck/hashiman/utils"
	"github.com/spf13/cobra"
)

func validateArgs(cmd *cobra.Command, args []string) error {
	if err := utils.ToolIsValid(args[0]); err != nil {
		return err
	}

	if len(args) > 2 {
		return errors.New("too many arguments")
	}

	if len(args) > 1 {
		if err := utils.SemverIsValid(args[1]); err != nil {
			return err
		}
	}

	return nil
}
