package main

import (
	"fmt"
	"github.com/mikebeller/golisp/lisp"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	envStr, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var progStr = "(main)"
	if len(os.Args) > 1 {
		progStr = os.Args[1]
	}

	v := lisp.Eval(lisp.ReadStr(progStr), lisp.ReadStr(string(envStr)))
	fmt.Println(lisp.WriteStr(v))
}
