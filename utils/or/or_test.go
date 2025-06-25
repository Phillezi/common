package or

import (
	"testing"
)

func TestOr(t *testing.T) {
	testOr(t, "strings", []string{"", "hello", "world"}, "hello")
	testOr(t, "ints", []int{0, 0, 5}, 5)
	testOr(t, "bools", []bool{false, true}, true)

	type Person struct {
		Name string
		Age  int
	}
	zero := Person{}
	alice := Person{"Alice", 30}
	bob := Person{"Bob", 25}
	testOr(t, "structs", []Person{zero, alice, bob}, alice)
	testOr(t, "structs all zero", []Person{zero, zero}, zero)
}

func TestCall(t *testing.T) {
	testCall(t, "strings", []func() string{
		func() string { return "" },
		func() string { return "hello" },
	}, "hello")

	testCall(t, "ints", []func() int{
		func() int { return 0 },
		func() int { return 42 },
	}, 42)

	testCall(t, "bools", []func() bool{
		func() bool { return false },
		func() bool { return true },
	}, true)

	type Person struct {
		Name string
		Age  int
	}
	zero := Person{}
	alice := Person{"Alice", 30}
	bob := Person{"Bob", 25}

	testCall(t, "structs", []func() Person{
		func() Person { return zero },
		func() Person { return alice },
		func() Person { return bob },
	}, alice)

	testCall(t, "structs all zero", []func() Person{
		func() Person { return zero },
		func() Person { return zero },
	}, zero)
}

func testOr[T comparable](t *testing.T, name string, input []T, want T) {
	t.Run(name, func(t *testing.T) {
		got := Or(input...)
		if got != want {
			t.Errorf("Or(%v) = %v, want %v", input, got, want)
		}
	})
}

func testCall[T comparable](t *testing.T, name string, funcs []func() T, want T) {
	t.Run(name, func(t *testing.T) {
		got := Call(funcs...)
		if got != want {
			t.Errorf("Call(...) = %v, want %v", got, want)
		}
	})
}
