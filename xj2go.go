package xj2go

import (
	"bytes"
	"encoding/xml"
	"log"
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
	filename, err := checkFile(xj.File, xj.Pkg)
	if err != nil {
		log.Fatal(err)
		return err
	}

	paths, err := xmlToPaths(xj.File)
	if err != nil {
		log.Fatal(err)
		return err
	}
	strcts := pathsToStrcts(&paths)

	return writeStruct(filename, xj.Pkg, &strcts)
}

// BytesToGo convert xml byte to struct
func BytesToGo(filename, pkg string, b *[]byte) error {
	filename, err := checkFile(filename, pkg)
	if err != nil {
		log.Fatal(err)
		return err
	}

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
