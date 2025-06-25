package or

// Iterate through all provided
// values, return the first one
// that isnt a "zero" value.
// For example with strings, it
// is the first one that isnt
// empty
func Or[T comparable](vals ...T) T {
	var zero T
	for _, val := range vals {
		if val != zero {
			return val
		}
	}
	// all are zero values
	// fallback to a zero value
	return zero
}

// Iterate through all provided
// getters, return the first one
// that doesnt return a "zero" value.
func Call[T comparable](getVals ...func() T) T {
	var zero T
	for _, getVal := range getVals {
		var val = getVal()
		if val != zero {
			return val
		}
	}
	// all functions return zero values
	// fallback to a zero value
	return zero
}
