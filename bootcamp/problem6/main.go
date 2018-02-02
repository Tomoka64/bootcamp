//create wordCount command

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func WordCount(rdr io.Reader) map[string]int {
	counts := map[string]int{}
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		counts[word]++
	}
	return counts
}

func main() {
	file := os.Args[1]
	srcFile, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()

	counts := WordCount(srcFile)
	a := os.Args[2]

	fmt.Printf("Number of %s:", a)
	fmt.Println(counts[a])
}
