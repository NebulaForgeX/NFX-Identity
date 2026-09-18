package slicex

import "github.com/google/uuid"

// ToStringSlice safely converts an arbitrary value into []string when possible.
//
// Supported inputs:
//   - []string: returned as-is.
//   - []any: only string elements are copied into the result; non-string elements are skipped.
//   - nil or unsupported types: returns nil.
//
// Parameters:
//   - v: value to convert.
//
// Returns:
//   - []string: converted string slice, or nil when v cannot be converted.
//
// Example:
//
//	values := slicex.ToStringSlice([]any{"a", 1, "b", true})
//
// Example result:
//   - values == []string{"a", "b"}
//
// Direct []string example:
//
//	values := slicex.ToStringSlice([]string{"x", "y"})
//
// Direct []string example result:
//   - values == []string{"x", "y"}
func ToStringSlice(v any) []string {
	if v == nil {
		return nil
	}
	switch vv := v.(type) {
	case []string:
		return vv
	case []any:
		out := make([]string, 0, len(vv))
		for _, x := range vv {
			if s, ok := x.(string); ok {
				out = append(out, s)
			}
		}
		return out
	default:
		return nil
	}
}

// ToAnySlice converts a typed slice into []any.
//
// Nil input returns nil. Non-nil empty input returns an empty []any.
//
// Type Parameters:
//   - T: input element type.
//
// Parameters:
//   - xs: typed slice to convert.
//
// Returns:
//   - []any: slice containing the same elements as interface values.
//
// Example:
//
//	values := slicex.ToAnySlice([]int{1, 2, 3})
//
// Example result:
//   - values == []any{1, 2, 3}
//
// Nil example:
//
//	var ids []int
//	values := slicex.ToAnySlice(ids)
//
// Nil example result:
//   - values == nil
func ToAnySlice[T any](xs []T) []any {
	if xs == nil {
		return nil
	}
	out := make([]any, len(xs))
	for i := range xs {
		out[i] = xs[i]
	}
	return out
}

// ToUUIDSlice converts string UUIDs into uuid.UUID values.
//
// Invalid UUID strings are converted to uuid.Nil and do not stop conversion.
// Nil input returns nil.
//
// Parameters:
//   - xs: UUID strings to parse.
//
// Returns:
//   - []uuid.UUID: parsed UUIDs; invalid entries become uuid.Nil.
//
// Example:
//
//	values := slicex.ToUUIDSlice([]string{
//		"550e8400-e29b-41d4-a716-446655440000",
//		"invalid",
//	})
//
// Example result:
//   - values[0] == uuid.MustParse("550e8400-e29b-41d4-a716-446655440000").
//   - values[1] == uuid.Nil.
func ToUUIDSlice(xs []string) []uuid.UUID {
	if xs == nil {
		return nil
	}
	out := make([]uuid.UUID, len(xs))
	for i := range xs {
		uid, err := uuid.Parse(xs[i])
		if err != nil {
			out[i] = uuid.Nil
			continue
		}
		out[i] = uid
	}
	return out
}

// ToUUIDSliceWithError converts string UUIDs into uuid.UUID values and fails on the first invalid value.
//
// Nil input returns nil, nil. Unlike ToUUIDSlice, invalid UUID strings are not
// replaced with uuid.Nil; the parse error is returned immediately.
//
// Parameters:
//   - xs: UUID strings to parse.
//
// Returns:
//   - []uuid.UUID: parsed UUIDs when every input is valid.
//   - error: first uuid.Parse error, or nil.
//
// Example:
//
//	values, err := slicex.ToUUIDSliceWithError([]string{
//		"550e8400-e29b-41d4-a716-446655440000",
//	})
//
// Example result:
//   - values == []uuid.UUID{uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")}.
//   - err == nil.
//
// Invalid example:
//
//	values, err := slicex.ToUUIDSliceWithError([]string{"invalid"})
//
// Invalid example result:
//   - values == nil.
//   - err != nil.
func ToUUIDSliceWithError(xs []string) ([]uuid.UUID, error) {
	if xs == nil {
		return nil, nil
	}
	out := make([]uuid.UUID, len(xs))
	for i := range xs {
		uid, err := uuid.Parse(xs[i])
		if err != nil {
			return nil, err
		}
		out[i] = uid
	}
	return out, nil
}

// UuidSliceToStrSlice converts UUID values into their canonical string form.
//
// Nil input returns nil. Non-nil empty input returns an empty []string.
//
// Parameters:
//   - xs: UUID values to stringify.
//
// Returns:
//   - []string: UUID strings in the same order as xs.
//
// Example:
//
//	values := slicex.UuidSliceToStrSlice([]uuid.UUID{
//		uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
//	})
//
// Example result:
//   - values == []string{"550e8400-e29b-41d4-a716-446655440000"}.
func UuidSliceToStrSlice(xs []uuid.UUID) []string {
	if xs == nil {
		return nil
	}
	out := make([]string, len(xs))
	for i := range xs {
		out[i] = xs[i].String()
	}
	return out
}
