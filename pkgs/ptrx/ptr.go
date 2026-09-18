package ptrx

import (
	"reflect"

	"github.com/google/uuid"
)

// Ptr returns a pointer to v.
//
// It is useful when building struct literals with optional pointer fields and
// you want to create the pointer inline.
//
// Type Parameters:
//   - T: value type.
//
// Parameters:
//   - v: value to copy into a new pointer.
//
// Returns:
//   - *T: pointer to v.
//
// Example:
//
//	age := ptrx.Ptr(30)
//
// Example result:
//   - age != nil
//   - *age == 30
//
// Struct literal example:
//
//	user := User{Name: ptrx.Ptr("Alice")}
//
// Struct literal example result:
//   - user.Name != nil
//   - *user.Name == "Alice"
func Ptr[T any](v T) *T {
	return &v
}

// To returns a pointer to v.
//
// To is an alias for Ptr with shorter naming. Use it when inline pointer
// construction reads better in struct literals or option values.
//
// Type Parameters:
//   - T: value type.
//
// Parameters:
//   - v: value to copy into a new pointer.
//
// Returns:
//   - *T: pointer to v.
//
// Example:
//
//	name := ptrx.To("Alice")
//
// Example result:
//   - name != nil
//   - *name == "Alice"
func To[T any](v T) *T {
	return &v
}

// Deref safely dereferences p and returns T's zero value when p is nil.
//
// Type Parameters:
//   - T: pointed value type.
//
// Parameters:
//   - p: pointer to dereference.
//
// Returns:
//   - T: *p when p is non-nil; otherwise the zero value of T.
//
// Example:
//
//	value := ptrx.Deref(ptrx.Ptr(42))
//
// Example result:
//   - value == 42
//
// Nil example:
//
//	var p *string
//	value := ptrx.Deref(p)
//
// Nil example result:
//   - value == ""
func Deref[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}

// DerefWithDefault safely dereferences p and returns def when p is nil.
//
// Type Parameters:
//   - T: pointed value type.
//
// Parameters:
//   - p: pointer to dereference.
//   - def: default value returned when p is nil.
//
// Returns:
//   - T: *p when p is non-nil; otherwise def.
//
// Example:
//
//	value := ptrx.DerefWithDefault(ptrx.Ptr("hello"), "default")
//
// Example result:
//   - value == "hello"
//
// Nil example:
//
//	var p *string
//	value := ptrx.DerefWithDefault(p, "default")
//
// Nil example result:
//   - value == "default"
func DerefWithDefault[T any](p *T, def T) T {
	if p == nil {
		return def
	}
	return *p
}

// PtrIfNotZero returns a pointer to v only when v is not T's zero value.
//
// Type Parameters:
//   - T: comparable value type.
//
// Parameters:
//   - v: value to optionally point to.
//
// Returns:
//   - *T: pointer to v when v is non-zero; otherwise nil.
//
// Example:
//
//	name := ptrx.PtrIfNotZero("Alice")
//
// Example result:
//   - name != nil
//   - *name == "Alice"
//
// Zero-value examples:
//   - PtrIfNotZero("") returns nil.
//   - PtrIfNotZero(0) returns nil.
//   - PtrIfNotZero(false) returns nil.
func PtrIfNotZero[T comparable](v T) *T {
	var zero T
	if v == zero {
		return nil
	}
	return &v
}

// Map transforms the value inside xs and returns R's zero value when xs is nil.
//
// Type Parameters:
//   - T: input pointed value type.
//   - R: output value type.
//
// Parameters:
//   - xs: input pointer.
//   - fn: transformation function that receives the dereferenced value.
//
// Returns:
//   - R: fn(*xs) when xs is non-nil; otherwise the zero value of R.
//
// Example:
//
//	result := ptrx.Map(ptrx.Ptr(42), func(n int) string {
//		return fmt.Sprint(n)
//	})
//
// Example result:
//   - result == "42"
//
// Nil example:
//
//	var p *int
//	result := ptrx.Map(p, func(n int) string { return fmt.Sprint(n) })
//
// Nil example result:
//   - result == ""
func Map[T any, R any](xs *T, fn func(T) R) R {
	if xs == nil {
		var zero R
		return zero
	}
	return fn(*xs)
}

// MapP transforms xs with a function that receives the original pointer.
//
// Unlike Map, fn receives *T directly, so it can read methods/fields that need a
// pointer receiver. When xs is nil, fn is not called and R's zero value is returned.
//
// Type Parameters:
//   - T: input pointed value type.
//   - R: output value type.
//
// Parameters:
//   - xs: input pointer.
//   - fn: transformation function that receives xs.
//
// Returns:
//   - R: fn(xs) when xs is non-nil; otherwise the zero value of R.
//
// Example:
//
//	type User struct{ Name string }
//	name := ptrx.MapP(&User{Name: "Alice"}, func(u *User) string {
//		return u.Name
//	})
//
// Example result:
//   - name == "Alice"
func MapP[T any, R any](xs *T, fn func(*T) R) R {
	if xs == nil {
		var zero R
		return zero
	}
	return fn(xs)
}

