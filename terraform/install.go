package terraform

import (
	"fmt"

	"github.com/itsjxck/tvm/config"
	"github.com/itsjxck/tvm/utils"
)

// Install =>
func (t *Terraform) Install() error {
	if t.SelectedVersion == nil {
		return fmt.Errorf("no version set")
	}

	file := fmt.Sprintf("%s/%s_%s.zip", config.DownloadDir, config.Tool, t.SelectedVersion)
	installLoc := fmt.Sprintf("%s/%s/%s/", config.HomeDir, config.InstallDir, config.Tool)
	defer t.Cleanup()

	t.Log("Installing", t.SelectedVersion)
	if err := utils.Unzip(file, installLoc, t.SelectedVersion.String()); err != nil {
		return err
	}

	t.Log("Installed", t.SelectedVersion)
	return nil
}
