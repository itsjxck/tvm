package tool

import (
	"fmt"

	"github.com/Masterminds/semver"

	"github.com/itsjxck/hashiman/utils"
)

// DetermineVersionToUse =>
func (t *Tool) DetermineVersionToUse() (err error) {
	if t.InstalledVersions == nil || len(t.InstalledVersions) == 0 {
		return fmt.Errorf("no installed versions")
	}

	if t.UseLatestVersion || len(t.InstalledVersions) == 1 {
		t.SelectedVersion = t.InstalledVersions[len(t.InstalledVersions)-1]
		return nil
	}

	if err := t.DetectCurrentVersion(); err != nil {
		t.Log("failed to detect current versions:", err)
	}

	t.SelectedVersion, err = utils.RequestVersionSelect(t.InstalledVersions, []*semver.Version{t.CurrentVersion}, "(Current)")

	return err
}
