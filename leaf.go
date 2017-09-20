package xj2go

import (
	"strconv"
	"strings"
)

func (xj *XJ) leafNodes(path, node string, m interface{}, l *[]leafNode, noattr bool) {
	if !noattr || node != "#text" {
		if path != "" && node[:1] != "[" {
			path += "."
		}
		path += node
	}

	switch m.(type) {
	case map[string]interface{}:
		for k, v := range m.(map[string]interface{}) {
			if noattr {
				continue
			}
			xj.leafNodes(path, k, v, l, noattr)
		}
	case []interface{}:
		for i, v := range m.([]interface{}) {
			xj.leafNodes(path, "["+strconv.Itoa(i)+"]", v, l, noattr)
		}
	default:
		n := leafNode{path, m}
		*l = append(*l, n)
	}
}

func (xj *XJ) leafPaths(lns []leafNode) []string {
	s := []string{}
	for i := 0; i < len(lns); i++ {
		s = append(s, lns[i].path)
	}

	return s
}

func (xj *XJ) leafPath(e, root string, paths []string) {
	for _, path := range paths {
		if strings.Index(path, e) == 0 {
			path = strings.TrimPrefix(path, e)
			if path == "" {
				continue
			}
			path = strings.TrimPrefix(path, ".")

			if strings.Index(path, "[") == 0 {
				path = re.ReplaceAllString(path, "")
				path = strings.TrimPrefix(path, ".")
			}

			s := strings.Split(path, ".")
			if len(s) >= 1 {
				name := re.ReplaceAllString(s[0], "")
				ek := e + "." + name //it is not same
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
							Tag:  "`xml:" + name + ",attr`",
						}
						strct[root] = append(strct[root], sn)
						exist[ek] = true
					}
				}
			}
		}
	}
}
