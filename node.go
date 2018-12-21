package xj2go

import (
	"strconv"
)

type leafNode struct {
	path  string
	value interface{}
}

func appendStrctNode(strct *strctMap, strcts *[]strctMap) {
	m := 0
	for key, vals := range *strct {
		for _, stct := range *strcts {
			if _, ok := stct[key]; !ok {
				continue
			}
			for j := 0; j < len(vals); j++ {
				n := 0
				for k := 0; k < len(stct[key]); k++ {
					if vals[j].Name == stct[key][k].Name && vals[j].Type == stct[key][k].Type {
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
		*strcts = append(*strcts, *strct)
	}
}

func leafNodes(m *map[string]interface{}) ([]leafNode, error) {
	l := []leafNode{}
	getLeafNodes("", "", *m, &l)

	// paths := []string{}
	// for i := 0; i < len(l); i++ {
	// 	paths = append(paths, l[i].path)
	// }

	return l, nil
}

func reLeafNodes(lns []leafNode, prefix string) ([]leafNode, error) {
	ls := []leafNode{}
	var l leafNode
	for _, ln := range lns {
		l = leafNode{
			path:  prefix + "." + ln.path,
			value: ln.value,
		}
		ls = append(ls, l)
	}

	return ls, nil
}

func getLeafNodes(path, node string, m interface{}, l *[]leafNode) {
	if node != "" {
		if path != "" && node[:1] != "[" {
			path += "."
		}
		path += node
	}

	switch mm := m.(type) {
	case map[string]interface{}:
		for k, v := range mm {
			getLeafNodes(path, k, v, l)
		}
	case []interface{}:
		for i, v := range mm {
			getLeafNodes(path, "["+strconv.Itoa(i)+"]", v, l)
		}
	default:
		if mm != nil {
			n := leafNode{path, mm}
			*l = append(*l, n)
		}
	}
}
