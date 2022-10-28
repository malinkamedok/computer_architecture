package comp

type IsaFormat struct {
	Opcode OpcodeType `json:"opcode"`
	Term   Term       `json:"term"`
}

func NewIsaFormat(opcode OpcodeType, term Term) IsaFormat {
	return IsaFormat{
		Opcode: opcode,
		Term:   term,
	}
}
