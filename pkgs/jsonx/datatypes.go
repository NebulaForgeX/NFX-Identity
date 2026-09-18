package jsonx

import (
	"encoding/json"

	"gorm.io/datatypes"
)

// Marshal encodes v as gorm datatypes.JSON.
func Marshal(v any) (datatypes.JSON, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return datatypes.JSON(b), nil
}

// MarshalOr is like Marshal but returns fallback when encoding fails.
func MarshalOr(v any, fallback datatypes.JSON) datatypes.JSON {
	j, err := Marshal(v)
	if err != nil {
		return fallback
	}
	return j
}
