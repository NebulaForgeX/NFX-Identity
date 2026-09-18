package constantx

// ValidAnyMember reports whether any value in candidates is present in allowed.
//
// Type Parameters:
//   - T: comparable value type.
//
// Parameters:
//   - candidates: values to test.
//   - allowed: whitelist of accepted values.
//
// Returns:
//   - bool: true when at least one candidate is in allowed; otherwise false.
//
// Example:
//
//	ok := constantx.ValidAnyMember([]string{"guest", "admin"}, "admin", "member")
//
// Example result:
//   - ok == true
func ValidAnyMember[T comparable](candidates []T, allowed ...T) bool {
	for _, v := range candidates {
		if ValidMember(v, allowed...) {
			return true
		}
	}
	return false
}
