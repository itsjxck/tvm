package tool

import (
	"fmt"
	"os"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/itsjxck/hashiman/config"
)

// DetectCurrentVersion =>
func (t *Tool) DetectCurrentVersion() error {
	file, err := os.Readlink(fmt.Sprintf("%s/%s", config.BinDir, t.Name))
	if err != nil {
		return fmt.Errorf("failed to read symlink: %s", err)
	}

	parts := strings.Split(file, "/")
	t.CurrentVersion, err = semver.NewVersion(parts[len(parts)-1])
	return err
}
