package xj2go

import (
	"strings"
)

func (xj *XJ) leafPaths(lns []leafNode) []string {
	s := []string{}
	for i := 0; i < len(lns); i++ {
		s = append(s, lns[i].path)
	}

	return s
}

func (xj *XJ) leafPath(e, root string, paths []string) map[string][]strctNode {
	strct := make(map[string][]strctNode)
	for _, path := range paths {
		var spath string
		if eld := strings.LastIndex(e, "."); eld > 0 {
			elp := e[eld:] // with .

			if pos := strings.Index(path, elp); pos > 0 {
				pre := path[:pos]
				rest := strings.TrimPrefix(path, pre)
				pre = re.ReplaceAllString(pre, "") // replace [1] with ""
				spath = pre + rest
			}
		} else {
			spath = path
		}

		chkpath := strings.TrimPrefix(spath, e)
		if !(chkpath == "" || chkpath[:1] == "[" || chkpath[:1] == ".") {
			continue
		}

		if len(spath) > len(e) && spath[:len(e)] == e {
			path = strings.TrimPrefix(spath, e)
			if path == "" {
				continue
			}
			path = strings.TrimPrefix(path, ".")
			if path[:1] == "[" {
				path = re.ReplaceAllString(path, "")
				path = strings.TrimPrefix(path, ".")
			}

			xj.leafStrctPath(e, root, path, strct)
		}
	}

	return strct
}

func (xj *XJ) leafStrctPath(e, root, path string, strct map[string][]strctNode) {
	s := strings.Split(path, ".")
	if len(s) >= 1 {
		name := re.ReplaceAllString(s[0], "")
		ek := e + "." + name
		if !exist[ek] {
			if len(s) > 1 {
				var sn strctNode
				if re.MatchString(s[0]) {
					sn = strctNode{
						Name: name,
						Type: "[]" + name,
					}
				} else {
					sn = strctNode{
						Name: name,
						Type: name,
					}
				}
				strct[root] = append(strct[root], sn)
				exist[ek] = true
			} else {
				sn := strctNode{
					Name: name,
					Type: "string",
					Tag:  "`xml:\"" + name + ",attr\"`",
				}
				strct[root] = append(strct[root], sn)
				exist[ek] = true
			}
		}
	}
}
