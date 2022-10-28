package main

import (
	"ca/internal/vm"
	"flag"
	"log"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "f", "", "op codes file")
	flag.Parse()

	if fileName == "" {
		log.Fatal("usage: ./vm -f <filename>")
	} else {
		util := vm.Util{}
		vl, err := util.ReadISAFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
		vrt := vm.NewVirtualMachine(1000)
		vrt.LoadProg(vl)
		vrt.Run()
	}

}
