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

func cons(car Value, cdr Value) Pair {
	return Pair{car, cdr}
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

func list(vs ...Value) Value {
	var ls Value = NIL
	for i := len(vs) - 1; i >= 0; i-- {
		ls = cons(vs[i], ls)
	}
	return ls
}

type Environment Value

func NewEnvironment() Environment {
	return NIL
}

func Set(e Environment, s Symbol, v Value) Value {
	return cons(Pair{s,v}, e)

}

func Lookup(s Symbol, e Environment) Value {
	if e == NIL {
		return NIL
	} else {
		p := car(e).(Pair)
		if p.a == s {
			return p.b
		} else {
			return Lookup(s, cdr(e))
		}
	}
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
