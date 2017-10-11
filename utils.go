package xj2go

import (
	"os"
	"strings"
)

func max(nodes *[]leafNode) int {
	n := 0
	for _, node := range *nodes {
		t := strings.Count(node.path, ".")
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
