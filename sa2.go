package sa2

import (
	"bytes"
	"encoding/binary"
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

	// run opcode
	opcodeFunctions := getOpcodeFunctionSet()
	for sa.pointer < len(sa.opcode) {
		opcodeFunctions[sa.opcode[sa.pointer]](sa)
	}

	// reset pointer so we can re-execute it
	sa.pointer = 0

	return sa.register
}

// Executes runs the opcode on a given byte slice
func (sa *SA2) ExecuteByteSeed(seed []byte) (int, error) {
	var seedInt int32
	buf := bytes.NewBuffer(seed)
	if err := binary.Read(buf, binary.BigEndian, &seedInt); err != nil {
		return 0, err
	}
	return sa.Execute(int(seedInt)), nil
}
