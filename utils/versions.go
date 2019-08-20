package utils

import (
	"fmt"
	"sort"

	"github.com/Masterminds/semver"
)

// ParseVersions =>
func ParseVersions(versionStrings []string) (versions []*semver.Version, err error) {
	for _, vs := range versionStrings {
		v, err := semver.NewVersion(vs)
		if err != nil {
			fmt.Println("error parsing version:", err)
			return nil, err
		}
		versions = append(versions, v)
	}
	sort.Sort(semver.Collection(versions))
	return
}

// SemverFilter =>
func SemverFilter(allVersions []*semver.Version, constraint *semver.Constraints) (versions []*semver.Version) {
	for _, v := range allVersions {
		if constraint.Check(v) {
			versions = append(versions, v)
		}
	}
	sort.Sort(semver.Collection(versions))
	return
}
