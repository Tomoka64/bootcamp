package main

import (
	"os"
)

func main() {
	e, err := os.Create("hello.txt")
	if err != nil {
		panic(err)
	}
	defer e.Close()

	//io.Copy(e, strings.NewReader("hello world"))
	e.Write([]byte("hello world"))

}
