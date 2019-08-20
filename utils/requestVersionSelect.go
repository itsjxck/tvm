package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Masterminds/semver"
)

// RequestVersionSelect =>
func RequestVersionSelect(versions []*semver.Version, match []*semver.Version, matchMessage string) (*semver.Version, error) {
	fmt.Println("Found these versions matching your SemVer input:")
	for i, v := range versions {
		fmt.Printf("[%d] %s	", i+1, v.String())
		if match != nil && matchMessage != "" && len(match) > 0 && IncludesVersion(match, v) {
			fmt.Println(matchMessage)
		} else {
			fmt.Println()
		}
	}
	fmt.Printf("Select a version [%d]: ", len(versions))
	reader, selection := bufio.NewReader(os.Stdin), len(versions)

	text, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	if text != "\n" {
		if selection, err = strconv.Atoi(text[:len(text)-1]); err != nil {
			return nil, err
		}
	}
	return versions[selection-1], err
}
