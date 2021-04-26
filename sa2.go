package sa2

import (
	"github.com/gammazero/deque"
)

// SA2 represents the VAG SA2 seed/key opcode algorithm
type SA2 struct {
	pointer int
	opcode  []byte

	forIterations deque.Deque
	forPointers   deque.Deque

	carryFlag int
	register  int
}

func New(opcode []byte) *SA2 {
	return &SA2{opcode: opcode}
}

// Execute executes the opcode on the seed
func (sa *SA2) Execute(seed int) int {
	sa.register = seed
	opcodeFunctions := getOpcodeFunctionSet()

	for sa.pointer < len(sa.opcode) {
		opcodeFunctions[sa.opcode[sa.pointer]](sa)
	}

	return int(sa.register)
}
