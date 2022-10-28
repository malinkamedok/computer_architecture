package comp_test

import (
	"ca/internal/comp"
	"testing"
)

func TestADI(t *testing.T) {
	prog := []string{"adi r1,3"}

	trans := comp.NewTrans(prog)
	result := trans.Split()
	if len(result) != 1 {
		t.Errorf("Error in isa format collection")
	}
	isaTkn := result[0]
	if isaTkn.Opcode != "ADI" {
		t.Errorf("Error in format of isa output")
	}
	if isaTkn.Term.Line != 0 {
		t.Errorf("Error in line number")
	}
	if isaTkn.Term.Arg0 != "r1" || isaTkn.Term.Arg1 != "3" {
		t.Errorf("Error in arguments of instruction")
	}
}

func TestADD(t *testing.T) {
	prog := []string{"add r0,r1"}
	trans := comp.NewTrans(prog)
	result := trans.Split()

	if len(result) != 1 {
		t.Errorf("Error in isa format collection")
	}
	isaTkn := result[0]
	if isaTkn.Opcode != "ADD" {
		t.Errorf("Error in format of isa output")
	}
	if isaTkn.Term.Line != 0 {
		t.Errorf("Error in line number")
	}
	if isaTkn.Term.Arg0 != "r0" || isaTkn.Term.Arg1 != "r1" {
		t.Errorf("Error in arguments of instruction")
	}
}

func TestMOV(t *testing.T) {
	prog := []string{"mov r0,r1"}
	trans := comp.NewTrans(prog)
	result := trans.Split()

	if len(result) != 1 {
		t.Errorf("Error in isa format collection")
	}
	isaTkn := result[0]
	if isaTkn.Opcode != "MOV" {
		t.Errorf("Error in format of isa output")
	}
	if isaTkn.Term.Line != 0 {
		t.Errorf("Error in line number")
	}
	if isaTkn.Term.Arg0 != "r0" || isaTkn.Term.Arg1 != "r1" {
		t.Errorf("Error in arguments of instruction")
	}
}

func TestMVI(t *testing.T) {
	prog := []string{"mvi r0,1"}
	trans := comp.NewTrans(prog)
	result := trans.Split()

	if len(result) != 1 {
		t.Errorf("Error in isa format collection")
	}
	isaTkn := result[0]
	if isaTkn.Opcode != "MVI" {
		t.Errorf("Error in format of isa output")
	}
	if isaTkn.Term.Line != 0 {
		t.Errorf("Error in line number")
	}
	if isaTkn.Term.Arg0 != "r0" || isaTkn.Term.Arg1 != "1" {
		t.Errorf("Error in arguments of instruction")
	}
}

func TestCMP(t *testing.T) {
	prog := []string{"cmp r0,r1"}
	trans := comp.NewTrans(prog)
	result := trans.Split()

	if len(result) != 1 {
		t.Errorf("Error in isa format collection")
	}
	isaTkn := result[0]
	if isaTkn.Opcode != "CMP" {
		t.Errorf("Error in format of isa output")
	}
	if isaTkn.Term.Line != 0 {
		t.Errorf("Error in line number")
	}
	if isaTkn.Term.Arg0 != "r0" || isaTkn.Term.Arg1 != "r1" {
		t.Errorf("Error in arguments of instruction")
	}
}

func TestCHK(t *testing.T) {
	prog := []string{"chk r0,1"}
	trans := comp.NewTrans(prog)
	result := trans.Split()

	if len(result) != 1 {
		t.Errorf("Error in isa format collection")
	}
	isaTkn := result[0]
	if isaTkn.Opcode != "CHK" {
		t.Errorf("Error in format of isa output")
	}
	if isaTkn.Term.Line != 0 {
		t.Errorf("Error in line number")
	}
	if isaTkn.Term.Arg0 != "r0" || isaTkn.Term.Arg1 != "1" {
		t.Errorf("Error in arguments of instruction")
	}
}

