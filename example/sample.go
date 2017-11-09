package main

import (
	"io/ioutil"
	"log"

	"github.com/stackerzzq/xj2go"
)

func main() {
	xmlfilename := "../testxml/xl/styles.xml"
	xj1 := xj2go.New(xmlfilename, "demoxml", "")
	xj1.XMLToGo()

	b1, err := ioutil.ReadFile(xmlfilename)
	if err != nil {
		log.Fatal(err)
	}

	if err := xj2go.XMLBytesToGo("test.go", "demoxml2", &b1); err != nil {
		log.Fatal(err)
	}

	jsonfilename := "../testjson/githubAPI.json"
	xj2 := xj2go.New(jsonfilename, "demojson", "sample")
	xj2.JSONToGo()

	b2, err := ioutil.ReadFile(jsonfilename)
	if err != nil {
		log.Fatal(err)
	}

	if err := xj2go.JSONBytesToGo("test.go", "demojson2", "github", &b2); err != nil {
		log.Fatal(err)
	}
}
