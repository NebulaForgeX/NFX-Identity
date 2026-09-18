package slicex

// Map transforms each element of xs with fn.
//
// Empty or nil input returns a non-nil empty []R. The transformation function is
// not called when len(xs) == 0.
//
// Type Parameters:
//   - T: input element type.
//   - R: output element type.
//
// Parameters:
//   - xs: input slice.
//   - fn: transformation function.
//
// Returns:
//   - []R: transformed values in the same order as xs.
//
// Example:
//
//	values := slicex.Map([]int{1, 2, 3}, func(n int) int {
//		return n * 2
//	})
//
// Example result:
//   - values == []int{2, 4, 6}
func Map[T any, R any](xs []T, fn func(T) R) []R {
	if len(xs) == 0 {
		return make([]R, 0)
	}
	out := make([]R, len(xs))
	for i := range xs {
		out[i] = fn(xs[i])
	}
	return out
}

// MapP transforms each element of xs with a function that receives a pointer to the element.
//
// fn receives &xs[i], so it can call pointer receiver methods or avoid copying
// large values. Empty or nil input returns a non-nil empty []R.
//
// Type Parameters:
//   - T: input element type.
//   - R: output element type.
//
// Parameters:
//   - xs: input slice.
//   - fn: transformation function that receives a pointer to each element.
//
// Returns:
//   - []R: transformed values in the same order as xs.
//
// Example:
//
//	type User struct{ Name string }
//	names := slicex.MapP([]User{{Name: "Alice"}, {Name: "Bob"}}, func(u *User) string {
//		return u.Name
//	})
//
// Example result:
//   - names == []string{"Alice", "Bob"}
func MapP[T any, R any](xs []T, fn func(*T) R) []R {
	if len(xs) == 0 {
		return make([]R, 0)
	}
	out := make([]R, len(xs))
	for i := range xs {
		out[i] = fn(&xs[i])
	}
	return out
}

// MapPP transforms each element with a pointer-in/pointer-out function and stores dereferenced results.
//
// Empty or nil input returns a non-nil empty []R. For non-empty input, fn must
// return a non-nil *R for every element; otherwise this function panics while
// dereferencing.
//
// Type Parameters:
//   - T: input element type.
//   - R: output element type.
//
// Parameters:
//   - xs: input slice.
//   - fn: transformation function that receives *T and returns *R.
//
// Returns:
//   - []R: dereferenced transformed values.
//
// Example:
//
//	type User struct{ ID int }
//	type UserDTO struct{ ID int }
//	dtos := slicex.MapPP([]User{{ID: 1}, {ID: 2}}, func(u *User) *UserDTO {
//		return &UserDTO{ID: u.ID}
//	})
//
// Example result:
//   - dtos == []UserDTO{{ID: 1}, {ID: 2}}
func MapPP[T any, R any](xs []T, fn func(*T) *R) []R {
	if len(xs) == 0 {
		return make([]R, 0)
	}
	out := make([]R, len(xs))
	for i := range xs {
		out[i] = *fn(&xs[i])
	}
	return out
}

// MapPtr transforms each element and returns pointers to the transformed values.
//
// Empty or nil input returns a non-nil empty []*R. Each returned pointer points to
// a distinct transformed value.
//
// Type Parameters:
//   - T: input element type.
//   - R: output element type.
//
// Parameters:
//   - xs: input slice.
//   - fn: transformation function.
//
// Returns:
//   - []*R: pointers to transformed values.
//
// Example:
//
//	values := slicex.MapPtr([]int{1, 2}, func(n int) int {
//		return n * 10
//	})
//
// Example result:
//   - *values[0] == 10.
//   - *values[1] == 20.
func MapPtr[T any, R any](xs []T, fn func(T) R) []*R {
	if len(xs) == 0 {
		return make([]*R, 0)
	}
	out := make([]*R, len(xs))
	for i := range xs {
		result := fn(xs[i])
		out[i] = &result
	}
	return out
}

// MapPtrP transforms each element with a pointer function and returns pointers to the results.
//
// Empty or nil input returns a non-nil empty []*R. fn receives &xs[i].
//
// Type Parameters:
//   - T: input element type.
//   - R: output element type.
//
// Parameters:
//   - xs: input slice.
//   - fn: transformation function that receives a pointer to each element.
//
// Returns:
//   - []*R: pointers to transformed values.
//
// Example:
//
//	type User struct {
//		Name string
//		Age  int
//	}
//	ages := slicex.MapPtrP([]User{{Name: "Alice", Age: 30}}, func(u *User) int {
//		return u.Age
//	})
//
// Example result:
//   - len(ages) == 1.
//   - *ages[0] == 30.
func MapPtrP[T any, R any](xs []T, fn func(*T) R) []*R {
	if len(xs) == 0 {
		return make([]*R, 0)
	}
	out := make([]*R, len(xs))
	for i := range xs {
		result := fn(&xs[i])
		out[i] = &result
	}
	return out
}

// MapPtrPP transforms each element with a pointer-in/pointer-out function.
//
// Empty or nil input returns a non-nil empty []*R. For non-empty input, each
// returned pointer is exactly the pointer returned by fn, including nil if fn
// returns nil for that element.
//
// Type Parameters:
//   - T: input element type.
//   - R: output element type.
//
// Parameters:
//   - xs: input slice.
//   - fn: transformation function that receives *T and returns *R.
//
// Returns:
//   - []*R: pointers returned by fn for each element.
//
// Example:
//
//	type User struct{ ID int }
//	type UserDTO struct{ ID int }
//	dtos := slicex.MapPtrPP([]User{{ID: 1}}, func(u *User) *UserDTO {
//		return &UserDTO{ID: u.ID}
//	})
//
// Example result:
//   - len(dtos) == 1.
//   - dtos[0].ID == 1.
func MapPtrPP[T any, R any](xs []T, fn func(*T) *R) []*R {
	if len(xs) == 0 {
		return make([]*R, 0)
	}
	out := make([]*R, len(xs))
	for i := range xs {
		out[i] = fn(&xs[i])
	}
	return out
}
