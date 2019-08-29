package terraform

import (
	"github.com/Masterminds/semver"
)

// Terraform =>
type Terraform struct {
	SelectedVersion   *semver.Version
	CurrentVersion    *semver.Version
	VersionConstraint *semver.Constraints
	InstalledVersions []*semver.Version
	AvailableVersions []*semver.Version
	UseLatestVersion  bool
}
