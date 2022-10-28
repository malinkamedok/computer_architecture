package vm_test

import (
	"ca/internal/comp"
	vm2 "ca/internal/vm"
	"encoding/json"
	"testing"
)

func TestAdi(t *testing.T) {
	prog := "[\n {\n  \"opcode\": \"ADI\",\n  \"term\": {\n   \"line\": 0,\n   \"arg0\": \"r0\",\n   \"arg1\": \"1\"\n  }\n },\n {\n  \"opcode\": \"BRK\",\n  \"term\": {\n   \"line\": 1,\n   \"arg0\": \"\",\n   \"arg1\": \"\"\n  }\n }\n]"

	var isa []comp.IsaFormat

	err := json.Unmarshal([]byte(prog), &isa)
	if err != nil {
		t.Errorf("Cannot umarshall isa format")
	}

	vm := vm2.NewVirtualMachine(1000)
	vm.LoadProg(isa)
	vm.Run()
	regs := vm.GetRegsStateForTest()
	flags := vm.GetFlagsStateForTest()
	if regs["r0"] != 1 {
		t.Errorf("Error in register values")
	}
	if flags["EQ"] || flags["DV"] {
		t.Errorf("Error in flags value")
	}
}

func TestMvi(t *testing.T) {
	prog := "[\n {\n  \"opcode\": \"MVI\",\n  \"term\": {\n   \"line\": 0,\n   \"arg0\": \"r0\",\n   \"arg1\": \"3\"\n  }\n },\n {\n  \"opcode\": \"BRK\",\n  \"term\": {\n   \"line\": 1,\n   \"arg0\": \"\",\n   \"arg1\": \"\"\n  }\n }\n]"

	var isa []comp.IsaFormat

	err := json.Unmarshal([]byte(prog), &isa)
	if err != nil {
		t.Errorf("Cannot umarshall isa format")
	}

	vm := vm2.NewVirtualMachine(1000)
	vm.LoadProg(isa)
	vm.Run()
	regs := vm.GetRegsStateForTest()
	flags := vm.GetFlagsStateForTest()
	if regs["r0"] != 3 {
		t.Errorf("Error in register values")
	}
	if flags["EQ"] || flags["DV"] {
		t.Errorf("Error in flags value")
	}
}

func TestAdd(t *testing.T) {
	prog := "[\n {\n  \"opcode\": \"MVI\",\n  \"term\": {\n   \"line\": 0,\n   \"arg0\": \"r1\",\n   \"arg1\": \"3\"\n  }\n },\n {\n  \"opcode\": \"ADD\",\n  \"term\": {\n   \"line\": 1,\n   \"arg0\": \"r0\",\n   \"arg1\": \"r1\"\n  }\n },\n {\n  \"opcode\": \"BRK\",\n  \"term\": {\n   \"line\": 2,\n   \"arg0\": \"\",\n   \"arg1\": \"\"\n  }\n }\n]"

	var isa []comp.IsaFormat

	err := json.Unmarshal([]byte(prog), &isa)
	if err != nil {
		t.Errorf("Cannot umarshall isa format")
	}

	vm := vm2.NewVirtualMachine(1000)
	vm.LoadProg(isa)
	vm.Run()
	regs := vm.GetRegsStateForTest()
	flags := vm.GetFlagsStateForTest()
	if regs["r0"] != 3 && regs["r1"] != 3 {
		t.Errorf("Error in register values")
	}
	if flags["EQ"] || flags["DV"] {
		t.Errorf("Error in flags value")
	}
}

func TestMov(t *testing.T) {
	prog := "[\n {\n  \"opcode\": \"MVI\",\n  \"term\": {\n   \"line\": 0,\n   \"arg0\": \"r0\",\n   \"arg1\": \"1\"\n  }\n },\n {\n  \"opcode\": \"MOV\",\n  \"term\": {\n   \"line\": 1,\n   \"arg0\": \"r1\",\n   \"arg1\": \"r0\"\n  }\n },\n {\n  \"opcode\": \"BRK\",\n  \"term\": {\n   \"line\": 2,\n   \"arg0\": \"\",\n   \"arg1\": \"\"\n  }\n }\n]"

	var isa []comp.IsaFormat

	err := json.Unmarshal([]byte(prog), &isa)
	if err != nil {
		t.Errorf("Cannot umarshall isa format")
	}

	vm := vm2.NewVirtualMachine(1000)
	vm.LoadProg(isa)
	vm.Run()
	regs := vm.GetRegsStateForTest()
	flags := vm.GetFlagsStateForTest()
	if regs["r0"] != 1 && regs["r1"] != 1 {
		t.Errorf("Error in register values")
	}
	if flags["EQ"] || flags["DV"] {
		t.Errorf("Error in flags value")
	}
}

