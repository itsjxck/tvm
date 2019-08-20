package tool_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Masterminds/semver"

	"github.com/itsjxck/hashiman/config"
	"github.com/itsjxck/hashiman/tool"
)

func TestDownload(t *testing.T) {
	tool := tool.Tool{
		Name: "terraform",
	}

	err := tool.Download()
	if err == nil {
		t.Errorf("[failed] expected error due to no version set")
	} else {
		t.Logf("[passed] expected error, got: %s", err)
	}

	tool.SelectedVersion, err = semver.NewVersion(testVersion)
	if err != nil {
		t.Errorf("failed to parse version: %s", err)
	}

	err = tool.Download()
	if err != nil {
		t.Errorf("[failed] expected no error, got: %s", err)
	} else {
		t.Logf("[passed] expected no errors, got no errors")
	}

	if _, err = os.Stat(fmt.Sprintf("%s/%s_%s.zip", config.DownloadDir, tool.Name, tool.SelectedVersion.String())); err != nil {
		t.Errorf("[failed] expected file to exist, got error: %s", err)
	} else {
		t.Logf("[passed] downloaded file exists")
	}

	tool.Cleanup()
}
