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
type SymTab map[Symbol]Value

func (Nil) IsType()    {}
func (Number) IsType() {}
func (Symbol) IsType() {}
func (Pair) IsType()   {}
func (SymTab) IsType() {}

var NIL = Nil(struct{}{})

type List interface {
	IsList()
}

func (Nil) IsList()  {}
func (Pair) IsList() {}

func cons(car Value, cdr Value) Pair {
	return Pair{car, cdr}
}

func car(p Pair) Value {
	return p.a
}

func cdr(p Pair) Value {
	return p.b
}

func list(vs ...Value) List {
	ls := NIL
	for i := len(vs) - 1; i >= 0; i-- {
		ls = cons(vs[i], ls)
	}
	return ls
}

type Environment List

func NewEnvironment() Environment {
	st := SymTab(make(map[Symbol]Value))
	return list(Value(st))
}

func (e Environment) Set(s Symbol, v Value) {
	e[0][s] = v
}

func Lookup(s Symbol, e Environment) Value {
	l := e
	for {
		switch l.(type) {
		case Nil:
			panic("lookup nil")
		case Pair:
			st := l.a.(SymTab)
			if v, ok := st[s]; ok {
				return v
			}
		}
	}
	panic("lookup wtf?")
}

func Eval(val Value, env Environment) Value {
	switch val.(type) {
	case Number:
		return val
	case Symbol:
		return Lookup(val.(Symbol), env)
	default:
		panic("type")
	}
}
