package tool

import (
	"fmt"
	"os"

	"github.com/itsjxck/hashiman/config"
)

// Use =>
func (t *Tool) Use() error {
	if t.SelectedVersion == nil {
		return fmt.Errorf("no version set")
	}

	source := fmt.Sprintf("%s/%s/%s/%s", config.HomeDir, config.InstallDir, t.Name, t.SelectedVersion)
	dest := fmt.Sprintf("%s/%s", config.BinDir, t.Name)

	t.Log("Switching to version", t.SelectedVersion)

	os.Remove(dest)

	if err := os.Symlink(source, dest); err != nil {
		return fmt.Errorf("error switching tool version: %s", err)
	}

	return nil
}
