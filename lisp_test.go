package golisp

import (
	"testing"
)

func TestList(t *testing.T) {
	l := cons(Value(Number(7)), NIL)
	if car(l) != Value(Number(7)) {
		t.Error("car failed")
	}

	if cdr(l) != NIL {
		t.Error("cdr failed")
	}
}

func TestEval(t *testing.T) {
	env := NewEnvironment()
	if Eval(Value(Number(7)), env) != Value(Number(7)) {
		t.Error("Eval number failed")
	}

	Set(env, "FOO", Number(7))
	if Eval(Symbol("FOO"), env) != Value(Number(7)) {
		t.Error("Eval basic lookup got", Eval(Symbol("FOO"), env))
	}
}
