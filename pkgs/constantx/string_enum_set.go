package constantx

import "strings"

// StringEnumSet 由允许取值列表构造，提供 Valid / ValidString / Parse* / Optional*（单一配置入口）。
type StringEnumSet[T ~string] struct {
	allowed []T
}

// NewStringEnumSet creates a reusable validator/parser for a string-backed enum.
//
// The allowed values are copied into the set, so later changes to the caller's
// slice do not affect validation.
//
// Type Parameters:
//   - T: string-backed enum type.
//
// Parameters:
//   - allowed: all valid enum values.
//
// Returns:
//   - *StringEnumSet[T]: enum helper with Valid, Parse, and Optional methods.
//
// Example:
//
//	type Role string
//	roles := constantx.NewStringEnumSet[Role]("admin", "member")
//
// Example result:
//   - roles.Valid("admin") == true.
//   - roles.Valid("guest") == false.
func NewStringEnumSet[T ~string](allowed ...T) *StringEnumSet[T] {
	return &StringEnumSet[T]{allowed: append([]T(nil), allowed...)}
}

// Valid reports whether v is one of the set's allowed enum values.
//
// Type Parameters:
//   - T: string-backed enum type.
//
// Parameters:
//   - v: enum value to check.
//
// Returns:
//   - bool: true when v is allowed; otherwise false.
//
// Example:
//
//	type Role string
//	roles := constantx.NewStringEnumSet[Role]("admin", "member")
//	ok := roles.Valid(Role("admin"))
//
// Example result:
//   - ok == true
func (s *StringEnumSet[T]) Valid(v T) bool {
	return ValidMember(v, s.allowed...)
}

// ValidString trims and validates a raw string against the set.
//
// Type Parameters:
//   - T: string-backed enum type.
//
// Parameters:
//   - str: raw string input.
//
// Returns:
//   - bool: true when strings.TrimSpace(str) is an allowed value.
//
// Example:
//
//	type Role string
//	roles := constantx.NewStringEnumSet[Role]("admin", "member")
//	ok := roles.ValidString(" admin ")
//
// Example result:
//   - ok == true
func (s *StringEnumSet[T]) ValidString(str string) bool {
	return ValidString(str, s.Valid)
}

// ParseString validates and converts a raw string into an enum value from the set.
//
// Type Parameters:
//   - T: string-backed enum type.
//
// Parameters:
//   - str: raw string input.
//   - invalid: error returned when str is not allowed.
//
// Returns:
//   - T: trimmed enum value when valid.
//   - error: nil when valid; otherwise invalid.
//
// Example:
//
//	type Role string
//	roles := constantx.NewStringEnumSet[Role]("admin", "member")
//	role, err := roles.ParseString(" member ", errors.New("invalid role"))
//
// Example result:
//   - role == Role("member")
//   - err == nil
func (s *StringEnumSet[T]) ParseString(str string, invalid error) (T, error) {
	return ParseString[T](str, invalid, s.ValidString)
}

// ParseStringsStrict validates multiple raw strings while skipping blank values.
//
// Type Parameters:
//   - T: string-backed enum type.
//
// Parameters:
//   - ss: raw string values, such as repeated query params.
//   - invalid: error returned when any non-blank value is not allowed.
//
// Returns:
//   - []T: parsed values, excluding blanks.
//   - error: nil when all non-blank values are valid; otherwise invalid.
//
// Example:
//
//	type Role string
//	roles := constantx.NewStringEnumSet[Role]("admin", "member")
//	values, err := roles.ParseStringsStrict([]string{"admin", "", " member "}, errors.New("invalid role"))
//
// Example result:
//   - values == []Role{"admin", "member"}
//   - err == nil
func (s *StringEnumSet[T]) ParseStringsStrict(ss []string, invalid error) ([]T, error) {
	return ParseStringsStrict[T](ss, invalid, s.ValidString)
}

