package xj2go

import (
	"encoding/xml"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

type xmlVal struct {
	val  string
	attr bool
}

func xmlToLeafNodes(filename string) ([]leafNode, error) {
	m, err := xmlToMap(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return leafNodes(&m)
}

func xmlToMap(filename string) (map[string]interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	m, err := decodeXML(xml.NewDecoder(file), "", nil)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func decodeXML(d *xml.Decoder, sk string, attr []xml.Attr) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	ma := make(map[string]interface{})
	if sk != "" {
		for _, v := range attr {
			ma[v.Name.Local] = xmlVal{v.Value, true}
		}
	}

	for {
		t, err := d.Token()
		if err != nil {
			if err != io.EOF {
				return nil, errors.New("xml.Decoder.Token() - " + err.Error())
			}
			return nil, err
		}
		switch element := t.(type) {
		case xml.StartElement:
			if sk == "" {
				return decodeXML(d, element.Name.Local, element.Attr)
			}
			mm, err := decodeXML(d, element.Name.Local, element.Attr)
			if err != nil {
				return nil, err
			}

			var k string
			var v interface{}
			for k, v = range mm {
				break
			}

			if vv, ok := ma[k]; ok {
				var a []interface{}
				switch vv.(type) {
				case []interface{}:
					a = vv.([]interface{})
				default:
					a = []interface{}{vv}
				}
				a = append(a, v)
				ma[k] = a
			} else {
				ma[k] = v
			}
		case xml.EndElement:
			if len(ma) > 0 {
				m[sk] = ma
			}
			return m, nil
		case xml.CharData:
			tt := strings.Trim(string(element), "\t\r\b\n ")
			if tt != "" {
				if sk != "" {
					m[sk] = tt
				}
			}
		default:
		}
	}
}
