package utils

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/itsjxck/hashiman/config"
)

// ToolIsValid => Checks if tool is supported
func ToolIsValid(tool string) error {
	if !Includes(config.SupportedTools, tool) {
		return fmt.Errorf("unsupported tool; supported tools: %s", strings.Join(config.SupportedTools, ", "))
	}
	return nil
}

// SemverIsValid => Checks the semver constraint is valid
func SemverIsValid(semverIdent string) error {
	if _, err := semver.NewConstraint(semverIdent); err != nil {
		return fmt.Errorf("incorrectly formatted semver identifier: %s", err)
	}
	return nil
}
