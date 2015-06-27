package main

import (
	"code.google.com/p/rsc/qr"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var src, dst string
	switch len(os.Args) {
	case 1:
		log.Fatal("Source string must be given!")
	case 2:
		src = os.Args[1]
		dst = "qr.png"
	case 3:
		src = os.Args[1]
		dst = os.Args[2]
	default:
		log.Fatal("Invalid arguments!")
	}

	code, err := qr.Encode(src, qr.M)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(dst, code.PNG(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
