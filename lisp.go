package golisp

import "fmt"

var _ = fmt.Printf

type Type int
type Value interface {
	Type() Type
}

const (
	NIL Type = iota
	NUMBER
	SYMBOL
	CONS
)

type Nil struct{}
type Number int
type Symbol string
type Cons struct {
	car Value
	cdr Value
}

func (Nil) Type() Type    { return NIL }
func (Number) Type() Type { return NUMBER }
func (Symbol) Type() Type { return SYMBOL }
func (Cons) Type() Type   { return CONS }

func cons(car Value, cdr Value) Cons {
	return Cons{car, cdr}
}

func car(c Cons) Value {
	return c.car
}

func cdr(c Cons) Value {
	return c.cdr
}

type SymTab map[Symbol]Value
type Environment []SymTab

func NewEnvironment() Environment {
	st := make(map[Symbol]Value)
	return []SymTab{st}
}

func (e Environment) Set(s Symbol, v Value) {
	e[0][s] = v
}

func Lookup(s Symbol, e Environment) Value {
	for i := len(e) - 1; i >= 0; i-- {
		if v, ok := e[i][s]; ok {
			return v
		}
	}
	panic("lookup")
}

func Eval(val Value, env Environment) Value {
	switch val.Type() {
	case NUMBER:
		return val
	case SYMBOL:
		return Lookup(val.(Symbol), env)
	default:
		panic("type")
	}
}
