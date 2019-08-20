package tool

import (
	"fmt"

	"github.com/itsjxck/hashiman/config"
	"github.com/itsjxck/hashiman/utils"
)

// Install =>
func (t *Tool) Install() error {
	if t.SelectedVersion == nil {
		return fmt.Errorf("no version set")
	}

	file := fmt.Sprintf("%s/%s_%s.zip", config.DownloadDir, t.Name, t.SelectedVersion)
	installLoc := fmt.Sprintf("%s/%s/%s/", config.HomeDir, config.InstallDir, t.Name)
	defer t.Cleanup()

	t.Log("Installing", t.SelectedVersion)
	if err := utils.Unzip(file, installLoc, t.SelectedVersion.String()); err != nil {
		return err
	}

	t.Log("Installed", t.SelectedVersion)
	return nil
}
