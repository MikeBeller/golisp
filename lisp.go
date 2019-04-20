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
var TRUE = Symbol("t")

func S(s string) Value {
	return Symbol(s)
}

/*
 * Basic primitives of lisp are:
 *  quote, atom, eq, car, cdr, cons, cond
 *
 * We implement cond using switch/if in go.
 */

func quote(v Value) Value {
	return v
}

func atom(v Value) Value {
	switch v.(type) {
	case Pair:
		return NIL
	default:
		return TRUE
	}
}

func isTrue(v Value) bool {
	return v != NIL
}

func eq(a, b Value) Value {
	if a == NIL && b == NIL {
		return TRUE
	}
	if isTrue(atom(a)) && isTrue(atom(b)) && a == b {
		return TRUE
	}
	return NIL
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
		panic("cdr")
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
	if isTrue(eq(caar(ps), k)) {
		return cadar(ps)
	} else {
		return assoc(k, cdr(ps))
	}
}

/* Assorted helper functions */
func null(x Value) Value {
	return eq(x, NIL)
}

func and(x, y Value) Value {
	if isTrue(x) {
		if isTrue(y) {
			return TRUE
		} else {
			return NIL
		}
	} else {
		return NIL
	}
}

func not(x Value) Value {
	if isTrue(x) {
		return NIL
	} else {
		return TRUE
	}
}

func append(x, y Value) Value {
	if isTrue(null(x)) {
		return y
	} else {
		return cons(car(x), append(cdr(x), y))
	}
}

func pair(x, y Value) Value {
	if isTrue(and(null(x), null(y))) {
		return NIL
	} else if isTrue(and(not(atom(x)), not(atom(y)))) {
		return cons(list(car(x), car(y)), pair(cdr(x), cdr(y)))
	} else {
		return NIL
	}
}

/* Eval -- the core lisp interpretation function */
func eval(e, a Value) Value {
	if isTrue(atom(e)) {
		return assoc(e, a)
	} else if isTrue(atom(car(e))) {
		if isTrue(eq(car(e), Symbol("quote"))) {
			return cadr(e)
		} else if isTrue(eq(car(e), Symbol("atom"))) {
			return atom(eval(cadr(e), a))
		} else if isTrue(eq(car(e), Symbol("eq"))) {
			return eq(eval(cadr(e), a), eval(caddr(e), a))
		} else if isTrue(eq(car(e), Symbol("car"))) {
			return car(eval(cadr(e), a))
		} else if isTrue(eq(car(e), Symbol("cdr"))) {
			return cdr(eval(cadr(e), a))
		} else if isTrue(eq(car(e), Symbol("cons"))) {
			return cons(eval(cadr(e), a), eval(caddr(e), a))
		} else if isTrue(eq(car(e), Symbol("cond"))) {
			return evcon(cdr(e), a)
		} else {
			return eval(cons(assoc(car(e), a), cdr(e)), a)
		}
	} else if isTrue(eq(caar(e), Symbol("label"))) {
		return eval(cons(caddar(e), cdr(e)),
			cons(list(cadar(e), car(e)), a))
	} else if isTrue(eq(caar(e), Symbol("lambda"))) {
		return eval(caddar(e), append(pair(cadar(e), evlis(cdr(e), a)), a))
	}
	panic("wtf?")
}

func evcon(c, a Value) Value {
	if isTrue(not(null(eval(caar(c), a)))) {
		return eval(cadar(c), a)
	} else {
		return evcon(cdr(c), a)
	}
}

func evlis(m, a Value) Value {
	if isTrue(null(m)) {
		return NIL
	} else {
		return cons(eval(car(m), a), evlis(cdr(m), a))
	}
}
