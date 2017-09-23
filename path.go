package xj2go

import (
	"regexp"
	"strings"
)

func leafPath(e, root string, paths *[]string, exist *map[string]bool, re *regexp.Regexp) strctMap {
	strct := make(strctMap)
	for _, path := range *paths {
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

			leafStrctPath(e, root, path, &strct, exist, re)
		}
	}

	return strct
}

func leafStrctPath(e, root, path string, strct *strctMap, exist *map[string]bool, re *regexp.Regexp) {
	s := strings.Split(path, ".")
	if len(s) >= 1 {
		name := re.ReplaceAllString(s[0], "")
		ek := e + "." + name
		if !(*exist)[ek] {
			var sn strctNode
			if len(s) > 1 {
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
			} else {
				sn = strctNode{
					Name: name,
					Type: "string",
					Tag:  "`xml:\"" + name + ",attr\"`",
				}
			}
			(*strct)[root] = append((*strct)[root], sn)
			(*exist)[ek] = true
		}
	}
}
