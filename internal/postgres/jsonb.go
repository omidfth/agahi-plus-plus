package postgres

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Hstore map[string]*string

// Jsonb Postgresql's JSONB data type
type Jsonb struct {
	json.RawMessage
}

// Value get value of Jsonb
func (j Jsonb) Value() (driver.Value, error) {
	if len(j.RawMessage) == 0 {
		return nil, nil
	}
	return j.MarshalJSON()
}

// Scan scan value into Jsonb
func (j *Jsonb) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	return json.Unmarshal(bytes, j)
}

func MakeJsonb(value interface{}) (Jsonb, error) {
	jb := Jsonb{}
	j, mErr := json.Marshal(value)
	if mErr != nil {
		return Jsonb{}, mErr
	}
	err := jb.Scan(j)
	if err != nil {
		return Jsonb{}, err
	}
	return jb, nil
}
