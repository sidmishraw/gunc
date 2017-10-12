/**
* gunc.go
* @author Sidharth Mishra
* @description A catalogue of various higher order functions to make functional programming easy with golang
* @created Wed Oct 11 2017 20:16:21 GMT-0700 (PDT)
* @copyright 2017 Sidharth Mishra
* @last-modified Wed Oct 11 2017 20:16:32 GMT-0700 (PDT)
 */

package gunc

// The A is a generic type
type A interface{}

// The B is a generic type
type B interface{}

// The C is a generic type
type C interface{}

// The D is a generic type
type D interface{}

// Head gives the head of the slice
func Head(xs []A) A {
	return xs[0]
}

// Tail gives the tail of the slice
func Tail(xs []A) []A {
	return xs[1:]
}

// Map maps a function f onto a slice xs
func Map(f func(A) B, xs []A) []B {
	if len(xs) == 0 {
		return make([]B, 0)
	}
	return append([]B{f(Head(xs))}, Map(f, Tail(xs))...)
}

// Foldl is the left associative folding operation on a slice xs
func Foldl(f func(A, B) B, acc B, xs []A) B {
	if len(xs) == 0 {
		return acc
	}
	return Foldl(f, f(Head(xs), acc), Tail(xs))
}

// Foldr is the right associative folding operation on a slice xs
func Foldr(f func(A, B) B, acc B, xs []A) B {
	if len(xs) == 0 {
		return acc
	}
	return f(Head(xs), Foldr(f, acc, Tail(xs)))
}

// Curry curries the functions
// currying like in Haskell
// curry(add,1)(2) = 3
func Curry(f func(A, B) C, a A) func(B) C {
	return func(b B) C {
		return f(a, b)
	}
}
