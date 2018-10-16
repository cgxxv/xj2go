package xj2go

import (
	"os"
	"reflect"
	"regexp"
	"strings"
)

var commonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XSRF":  true,
	"XSS":   true,
}

func max(nodes *[]leafNode) int {
	n := 0
	for _, node := range *nodes {
		t := strings.Count(node.path, ".")
		if n < t {
			n = t
		}
	}

	return n + 1
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

var toProperCaseRE = regexp.MustCompile(`([A-Z])([a-z]+)`)

// https://github.com/golang/lint/blob/39d15d55e9777df34cdffde4f406ab27fd2e60c0/lint.go#L695-L731
func toProperCase(str string) string {
	subProperCase := func(v string) string {
		if commonInitialisms[strings.ToTitle(v)] {
			v = strings.ToTitle(v)
		} else {
			v = strings.Title(v)
		}

		return v
	}
	str = toProperCaseRE.ReplaceAllStringFunc(str, subProperCase)
	s := strings.Split(str, "_")
	str = ""
	for _, v := range s {
		str += subProperCase(v)
	}

	return str
}

var toProperTypeRE = regexp.MustCompile(`\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d(\.\d+)?(\+\d\d:\d\d|Z)`)

//TODO: should be optimize for time type
func toProperType(v interface{}) string {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.String {
		if toProperTypeRE.MatchString(v.(string)) {
			return "time.Time"
		}
	}

	if t.Kind() == reflect.Struct {
		//detect xmlVal struct val field
		for i := 0; i < t.NumField(); i++ {
			if t.Field(i).Name == "val" {
				return t.Field(i).Type.String()
			}
		}
	}
	// if _, isTime := v.(time.Time); isTime {
	// 	return "time.Time"
	// }

	return t.String()
}
