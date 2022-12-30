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

// Marshals a JSON object using the given key order and values from a slice or map.
// The "Values" and "Map" fields are mutually exclusive.
type OrderedJsonObj struct {
	Keys   []string
	Values []any
	Map    map[string]any
	Nulls  bool
}

func (o OrderedJsonObj) MarshalJSON() ([]byte, error) {
	if o.Values != nil && o.Map != nil {
		panic(fmt.Errorf("OrderedJsonObj cannot have both Values and Map"))
	}
	if o.Values == nil && o.Map == nil {
		panic(fmt.Errorf("OrderedJsonObj must have either Values or Map"))
	}

	buf := &bytes.Buffer{}
	buf.WriteByte('{')
	for i, key := range o.Keys {
		var value any
		if o.Map != nil {
			value = o.Map[key]
		} else if i < len(o.Values) {
			value = o.Values[i]
		}
		if value == nil && !o.Nulls {
			continue
		}

		if bs, err := json.Marshal(value); err != nil {
			return nil, err
		} else {
			fmt.Fprintf(buf, "%q:%v,", key, string(bs))
		}
	}
	bs := buf.Bytes()
	bs[len(bs)-1] = '}'
	return bs, nil
}
