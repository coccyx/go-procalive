package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: runner <path-to-fifo> <timeout>\n")
		os.Exit(1)
	}

	fifoPath := os.Args[1]
	err := syscall.Mkfifo(fifoPath, 0666)
	if err != nil {
		fmt.Printf("error creating fifo: %v", err)
		os.Exit(1)
	}
	defer os.Remove(fifoPath)

	fmt.Printf("opening fifo for read: path=%s\n", fifoPath)
	// opening a fifo in read-only non-blocking mode will
	// succeed even if noone is attached on the other end
	// http://man7.org/linux/man-pages/man7/fifo.7.html
	f, err := os.OpenFile(fifoPath, os.O_RDONLY|syscall.O_NONBLOCK, 0666)
	defer f.Close()
	if err != nil {
		fmt.Printf("error opening fifo: path=%s, err: %v", fifoPath, err)
		os.Exit(1)
	}

	fmt.Printf("opened fifo for read: path=%s\n", fifoPath)

	sleepSecs, _ := strconv.Atoi(os.Args[2])
	if sleepSecs < 0 {
		sleepSecs = 120
	}

	fmt.Printf("sleeping for %d seconds ...\n", sleepSecs)
	time.Sleep(time.Duration(sleepSecs) * time.Second)
	fmt.Printf("cleaning and exiting")
}
