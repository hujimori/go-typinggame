package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
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

func init() {
	rand.Seed(time.Now().UnixNano())
}

var rs1Lettesr = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Lettesr[rand.Intn(len(rs1Lettesr))]
	}

	return string(b)
}

func askQuestion() string {
	return randString(3)
}

func main() {
	bc := context.Background()
	t := 20 * time.Second
	ctx, cancel := context.WithTimeout(bc, t)
	ch := input(os.Stdin)
	defer cancel()
	correctAns := 0
	for {
		q := askQuestion()
		fmt.Println("Q.", q)
		select {
		case v := <-ch:
			fmt.Println(">", v)
			if v == q {
				correctAns++
				fmt.Println("correct!")
			} else {
				fmt.Println("incorrect!")
			}

		case <-ctx.Done():
			// 入力途中で終了した場合のために改行する
			fmt.Println()
			fmt.Println("TotalCorrect:", correctAns)
			fmt.Println("Time UP!")
			return
		}
	}

}
