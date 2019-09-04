package terraform

import (
	"fmt"
	"os"

	"github.com/Masterminds/semver"
	"github.com/itsjxck/tvm/config"
)

// Remove =>
func (t *Terraform) Remove() error {
	if !t.RemoveAll && t.SelectedVersion == nil {
		return fmt.Errorf("no version to remove")
	}

	if t.RemoveAll {
		for _, version := range t.InstalledVersions {
			if err := t.deleteVersion(version); err != nil {
				return fmt.Errorf("error removing version: %s", err)
			}
		}
	} else {
		if err := t.deleteVersion(t.SelectedVersion); err != nil {
			return fmt.Errorf("error removing version: %s", err)
		}
	}

	return nil
}

func (t *Terraform) deleteVersion(version *semver.Version) error {
	t.Log("Removing", version.String())
	file := fmt.Sprintf("%s/%s/%s/%s", config.HomeDir, config.InstallDir, config.Tool, version)
	if err := os.Remove(file); err != nil {
		return fmt.Errorf("error removing file: %s", err)
	}
	t.Log("Removed", version.String())
	return nil
}
