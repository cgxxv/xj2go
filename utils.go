package xj2go

import (
	"os"
	"strings"
)

func max(paths *[]string) int {
	n := 0
	for _, path := range *paths {
		t := strings.Count(path, ".")
		if n < t {
			n = t
		}
	}

	return n + 1
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
