package goutil

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Map map[string]any

func (m Map) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *Map) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	bs, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("expected byte []byte, got %T", src)
	}

	var mx Map
	err := json.Unmarshal(bs, &mx)
	if err != nil {
		return err
	}

	*m = Map(mx)
	return nil
}
