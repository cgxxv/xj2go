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
	Root string
}

// New return a xj2go instance
func New(xmlfile, pkgname, root string) *XJ {
	return &XJ{
		File: xmlfile,
		Pkg:  pkgname,
		Root: root,
	}
}

// XMLToGo convert xml to go struct and write this struct to a go file
func (xj *XJ) XMLToGo() error {
	filename, err := checkFile(xj.File, xj.Pkg)
	if err != nil {
		log.Fatal(err)
		return err
	}

	nodes, err := xmlToLeafNodes(xj.File)
	if err != nil {
		log.Fatal(err)
		return err
	}
	strcts := leafNodesToStrcts("xml", &nodes)

	return writeStruct(filename, xj.Pkg, &strcts)
}

// XMLBytesToGo convert xml byte to struct
func XMLBytesToGo(filename, pkg string, b *[]byte) error {
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

	nodes, err := leafNodes(&m)
	if err != nil {
		log.Fatal(err)
		return err
	}
	strcts := leafNodesToStrcts("xml", &nodes)

	return writeStruct(filename, pkg, &strcts)
}

func (xj *XJ) JSONToGo() error {
	filename, err := checkFile(xj.File, xj.Pkg)
	if err != nil {
		log.Fatal(err)
		return err
	}

	nodes, err := jsonToLeafNodes(xj.Root, xj.File)
	if err != nil {
		log.Fatal(err)
		return err
	}
	strcts := leafNodesToStrcts("json", &nodes)

	return writeStruct(filename, xj.Pkg, &strcts)
}

func JSONBytesToGo(filename, pkg, root string, b *[]byte) error {
	filename, err := checkFile(filename, pkg)
	if err != nil {
		log.Fatal(err)
		return err
	}

	m, err := jsonBytesToMap(pkg, root, b)
	if err != nil {
		log.Fatal(err)
		return err
	}

	ns, err := leafNodes(&m)
	if err != nil {
		log.Fatal(err)
		return err
	}

	nodes, err := reLeafNodes(&ns, root)
	if err != nil {
		log.Fatal(err)
		return err
	}

	strcts := leafNodesToStrcts("json", &nodes)

	return writeStruct(filename, pkg, &strcts)
}
