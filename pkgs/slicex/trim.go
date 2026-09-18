package slicex

import "sort"

// TrimByMax keeps up to max items with the lowest order according to less.
// When less(a, b) is true, a is ranked before b. Ties are resolved by less itself.
func TrimByMax[T any](items []T, max int, less func(a, b T) bool) []T {
	if max <= 0 || len(items) <= max {
		return items
	}

	sorted := append([]T(nil), items...)
	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted[:max]
}
