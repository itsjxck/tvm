package cmd

import (
	"errors"

	"github.com/itsjxck/tvm/utils"
	"github.com/spf13/cobra"
)

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return errors.New("too many arguments")
	}

	if len(args) > 0 {
		if err := utils.SemverIsValid(args[0]); err != nil {
			return err
		}
	}

	return nil
}
