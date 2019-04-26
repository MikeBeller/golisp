package golisp

import (
	"fmt"
	"testing"
)

var _ = fmt.Println

var v1 Value = Symbol("A")
var v2 Value = Symbol("B")
var v3 Value = Symbol("C")
var v4 Value = Symbol("D")
var v5 Value = Symbol("E")
var v6 Value = Symbol("F")

func S(s string) Value {
	return Symbol(s)
}

func TestQuote(t *testing.T) {
	for _, v := range []Value{NIL, v3, Symbol("FOO")} {
		if quote(v) != v {
			t.Error("quote:", v)
		}
	}
	if quote(v3) != v3 {
		t.Error("quote outright")
	}
}

func TestAtom(t *testing.T) {
	for _, p := range []struct {
		a, b Value
	}{{NIL, TRUE},
		{v3, TRUE},
		{Symbol("FOO"), TRUE},
		{Pair{NIL, NIL}, NIL}} {
		if atom(p.a) != p.b {
			t.Error("atom:", p.a)
		}
	}
}

func TestEq(t *testing.T) {
	for _, p := range []struct {
		a, b, c Value
	}{{NIL, NIL, TRUE},
		{NIL, v3, NIL},
		{v3, v3, TRUE},
		{v3, v4, NIL},
		{Symbol("FOO"), Symbol("FOO"), TRUE},
		{Pair{NIL, NIL}, Pair{NIL, NIL}, NIL},
		{Pair{NIL, NIL}, Symbol("FOO"), NIL},
	} {
		if eq(p.a, p.b) != p.c {
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

	car(v3)
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

	cdr(v3)
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
	assoc(Symbol("FOO"), al)
}

func TestAssoc(t *testing.T) {
	al := list(list(v1, v2), list(v3, v4), list(v5, v6))
	if assoc(v1, al) != v2 {
		t.Error("assoc1")
	}
	if assoc(v3, al) != v4 {
		t.Error("assoc2")
	}
	al = cons(list(v3, v6), al)
	if assoc(v3, al) != v6 {
		t.Error("assoc3")
	}
	al = cdr(al)
	if assoc(v3, al) != v4 {
		t.Error("assoc4")
	}
}

func TestAppend(t *testing.T) {
	if append(NIL, NIL) != NIL {
		t.Error("append NIL NIL")
	}
	if append(NIL, list(v1, v2)) != list(v1, v2) {
		t.Error("append NIL list")
	}
	if append(list(v1, v2), list(v3, v4)) != list(v1, v2, v3, v4) {
		t.Error("append 4 items")
	}
}

func TestPair(t *testing.T) {
	if pair(NIL, NIL) != NIL {
		t.Error("pair nil nil")
	}
	if pair(list(v1, v2), list(v3, v4)) != list(list(v1, v3), list(v2, v4)) {
		t.Error("pair normal")
	}
}

var env Value = list(list(v1, v2), list(v3, v4))

func TestEvalAnAtom(t *testing.T) {
	if eval(v1, env) != v2 {
		t.Error("eval an atom")
	}
}

func TestEvalQuote(t *testing.T) {
	if eval(list(S("quote"), v2), env) != v2 {
		t.Error("eval 'quote")
	}
}

func TestEvalAtom(t *testing.T) {
	if eval(list(S("atom"), list(S("quote"), v2)), env) != TRUE {
		t.Error("eval 'atom of atom")
	}
}

func TestEvalNumber(t *testing.T) {
	if eval(Number(5), NIL) != Number(5) {
		t.Error("eval number")
	}
}

func TestRead(t *testing.T) {
	if readStr("FOO") != Symbol("FOO") {
		t.Error("read FOO")
	}
	if readStr(" FOO ") != Symbol("FOO") {
		t.Error("read FOO")
	}
	if readStr("(FOO BAR)") != list(S("FOO"), S("BAR")) {
		t.Error("read (FOO BAR)")
	}
	if readStr("(FOO BAR BAZ)") != list(S("FOO"), S("BAR"), S("BAZ")) {
		t.Error("read (FOO BAR BAZ)")
	}
	if readStr("(FOO (BAR 37) BEE)") != list(S("FOO"), list(S("BAR"), Number(37)), S("BEE")) {
		t.Error("read nested list")
	}
	if readStr("'FOO") != list(S("quote"), S("FOO")) {
		t.Error("quoted symbol")
	}
	if readStr("'(A B)") != list(S("quote"), list(S("A"), S("B"))) {
		t.Error("quoted list")
	}
	if readStr("234") != Number(234) {
		t.Error("Read number")
	}
	if readStr("-234") != Number(-234) {
		t.Error("Read negative number")
	}
	if readStr("0") != Number(0) {
		t.Error("Read 0")
	}
}

func TestArithmetic(t *testing.T) {
	if eval(list(S("add"), Number(3), Number(5)), env) != Number(8) {
		t.Error("add 3 5")
	}
	if eval(list(S("sub"), Number(3), Number(5)), env) != Number(-2) {
		t.Error("sub 3 5")
	}
	if eval(list(S("lt"), Number(3), Number(5)), env) != TRUE {
		t.Error("lt 3 5")
	}
	if eval(readStr("(add a b)"), readStr("((a 3) (b 5))")) != Number(8) {
		t.Error("add a b")
	}
}

func TestWrite(t *testing.T) {
	if toStr(readStr("3")) != "3" {
		t.Error("write 3", toStr(readStr("3")))
	}
	if toStr(readStr("(FOO 3 (7 9))")) != "(FOO 3 (7 9))" {
		t.Error("write list")
	}
}

func TestLambdaProg(t *testing.T) {
	prog := `(
(f (lambda (n m)
   (cond
       ((eq n 0) m)
       ('t (f (sub n 1) (add m 2)))
       )))
(main (lambda () (f 10 0)))
)`
	r := eval(readStr("(main)"), readStr(prog))
	if r != Number(20) {
		t.Error("lambda prog")
	}
}

func TestLambdaFib(t *testing.T) {
	prog := `(
(fib (lambda (a b n)
   (cond
       ((eq n 0) a)
       ('t (fib b (add a b) (sub n 1) ))
       )))
(main (lambda () (fib 0 1 10)))
)`
	r := eval(readStr("(main)"), readStr(prog))
	if r != Number(55) {
		t.Error("lambda fib")
	}
}
