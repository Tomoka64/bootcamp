//write a program which uppcases the first letter of teh contents
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"unsafe"
)

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func UpperCase(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	words, err := ioutil.ReadAll(f)
	w := BytesToString(words)
	w = string(w)
	for _, word := range w {
		word = strings.ToUpper(word)
		return word
	}
	if err != nil {
		log.Fatalln("could not uppercase")
	}
	return nil

}

func main() {
	arg := os.Args[1]
	fmt.Println(UpperCase(arg))
}
