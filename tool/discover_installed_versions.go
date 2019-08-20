package tool

import (
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/Masterminds/semver"
	"github.com/itsjxck/hashiman/config"
)

// DiscoverInstalledVersions =>
func (t *Tool) DiscoverInstalledVersions() error {
	files, err := ioutil.ReadDir(fmt.Sprintf("%s/%s/%s", config.HomeDir, config.InstallDir, t.Name))
	if err != nil {
		return err
	}

	for _, file := range files {
		v, err := semver.NewVersion(file.Name())
		if err != nil {
			fmt.Println("error parsing version:", err)
			return err
		}
		t.InstalledVersions = append(t.InstalledVersions, v)
	}

	sort.Sort(semver.Collection(t.InstalledVersions))

	return nil
}
