package comp

type OpcodeType = string

const (

	//instructions
	adi OpcodeType = "ADI"
	add            = "ADD"
	mov            = "MOV"
	mvi            = "MVI"
	cmp            = "CMP"
	chk            = "CHK"
	jmp            = "JMP"
	jne            = "JNE"
	phi            = "PHI"
	brk            = "BRK"
	inp            = "INP"
	prt            = "PRT"

	// regs
	r0 = "r0"
	r1 = "r1"
	r2 = "r2"
	r3 = "r3" // for input/output
)

func CheckReg(mbReg string) bool {
	return mbReg == r0 ||
		mbReg == r1 ||
		mbReg == r2 ||
		mbReg == r3
}
