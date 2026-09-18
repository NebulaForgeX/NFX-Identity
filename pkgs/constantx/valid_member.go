package constantx

import "slices"

// ValidMember reports whether v equals one of the allowed values.
//
// It is useful for enum-like constants, especially string-backed enums. The
// comparison uses Go's == operator, so T must be comparable.
//
// Type Parameters:
//   - T: comparable value type.
//
// Parameters:
//   - v: value to validate.
//   - allowed: whitelist of accepted values.
//
// Returns:
//   - bool: true when v is present in allowed; otherwise false.
//
// Example:
//
//	ok := constantx.ValidMember("admin", "admin", "member")
//
// Example result:
//   - ok == true
//
// Invalid example:
//
//	ok := constantx.ValidMember("guest", "admin", "member")
//
// Invalid example result:
//   - ok == false
func ValidMember[T comparable](v T, allowed ...T) bool {
	return slices.Contains(allowed, v)
}
