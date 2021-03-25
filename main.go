package main

import (
	"golang.org/x/xerrors"
	"io/ioutil"
	"log"
	"os"
	"rsc.io/qr"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var (
		src  string
		dst  string = "qr.png"
		argn int    = 2 // max number of arguments
	)

	stat, err := os.Stdin.Stat()
	if err != nil {
		return err
	}

	if stat.Mode()&os.ModeCharDevice == 0 {
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		src = string(input)
		argn-- // if the input text exists, "src" argument will be not required.
	} else {
		if len(os.Args) < 2 {
			return xerrors.New("The source text must be given!")
		}
		src = os.Args[1]
	}

	if len(src) == 0 {
		return xerrors.New("The source text must be given!")
	}

	switch len(os.Args) {
	case argn:
		break
	case argn + 1:
		dst = os.Args[2]
	default:
		return xerrors.New("Invalid arguments!")
	}

	code, err := qr.Encode(src, qr.M)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(dst, code.PNG(), 0644)
}
