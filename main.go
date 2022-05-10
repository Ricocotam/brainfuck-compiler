package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	buffer_size          = 66000
	max_nested_loop      = 10
	JUMP_LINE       byte = '\n'
	INC_POINTER     byte = '>'
	DEC_POINTER     byte = '<'
	INC_VALUE       byte = '+'
	DEC_VALUE       byte = '-'
	START_LOOP      byte = '['
	END_LOOP        byte = ']'
	GETCHAR         byte = ','
	PUTCHAR         byte = '.'
)

var (
	reader = bufio.NewReader(os.Stdin)
)

func main() {
	fileToCompile := os.Args[1]

	data, err := os.ReadFile(fileToCompile)
	if err != nil {
		log.Panic(err)
	}

	memory := new(Memory)
	interpret(data, memory)
	log.Println(memory.buffer)
}

type Memory struct {
	buffer    [buffer_size]byte
	pos       int
	loops     [max_nested_loop]Loop
	loopLevel int // Number of loops. Value 0 means no loop
}

type Loop struct {
	startPos int
}

// > = increases memory pointer, or moves the pointer to the right 1 block.
// < = decreases memory pointer, or moves the pointer to the left 1 block.
// + = increases value stored at the block pointed to by the memory pointer
// - = decreases value stored at the block pointed to by the memory pointer
// [ = like c while(cur_block_value != 0) loop.
// ] = if block currently pointed to's value is not zero, jump back to [
// , = like c getchar(). input 1 character.
// . = like c putchar(). print 1 character to the console
func interpret(fileContent []byte, memory *Memory) {
	var r byte
	for i := 0; i < len(fileContent); i++ {
		r = fileContent[i]
		switch r {
		case INC_POINTER:
			memory.inc_pointer()
		case DEC_POINTER:
			memory.dec_pointer()
		case INC_VALUE:
			memory.inc_value()
		case DEC_VALUE:
			memory.dec_value()
		case GETCHAR:
			memory.get_char()
		case PUTCHAR:
			memory.put_char()
		case START_LOOP:
			memory.start_loop(i)
		case END_LOOP:
			i = memory.end_loop(i)
		case JUMP_LINE:
			continue
		default:
			log.Println("Wrong rune", r)
		}
	}
	fmt.Println()
}

func (memory *Memory) inc_pointer() {
	memory.pos += 1
}

func (memory *Memory) dec_pointer() {
	memory.pos -= 1
}

func (memory *Memory) inc_value() {
	memory.buffer[memory.pos] += 1
}

func (memory *Memory) dec_value() {
	memory.buffer[memory.pos] -= 1
}

func (memory *Memory) get_char() {
	input, err := reader.ReadByte()
	if err != nil {
		log.Panic(err)
	}
	memory.buffer[memory.pos] = input
}

func (memory *Memory) put_char() {
	fmt.Printf("%s", string(memory.buffer[memory.pos]))
}

func (memory *Memory) start_loop(readingPos int) {
	memory.loopLevel += 1
	memory.loops[memory.loopLevel-1] = Loop{readingPos}
}

func (memory *Memory) end_loop(readingPos int) int {
	loop := memory.loops[memory.loopLevel-1]
	memory.loopLevel -= 1

	if memory.buffer[memory.pos] == 0 {
		return readingPos
	}
	return loop.startPos - 1 // Decrease one time because the loop will do ++
}