// MapPP transforms xs with a pointer-in/pointer-out function and dereferences the result.
//
// When xs is nil, fn is not called and R's zero value is returned. When xs is
// non-nil, fn must return a non-nil *R; otherwise this function will panic while
// dereferencing the returned pointer.
//
// Type Parameters:
//   - T: input pointed value type.
//   - R: output pointed value type.
//
// Parameters:
//   - xs: input pointer.
//   - fn: transformation function that receives *T and returns *R.
//
// Returns:
//   - R: *fn(xs) when xs is non-nil; otherwise the zero value of R.
//
// Example:
//
//	result := ptrx.MapPP(ptrx.Ptr(42), func(n *int) *string {
//		s := fmt.Sprint(*n)
//		return &s
//	})
//
// Example result:
//   - result == "42"
func MapPP[T any, R any](xs *T, fn func(*T) *R) R {
	if xs == nil {
		var zero R
		return zero
	}
	return *fn(xs)
}

// MapPtr transforms the value inside xs and returns a pointer to the transformed result.
//
// When xs is nil, fn is not called and nil is returned.
//
// Type Parameters:
//   - T: input pointed value type.
//   - R: output value type.
//
// Parameters:
//   - xs: input pointer.
//   - fn: transformation function that receives the dereferenced value.
//
// Returns:
//   - *R: pointer to fn(*xs) when xs is non-nil; otherwise nil.
//
// Example:
//
//	result := ptrx.MapPtr(ptrx.Ptr(42), func(n int) string {
//		return fmt.Sprint(n)
//	})
//
// Example result:
//   - result != nil
//   - *result == "42"
func MapPtr[T any, R any](xs *T, fn func(T) R) *R {
	if xs == nil {
		return nil
	}
	r := fn(*xs)
	return &r
}

// MapPtrP transforms xs with a pointer function and returns a pointer to the result.
//
// When xs is nil, fn is not called and nil is returned.
//
// Type Parameters:
//   - T: input pointed value type.
//   - R: output value type.
//
// Parameters:
//   - xs: input pointer.
//   - fn: transformation function that receives xs.
//
// Returns:
//   - *R: pointer to fn(xs) when xs is non-nil; otherwise nil.
//
// Example:
//
//	type User struct{ Name string }
//	name := ptrx.MapPtrP(&User{Name: "Alice"}, func(u *User) string {
//		return u.Name
//	})
//
// Example result:
//   - name != nil
//   - *name == "Alice"
func MapPtrP[T any, R any](xs *T, fn func(*T) R) *R {
	if xs == nil {
		return nil
	}
	r := fn(xs)
	return &r
}

// MapPtrPP transforms xs with a pointer-in/pointer-out function.
//
// When xs is nil, fn is not called and nil is returned. When xs is non-nil, the
// pointer returned by fn is returned directly, including nil if fn returns nil.
//
// Type Parameters:
//   - T: input pointed value type.
//   - R: output pointed value type.
//
// Parameters:
//   - xs: input pointer.
//   - fn: transformation function that receives *T and returns *R.
//
// Returns:
//   - *R: fn(xs) when xs is non-nil; otherwise nil.
//
// Example:
//
//	result := ptrx.MapPtrPP(ptrx.Ptr(42), func(n *int) *string {
//		s := fmt.Sprint(*n)
//		return &s
//	})
//
// Example result:
//   - result != nil
//   - *result == "42"
func MapPtrPP[T any, R any](xs *T, fn func(*T) *R) *R {
	if xs == nil {
		return nil
	}
	return fn(xs)
}

// IsNil reports whether x is nil, including typed nil values stored in interfaces.
//
// It handles nil-capable kinds such as channels, functions, interfaces, maps,
// pointers, and slices. Non-nil-capable values such as int, string, and structs
// return false.
//
// Parameters:
//   - x: value to check.
//
// Returns:
//   - bool: true when x is nil or a typed nil; otherwise false.
//
// Example:
//
//	var p *int
//	ok := ptrx.IsNil(p)
//
// Example result:
//   - ok == true
//
// More examples:
//   - IsNil(map[string]int(nil)) returns true.
//   - IsNil([]int(nil)) returns true.
//   - IsNil(42) returns false.
//   - IsNil(ptrx.Ptr(42)) returns false.
func IsNil(x any) bool {
	if x == nil {
		return true
	}
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

// UuidToStrPtr converts a UUID pointer into a string pointer.
//
// Nil UUID pointers return nil. Non-nil UUIDs are converted with UUID.String()
// and then passed through PtrIfNotZero.
//
// Parameters:
//   - s: UUID pointer to convert.
//
// Returns:
//   - *string: pointer to the UUID string when s is non-nil; otherwise nil.
//
// Example:
//
//	id := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
//	value := ptrx.UuidToStrPtr(&id)
//
// Example result:
//   - value != nil
//   - *value == "550e8400-e29b-41d4-a716-446655440000"
//
// Nil example:
//
//	var id *uuid.UUID
//	value := ptrx.UuidToStrPtr(id)
//
// Nil example result:
//   - value == nil
func UuidToStrPtr(s *uuid.UUID) *string {
	if s == nil {
		return nil
	}
	return PtrIfNotZero(s.String())
}
