package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type registers struct {
}

type instruction struct {
	op      int
	a, b, c int
}

var opcodes = map[string]func(in instruction, registers *[4]int){

	"addr": func(in instruction, registers *[4]int) {
		registers[in.c] = registers[in.a] + registers[in.b]
	},

	"addi": func(in instruction, registers *[4]int) {
		registers[in.c] = registers[in.a] + in.b
	},

	"mulr": func(in instruction, registers *[4]int) {
		registers[in.c] = registers[in.a] * registers[in.b]
	},

	"muli": func(in instruction, registers *[4]int) {
		registers[in.c] = registers[in.a] * in.b
	},

	"banr": func(in instruction, registers *[4]int) {
		registers[in.c] = registers[in.a] & registers[in.b]
	},

	"bani": func(in instruction, registers *[4]int) {
		registers[in.c] = registers[in.a] & in.b
	},

	"borr": func(in instruction, registers *[4]int) {
		registers[in.c] = registers[in.a] | registers[in.b]
	},

	"bori": func(in instruction, registers *[4]int) {
		registers[in.c] = registers[in.a] | in.b
	},

	"setr": func(in instruction, registers *[4]int) {
		registers[in.c] = registers[in.a]
	},

	"seti": func(in instruction, registers *[4]int) {
		registers[in.c] = in.a
	},

	"gtir": func(in instruction, registers *[4]int) {
		if in.a > registers[in.b] {
			registers[in.c] = 1
		} else {
			registers[in.c] = 0
		}
	},

	"gtri": func(in instruction, registers *[4]int) {
		if registers[in.a] > in.b {
			registers[in.c] = 1
		} else {
			registers[in.c] = 0
		}
	},

	"gtrr": func(in instruction, registers *[4]int) {
		if registers[in.a] > registers[in.b] {
			registers[in.c] = 1
		} else {
			registers[in.c] = 0
		}
	},

	"eqir": func(in instruction, registers *[4]int) {
		if in.a == registers[in.b] {
			registers[in.c] = 1
		} else {
			registers[in.c] = 0
		}
	},

	"eqri": func(in instruction, registers *[4]int) {
		if registers[in.a] == in.b {
			registers[in.c] = 1
		} else {
			registers[in.c] = 0
		}
	},

	"eqrr": func(in instruction, registers *[4]int) {
		if registers[in.a] == registers[in.b] {
			registers[in.c] = 1
		} else {
			registers[in.c] = 0
		}
	},
}

type sample struct {
	registers [4]int
	in        instruction
	expected  [4]int
}

const (
	STATE_SAMPLES = iota
	STATE_PROGRAM
)

func main() {

	samples := []sample{}
	program := []instruction{}

	var s = sample{}
	var state = STATE_SAMPLES

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		if state == STATE_PROGRAM {

			in := instruction{}

			if _, err := fmt.Sscanf(
				line,
				"%d %d %d %d",
				&in.op, &in.a, &in.b, &in.c,
			); err != nil {
				log.Fatal(err)
			}

			program = append(program, in)

			continue
		}

		if n, _ := fmt.Sscanf(
			line,
			"Before: [%d, %d, %d, %d]",
			&s.registers[0], &s.registers[1], &s.registers[2], &s.registers[3],
		); n > 0 {

			scanner.Scan()

			line = strings.TrimSpace(scanner.Text())

			if _, err := fmt.Sscanf(
				line,
				"%d %d %d %d",
				&s.in.op, &s.in.a, &s.in.b, &s.in.c,
			); err != nil {
				log.Fatal(err)
			}

			scanner.Scan()

			line = strings.TrimSpace(scanner.Text())

			if _, err := fmt.Sscanf(
				line,
				"After: [%d, %d, %d, %d]",
				&s.expected[0], &s.expected[1], &s.expected[2], &s.expected[3],
			); err != nil {
				log.Fatal(err)
			}

			samples = append(samples, s)
			s = sample{}

			scanner.Scan()

		}

		if line == "" {
			scanner.Scan()
			state = STATE_PROGRAM
		}

	}

	// Loop through parsed samples and note which opcodes numbers match
	// the instruction

	matched3 := 0
	opcodesByNum := [16]map[string]bool{}

	for _, s := range samples {

		if opcodesByNum[s.in.op] == nil {
			opcodesByNum[s.in.op] = map[string]bool{}
		}

		matched := 0

		for name, opcode := range opcodes {

			regs := [4]int{}

			copy(regs[:], s.registers[:])

			opcode(s.in, &regs)

			if regs == s.expected {
				matched++
				opcodesByNum[s.in.op][name] = true
			}

		}

		if matched >= 3 {
			matched3++
		}

	}

	fmt.Printf("samples in the input that behave like three or more opcodes: %v\n", matched3)

	// Attempt to find opcode numbers

	opcodeByNum := [16]string{}

	for i := 0; i < 16; i++ {

		for opcode, candidates := range opcodesByNum {

			if len(candidates) > 1 {
				continue
			}

			if len(candidates) == 0 {
				continue
			}

			var name string
			for k := range candidates {
				name = k
			}

			opcodeByNum[opcode] = name

			delete(opcodesByNum[opcode], name)

			for opcode := range opcodesByNum {
				delete(opcodesByNum[opcode], name)
			}

		}

	}

	// Run sample program

	registers := [4]int{}
	for _, in := range program {
		opcodes[opcodeByNum[in.op]](in, &registers)
	}

	fmt.Printf("after executing the test program, register 0 contains: %v\n", registers[0])

}
