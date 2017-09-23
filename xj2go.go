package xj2go

import (
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

// XJ define xj2go struct
type XJ struct {
	File string
	Pkg  string
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

type strctMap map[string][]strctNode

// New return a xj2go instance
func New(xmlfile, pkgname string) *XJ {
	return &XJ{
		File: xmlfile,
		Pkg:  pkgname,
	}
}

// XMLToStruct convert xml to go struct and write this struct to a go file
func (xj *XJ) XMLToStruct() error {
	paths, err := xj.xmlToPaths()
	if err != nil {
		log.Fatal(err)
		return err
	}
	strcts := xj.pathsToNodes(&paths)

	if ok, _ := pathExists(xj.Pkg); !ok {
		os.Mkdir(xj.Pkg, 0755)
	}

	filename := xj.Pkg + "/" + path.Base(xj.File) + ".go"
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

	file.WriteString("package " + xj.Pkg + "\n\n")
	for _, strct := range strcts {
		for root, sns := range strct {
			file.WriteString("type " + strings.Title(root) + " struct {\n")
			for i := 0; i < len(sns); i++ {
				if sns[i].Type != "string" {
					file.WriteString("\t" + strings.Title(sns[i].Name) + "\t" + strings.Title(sns[i].Type) + "\t" + sns[i].Tag + "\n")
				} else {
					file.WriteString("\t" + strings.Title(sns[i].Name) + "\t" + sns[i].Type + "\t" + sns[i].Tag + "\n")
				}
			}
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

	return nil
}

func (xj *XJ) xmlToPaths() ([]string, error) {
	m, err := xmlToMap(xj.File)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	l := []leafNode{}
	leafNodes("", "", m, &l)

	paths := []string{}
	for i := 0; i < len(l); i++ {
		paths = append(paths, l[i].path)
	}

	return paths, nil
}

func (xj *XJ) pathsToNodes(paths *[]string) []strctMap {
	n := max(paths)
	root := strings.Split((*paths)[0], ".")[0]

	re := regexp.MustCompile(`\[\d+\]`)
	exist := make(map[string]bool)

	strct := leafPath(root, root, paths, &exist, re)
	strcts := []strctMap{}
	strcts = append(strcts, strct)

	es := []string{}
	for i := 0; i < n; i++ {
		for e := range exist {
			es = strings.Split(e, ".")
			root = es[len(es)-1]
			strct = leafPath(e, root, paths, &exist, re)
			appendStrctNode(&strct, &strcts)
		}
	}

	return strcts
}
