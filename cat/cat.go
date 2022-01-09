package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
		return
	}

	for _, val := range flag.Args() {
		f, err := os.Open(val)
		if err != nil {
			log.Print("open file error " + val)
			continue
		}
		cat(bufio.NewReader(f))

	}

}

func cat(r *bufio.Reader) {
	for {
		bytes, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		fmt.Fprintf(os.Stdout, "%s", bytes)
	}
}
