package patch

// ApplyVal writes f into target when f is set to a non-null value.
// Unset and SetNull are no-ops (target unchanged).
func ApplyVal[T any](target *T, f PatchField[T]) {
	if !f.IsSet() || f.IsNull() {
		return
	}
	if v, ok := f.Value(); ok {
		*target = v
	}
}

// ApplyPtr writes f into *target when f is set.
// Unset is a no-op; SetNull clears *target to nil; Set assigns a new pointer.
func ApplyPtr[T any](target **T, f PatchField[T]) {
	ApplyPtrWith(target, f, nil)
}

// ApplyPtrWith is ApplyPtr with an optional value transform (e.g. strings.TrimSpace).
func ApplyPtrWith[T any](target **T, f PatchField[T], mapFn func(T) T) {
	if !f.IsSet() {
		return
	}
	if f.IsNull() {
		*target = nil
		return
	}
	if v, ok := f.Value(); ok {
		if mapFn != nil {
			v = mapFn(v)
		}
		*target = &v
	}
}
