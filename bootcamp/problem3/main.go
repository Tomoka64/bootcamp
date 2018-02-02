//create your own cp command

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func cp(filename1, filename2 string) error {
	srcFilename, err := os.Open(filename1)
	if err != nil {
		fmt.Errorf("could not open the file: %v", err)
	}
	defer srcFilename.Close()

	newFilename, err := os.Create(filename2)
	if err != nil {
		fmt.Errorf("could not create a file: %v", err)
	}
	defer newFilename.Close()

	_, err = io.Copy(newFilename, srcFilename)
	if err != nil {
		fmt.Errorf("could not copy: %v", err)
	}
	return nil
}
func main() {
	srcFileName := os.Args[1]
	newFileName := os.Args[2]
	cp(srcFileName, newFileName)
	if len(os.Args) > 3 {
		log.Fatalln("usage: problem3 <the name of the file you want to copy> <the name of the file you want to create>")
	}

}
