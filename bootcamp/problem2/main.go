//problem2: Read "hello world" from the file from problem1
//and write it to stdout.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("../problem1/hello.txt")
	if err != nil {
		fmt.Errorf("could not open the file: %v", err)
	}
	a, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Errorf("could not read the file: %v", err)
	}
	fmt.Printf("%s\n", a)
}
