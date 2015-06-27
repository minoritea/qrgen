package main
import (
	"code.google.com/p/rsc/qr"
	"os"
	"log"
	"io/ioutil"
)

func main() {
	src := os.Args[1]
	if len(src) < 1 {
		log.Fatal("Source string must be given!")
	}

	dst := os.Args[2]
	if len(dst) < 1 {
		dst = "qr.png"
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
