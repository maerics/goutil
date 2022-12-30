package goutil

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func MustJson(x any, pretty ...bool) string {
	var bs []byte
	var err error
	if len(pretty) > 0 && pretty[0] {
		bs, err = json.MarshalIndent(x, "", "  ")
	} else {
		bs, err = json.Marshal(x)
	}
	if err != nil {
		panic(err)
	}
	return string(bs)
}

// Marshals maps to JSON with specified key order.
type OrderedJsonObj struct {
	Map  map[string]any
	Keys []string
}

func (o OrderedJsonObj) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.WriteByte('{')
	for _, key := range o.Keys {
		if bs, err := json.Marshal(o.Map[key]); err != nil {
			return nil, err
		} else {
			fmt.Fprintf(buf, "%q:%v,", key, string(bs))
		}
	}
	bs := buf.Bytes()
	bs[len(bs)-1] = '}'
	return bs, nil
}
