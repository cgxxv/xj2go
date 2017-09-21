package xj2go

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func (xj *XJ) pathsToNodes(paths []string) []map[string][]strctNode {
	n := max(paths)
	root := strings.Split(paths[0], ".")[0]

	re = regexp.MustCompile(`\[\d+\]`)
	exist = make(map[string]bool)

	strcts := []map[string][]strctNode{}
	strct := xj.leafPath(root, root, paths)
	strcts = append(strcts, strct)

	for i := 0; i < n; i++ {
		for e := range exist {
			es := strings.Split(e, ".")
			root := es[len(es)-1]
			strct := xj.leafPath(e, root, paths)
			if strct != nil {
				strcts = append(strcts, strct)
			}
		}
	}

	fmt.Println()

	for e := range exist {
		fmt.Println(e)
	}

	return strcts
}

func (xj *XJ) leafNodes(path, node string, m interface{}, l *[]leafNode, noattr bool) {
	// fmt.Println("path =>", path) //, "\tnode =>", node, "\tnoattr =>", noattr
	if !noattr || node != "#text" {
		if node != "" {
			if path != "" && node[:1] != "[" {
				path += "."
			}
			path += node
		}
	}

	switch m.(type) {
	case map[string]interface{}:
		i := 0
		for k, v := range m.(map[string]interface{}) {
			i++
			if noattr {
				continue
			}
			xj.leafNodes(path, k, v, l, noattr)
		}
		// to fix when m is empty, TODO: need better code
		// if i == 0 {
		// 	n := leafNode{path, m}
		// 	*l = append(*l, n)
		// }
	case []interface{}:
		for i, v := range m.([]interface{}) {
			xj.leafNodes(path, "["+strconv.Itoa(i)+"]", v, l, noattr)
		}
	default:
		n := leafNode{path, m}
		*l = append(*l, n)
	}
}
