package main

import (
	"ca/internal/comp"
	"flag"
	"log"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "f", "", "assembly code file")
	flag.Parse()

	if fileName == "" {
		log.Fatal("usage: ./comp -f <filename>")
	} else {
		util := comp.NewUtil()
		code, err := util.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
		lexer := comp.NewTrans(code)
		vl := lexer.Split()
		err = util.WriteFile(vl)
		if err != nil {
			log.Fatal(err)
		}
	}
}
