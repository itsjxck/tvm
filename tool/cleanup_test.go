package tool_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Masterminds/semver"

	"github.com/itsjxck/hashiman/config"
	"github.com/itsjxck/hashiman/tool"
)

func TestCleanup(t *testing.T) {
	tool := tool.Tool{
		Name: "terraform",
	}

	err := tool.Cleanup()
	if err == nil {
		t.Errorf("[failed] expected error due to version not being set")
	} else {
		t.Logf("[passed] expected error, got: %s", err)
	}

	tool.SelectedVersion, err = semver.NewVersion(testVersion)
	if err != nil {
		t.Errorf("failed to parse version")
	}

	file, err := os.Create(fmt.Sprintf("%s/%s_%s", config.DownloadDir, tool.Name, tool.SelectedVersion.String()))
	if err != nil {
		t.Errorf("failed to create test file")
	}
	file.Close()

	err = tool.Cleanup()
	if err != nil {
		t.Errorf("[failed] expected no error, got: %s", err)
	} else {
		t.Logf("[passed]")
	}
}
