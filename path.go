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

// sort.Interface implementation
type byName []strctNode

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return strings.Compare(a[i].Name, a[j].Name) < 0 }

type strctMap map[string][]strctNode

func leafNodesToStrcts(typ string, nodes *[]leafNode) []strctMap {
	n := max(nodes)
	root := strings.Split((*nodes)[0].path, ".")[0]

	exist := make(map[string]bool)

	strct := leafNodesToStruct(typ, root, root, nodes, &exist)
	strcts := []strctMap{}
	strcts = append(strcts, strct)

	es := []string{}
	for i := 0; i < n; i++ {
		for e := range exist {
			es = strings.Split(e, ".")
			root = es[len(es)-1]
			strct = leafNodesToStruct(typ, e, root, nodes, &exist)
			appendStrctNode(&strct, &strcts)
		}
	}

	return strcts
}

func leafNodesToStruct(typ, e, root string, nodes *[]leafNode, exist *map[string]bool) strctMap {
	strct := make(strctMap)
	var spath string
	re := regexp.MustCompile(`\[\d+\]`)
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

			if node.path != "" {
				leafNodeToStruct(typ, e, root, &node, &strct, exist, re)
			}
		}
	}

	return strct
}

func leafNodeToStruct(typ, e, root string, node *leafNode, strct *strctMap, exist *map[string]bool, re *regexp.Regexp) {
	s := strings.Split(node.path, ".")
	if len(s) >= 1 {
		name := re.ReplaceAllString(s[0], "")
		ek := e + "." + name
		sname := toProperCase(name)
		if !(*exist)[ek] {
			sn := strctNode{
				Name: sname,
			}
			if re.MatchString(s[0]) {
				if len(s) > 1 {
					sn.Type = "[]" + sname
				} else {
					sn.Type = "[]" + toProperType(node.value)
				}
			} else {
				if len(s) > 1 {
					sn.Type = sname
				} else {
					sn.Type = toProperType(node.value)
				}
			}
			switch node.value.(type) {
			case xmlVal:
				sn.Tag = "`" + typ + ":\"" + name + ",attr\"`"
			default:
				sn.Tag = "`" + typ + ":\"" + name + "\"`"
			}

			/*
				if len(s) > 1 {
					if re.MatchString(s[0]) {
						sn.Type = "[]" + sname
					} else {
						sn.Type = sname
					}
					sn.Tag = "`" + typ + ":\"" + name + "\"`"
				} else {
					sn.Type = toProperType(node.value)
					switch node.value.(type) {
					case xmlVal:
						sn.Tag = "`" + typ + ":\"" + name + ",attr\"`"
					// case string:
					// sn.Tag = "`" + typ + ":\"" + name + "\"`"
					default:
						sn.Tag = "`" + typ + ":\"" + name + "\"`"
					}
				}
			*/
			(*strct)[root] = append((*strct)[root], sn)
			(*exist)[ek] = true
		}
	}
}
