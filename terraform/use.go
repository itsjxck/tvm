package terraform

import (
	"fmt"
	"os"

	"github.com/itsjxck/tvm/config"
)

// Use =>
func (t *Terraform) Use() error {
	if t.SelectedVersion == nil {
		return fmt.Errorf("no version set")
	}

	source := fmt.Sprintf("%s/%s/%s/%s", config.HomeDir, config.InstallDir, config.Tool, t.SelectedVersion)
	dest := fmt.Sprintf("%s/%s", config.BinDir, config.Tool)

	t.Log("Switching to version", t.SelectedVersion)

	os.Remove(dest)

	if err := os.Symlink(source, dest); err != nil {
		return fmt.Errorf("error switching tool version: %s", err)
	}

	return nil
}
