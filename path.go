package xj2go

import (
	"regexp"
	"strings"
)

type strctNode struct {
	Name string
	Type string
	Tag  string
}

type strctMap map[string][]strctNode

func leafPaths(m *map[string]interface{}) ([]leafNode, error) {
	l := []leafNode{}
	leafNodes("", "", *m, &l)

	// paths := []string{}
	// for i := 0; i < len(l); i++ {
	// 	paths = append(paths, l[i].path)
	// }

	return l, nil
}

func leafPath(e, root string, nodes *[]leafNode, exist *map[string]bool, re *regexp.Regexp) strctMap {
	strct := make(strctMap)
	var spath string
	for _, node := range *nodes {
		if eld := strings.LastIndex(e, "."); eld > 0 {
			elp := e[eld:] // with .
			if pos := strings.Index(node.path, elp); pos > 0 {
				pre := node.path[:pos]
				rest := strings.TrimPrefix(node.path, pre)
				pre = re.ReplaceAllString(pre, "") // replace [1] with ""
				spath = pre + rest
			}
		} else {
			spath = node.path
		}

		chkpath := strings.TrimPrefix(spath, e)
		if !(chkpath == "" || chkpath[:1] == "[" || chkpath[:1] == ".") {
			continue
		}

		if len(spath) > len(e) && spath[:len(e)] == e {
			node.path = strings.TrimPrefix(spath, e)
			if node.path == "" {
				continue
			}
			node.path = strings.TrimPrefix(node.path, ".")
			if node.path[:1] == "[" {
				loc := re.FindStringIndex(node.path)
				node.path = strings.Replace(node.path, node.path[loc[0]:loc[1]], "", 1)
				node.path = strings.TrimPrefix(node.path, ".")
			}

			leafStrctPath(e, root, &node, &strct, exist, re)
		}
	}

	return strct
}

func leafStrctPath(e, root string, node *leafNode, strct *strctMap, exist *map[string]bool, re *regexp.Regexp) {
	s := strings.Split(node.path, ".")
	if len(s) >= 1 {
		name := re.ReplaceAllString(s[0], "")
		ek := e + "." + name
		if !(*exist)[ek] {
			sn := strctNode{
				Name: name,
			}
			if len(s) > 1 {
				if re.MatchString(s[0]) {
					sn.Type = "[]" + name
				} else {
					sn.Type = name
				}
				sn.Tag = "`xml:\"" + name + "\"`"
			} else {
				sn.Type = "string"
				switch node.value.(type) {
				case xmlVal:
					sn.Tag = "`xml:\"" + name + ",attr\"`"
				case string:
					sn.Tag = "`xml:\"" + name + "\"`"
				}
			}
			(*strct)[root] = append((*strct)[root], sn)
			(*exist)[ek] = true
		}
	}
}

func pathsToStrcts(nodes *[]leafNode) []strctMap {
	n := max(nodes)
	root := strings.Split((*nodes)[0].path, ".")[0]

	re := regexp.MustCompile(`\[\d+\]`)
	exist := make(map[string]bool)

	strct := leafPath(root, root, nodes, &exist, re)
	strcts := []strctMap{}
	strcts = append(strcts, strct)

	es := []string{}
	for i := 0; i < n; i++ {
		for e := range exist {
			es = strings.Split(e, ".")
			root = es[len(es)-1]
			strct = leafPath(e, root, nodes, &exist, re)
			appendStrctNode(&strct, &strcts)
		}
	}

	return strcts
}
