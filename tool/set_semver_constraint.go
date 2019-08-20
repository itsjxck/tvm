package tool

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
)

// SetSemverConstraint =>
func (t *Tool) SetSemverConstraint(semverString string) (err error) {
	if len(strings.Split(semverString, ".")) < 3 && semverString[0:1] != "~" {
		semverString = fmt.Sprintf("~%s", semverString)
	}

	t.VersionConstraint, err = semver.NewConstraint(semverString)
	return err
}
