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

// New creates a new SA2 opcode executor
func New(opcode []byte) *SA2 {
	return &SA2{opcode: opcode}
}

// Execute executes the opcode on the give seed
func (sa *SA2) Execute(seed int) int {
	sa.register = seed
	opcodeFunctions := getOpcodeFunctionSet()

	for sa.pointer < len(sa.opcode) {
		opcodeFunctions[sa.opcode[sa.pointer]](sa)
	}

	return sa.register
}
