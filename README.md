# xj2go

[![Go Report Card](https://goreportcard.com/badge/github.com/stackerzzq/xj2go)](https://goreportcard.com/badge/github.com/stackerzzq/xj2go)
[![Build Status](https://www.travis-ci.org/stackerzzq/xj2go.svg?branch=master)](https://www.travis-ci.org/stackerzzq/xj2go)
[![codecov](https://codecov.io/gh/stackerzzq/xj2go/branch/master/graph/badge.svg)](https://codecov.io/gh/stackerzzq/xj2go)
[![codebeat badge](https://codebeat.co/badges/39396cd9-aa4d-4801-a732-d3ef7dda7580)](https://codebeat.co/projects/github-com-stackerzzq-xj2go-master)
[![apache 2.0 license](https://img.shields.io/badge/license-apache-2.0.svg)](https://img.shields.io/badge/license-apache-2.0.svg)

The goal is to convert xml or json file to go struct file.

## usage

1. convert xml to go struct

```go
package main

import "github.com/stackerzzq/xj2go"

func main() {
  filename := "test.xml"
  pkg := "demo"
  xj := xj2go.New(filename)
  xj.XMLToStruct(pkg+"/"+filename+".go", pkg)
}
```

2. convert json to go struct

TBD