func TestCmp(t *testing.T) {
	prog := "[\n {\n  \"opcode\": \"MVI\",\n  \"term\": {\n   \"line\": 0,\n   \"arg0\": \"r0\",\n   \"arg1\": \"2\"\n  }\n },\n {\n  \"opcode\": \"MOV\",\n  \"term\": {\n   \"line\": 1,\n   \"arg0\": \"r1\",\n   \"arg1\": \"r0\"\n  }\n },\n {\n  \"opcode\": \"CMP\",\n  \"term\": {\n   \"line\": 2,\n   \"arg0\": \"r0\",\n   \"arg1\": \"r1\"\n  }\n },\n {\n  \"opcode\": \"BRK\",\n  \"term\": {\n   \"line\": 3,\n   \"arg0\": \"\",\n   \"arg1\": \"\"\n  }\n }\n]"

	var isa []comp.IsaFormat

	err := json.Unmarshal([]byte(prog), &isa)
	if err != nil {
		t.Errorf("Cannot umarshall isa format")
	}

	vm := vm2.NewVirtualMachine(1000)
	vm.LoadProg(isa)
	vm.Run()
	regs := vm.GetRegsStateForTest()
	flags := vm.GetFlagsStateForTest()
	if regs["r0"] != 2 && regs["r1"] != 2 {
		t.Errorf("Error in register values")
	}
	if !flags["EQ"] || flags["DV"] {
		t.Errorf("Error in flags value")
	}
}

func TestChk(t *testing.T) {
	prog := "[\n {\n  \"opcode\": \"ADI\",\n  \"term\": {\n   \"line\": 0,\n   \"arg0\": \"r0\",\n   \"arg1\": \"21\"\n  }\n },\n {\n  \"opcode\": \"CHK\",\n  \"term\": {\n   \"line\": 1,\n   \"arg0\": \"r0\",\n   \"arg1\": \"7\"\n  }\n },\n {\n  \"opcode\": \"BRK\",\n  \"term\": {\n   \"line\": 2,\n   \"arg0\": \"\",\n   \"arg1\": \"\"\n  }\n }\n]"

	var isa []comp.IsaFormat

	err := json.Unmarshal([]byte(prog), &isa)
	if err != nil {
		t.Errorf("Cannot umarshall isa format")
	}

	vm := vm2.NewVirtualMachine(1000)
	vm.LoadProg(isa)
	vm.Run()
	regs := vm.GetRegsStateForTest()
	flags := vm.GetFlagsStateForTest()
	if regs["r0"] != 21 {
		t.Errorf("Error in register values")
	}
	if flags["EQ"] || !flags["DV"] {
		t.Errorf("Error in flags value")
	}
}

func TestJmpJne(t *testing.T) {
	prog := "[\n {\n  \"opcode\": \"MVI\",\n  \"term\": {\n   \"line\": 0,\n   \"arg0\": \"r0\",\n   \"arg1\": \"10\"\n  }\n },\n {\n  \"opcode\": \"MVI\",\n  \"term\": {\n   \"line\": 1,\n   \"arg0\": \"r1\",\n   \"arg1\": \"0\"\n  }\n },\n {\n  \"opcode\": \"JMP\",\n  \"term\": {\n   \"line\": 2,\n   \"arg0\": \"\",\n   \"arg1\": \"\"\n  }\n },\n {\n  \"opcode\": \"CMP\",\n  \"term\": {\n   \"line\": 3,\n   \"arg0\": \"r0\",\n   \"arg1\": \"r1\"\n  }\n },\n {\n  \"opcode\": \"ADI\",\n  \"term\": {\n   \"line\": 4,\n   \"arg0\": \"r1\",\n   \"arg1\": \"1\"\n  }\n },\n {\n  \"opcode\": \"JNE\",\n  \"term\": {\n   \"line\": 5,\n   \"arg0\": \"\",\n   \"arg1\": \"\"\n  }\n },\n {\n  \"opcode\": \"BRK\",\n  \"term\": {\n   \"line\": 6,\n   \"arg0\": \"\",\n   \"arg1\": \"\"\n  }\n }\n]"

	var isa []comp.IsaFormat

	err := json.Unmarshal([]byte(prog), &isa)
	if err != nil {
		t.Errorf("Cannot umarshall isa format")
	}

	vm := vm2.NewVirtualMachine(1000)
	vm.LoadProg(isa)
	vm.Run()
	flags := vm.GetFlagsStateForTest()
	if flags["EQ"] {
		t.Errorf("Error in flags value")
	}
}

