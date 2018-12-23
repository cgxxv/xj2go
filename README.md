# xj2go

[![Go Report Card](https://goreportcard.com/badge/github.com/stackerzzq/xj2go)](https://goreportcard.com/badge/github.com/stackerzzq/xj2go) [![Build Status](https://www.travis-ci.org/stackerzzq/xj2go.svg?branch=master)](https://www.travis-ci.org/stackerzzq/xj2go) [![codecov](https://codecov.io/gh/stackerzzq/xj2go/branch/master/graph/badge.svg)](https://codecov.io/gh/stackerzzq/xj2go) [![codebeat badge](https://codebeat.co/badges/baec2a13-1f35-4032-bbf4-66cbead635c4)](https://codebeat.co/projects/github-com-stackerzzq-xj2go-master)

The goal is to convert xml or json file to go struct file.

## Usage

Download and install it:
```sh
$ go get -u -v github.com/stackerzzq/xj2go/cmd/...

$ xj [-t json/xml] [-p sample] [-r result] sample.json
```
Import it in your code:
```go
import "github.com/stackerzzq/xj2go"
```
## Example

Please see [the example file](example/sample.go).

[embedmd]:# (example/sample.go go)
```go
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

	if err := xj2go.XMLBytesToGo("test", "demoxml2", &b1); err != nil {
		log.Fatal(err)
	}

	jsonfilename := "../testjson/githubAPI.json"
	xj2 := xj2go.New(jsonfilename, "demojson", "sample")
	xj2.JSONToGo()

	b2, err := ioutil.ReadFile(jsonfilename)
	if err != nil {
		log.Fatal(err)
	}

	if err := xj2go.JSONBytesToGo("test", "demojson2", "github", &b2); err != nil {
		log.Fatal(err)
	}
}
```
