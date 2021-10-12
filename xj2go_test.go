package xj2go

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
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
		"./testxml/test.xml",
		"./testxml/sample.xml",
	}

	for k, v := range fs {
		t.Run("xml to paths"+strconv.Itoa(k), func(t *testing.T) {
			_, err := xmlToLeafNodes(v)
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
		"./testxml/test.xml",
		"./testxml/sample.xml",
	}

	pkgname := "excel"
	for k, v := range fs {
		t.Run("xml to go "+strconv.Itoa(k), func(t *testing.T) {
			xj := New(v, pkgname, "")
			err := xj.XMLToGo()
			if err != nil {
				t.Errorf("XMLToGo() error = %v", err)
				return
			}
		})
	}
}

func Test_XMLBytesToGo(t *testing.T) {
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
		if err := XMLBytesToGo(file.Name, pkgname, &b); err != nil {
			t.Errorf("XMLBytesToGo() error = %v", err)
		}
	}
}

func TestXJ_JSONToGo(t *testing.T) {
	fs := []string{
		"./testjson/topics.json",
		"./testjson/smartStreetsAPI.json",
		"./testjson/githubAPI.json",
	}

	pkgname := "xjson"
	for k, v := range fs {
		t.Run("json to go "+strconv.Itoa(k), func(t *testing.T) {
			root := strings.TrimSuffix(path.Base(v), path.Ext(v))
			xj := New(v, pkgname, root)
			err := xj.JSONToGo()
			if err != nil {
				t.Errorf("JSONToGo() error = %v", err)
				return
			}
		})
	}
}

func Test_JSONBytesToGo(t *testing.T) {
	fs := []string{
		"./testjson/topics.json",
		"./testjson/smartStreetsAPI.json",
		"./testjson/githubAPI.json",
	}

	pkgname := "xjson2"
	for k, v := range fs {
		t.Run("json bytes to go "+strconv.Itoa(k), func(t *testing.T) {
			root := strings.TrimSuffix(path.Base(v), path.Ext(v))
			b, err := ioutil.ReadFile(v)
			if err != nil {
				t.Errorf("ioutil.ReadFile() error = %v", err)
			}
			if err := JSONBytesToGo(path.Base(v), pkgname, root, &b); err != nil {
				t.Errorf("JSONBytesToGo() error = %v", err)
				return
			}
		})
	}
}
