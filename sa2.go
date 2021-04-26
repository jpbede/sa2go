package sa2

import (
	"github.com/gammazero/deque"
)

type SA2 struct {
	pointer int
	Tape    []byte

	forIterations deque.Deque
	forPointers   deque.Deque

	carryFlag int
	register  int
}

func New(tape []byte, seed int) *SA2 {
	return &SA2{Tape: tape, register: seed}
}

func (sa *SA2) Execute() int {
	opcodeFunctions := getOpcodeFunctionSet()

	for sa.pointer < len(sa.Tape) {
		opcodeFunctions[sa.Tape[sa.pointer]](sa)
	}

	return int(sa.register)
}
