package patch

import (
	"bytes"
	"encoding/json"
)

// Field is a PatchField that supports JSON PATCH semantics:
// omitted key → Unset; null → SetNull; value → Set.
type Field[T any] struct {
	PatchField[T]
}

// UnmarshalJSON implements json.Unmarshaler for PATCH request bodies.
func (f *Field[T]) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		f.PatchField = Unset[T]()
		return nil
	}
	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		f.PatchField = SetNull[T]()
		return nil
	}
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	f.PatchField = Set(v)
	return nil
}
