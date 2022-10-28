package comp

import (
	"fmt"
	"log"
	"strings"
)

type Trans struct {
	Code []string
}

func NewTrans(code []string) *Trans {
	return &Trans{
		Code: code,
	}
}

func (l *Trans) Split() []IsaFormat {
	var encoded []IsaFormat
	for line, str := range l.Code {
		instruction := str[0:3]
		var args string
		if len(str) > 3 {
			args = str[4:]
		}
		switch instruction {
		case "adi":
			regAndImm := strings.Split(args, ",")
			if len(regAndImm) != 2 || !CheckReg(regAndImm[0]) {
				fmt.Println(adi)
				log.Fatal("error format of input text")
			}
			encoded = append(encoded, NewIsaFormat(adi, Term{
				Line: line,
				Arg0: regAndImm[0],
				Arg1: regAndImm[1],
			}))
		case "add":
			regAndReg := strings.Split(args, ",")
			if len(regAndReg) != 2 || !CheckReg(regAndReg[0]) || !CheckReg(regAndReg[1]) {
				fmt.Println(adi)
				log.Fatal("error format of input text")
			}
			encoded = append(encoded, NewIsaFormat(add, Term{
				Line: line,
				Arg0: regAndReg[0],
				Arg1: regAndReg[1],
			}))
		case "mov":
			regAndReg := strings.Split(args, ",")
			if len(regAndReg) != 2 || !CheckReg(regAndReg[0]) || !CheckReg(regAndReg[1]) {
				fmt.Println(mov)
				log.Fatal("error format of input text")
			}
			encoded = append(encoded, NewIsaFormat(mov, Term{
				Line: line,
				Arg0: regAndReg[0],
				Arg1: regAndReg[1],
			}))
		case "mvi":
			regAndImm := strings.Split(args, ",")
			if len(regAndImm) != 2 || !CheckReg(regAndImm[0]) {
				fmt.Println(mvi)
				log.Fatal("error format of input text")
			}
			encoded = append(encoded, NewIsaFormat(mvi, Term{
				Line: line,
				Arg0: regAndImm[0],
				Arg1: regAndImm[1],
			}))
		case "cmp":
			regAndReg := strings.Split(args, ",")
			if len(regAndReg) != 2 || !CheckReg(regAndReg[0]) || !CheckReg(regAndReg[1]) {
				fmt.Println(cmp)
				log.Fatal("error format of input text")
			}
			encoded = append(encoded, NewIsaFormat(cmp, Term{
				Line: line,
				Arg0: regAndReg[0],
				Arg1: regAndReg[1],
			}))
		case "chk":
			regAndImm := strings.Split(args, ",")
			if len(regAndImm) != 2 || !CheckReg(regAndImm[0]) {
				fmt.Println(chk)
				log.Fatal("error format of input text")
			}
			encoded = append(encoded, NewIsaFormat(chk, Term{
				Line: line,
				Arg0: regAndImm[0],
				Arg1: regAndImm[1],
			}))
		case "jmp":
			encoded = append(encoded, NewIsaFormat(jmp, Term{
				Line: line,
			}))
		case "jne":
			jmpCounter := 0
			for _, ops := range encoded {
				if ops.Opcode == jmp {
					jmpCounter++
				}
			}
			if jmpCounter == 0 {
				log.Fatal("error in loops")
			}
			encoded = append(encoded, NewIsaFormat(jne, Term{
				Line: line,
			}))
		case "phi":
			encoded = append(encoded, NewIsaFormat(phi, Term{
				Line: line,
			}))
		case "brk":
			encoded = append(encoded, NewIsaFormat(brk, Term{
				Line: line,
			}))
		case "inp":
			encoded = append(encoded, NewIsaFormat(inp, Term{
				Line: line,
			}))
		case "prt":
			encoded = append(encoded, NewIsaFormat(prt, Term{
				Line: line,
			}))
		default:
			log.Fatal("error format of input text")
		}
	}
	jmpCount := 0
	for _, ops := range encoded {
		if ops.Opcode == jmp {
			jmpCount++
		}
		if ops.Opcode == jne {
			jmpCount--
		}
	}
	if jmpCount != 0 {
		log.Fatal("error in loops")
	}
	return encoded
}
