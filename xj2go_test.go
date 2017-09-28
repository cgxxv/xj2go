package xj2go

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"path"
	"testing"
)

func Test_xmlToPaths(t *testing.T) {
	fs := []string{
		"./testxml/[Content_Types].xml",
		"./testxml/xl/workbook.xml",
		"./testxml/xl/styles.xml",
		"./testxml/xl/sharedStrings.xml",
		"./testxml/xl/_rels/workbook.xml.rels",
		"./testxml/xl/theme/theme1.xml",
		"./testxml/xl/worksheets/sheet1.xml",
		"./testxml/docProps/app.xml",
		"./testxml/docProps/core.xml",
	}

	for k, v := range fs {
		t.Run("xml to map"+string(k), func(t *testing.T) {
			// paths, err := xj.xmlToPaths()
			// fmt.Println(paths)
			_, err := xmlToPaths(v)
			if err != nil {
				t.Errorf("xmlToMap() error = %v", err)
				return
			}
		})
	}
}

func Test_XMLToGo(t *testing.T) {
	fs := []string{
		"./testxml/[Content_Types].xml",
		"./testxml/xl/workbook.xml",
		"./testxml/xl/styles.xml",
		"./testxml/xl/sharedStrings.xml",
		"./testxml/xl/_rels/workbook.xml.rels",
		"./testxml/xl/theme/theme1.xml",
		"./testxml/xl/worksheets/sheet1.xml",
		"./testxml/docProps/app.xml",
		"./testxml/docProps/core.xml",
	}

	pkgname := "excel"
	for k, v := range fs {
		t.Run("xml to go "+string(k), func(t *testing.T) {
			// filename := pkgname + "/" + path.Base(v) + ".go"
			xj := New(v, pkgname)
			err := xj.XMLToGo()
			if err != nil {
				t.Errorf("XMLToGo() error = %v", err)
				return
			}
		})
	}
}

func Test_BytesToGo(t *testing.T) {
	filename := "./testxlsx/testxml.xlsx"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("ioutil.ReadFile() error = %v", err)
	}

	zr, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		t.Errorf("zip.NewReader() error = %v", err)
	}

	pkgname := "excel2"
	for _, file := range zr.File {
		if path.Base(file.Name)[:1] == "." {
			continue
		}

		rc, err := file.Open()
		if err != nil {
			t.Errorf("file.Open() error = %v", err)
		}
		buff := bytes.NewBuffer(nil)
		io.Copy(buff, rc)
		rc.Close()
		b := buff.Bytes()
		if err := BytesToGo(file.Name, pkgname, &b); err != nil {
			t.Errorf("BytesToGo() error = %v", err)
		}
	}
}
