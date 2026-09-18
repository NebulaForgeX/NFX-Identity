package slicex

// Deduplicate removes duplicate comparable values while preserving first-seen order.
//
// For slices with length 0 or 1, the original slice is returned directly.
// For longer slices, a new output slice is built.
//
// Type Parameters:
//   - T: comparable element type.
//
// Parameters:
//   - items: input slice that may contain duplicates.
//
// Returns:
//   - []T: values with duplicates removed, keeping the first occurrence.
//
// Example:
//
//	unique := slicex.Deduplicate([]string{"a", "b", "a", "c", "b"})
//
// Example result:
//   - unique == []string{"a", "b", "c"}
func Deduplicate[T comparable](items []T) []T {
	if len(items) <= 1 {
		return items
	}

	seen := make(map[T]struct{}, len(items))
	out := make([]T, 0, len(items))

	for _, item := range items {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		out = append(out, item)
	}

	return out
}

// DeduplicateBy removes duplicate values using a key function while preserving first-seen order.
//
// The first item for each key is kept. Later items with the same key are skipped.
// For slices with length 0 or 1, the original slice is returned directly.
//
// Type Parameters:
//   - T: input element type.
//   - K: comparable key type used to decide uniqueness.
//
// Parameters:
//   - items: input slice that may contain duplicate keys.
//   - keyFn: function that extracts the uniqueness key from each item.
//
// Returns:
//   - []T: values with duplicate keys removed, keeping the first item for each key.
//
// Example:
//
//	type User struct {
//		ID   string
//		Name string
//	}
//	users := []User{
//		{ID: "1", Name: "Alice"},
//		{ID: "2", Name: "Bob"},
//		{ID: "1", Name: "Alice2"},
//	}
//	unique := slicex.DeduplicateBy(users, func(u User) string { return u.ID })
//
// Example result:
//   - unique == []User{{ID:"1", Name:"Alice"}, {ID:"2", Name:"Bob"}}.
func DeduplicateBy[T any, K comparable](items []T, keyFn func(T) K) []T {
	if len(items) <= 1 {
		return items
	}

	seen := make(map[K]struct{}, len(items))
	out := make([]T, 0, len(items))

	for _, item := range items {
		key := keyFn(item)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, item)
	}

	return out
}
