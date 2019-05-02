package main

import (
	"flag"
	"fmt"
	"github.com/mikebeller/golisp/lisp"
	"io/ioutil"
	"log"
)

func main() {
	dbgPtr := flag.Bool("debug", false, "Enable debug printing")
	flag.Parse()

	if *dbgPtr {
		lisp.Debug(true)
	}

	if len(flag.Args()) < 1 {
		log.Fatal("Usage: golisp [--debug] expr [envfile]")
	}
	progStr := flag.Args()[0]

	envStr := "()"

	if len(flag.Args()) > 1 {
		envFile := flag.Args()[1]
		envBytes, err := ioutil.ReadFile(envFile)
		if err != nil {
			log.Fatal(err)
		}
		envStr = string(envBytes)
	}

	v := lisp.Eval(lisp.ReadStr(progStr), lisp.ReadStr(envStr))
	fmt.Println(lisp.WriteStr(v))
}
