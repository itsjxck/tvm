package tool

import (
	"github.com/Masterminds/semver"
)

// Tool =>
type Tool struct {
	Name              string
	SelectedVersion   *semver.Version
	CurrentVersion    *semver.Version
	VersionConstraint *semver.Constraints
	InstalledVersions []*semver.Version
	AvailableVersions []*semver.Version
	UseLatestVersion  bool
}
