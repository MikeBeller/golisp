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

func listVal(ls List) Value {
	switch p := ls.(type) {
	case Nil:
		return NIL
	case Pair:
		return p
	default:
		panic("invalid list")
	}
}

func list(vs ...Value) List {
	ls := List(NIL)
	for i := len(vs) - 1; i >= 0; i-- {
		ls = cons(vs[i], listVal(ls))
	}
	return ls
}

type Environment List

func NewEnvironment() Environment {
	st := SymTab(make(map[Symbol]Value))
	return list(Value(st))
}

func Set(e Environment, s Symbol, v Value) {
	switch p := e.(type) {
	case Nil:
		panic("set empty environment")
	case Pair:
		st := car(p).(SymTab)
		st[s] = v
	}
}

func Lookup(s Symbol, e Environment) Value {
	l := e
	for {
		switch p := l.(type) {
		case Nil:
			panic("lookup nil")
		case Pair:
			st := p.a.(SymTab)
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
