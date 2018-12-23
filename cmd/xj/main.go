package main

import (
	"flag"
	"fmt"

	"github.com/stackerzzq/xj2go"
)

func usage() {
	fmt.Println("Usage: xj [-t json/xml] [-p sample] [-r result] sample.json")
	flag.PrintDefaults()
}

func main() {
	t := flag.String("t", "xml", "Type to parse\navaliable type:xml,json")
	p := flag.String("p", "sample", "Package name to generate")
	r := flag.String("r", "", "Root name for struct") //TODO: useless for xml file
	flag.Parse()

	if flag.NArg() > 0 {
		xj := xj2go.New(flag.Arg(0), *p, *r)
		switch *t {
		case "xml":
			if err := xj.XMLToGo(); err != nil {
				panic(err)
			}
		case "json":
			if err := xj.JSONToGo(); err != nil {
				panic(err)
			}
		}
	} else {
		usage()
	}
}
