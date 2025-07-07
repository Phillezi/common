package ptr

import "github.com/Phillezi/common/utils/or"

func Of[T any](v T) *T {
	return &v
}

func DerefNonNilOr[T comparable](v *T, alt ...T) T {
	if v == nil {
		return or.Or(alt...)
	}
	return (*v)
}
