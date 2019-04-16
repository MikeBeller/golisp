package golisp

import (
	"fmt"
	"testing"
)

var _ = fmt.Println

var v1 Value = Number(1)
var v2 Value = Number(2)
var v3 Value = Number(3)
var v4 Value = Number(4)
var v5 Value = Number(5)
var v6 Value = Number(6)

func TestQuote(t *testing.T) {
	for _, v := range []Value{NIL, Number(3), Symbol("FOO")} {
		if quote(v) != v {
			t.Error("quote:", v)
		}
	}
	if quote(Number(3)) != Number(3) {
		t.Error("quote outright")
	}
}

func TestAtom(t *testing.T) {
	for _, p := range []struct {
		Value
		bool
	}{{NIL, true},
		{Number(3), true},
		{Symbol("FOO"), true},
		{Pair{NIL, NIL}, false}} {
		if atom(p.Value) != p.bool {
			t.Error("atom:", p.Value)
		}
	}
}

func TestEq(t *testing.T) {
	for _, p := range []struct {
		a, b Value
		bool
	}{{NIL, NIL, true},
		{NIL, Number(3), false},
		{Number(3), Number(3), true},
		{Number(3), Number(4), false},
		{Symbol("FOO"), Symbol("FOO"), true},
		{Pair{NIL, NIL}, Pair{NIL, NIL}, false},
		{Pair{NIL, NIL}, Symbol("FOO"), false},
	} {
		if eq(p.a, p.b) != p.bool {
			t.Error("eq failed:", p.a, p.b)
		}
	}
}

func TestCons(t *testing.T) {
	if (Pair{v1, v2}) != cons(v1, v2) {
		t.Error("cons1")
	}

	if cons(v1, cons(v2, cons(v3, NIL))) != (Pair{v1, Pair{v2, Pair{v3, NIL}}}) {
		t.Error("cons2")
	}
}

func TestCarInvalidValue(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Expected car to panic")
		}
	}()

	car(Number(3))
}

func TestCarNil(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Expected car to panic")
		}
	}()

	car(NIL)
}

func TestCarValid(t *testing.T) {
	for _, p := range []struct {
		a, b Value
	}{{Pair{v1, v2}, v1},
		{Pair{v1, Pair{v2, NIL}}, v1},
		{Pair{Pair{v1, v2}, NIL}, Pair{v1, v2}},
	} {
		if car(p.a) != p.b {
			t.Error("atom:", p.a)
		}
	}
}

func TestCdrInvalidValue(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Expected cdr to panic")
		}
	}()

	cdr(Number(3))
}

func TestCdrNil(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Expected cdr to panic")
		}
	}()

	cdr(NIL)
}

func TestCdrValid(t *testing.T) {
	for _, p := range []struct {
		a, b Value
	}{{Pair{v1, v2}, v2},
		{Pair{v1, Pair{v2, NIL}}, Pair{v2, NIL}},
		{Pair{Pair{v1, v2}, NIL}, NIL},
	} {
		if cdr(p.a) != p.b {
			t.Error("cdr:", p.a)
		}
	}
}

func TestList(t *testing.T) {
	if list(v1, v2, v3) != (Pair{v1, Pair{v2, Pair{v3, NIL}}}) {
		t.Error("list")
	}

	if list() != NIL {
		t.Error("empty list")
	}
}

func TestConvenience(t *testing.T) {
	ls := list(list(v1, v2, v3), list(v4, v5), v6)

	if caar(ls) != v1 {
		t.Error("caar")
	}
	if cadr(ls) != list(v4, v5) {
		t.Error("cadr")
	}
	if cadar(ls) != v2 {
		t.Error("cadar")
	}
	if caddr(ls) != v6 {
		t.Error("caddr")
	}
	if caddar(ls) != v3 {
		t.Error("caddar")
	}
}

func TestAssocInvalid(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("assoc should have paniced")
		}
	}()

	al := list(list(v1, v2), list(v3, v4), list(v5, v6))
	assoc(Number(7), al)
}

func TestAssoc(t *testing.T) {
	al := list(list(v1, v2), list(v3, v4), list(v5, v6))
	if assoc(v1, al) != v2 {
		t.Errorf("assoc1")
	}
	if assoc(v3, al) != v4 {
		t.Errorf("assoc2")
	}
	al = cons(list(v3, v6), al)
	if assoc(v3, al) != v6 {
		t.Errorf("assoc3")
	}
	al = cdr(al)
	if assoc(v3, al) != v4 {
		t.Errorf("assoc4")
	}
}
