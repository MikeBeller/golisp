package main

import (
	"flag"
	"fmt"
	"github.com/mikebeller/golisp/lisp"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	dbgPtr := flag.Bool("debug", false, "Enable debug printing")
	envPtr := flag.String("env", "'()", "Environment")
	envFilePtr := flag.String("envfile", "", "Environment file")
	flag.Parse()

	if *dbgPtr {
		lisp.Debug(true)
	}

	if len(flag.Args()) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] expr\n", os.Args[0])
		flag.PrintDefaults()
		log.Fatal("exiting")
	}
	progStr := flag.Args()[0]

	envStr := *envPtr
	if *envFilePtr != "" {
		envBytes, err := ioutil.ReadFile(*envFilePtr)
		if err != nil {
			log.Fatal(err)
		}
		envStr = string(envBytes)
	}

	v := lisp.Eval(lisp.ReadStr(progStr), lisp.ReadStr(envStr))
	fmt.Println(lisp.WriteStr(v))
}