func TestJMPJNE(t *testing.T) {
	prog := []string{"jmp", "jne"}
	trans := comp.NewTrans(prog)
	result := trans.Split()

	if len(result) != 2 {
		t.Errorf("Error in isa format collection")
	}
	isaTknJmp := result[0]
	if isaTknJmp.Opcode != "JMP" {
		t.Errorf("Error in format of isa output JMP")
	}
	isaTknJne := result[1]
	if isaTknJne.Opcode != "JNE" {
		t.Errorf("Error in format of isa output JNE")
	}
	if isaTknJmp.Term.Line != 0 {
		t.Errorf("Error in line number JMP")
	}
	if isaTknJne.Term.Line != 1 {
		t.Errorf("Error in line number JNE")
	}
	if isaTknJmp.Term.Line > isaTknJne.Term.Line {
		t.Errorf("Error in line order")
	}
}

func TestADIADD(t *testing.T) {
	prog := []string{"adi r0,3", "add r0,r1"}
	trans := comp.NewTrans(prog)
	result := trans.Split()

	if len(result) != 2 {
		t.Errorf("Error in isa format collection")
	}
	isaTknAdi := result[0]
	if isaTknAdi.Opcode != "ADI" {
		t.Errorf("Error in format of isa output ADI")
	}
	isaTknAdd := result[1]
	if isaTknAdd.Opcode != "ADD" {
		t.Errorf("Error in format of isa output ADD")
	}
	if isaTknAdi.Term.Line != 0 {
		t.Errorf("Error in line number ADI")
	}
	if isaTknAdd.Term.Line != 1 {
		t.Errorf("Error in line number ADD")
	}
	if isaTknAdi.Term.Arg0 != "r0" || isaTknAdi.Term.Arg1 != "3" {
		t.Errorf("Error in arguments of instruction ADI")
	}
	if isaTknAdd.Term.Arg0 != "r0" || isaTknAdd.Term.Arg1 != "r1" {
		t.Errorf("Error in arguments of instruction ADD")
	}
}

func TestMOVMVI(t *testing.T) {
	prog := []string{"mov r0,r1", "mvi r0,1"}
	trans := comp.NewTrans(prog)
	result := trans.Split()

	if len(result) != 2 {
		t.Errorf("Error in isa format collection")
	}
	isaTknMov := result[0]
	if isaTknMov.Opcode != "MOV" {
		t.Errorf("Error in format of isa output MOV")
	}
	isaTknMvi := result[1]
	if isaTknMvi.Opcode != "MVI" {
		t.Errorf("Error in format of isa output MVI")
	}
	if isaTknMov.Term.Line != 0 {
		t.Errorf("Error in line number MOV")
	}
	if isaTknMvi.Term.Line != 1 {
		t.Errorf("Error in line number MVI")
	}
	if isaTknMov.Term.Arg0 != "r0" || isaTknMov.Term.Arg1 != "r1" {
		t.Errorf("Error in arguments of instruction MOV")
	}
	if isaTknMvi.Term.Arg0 != "r0" || isaTknMvi.Term.Arg1 != "1" {
		t.Errorf("Error in arguments of instruction MVI")
	}
}

func TestADDMOV(t *testing.T) {
	prog := []string{"add r0,r1", "mov r0,r2"}
	trans := comp.NewTrans(prog)
	result := trans.Split()

	if len(result) != 2 {
		t.Errorf("Error in isa format collection")
	}
	isaTknAdd := result[0]
	if isaTknAdd.Opcode != "ADD" {
		t.Errorf("Error in format of isa output ADD")
	}
	isaTknMov := result[1]
	if isaTknMov.Opcode != "MOV" {
		t.Errorf("Error in format of isa output MOV")
	}
	if isaTknAdd.Term.Line != 0 {
		t.Errorf("Error in line number ADD")
	}
	if isaTknMov.Term.Line != 1 {
		t.Errorf("Error in line number MOV")
	}
	if isaTknAdd.Term.Arg0 != "r0" || isaTknAdd.Term.Arg1 != "r1" {
		t.Errorf("Error in arguments of instruction ADD")
	}
	if isaTknMov.Term.Arg0 != "r0" || isaTknMov.Term.Arg1 != "r2" {
		t.Errorf("Error in arguments of instruction MOV")
	}
}

