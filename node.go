package xj2go

import (
	"regexp"
	"strings"
)

func (xj *XJ) pathsToNodes(paths []string) map[string][]strctNode {
	strct = make(map[string][]strctNode)
	exist = make(map[string]bool)
	n := max(paths)
	root := strings.Split(paths[0], ".")[0]

	re = regexp.MustCompile(`\[\d+\]`)
	xj.leafPath(root, root, paths)

	for i := 0; i < n; i++ {
		for e := range exist {
			es := strings.Split(e, ".")
			root := es[len(es)-1]
			xj.leafPath(e, root, paths)
		}
	}

	return strct
}
