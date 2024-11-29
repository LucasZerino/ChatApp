package models

import (
	"database/sql/driver"
	"encoding/json"
)

// JSONB é uma estrutura personalizada para lidar com campos JSONB no GORM
type JSONB map[string]interface{}

// Scan implementa a interface Scanner para a estrutura JSONB
func (jb *JSONB) Scan(value interface{}) error {
	if value == nil {
		*jb = make(map[string]interface{})
		return nil
	}
	return json.Unmarshal(value.([]byte), jb)
}

// Value implementa a interface Valuer para a estrutura JSONB
func (jb JSONB) Value() (driver.Value, error) {
	return json.Marshal(jb)
}
