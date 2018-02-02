//generalize #2 and create your own version of cat which reads
//a file and dumps it to stdout.

package main

import (
	"io"
	"log"
	"os"
)

func cat(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = io.Copy(os.Stdout, f)
	if err != nil {
		log.Fatalln("usage: problem4 <the name of the file you want to cat> ")
	}
	return nil
}

func main() {
	filename := os.Args[1]
	cat(filename)

}
