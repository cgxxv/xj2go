package xj2go

import (
	"bytes"
	"encoding/xml"
	"log"
	"os"
	"path"
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

// XMLToGo convert xml to go struct and write this struct to a go file
func (xj *XJ) XMLToGo() error {
	paths, err := xmlToPaths(xj.File)
	if err != nil {
		log.Fatal(err)
		return err
	}
	strcts := pathsToStrcts(&paths)

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

	return writeStruct(filename, xj.Pkg, &strcts)
}

// ByteToGo convert xml to struct
func ByteToGo(filename, pkg string, b *[]byte) error {
	r := bytes.NewReader(*b)
	m, err := decodeXML(xml.NewDecoder(r), "", nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	paths, err := leafPaths(&m)
	if err != nil {
		log.Fatal(err)
		return err
	}
	strcts := pathsToStrcts(&paths)

	return writeStruct(filename, pkg, &strcts)
}