func TestADIMVI(t *testing.T) {
	prog := []string{"adi r0,1", "mvi r0,0"}
	trans := comp.NewTrans(prog)
	result := trans.Split()

	if len(result) != 2 {
		t.Errorf("Error in isa format collection")
	}
	isaTknAdi := result[0]
	if isaTknAdi.Opcode != "ADI" {
		t.Errorf("Error in format of isa output ADI")
	}
	isaTknMvi := result[1]
	if isaTknMvi.Opcode != "MVI" {
		t.Errorf("Error in format of isa output MVI")
	}
	if isaTknAdi.Term.Line != 0 {
		t.Errorf("Error in line number ADI")
	}
	if isaTknMvi.Term.Line != 1 {
		t.Errorf("Error in line number MVI")
	}
	if isaTknAdi.Term.Arg0 != "r0" || isaTknAdi.Term.Arg1 != "1" {
		t.Errorf("Error in arguments of instruction ADI")
	}
	if isaTknMvi.Term.Arg0 != "r0" || isaTknMvi.Term.Arg1 != "0" {
		t.Errorf("Error in arguments of instruction MVI")
	}
}

func TestADDCMP(t *testing.T) {
	prog := []string{"add r0,r1", "cmp r0,r1"}
	trans := comp.NewTrans(prog)
	result := trans.Split()

	if len(result) != 2 {
		t.Errorf("Error in isa format collection")
	}
	isaTknAdd := result[0]
	if isaTknAdd.Opcode != "ADD" {
		t.Errorf("Error in format of isa output ADD")
	}
	isaTknCmp := result[1]
	if isaTknCmp.Opcode != "CMP" {
		t.Errorf("Error in format of isa output CMP")
	}
	if isaTknAdd.Term.Line != 0 {
		t.Errorf("Error in line number ADD")
	}
	if isaTknCmp.Term.Line != 1 {
		t.Errorf("Error in line number CMP")
	}
	if isaTknAdd.Term.Arg0 != "r0" || isaTknAdd.Term.Arg1 != "r1" {
		t.Errorf("Error in arguments of instruction ADD")
	}
	if isaTknCmp.Term.Arg0 != "r0" || isaTknCmp.Term.Arg1 != "r1" {
		t.Errorf("Error in arguments of instruction CMP")
	}
}

func TestCMPCHK(t *testing.T) {
	prog := []string{"cmp r0,r1", "chk r0,1"}
	trans := comp.NewTrans(prog)
	result := trans.Split()

	if len(result) != 2 {
		t.Errorf("Error in isa format collection")
	}
	isaTknCmp := result[0]
	if isaTknCmp.Opcode != "CMP" {
		t.Errorf("Error in format of isa output CMP")
	}
	isaTknChk := result[1]
	if isaTknChk.Opcode != "CHK" {
		t.Errorf("Error in format of isa output CHK")
	}
	if isaTknCmp.Term.Line != 0 {
		t.Errorf("Error in line number CMP")
	}
	if isaTknChk.Term.Line != 1 {
		t.Errorf("Error in line number CHK")
	}
	if isaTknChk.Term.Arg0 != "r0" || isaTknChk.Term.Arg1 != "1" {
		t.Errorf("Error in arguments of instruction CHK")
	}
	if isaTknCmp.Term.Arg0 != "r0" || isaTknCmp.Term.Arg1 != "r1" {
		t.Errorf("Error in arguments of instruction CMP")
	}
}

func TestMOVCHK(t *testing.T) {
	prog := []string{"mov r0,r1", "chk r1,1"}
	trans := comp.NewTrans(prog)
	result := trans.Split()

	if len(result) != 2 {
		t.Errorf("Error in isa format collection")
	}
	isaTknMov := result[0]
	if isaTknMov.Opcode != "MOV" {
		t.Errorf("Error in format of isa output MOV")
	}
	isaTknChk := result[1]
	if isaTknChk.Opcode != "CHK" {
		t.Errorf("Error in format of isa output CHK")
	}
	if isaTknMov.Term.Line != 0 {
		t.Errorf("Error in line number MOV")
	}
	if isaTknChk.Term.Line != 1 {
		t.Errorf("Error in line number CHK")
	}
	if isaTknChk.Term.Arg0 != "r1" || isaTknChk.Term.Arg1 != "1" {
		t.Errorf("Error in arguments of instruction CHK")
	}
	if isaTknMov.Term.Arg0 != "r0" || isaTknMov.Term.Arg1 != "r1" {
		t.Errorf("Error in arguments of instruction MOV")
	}
}
