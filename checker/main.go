package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: checker <path-to-fifo>\n")
		os.Exit(1)
	}

	fifoPath := os.Args[1]

	fmt.Printf("opening fifo for write: path=%s\n", fifoPath)

	// opening a fifo in non-blocking write only mode will
	// fail if there is no reader attached on the other end
	f, err := os.OpenFile(fifoPath, os.O_WRONLY|syscall.O_NONBLOCK, 0666)
	defer f.Close()

	if err != nil {
		fmt.Printf("RUNNER PROCESS IS DEAD, err: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("RUNNER PROCESS IS **ALIVE**\n")
	os.Exit(0)
}
