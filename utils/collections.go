package utils

import (
	"github.com/Masterminds/semver"
)

// Index =>
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Includes =>
func Includes(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// VersionIndex =>
func VersionIndex(versions []*semver.Version, version *semver.Version) int {
	for i, v := range versions {
		if v.String() == version.String() {
			return i
		}
	}
	return -1
}

// IncludesVersion =>
func IncludesVersion(versions []*semver.Version, version *semver.Version) bool {
	return VersionIndex(versions, version) >= 0
}
