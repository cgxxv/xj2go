package xj2go

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

func jsonToLeafNodes(root, filename string) ([]leafNode, error) {
	m, err := jsonFileToMap(root+"Result", filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	lns, err := leafNodes(&m)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return reLeafNodes(lns, root)
}

func jsonFileToMap(top, filename string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	val, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if len(val) == 0 {
		return m, nil
	}

	if val[0] == '[' {
		val = []byte(`{"` + top + `":` + string(val) + `}`)
	}

	return jsonDecode(&m, &val)
}

func jsonBytesToMap(pkg, root string, b *[]byte) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if (*b)[0] == '[' {
		*b = []byte(`{"` + root + "Result" + `":` + string(*b) + `}`)
	}

	return jsonDecode(&m, b)
}

func jsonDecode(m *map[string]interface{}, b *[]byte) (map[string]interface{}, error) {
	buf := bytes.NewReader(*b)
	dec := json.NewDecoder(buf)
	err := dec.Decode(m)
	return *m, err
}
