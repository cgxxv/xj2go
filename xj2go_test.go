package xj2go

import (
	"path"
	"testing"
)

func Test_xmlToMap(t *testing.T) {
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
			xj := New(v)
			_, err := xj.xmlToMap("", nil)
			if err != nil {
				t.Errorf("xmlToMap() error = %v", err)
				return
			}
		})
	}
}

func Test_XMLToStruct(t *testing.T) {
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

	pkg := "excel"
	for k, v := range fs {
		t.Run("leaf paths "+string(k), func(t *testing.T) {
			xj := New(v)
			filename := path.Base(v)
			err := xj.XMLToStruct(pkg+"/"+filename+".go", pkg)
			if err != nil {
				t.Errorf("XMLToStruct() error = %v", err)
				return
			}
		})
	}
}
