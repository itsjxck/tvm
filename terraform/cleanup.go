package terraform

import (
	"fmt"
	"os"

	"github.com/itsjxck/tvm/config"
)

// Cleanup =>
func (t *Terraform) Cleanup() error {
	if t.SelectedVersion == nil {
		return fmt.Errorf("no version set")
	}

	//t.Log("Cleaning up")
	//defer t.Log("Cleaned")
	downloadedFile := fmt.Sprintf("%s/%s_%s.zip", config.DownloadDir, config.Tool, t.SelectedVersion)
	if err := os.Remove(downloadedFile); err != nil {
		return fmt.Errorf("error removing file: %s", err)
	}
	return nil
}