func TestPhi(t *testing.T) {
	prog := "[\n {\n  \"opcode\": \"MVI\",\n  \"term\": {\n   \"line\": 0,\n   \"arg0\": \"r0\",\n   \"arg1\": \"21\"\n  }\n },\n {\n  \"opcode\": \"MVI\",\n  \"term\": {\n   \"line\": 1,\n   \"arg0\": \"r1\",\n   \"arg1\": \"0\"\n  }\n },\n {\n  \"opcode\": \"CHK\",\n  \"term\": {\n   \"line\": 2,\n   \"arg0\": \"r0\",\n   \"arg1\": \"3\"\n  }\n },\n {\n  \"opcode\": \"PHI\",\n  \"term\": {\n   \"line\": 3,\n   \"arg0\": \"\",\n   \"arg1\": \"\"\n  }\n },\n {\n  \"opcode\": \"ADI\",\n  \"term\": {\n   \"line\": 4,\n   \"arg0\": \"r1\",\n   \"arg1\": \"1\"\n  }\n },\n {\n  \"opcode\": \"BRK\",\n  \"term\": {\n   \"line\": 5,\n   \"arg0\": \"\",\n   \"arg1\": \"\"\n  }\n }\n]"

	var isa []comp.IsaFormat

	err := json.Unmarshal([]byte(prog), &isa)
	if err != nil {
		t.Errorf("Cannot umarshall isa format")
	}

	vm := vm2.NewVirtualMachine(1000)
	vm.LoadProg(isa)
	vm.Run()
	flags := vm.GetFlagsStateForTest()
	if flags["DV"] {
		t.Errorf("Error in flags value")
	}
}

func TestAdiMvi(t *testing.T) {
	prog := "[\n {\n  \"opcode\": \"ADI\",\n  \"term\": {\n   \"line\": 0,\n   \"arg0\": \"r0\",\n   \"arg1\": \"1\"\n  }\n },\n {\n  \"opcode\": \"MVI\",\n  \"term\": {\n   \"line\": 1,\n   \"arg0\": \"r0\",\n   \"arg1\": \"0\"\n  }\n },\n {\n  \"opcode\": \"BRK\",\n  \"term\": {\n   \"line\": 2,\n   \"arg0\": \"\",\n   \"arg1\": \"\"\n  }\n }\n]"

	var isa []comp.IsaFormat

	err := json.Unmarshal([]byte(prog), &isa)
	if err != nil {
		t.Errorf("Cannot umarshall isa format")
	}

	vm := vm2.NewVirtualMachine(1000)
	vm.LoadProg(isa)
	vm.Run()
	regs := vm.GetRegsStateForTest()
	flags := vm.GetFlagsStateForTest()
	if regs["r0"] != 0 {
		t.Errorf("Error in register values")
	}
	if flags["EQ"] || flags["DV"] {
		t.Errorf("Error in flags value")
	}
}

func TestPrt(t *testing.T) {
	prog := "[\n {\n  \"opcode\": \"MVI\",\n  \"term\": {\n   \"line\": 0,\n   \"arg0\": \"r0\",\n   \"arg1\": \"7\"\n  }\n },\n {\n  \"opcode\": \"MOV\",\n  \"term\": {\n   \"line\": 1,\n   \"arg0\": \"r3\",\n   \"arg1\": \"r0\"\n  }\n },\n {\n  \"opcode\": \"PRT\",\n  \"term\": {\n   \"line\": 2,\n   \"arg0\": \"\",\n   \"arg1\": \"\"\n  }\n },\n {\n  \"opcode\": \"BRK\",\n  \"term\": {\n   \"line\": 3,\n   \"arg0\": \"\",\n   \"arg1\": \"\"\n  }\n }\n]"

	var isa []comp.IsaFormat

	err := json.Unmarshal([]byte(prog), &isa)
	if err != nil {
		t.Errorf("Cannot umarshall isa format")
	}

	vm := vm2.NewVirtualMachine(1000)
	vm.LoadProg(isa)
	vm.Run()
	regs := vm.GetRegsStateForTest()
	if regs["r3"] != 7 {
		t.Errorf("Error in register values")
	}
}
