package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	buffer_size = 256
)

var (
	reader = bufio.NewReader(os.Stdin)
)

func main() {
	cmdName := os.Args
	fileToCompile := os.Args[1]
	memory := new(Memory)
	interpret(fileToCompile, memory)
}

type Memory struct {
	buffer [buffer_size]byte
	pos    int
}

// > = increases memory pointer, or moves the pointer to the right 1 block.
// < = decreases memory pointer, or moves the pointer to the left 1 block.
// + = increases value stored at the block pointed to by the memory pointer
// - = decreases value stored at the block pointed to by the memory pointer
// [ = like c while(cur_block_value != 0) loop.
// ] = if block currently pointed to's value is not zero, jump back to [
// , = like c getchar(). input 1 character.
// . = like c putchar(). print 1 character to the console
func interpret(fileContent string, memory *Memory) {

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
	input, _, err := reader.ReadByte()
	if err != nil {
		log.panic(err)
	}
	memory.buffer[memory.pos] = input
}

func (memory *Memory) show_char() {
	fmt.Printf("%c", memory.buffer[memory.pos])
}

func (memory *Memory) start_loop() {

}

func (memory *Memory) end_loop() {

}
