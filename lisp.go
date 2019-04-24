package golisp

import (
	"fmt"
	"io"
	"strings"
)

var _ = fmt.Printf

type Value interface {
	IsType()
}

//type Number int
type Nil struct{}
type Symbol string
type Number int
type Pair struct {
	a Value
	b Value
}

//func (Number) IsType() {}
func (Nil) IsType()    {}
func (Symbol) IsType() {}
func (Number) IsType() {}
func (Pair) IsType()   {}

var NIL = Nil(struct{}{})
var TRUE = Symbol("t")

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
	if ps == NIL {
		panic(fmt.Sprint("assoc:", k, ps))
	}
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

func add(x, y Value) Value {
	xn, ok := x.(Number)
	if !ok {
		panic(fmt.Sprint("Invalid number in add:", x))
	}
	yn, ok := y.(Number)
	if !ok {
		panic(fmt.Sprint("Invalid number in add:", y))
	}
	return Number(xn + yn)
}
func sub(x, y Value) Value {
	xn, ok := x.(Number)
	if !ok {
		panic(fmt.Sprint("Invalid number in sub:", x))
	}
	yn, ok := y.(Number)
	if !ok {
		panic(fmt.Sprint("Invalid number in sub:", y))
	}
	return Number(xn - yn)
}
func lt(x, y Value) Value {
	xn, ok := x.(Number)
	if !ok {
		panic(fmt.Sprint("Invalid number in lt:", x))
	}
	yn, ok := y.(Number)
	if !ok {
		panic(fmt.Sprint("Invalid number in lt:", y))
	}
	if xn < yn {
		return TRUE
	} else {
		return NIL
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
		switch e.(type) {
		case Number:
			return e
		default:
			return assoc(e, a)
		}
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
		} else if isTrue(eq(car(e), Symbol("add"))) {
			return add(eval(cadr(e), a), eval(caddr(e), a))
		} else if isTrue(eq(car(e), Symbol("sub"))) {
			return sub(eval(cadr(e), a), eval(caddr(e), a))
		} else if isTrue(eq(car(e), Symbol("lt"))) {
			return lt(eval(cadr(e), a), eval(caddr(e), a))
		} else {
			return eval(cons(assoc(car(e), a), cdr(e)), a)
		}
	} else if isTrue(eq(caar(e), Symbol("label"))) {
		return eval(cons(caddar(e), cdr(e)),
			cons(list(cadar(e), car(e)), a))
	} else if isTrue(eq(caar(e), Symbol("lambda"))) {
		return eval(caddar(e), append(pair(cadar(e), evlis(cdr(e), a)), a))
	}
	panic(fmt.Sprint("eval: invalid construct:", e))
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

func readSym(rdr io.ByteScanner) Value {
	var b strings.Builder
	for {
		c, err := rdr.ReadByte()
		//fmt.Println("readSym", c, err)
		if err != nil {
			break
		}
		if c == ' ' || c == '\n' {
			break
		}
		if c == ')' {
			rdr.UnreadByte()
			break
		}
		b.WriteByte(c)
	}
	return Symbol(b.String())
}

func readNum(rdr io.ByteScanner) Value {
	n := 0
	negative := false
	c, err := rdr.ReadByte()
	if err != nil {
		panic("unexpected EOF")
	}
	if c == '-' {
		negative = true
	} else {
		rdr.UnreadByte()
	}
	ndig := 0

	for {
		c, err := rdr.ReadByte()
		//fmt.Println("readNum", c, err)
		if err != nil {
			break
		}
		if c == ' ' || c == '\n' {
			break
		}
		if c == ')' {
			rdr.UnreadByte()
			break
		}
		if !isDig(c) {
			panic(fmt.Sprint("Invalid digit in number", c))
		}
		n *= 10
		n += int(c - 48)
		ndig++
	}
	if ndig == 0 {
		panic("No digits in number")
	}
	if negative {
		return Number(-n)
	} else {
		return Number(n)
	}
}

func isDig(c byte) bool {
	return c >= 48 && c <= 57
}

func readList(rdr io.ByteScanner) Value {
	var r Value = NIL
	for {
		c, err := rdr.ReadByte()
		//fmt.Println("readList", c, err)
		if err != nil {
			panic("readList")
		}
		if c == ')' {
			return r
		}
		if c == ' ' || c == '\n' {
			continue
		}
		rdr.UnreadByte()
		v := read(rdr)
		r = cons(v, r)
	}
}

func reverse(l Value) Value {
	var r Value = NIL
	for l != NIL {
		r = cons(car(l), r)
		l = cdr(l)
	}
	return r
}

func read(rdr io.ByteScanner) Value {
	for {
		c, err := rdr.ReadByte()
		//fmt.Println("read", c, err)
		if err != nil {
			panic("read")
		}
		if c == '(' {
			return reverse(readList(rdr))
		} else if c == ' ' || c == '\n' {
			continue
		} else if c == '\'' {
			return list(Symbol("quote"), read(rdr))
		} else if c == '-' || (c >= '0' && c <= '9') {
			rdr.UnreadByte()
			return readNum(rdr)
		} else {
			rdr.UnreadByte()
			return readSym(rdr)
		}
	}
}

func readStr(s string) Value {
	return read(strings.NewReader(s))
}
