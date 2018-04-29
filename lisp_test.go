package golisp

import (
	"testing"
)

func TestList(t *testing.T) {
	l := cons(Value(Number(7)), Nil{})
	if car(l) != Value(Number(7)) {
		t.Error("car failed")
	}

	if cdr(l) != (Nil{}) {
		t.Error("cdr failed")
	}
}

func TestEval(t *testing.T) {
	env := NewEnvironment()
	if Eval(Value(Number(7)), env) != Value(Number(7)) {
		t.Error("Eval number failed")
	}

	env.Set("FOO", Number(7))
	if Eval(Symbol("FOO"), env) != Value(Number(7)) {
		t.Error("Eval basic lookup")
	}
}
