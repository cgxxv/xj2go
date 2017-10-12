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

func leafNodes(path, node string, m interface{}, l *[]leafNode) {
	if node != "" {
		if path != "" && node[:1] != "[" {
			path += "."
		}
		path += node
	}

	switch mm := m.(type) {
	case map[string]interface{}:
		for k, v := range mm {
			leafNodes(path, k, v, l)
		}
	case []interface{}:
		for i, v := range mm {
			leafNodes(path, "["+strconv.Itoa(i)+"]", v, l)
		}
	default:
		if mm != nil {
			n := leafNode{path, mm}
			*l = append(*l, n)
		}
	}
}
