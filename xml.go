package xj2go

import (
	"encoding/xml"
	"errors"
	"io"
	"strings"
)

func (xj *XJ) xmlToMap(sk string, attr []xml.Attr) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	ma := make(map[string]interface{})
	if sk != "" {
		for _, v := range attr {
			ma[v.Name.Local] = v.Value
		}
	}

	for {
		t, err := xj.d.Token()
		if err != nil {
			if err != io.EOF {
				return nil, errors.New("xml.Decoder.Token() - " + err.Error())
			}
			return nil, err
		}
		switch t.(type) {
		case xml.StartElement:
			tt := t.(xml.StartElement)
			if sk == "" {
				return xj.xmlToMap(tt.Name.Local, tt.Attr)
			}
			mm, err := xj.xmlToMap(tt.Name.Local, tt.Attr)
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
			tt := strings.Trim(string(t.(xml.CharData)), "\t\r\b\n ")
			if tt != "" {
				if sk != "" {
					m[sk] = tt
				}
			}
		default:
		}
	}
}
