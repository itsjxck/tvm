package terraform

import (
	"fmt"

	"github.com/itsjxck/tvm/utils"
)

// DetermineVersionToInstall =>
func (t *Terraform) DetermineVersionToInstall() (err error) {
	if t.AvailableVersions == nil || len(t.AvailableVersions) == 0 {
		return fmt.Errorf("no available versions")
	}

	if t.UseLatestVersion || len(t.AvailableVersions) == 1 {
		t.SelectedVersion = t.AvailableVersions[len(t.AvailableVersions)-1]
		return nil
	}

	installed, err := t.discoverVersionsFromInstallDir()
	if err != nil {
		t.Log("failed to discover installed versions:", err)
	}

	t.SelectedVersion, err = utils.RequestVersionSelect(t.AvailableVersions, installed, "(Installed)")

	return err
}
