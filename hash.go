package goutil

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Hash map[string]any

func (m Hash) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *Hash) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	bs, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("expected byte []byte, got %T", src)
	}

	var mx Hash
	err := json.Unmarshal(bs, &mx)
	if err != nil {
		return err
	}

	*m = Hash(mx)
	return nil
}
