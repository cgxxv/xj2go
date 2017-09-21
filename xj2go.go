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
func (xj *XJ) XMLToStruct(filename, pkg string) error {
	m, _ := xj.xmlToMap("", nil)
	l := &[]leafNode{}
	xj.leafNodes("", "", m, l)
	paths := xj.leafPaths(*l)
	strcts := xj.pathsToNodes(paths)

	if ok, _ := pathExists(pkg); !ok {
		os.Mkdir(pkg, 0755)
	}

	if ok, _ := pathExists(filename); ok {
		if err := os.Remove(filename); err != nil {
			log.Fatal(err)
			return err
		}
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	// s := "package " + pkg + "\n\n"
	file.WriteString("package " + pkg + "\n\n")
	for _, strct := range strcts {
		for root, snodes := range strct {
			// s += "type " + strings.Title(root) + " struct {\n"
			file.WriteString("type " + strings.Title(root) + " struct {\n")
			for i := 0; i < len(snodes); i++ {
				typ := snodes[i].Type
				if typ != "string" {
					typ = strings.Title(snodes[i].Type)
				}

				// s += "\t" + strings.Title(snodes[i].Name) + "\t" + typ + "\t" + snodes[i].Tag + "\n"
				file.WriteString("\t" + strings.Title(snodes[i].Name) + "\t" + typ + "\t" + snodes[i].Tag + "\n")
			}
			// s += "}\n"
			file.WriteString("}\n")
		}
	}
	file.WriteString("\n")
	ft := exec.Command("go", "fmt", filename)
	if err := ft.Run(); err != nil {
		log.Fatal(err)
		return err
	}

	vt := exec.Command("go", "vet", filename)
	if err := vt.Run(); err != nil {
		log.Fatal(err)
		return err
	}

	// log.Println(s)

	return nil
}
