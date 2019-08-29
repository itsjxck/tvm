package terraform

import (
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/Masterminds/semver"
	"github.com/itsjxck/tvm/config"
	"github.com/itsjxck/tvm/utils"
)

// DiscoverInstalledVersions =>
func (t *Terraform) DiscoverInstalledVersions() error {
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

func (t *Terraform) discoverVersionsFromInstallDir() (versions []*semver.Version, err error) {
	files, err := ioutil.ReadDir(fmt.Sprintf("%s/%s/%s", config.HomeDir, config.InstallDir, config.Tool))
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
