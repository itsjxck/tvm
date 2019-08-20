package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Masterminds/semver"
	"github.com/itsjxck/hashiman/utils"
)

// DiscoverAvailableVersions =>
func (t *Tool) DiscoverAvailableVersions() error {
	if t.VersionConstraint == nil {
		return fmt.Errorf("version constraint not set")
	}

	t.Log("Fetching versions")

	versions, err := t.fetchGithubReleases()
	if err != nil {
		return fmt.Errorf("cant get releases: %s", err)
	}

	if len(versions) == 0 {
		// TODO: Go get some another way
	}

	t.AvailableVersions = utils.SemverFilter(versions, t.VersionConstraint)

	return nil
}

func (t *Tool) fetchGithubReleases() (versions []*semver.Version, err error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/hashicorp/%s/releases", t.Name)

	resp, err := http.Get(apiURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	type releasesJSON []map[string]interface{}
	var releases releasesJSON
	err = json.Unmarshal(body, &releases)
	if err != nil {
		return
	}

	for _, r := range releases {
		vs, ok := r["tag_name"].(string)
		if !ok {
			return nil, fmt.Errorf("cant read version")
		}
		v, err := semver.NewVersion(vs)
		if err != nil {
			return nil, fmt.Errorf("cant parse version: %s", err)
		}
		versions = append(versions, v)
	}

	return
}
