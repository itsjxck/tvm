package utils

import (
	"fmt"

	"github.com/Masterminds/semver"
)

// SemverIsValid => Checks the semver constraint is valid
func SemverIsValid(semverIdent string) error {
	if _, err := semver.NewConstraint(semverIdent); err != nil {
		return fmt.Errorf("incorrectly formatted semver identifier: %s", err)
	}
	return nil
}
