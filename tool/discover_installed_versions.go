package tool

import (
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/Masterminds/semver"
	"github.com/itsjxck/hashiman/config"
	"github.com/itsjxck/hashiman/utils"
)

// DiscoverInstalledVersions =>
func (t *Tool) DiscoverInstalledVersions() error {
	versions, err := t.discoverVersionsFromInstallDir()
	if err != nil {
		return fmt.Errorf("cant load installed versions: %s", err)
	}

	if t.VersionConstraint != nil {
		versions = utils.SemverFilter(versions, t.VersionConstraint)
	}

	t.InstalledVersions = versions

	return nil
}

func (t *Tool) discoverVersionsFromInstallDir() (versions []*semver.Version, err error) {
	files, err := ioutil.ReadDir(fmt.Sprintf("%s/%s/%s", config.HomeDir, config.InstallDir, t.Name))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		v, err := semver.NewVersion(file.Name())
		if err != nil {
			fmt.Println("error parsing version:", err)
			return nil, err
		}
		versions = append(versions, v)
	}

	sort.Sort(semver.Collection(versions))

	return
}
