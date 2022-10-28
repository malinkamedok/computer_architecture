package vm

import (
	"ca/internal/comp"
	"fmt"
	"log"
	"strconv"
)

type VirtualMachine struct {
	Memory      []string
	Registers   map[string]int
	Flags       map[string]bool
	JumpStorage []int
}

func initRegs() map[string]int {
	return map[string]int{
		"rc": 0,
		"r0": 0,
		"r1": 0,
		"r2": 0,
		"r3": 0,
	}
}

func initFlags() map[string]bool {
	return map[string]bool{
		"EQ": false,
		"DV": false,
	}
}

func NewVirtualMachine(memorySize int) *VirtualMachine {
	return &VirtualMachine{
		Memory:      make([]string, memorySize),
		Registers:   initRegs(),
		Flags:       initFlags(),
		JumpStorage: make([]int, 4),
	}
}

func (v *VirtualMachine) LoadProg(prog []comp.IsaFormat) {
	start := 10
	for _, opt := range prog {
		v.Memory[start] = opt.Opcode
		start++
		v.Memory[start] = strconv.Itoa(opt.Term.Line)
		start++
		if opt.Term.Arg0 != "" {
			v.Memory[start] = opt.Term.Arg0
			start++
		}
		if opt.Term.Arg1 != "" {
			v.Memory[start] = opt.Term.Arg1
			start++
		}
	}
}

func (v *VirtualMachine) GetRegsStateForTest() map[string]int {
	return v.Registers
}

func (v *VirtualMachine) GetFlagsStateForTest() map[string]bool {
	return v.Flags
}

func (v *VirtualMachine) Run() {
	v.Registers["rc"] = 10
	running := true
	for running {
		isaTkn := v.Memory[v.Registers["rc"]]
		switch isaTkn {
		case "ADD":
			v.Registers["rc"] += 2
			reg1 := v.Memory[v.Registers["rc"]]
			v.Registers["rc"]++
			reg2 := v.Memory[v.Registers["rc"]]
			v.Registers[reg1] += v.Registers[reg2]
			v.Registers["rc"]++
		case "ADI":
			v.Registers["rc"] += 2
			reg1 := v.Memory[v.Registers["rc"]]
			v.Registers["rc"]++
			strImm := v.Memory[v.Registers["rc"]]
			imm, err := strconv.Atoi(strImm)
			if err != nil {
				log.Fatal(err)
			}
			v.Registers[reg1] += imm
			v.Registers["rc"]++
		case "MOV":
			v.Registers["rc"] += 2
			reg1 := v.Memory[v.Registers["rc"]]
			v.Registers["rc"]++
			reg2 := v.Memory[v.Registers["rc"]]
			v.Registers[reg1] = v.Registers[reg2]
			v.Registers["rc"]++
		case "MVI":
			v.Registers["rc"] += 2
			reg1 := v.Memory[v.Registers["rc"]]
			v.Registers["rc"]++
			strImm := v.Memory[v.Registers["rc"]]
			imm, err := strconv.Atoi(strImm)
			if err != nil {
				log.Fatal(err)
			}
			v.Registers[reg1] = imm
			v.Registers["rc"]++
		case "CMP":
			v.Registers["rc"] += 2
			reg1 := v.Memory[v.Registers["rc"]]
			v.Registers["rc"]++
			reg2 := v.Memory[v.Registers["rc"]]
			if v.Registers[reg1] == v.Registers[reg2] {
				v.Flags["EQ"] = true
			} else {
				v.Flags["EQ"] = false
			}
			v.Registers["rc"]++
		case "CHK":
			v.Registers["rc"] += 2
			reg1 := v.Memory[v.Registers["rc"]]
			v.Registers["rc"]++
			strImm := v.Memory[v.Registers["rc"]]
			imm, err := strconv.Atoi(strImm)
			if err != nil {
				log.Fatal(err)
			}
			if v.Registers[reg1]%imm == 0 {
				v.Flags["DV"] = true
			}
			v.Registers["rc"]++
		case "JMP":
			v.Registers["rc"] += 2
			v.JumpStorage = append(v.JumpStorage, v.Registers["rc"])
		case "JNE":
			v.Registers["rc"] += 2
			if !v.Flags["EQ"] {
				value := v.JumpStorage[len(v.JumpStorage)-1]
				v.Registers["rc"] = value
			} else {
				v.JumpStorage = v.JumpStorage[:len(v.JumpStorage)-1]
			}
			v.Flags["EQ"] = false
		case "PHI":
			v.Registers["rc"] += 2
			if !v.Flags["DV"] {
				// Use this only for two args instructions
				v.Registers["rc"] += 4
			}
			v.Flags["DV"] = false
		case "PRT":
			v.Registers["rc"] += 2
			fmt.Printf("Trap for output value: \n%d\n", v.Registers["r3"])
		case "INP":
			v.Registers["rc"] += 2
			var value int
			fmt.Println("Trap for input value: ")
			_, err := fmt.Scanf("%d", &value)
			if err != nil {
				log.Fatal("cannot read int input symbol")
			}
			v.Registers["r3"] = value
		case "BRK":
			running = false
		default:
			fmt.Println(v.Memory[v.Registers["rc"]])
			log.Fatal("Unknown instruction")
		}
	}
}
