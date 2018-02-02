//create a program which converts all the text in a file to
//uppercase and writes it to Stdout
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func convertion(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("could not open: %v", err)
	}
	defer f.Close()
	a, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	b := strings.ToUpper(string(a))

	fmt.Printf("%s\n", b)
}
func main() {
	filename := os.Args[1]
	convertion(filename)
}
