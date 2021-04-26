package sa2

func getOpcodeFunctionSet() map[byte]func(sa *SA2) {
	return map[byte]func(sa *SA2){
		0x81: rsl,
		0x82: rsr,
		0x93: add,
		0x84: sub,
		0x87: eor,
		0x68: for_loop,
		0x49: next_loop,
		0x4A: bcc,
		0x6B: bra,
		0x4C: finish,
	}
}

func rsl(sa *SA2) {
	sa.carryFlag = sa.register & 0x80000000
	sa.register = sa.register << 1
	if sa.carryFlag > 0 {
		sa.register = sa.register | 0x1
	}
	sa.register = sa.register & 0xFFFFFFFF
	sa.pointer += 1
}

func rsr(sa *SA2) {
	sa.carryFlag = sa.register & 0x1
	sa.register = sa.register >> 1
	if sa.carryFlag > 0 {
		sa.register = sa.register | 0x80000000
	}
	sa.pointer += 1
}

func add(sa *SA2) {
	sa.carryFlag = 0
	operands := sa.Tape[sa.pointer+1 : sa.pointer+5]
	addInt := int(operands[0])<<24 | int(operands[1])<<16 | int(operands[2])<<8 | int(operands[3])
	outputRegister := sa.register + addInt
	if outputRegister > 0xffffffff {
		sa.carryFlag = 1
		outputRegister = outputRegister & 0xffffffff
	}
	sa.register = outputRegister
	sa.pointer += 5
}

func sub(sa *SA2) {
	sa.carryFlag = 0
	operands := sa.Tape[sa.pointer+1 : sa.pointer+5]
	subInt := int(operands[0])<<24 | int(operands[1])<<16 | int(operands[2])<<8 | int(operands[3])
	outputRegister := sa.register - subInt
	if outputRegister < 0 {
		sa.carryFlag = 1
		outputRegister = outputRegister & 0xffffffff
	}
	sa.register = outputRegister
	sa.pointer += 5
}

func eor(sa *SA2) {
	operands := sa.Tape[sa.pointer+1 : sa.pointer+5]
	xorInt := int(operands[0])<<24 | int(operands[1])<<16 | int(operands[2])<<8 | int(operands[3])
	sa.register = sa.register ^ xorInt
	sa.pointer += 5
}

func for_loop(sa *SA2) {
	operands := sa.Tape[sa.pointer+1 : sa.pointer+2]
	sa.forIterations.PushFront(int(operands[0] - 1))
	sa.pointer += 2
	sa.forPointers.PushFront(sa.pointer)
}

func next_loop(sa *SA2) {
	if sa.forIterations.At(0).(int) > 0 {
		first := sa.forIterations.Front().(int) - 1
		sa.forIterations.Set(0, first)
		sa.pointer = sa.forPointers.Front().(int)
	} else {
		sa.forIterations.PopFront()
		sa.forPointers.PopFront()
		sa.pointer += 1
	}
}

func bcc(sa *SA2) {
	operands := sa.Tape[sa.pointer+1 : sa.pointer+2]
	skipCount := operands[0] + 2
	if sa.carryFlag == 0 {
		sa.pointer += int(skipCount)
	} else {
		sa.pointer += 2
	}
}

func bra(sa *SA2) {
	operands := sa.Tape[sa.pointer+1 : sa.pointer+2]
	skipCount := operands[0] + 2
	sa.pointer += int(skipCount)
}

func finish(sa *SA2) {
	sa.pointer += 1
}
