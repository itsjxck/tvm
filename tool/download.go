package tool

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"github.com/itsjxck/hashiman/config"
)

// Download =>
func (t *Tool) Download() error {
	if t.SelectedVersion == nil {
		return fmt.Errorf("no version set")
	}

	t.Log("Downloading", t.SelectedVersion)

	fileURL := fmt.Sprintf("https://releases.hashicorp.com/%s/%s/%s_%s_%s_%s.zip", t.Name, t.SelectedVersion, t.Name, t.SelectedVersion, runtime.GOOS, runtime.GOARCH)
	downloadTo := fmt.Sprintf("%s/%s_%s.zip", config.DownloadDir, t.Name, t.SelectedVersion)

	resp, err := http.Get(fileURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(downloadTo)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	defer t.Log("Downloaded", t.SelectedVersion)
	return err
}
