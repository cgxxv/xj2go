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
		sp := strings.TrimPrefix(path, e)
		if len(sp) <= 0 || (sp[:1] != "[" && sp[:1] != ".") {
			continue
		}
		if len(path) > len(e) && path[:len(e)] == e {
			path = strings.TrimPrefix(path, e)
			if path == "" {
				continue
			}
			path = strings.TrimPrefix(path, ".")
			if path[:1] == "[" {
				path = re.ReplaceAllString(path, "")
				path = strings.TrimPrefix(path, ".")
			}

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
	}

	return strct
}
