// Package constantx 提供可比较枚举的白名单校验（ValidMember）、原始字符串校验（ValidString）与解析辅助。
package constantx

import (
	"strings"
)

// ValidString trims a raw string, converts it to T, and validates it.
//
// Use this when input arrives as a string, but your domain enum is a custom
// string type such as type Role string.
//
// Type Parameters:
//   - T: string-backed enum type.
//
// Parameters:
//   - s: raw string input. Leading/trailing spaces are ignored.
//   - valid: validator that accepts T and returns whether it is allowed.
//
// Returns:
//   - bool: true when the trimmed value is valid.
//
// Example:
//
//	type Role string
//	valid := func(r Role) bool { return r == "admin" || r == "member" }
//	ok := constantx.ValidString[Role](" admin ", valid)
//
// Example result:
//   - ok == true
func ValidString[T ~string](s string, valid func(T) bool) bool {
	return valid(T(strings.TrimSpace(s)))
}

// ParseString validates and converts one raw string into a string-backed enum value.
//
// The returned value is always strings.TrimSpace(s). If valid(s) returns false,
// the zero value of T and invalid are returned.
//
// Type Parameters:
//   - T: string-backed enum type.
//
// Parameters:
//   - s: raw string input.
//   - invalid: error returned when s is not valid.
//   - valid: validator for the raw string. It should normally trim internally.
//
// Returns:
//   - T: trimmed enum value when valid.
//   - error: nil when valid; otherwise invalid.
//
// Example:
//
//	type Role string
//	errInvalid := errors.New("invalid role")
//	valid := func(s string) bool { return s == "admin" || strings.TrimSpace(s) == "admin" }
//	role, err := constantx.ParseString[Role](" admin ", errInvalid, valid)
//
// Example result:
//   - role == Role("admin")
//   - err == nil
//
// Invalid example result:
//   - ParseString[Role]("guest", errInvalid, valid) returns "", errInvalid.
func ParseString[T ~string](s string, invalid error, valid func(string) bool) (T, error) {
	if !valid(s) {
		var z T
		return z, invalid
	}
	return T(strings.TrimSpace(s)), nil
}

// ParseStringsStrict validates multiple string values while skipping empty items.
//
// This is useful for repeated query parameters. Blank strings are ignored; any
// non-blank invalid value stops parsing and returns invalid.
//
// Type Parameters:
//   - T: string-backed enum type.
//
// Parameters:
//   - ss: raw string values.
//   - invalid: error returned when any non-empty value is invalid.
//   - valid: validator for each raw string.
//
// Returns:
//   - []T: parsed trimmed values, excluding blanks.
//   - error: nil when every non-blank value is valid; otherwise invalid.
//
// Example:
//
//	type Role string
//	errInvalid := errors.New("invalid role")
//	valid := func(s string) bool {
//		v := strings.TrimSpace(s)
//		return v == "admin" || v == "member"
//	}
//	roles, err := constantx.ParseStringsStrict[Role](
//		[]string{"admin", "", " member "},
//		errInvalid,
//		valid,
//	)
//
// Example result:
//   - roles == []Role{"admin", "member"}
//   - err == nil
//
// Invalid example result:
//   - ParseStringsStrict[Role]([]string{"admin", "guest"}, errInvalid, valid)
//     returns nil, errInvalid.
func ParseStringsStrict[T ~string](ss []string, invalid error, valid func(string) bool) ([]T, error) {
	out := make([]T, 0, len(ss))
	for _, s := range ss {
		if strings.TrimSpace(s) == "" {
			continue
		}
		if !valid(s) {
			return nil, invalid
		}
		out = append(out, T(strings.TrimSpace(s)))
	}
	return out, nil
}

// ParseStringSlice validates and converts every item in a string slice without skipping blanks.
//
// Unlike ParseStringsStrict, this function passes every element to ParseString,
// including empty strings. Use it for JSON arrays where an empty string should be
// treated as invalid unless valid("") explicitly allows it.
//
// Type Parameters:
//   - T: string-backed enum type.
//
// Parameters:
//   - ss: raw string values.
//   - invalid: error returned when any value is invalid.
//   - valid: validator for each raw string.
//
// Returns:
//   - []T: parsed trimmed values.
//   - error: nil when every value is valid; otherwise invalid.
//
// Example:
//
//	type Role string
//	errInvalid := errors.New("invalid role")
//	valid := func(s string) bool {
//		v := strings.TrimSpace(s)
//		return v == "admin" || v == "member"
//	}
//	roles, err := constantx.ParseStringSlice[Role](
//		[]string{" admin ", "member"},
//		errInvalid,
//		valid,
//	)
//
// Example result:
//   - roles == []Role{"admin", "member"}
//   - err == nil
//
// Empty-value example result:
//   - ParseStringSlice[Role]([]string{""}, errInvalid, valid) returns nil, errInvalid.
func ParseStringSlice[T ~string](ss []string, invalid error, valid func(string) bool) ([]T, error) {
	out := make([]T, 0, len(ss))
	for _, s := range ss {
		v, err := ParseString[T](s, invalid, valid)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, nil
}

// OptionalPtr parses an optional string-backed enum and returns a pointer when present.
//
// Blank input means "not provided" and returns nil, nil. Non-blank input must be
// valid and is returned as a pointer to the trimmed enum value.
//
// Type Parameters:
//   - T: string-backed enum type.
//
// Parameters:
//   - s: raw optional input.
//   - invalid: error returned when non-blank input is invalid.
//   - valid: validator for the raw string.
//
// Returns:
//   - *T: nil when s is blank; pointer to parsed value when valid.
//   - error: nil when blank/valid; otherwise invalid.
//
// Example:
//
//	type Role string
//	errInvalid := errors.New("invalid role")
//	valid := func(s string) bool { return strings.TrimSpace(s) == "admin" }
//	role, err := constantx.OptionalPtr[Role](" admin ", errInvalid, valid)
//
// Example result:
//   - *role == Role("admin")
//   - err == nil
//
// Blank example result:
//   - OptionalPtr[Role](" ", errInvalid, valid) returns nil, nil.
func OptionalPtr[T ~string](s string, invalid error, valid func(string) bool) (*T, error) {
	if strings.TrimSpace(s) == "" {
		return nil, nil
	}
	if !valid(s) {
		return nil, invalid
	}
	v := T(strings.TrimSpace(s))
	return &v, nil
}

// OptionalValue parses an optional string-backed enum and returns the zero value when absent.
//
// Blank input means "not provided" and returns the zero value of T with nil
// error. Non-blank input must be valid and is returned as the trimmed enum value.
//
// Type Parameters:
//   - T: string-backed enum type.
//
// Parameters:
//   - s: raw optional input.
//   - invalid: error returned when non-blank input is invalid.
//   - valid: validator for the raw string.
//
// Returns:
//   - T: zero value when blank; parsed value when valid.
//   - error: nil when blank/valid; otherwise invalid.
//
// Example:
//
//	type Role string
//	errInvalid := errors.New("invalid role")
//	valid := func(s string) bool { return strings.TrimSpace(s) == "admin" }
//	role, err := constantx.OptionalValue[Role](" admin ", errInvalid, valid)
//
// Example result:
//   - role == Role("admin")
//   - err == nil
//
// Blank example result:
//   - OptionalValue[Role](" ", errInvalid, valid) returns Role(""), nil.
func OptionalValue[T ~string](s string, invalid error, valid func(string) bool) (T, error) {
	if strings.TrimSpace(s) == "" {
		var z T
		return z, nil
	}
	if !valid(s) {
		var z T
		return z, invalid
	}
	return T(strings.TrimSpace(s)), nil
}
