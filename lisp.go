package golisp

import "fmt"

var _ = fmt.Printf

type Value interface {
	IsType()
}

type Nil struct{}
type Number int
type Symbol string
type Pair struct {
	a Value
	b Value
}

func (Nil) IsType()    {}
func (Number) IsType() {}
func (Symbol) IsType() {}
func (Pair) IsType()   {}

var NIL = Nil(struct{}{})

/*
 * Basic primitives of lisp are:
 *  quote, atom, eq, car, cdr, cons, cond
 *
 * We implement cond using switch/if in go.
 */

func quote(v Value) Value {
	return v
}

func atom(v Value) bool {
	switch v.(type) {
	case Pair:
		return false
	default:
		return true
	}
}

func eq(a, b Value) bool {
	if a == NIL && b == NIL {
		return true
	}
	if atom(a) && atom(b) && a == b {
		return true
	}
	return false
}

func car(v Value) Value {
	switch p := v.(type) {
	case Pair:
		return p.a
	default:
		panic("car")
	}
}

func cdr(v Value) Value {
	switch p := v.(type) {
	case Pair:
		return p.b
	default:
		panic("car")
	}
}

/* Convenience functions of car/cdr */
func caar(v Value) Value   { return car(car(v)) }
func cadr(v Value) Value   { return car(cdr(v)) }
func cadar(v Value) Value  { return car(cdr(car(v))) }
func caddr(v Value) Value  { return car(cdr(cdr(v))) }
func caddar(v Value) Value { return car(cdr(cdr(car(v)))) }

func cons(a Value, d Value) Pair {
	return Pair{a, d}
}

/* Convenience function for creating lists */
func list(vs ...Value) Value {
	var ls Value = NIL
	for i := len(vs) - 1; i >= 0; i-- {
		ls = cons(vs[i], ls)
	}
	return ls
}

/* This is a recreation of the McCarthy assoc -- panics if k is not found */
func assoc(k, ps Value) Value {
	if eq(caar(ps), k) {
		return cadar(ps)
	} else {
		return assoc(k, cdr(ps))
	}
}