// ParseStringSlice validates every raw string without skipping blank values.
//
// Use this for JSON arrays or strict request bodies where "" should be rejected
// unless it is explicitly included in the enum set.
//
// Type Parameters:
//   - T: string-backed enum type.
//
// Parameters:
//   - ss: raw string values.
//   - invalid: error returned when any value is not allowed.
//
// Returns:
//   - []T: parsed values.
//   - error: nil when all values are valid; otherwise invalid.
//
// Example:
//
//	type Role string
//	roles := constantx.NewStringEnumSet[Role]("admin", "member")
//	values, err := roles.ParseStringSlice([]string{"admin", " member "}, errors.New("invalid role"))
//
// Example result:
//   - values == []Role{"admin", "member"}
//   - err == nil
func (s *StringEnumSet[T]) ParseStringSlice(ss []string, invalid error) ([]T, error) {
	return ParseStringSlice[T](ss, invalid, s.ValidString)
}

// OptionalPtr parses an optional raw string into an enum pointer.
//
// Blank input returns nil, nil. Non-blank input must be allowed by the set.
//
// Type Parameters:
//   - T: string-backed enum type.
//
// Parameters:
//   - str: raw optional string.
//   - invalid: error returned when non-blank input is not allowed.
//
// Returns:
//   - *T: nil when blank; pointer to parsed value when valid.
//   - error: nil when blank/valid; otherwise invalid.
//
// Example:
//
//	type Role string
//	roles := constantx.NewStringEnumSet[Role]("admin", "member")
//	role, err := roles.OptionalPtr(" admin ", errors.New("invalid role"))
//
// Example result:
//   - *role == Role("admin")
//   - err == nil
//
// Blank example result:
//   - roles.OptionalPtr(" ", errInvalid) returns nil, nil.
func (s *StringEnumSet[T]) OptionalPtr(str string, invalid error) (*T, error) {
	return OptionalPtr[T](str, invalid, s.ValidString)
}

// OptionalValue parses an optional raw string into an enum value.
//
// Blank input returns the zero value of T and nil error. Non-blank input must be
// allowed by the set.
//
// Type Parameters:
//   - T: string-backed enum type.
//
// Parameters:
//   - str: raw optional string.
//   - invalid: error returned when non-blank input is not allowed.
//
// Returns:
//   - T: zero value when blank; parsed value when valid.
//   - error: nil when blank/valid; otherwise invalid.
//
// Example:
//
//	type Role string
//	roles := constantx.NewStringEnumSet[Role]("admin", "member")
//	role, err := roles.OptionalValue(" member ", errors.New("invalid role"))
//
// Example result:
//   - role == Role("member")
//   - err == nil
//
// Blank example result:
//   - roles.OptionalValue(" ", errInvalid) returns Role(""), nil.
func (s *StringEnumSet[T]) OptionalValue(str string, invalid error) (T, error) {
	return OptionalValue[T](str, invalid, s.ValidString)
}

// Array returns a copy of the set's allowed values in their declared order.
//
// The result is a fresh slice, so callers may freely range over or mutate it
// without affecting the set.
//
// Type Parameters:
//   - T: string-backed enum type.
//
// Returns:
//   - []T: allowed enum values in declaration order.
//
// Example:
//
//	type Role string
//	roles := constantx.NewStringEnumSet[Role]("admin", "member")
//	all := roles.Array()
//
// Example result:
//   - all == []Role{"admin", "member"}
func (s *StringEnumSet[T]) Array() []T {
	return append([]T(nil), s.allowed...)
}

// Parse returns a valid enum value, or fallback when v is not allowed.
//
// Input is trimmed before validation to align with ParseString behavior.
func (s *StringEnumSet[T]) Parse(v T, fallback T) T {
	trimmed := T(strings.TrimSpace(string(v)))
	if s.Valid(trimmed) {
		return trimmed
	}
	return fallback
}
