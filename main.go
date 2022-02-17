package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()

	return ch
}

func main() {
	bc := context.Background()
	t := 10 * time.Second
	ctx, cancel := context.WithTimeout(bc, t)
	ch := input(os.Stdin)
	defer cancel()

	for {
		select {
		case v := <-ch:
			fmt.Println(">", v)
		case <-ctx.Done():
			fmt.Println("Time UP!")
			return
		}
	}

}
