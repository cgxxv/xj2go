package xj2go

import (
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
				m := 0
				for key, vals := range strct {
					for _, stct := range strcts {
						if _, ok := stct[key]; !ok {
							continue
						}
						for j := 0; j < len(vals); j++ {
							n := 0
							for k := 0; k < len(stct[key]); k++ {
								if vals[j].Name == stct[key][k].Name {
									n++
								}
							}
							if n == 0 {
								stct[key] = append(stct[key], vals[j])
							}
						}
						m++
					}
				}
				if m == 0 {
					strcts = append(strcts, strct)
				}
			}
		}
	}

	return strcts
}

func (xj *XJ) leafNodes(path, node string, m interface{}, l *[]leafNode) {
	if node != "" {
		if path != "" && node[:1] != "[" {
			path += "."
		}
		path += node
	}

	switch m.(type) {
	case map[string]interface{}:
		for k, v := range m.(map[string]interface{}) {
			xj.leafNodes(path, k, v, l)
		}
	case []interface{}:
		for i, v := range m.([]interface{}) {
			xj.leafNodes(path, "["+strconv.Itoa(i)+"]", v, l)
		}
	default:
		if m != nil {
			n := leafNode{path, m}
			*l = append(*l, n)
		}
	}
}
