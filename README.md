# xj2go

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