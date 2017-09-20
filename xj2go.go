package xj2go

import (
	"encoding/xml"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// XJ define xj2go struct
type XJ struct {
	d *xml.Decoder
}

type leafNode struct {
	path  string
	value interface{}
}

type strctNode struct {
	Name string
	Type string
	Tag  string
}

var (
	re    *regexp.Regexp
	strct map[string][]strctNode
	exist map[string]bool
)

// New return a xj2go instance
func New(filename string) *XJ {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &XJ{
		d: xml.NewDecoder(file),
	}
}

// XMLToStruct convert xml to go struct and write this struct to a go file
func (xj *XJ) XMLToStruct(filename, pkg string) {
	m, _ := xj.xmlToMap("", nil)
	l := &[]leafNode{}
	xj.leafNodes("", "", m, l, false)
	paths := xj.leafPaths(*l)
	strct := xj.pathsToNodes(paths)

	if ok, _ := pathExists(filename); ok {
		if err := os.Remove(filename); err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// s := "package " + pkg + "\n\n"
	file.WriteString("package " + pkg + "\n\n")
	for root, snodes := range strct {
		// s += "type " + strings.Title(root) + " struct {\n"
		file.WriteString("type " + strings.Title(root) + " struct {\n")
		for i := 0; i < len(snodes); i++ {
			typ := snodes[i].Type
			if typ != "string" {
				typ = strings.Title(snodes[i].Type)
			}

			tag := snodes[i].Tag
			if tag != "" {
				tag = "\t`" + tag + "`"
			}

			// s += "\t" + strings.Title(snodes[i].Name) + "\t" + typ + tag + "\n"
			file.WriteString("\t" + strings.Title(snodes[i].Name) + "\t" + typ + tag + "\n")
		}
		// s += "}\n"
		file.WriteString("}\n")
	}
	file.WriteString("\n")
	cmd := exec.Command("go", "fmt", filename)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	// log.Println(s)
}
