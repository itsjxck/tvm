package tool

import (
	"fmt"

	"github.com/itsjxck/hashiman/utils"
)

// DetermineVersionToInstall =>
func (t *Tool) DetermineVersionToInstall() (err error) {
	if t.AvailableVersions == nil || len(t.AvailableVersions) == 0 {
		return fmt.Errorf("no available versions")
	}

	if t.UseLatestVersion || len(t.AvailableVersions) == 1 {
		t.SelectedVersion = t.AvailableVersions[len(t.AvailableVersions)-1]
		return nil
	}

	if err := t.DiscoverInstalledVersions(); err != nil {
		t.Log("failed to discover installed versions:", err)
	}

	t.SelectedVersion, err = utils.RequestVersionSelect(t.AvailableVersions, t.InstalledVersions, "(Installed)")

	return err
}
