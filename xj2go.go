package xj2go

import (
	"bytes"
	"encoding/xml"
	"log"
)

// XJ define xj2go struct
type XJ struct {
	// xml or json file
	Filepath string
	// the pkg name for struct
	Pkgname string
	// the root name for json bytes
	Rootname string
}

// New return a xj2go instance
func New(filepath, pkgname, rootname string) *XJ {
	return &XJ{
		Filepath: filepath,
		Pkgname:  pkgname,
		Rootname: rootname,
	}
}

// XMLToGo convert xml to go struct, then write this struct to a go file
func (xj *XJ) XMLToGo() error {
	filename, err := checkFile(xj.Filepath, xj.Pkgname)
	if err != nil {
		log.Fatal(err)
		return err
	}

	nodes, err := xmlToLeafNodes(xj.Filepath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	strcts := leafNodesToStrcts("xml", nodes)

	return writeStruct(filename, xj.Pkgname, strcts)
}

// XMLBytesToGo convert xml bytes to struct, then the struct will be writed to ./{pkg}/{filename}.go
func XMLBytesToGo(filename, pkgname string, b *[]byte) error {
	filename, err := checkFile(filename, pkgname)
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
	strcts := leafNodesToStrcts("xml", nodes)

	return writeStruct(filename, pkgname, strcts)
}

// JSONToGo convert json to go struct, then write this struct to a go file
func (xj *XJ) JSONToGo() error {
	filename, err := checkFile(xj.Filepath, xj.Pkgname)
	if err != nil {
		log.Fatal(err)
		return err
	}

	nodes, err := jsonToLeafNodes(xj.Rootname, xj.Filepath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	strcts := leafNodesToStrcts("json", nodes)

	return writeStruct(filename, xj.Pkgname, strcts)
}

// JSONBytesToGo convert json bytes to struct, then the struct will be writed to ./{pkg}/{filename}.go
func JSONBytesToGo(filename, pkgname, rootname string, b *[]byte) error {
	filename, err := checkFile(filename, pkgname)
	if err != nil {
		log.Fatal(err)
		return err
	}

	m, err := jsonBytesToMap(pkgname, rootname, b)
	if err != nil {
		log.Fatal(err)
		return err
	}

	ns, err := leafNodes(&m)
	if err != nil {
		log.Fatal(err)
		return err
	}

	nodes, err := reLeafNodes(ns, rootname)
	if err != nil {
		log.Fatal(err)
		return err
	}

	strcts := leafNodesToStrcts("json", nodes)

	return writeStruct(filename, pkgname, strcts)
}
