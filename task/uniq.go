package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type Options struct {
	countStr           bool
	onlyRepeated       bool
	onlyUnique         bool
	ignoreFirstNFields int
	ignoreFirstNChars  int
	ignoreCase         bool
}

func init() {
	var options Options
	flag.BoolVar(&options.countStr, "c", false, "count repeated string")
	flag.BoolVar(&options.onlyRepeated, "d", false, "print only repeated string")
	flag.BoolVar(&options.onlyUnique, "u", false, "print only unique string")
	flag.IntVar(&options.ignoreFirstNFields, "f", 0, "ignore the first num fields")
	flag.IntVar(&options.ignoreFirstNChars, "s", 0, "ignore the first num chars")
	flag.BoolVar(&options.ignoreCase, "i", false, "ignore case")
}

func read(r io.Reader) (content []string, err error) {
	buf := make([]byte, 60)
	for {
		if _, err := r.Read(buf); err == io.EOF {
			break
		} else {
			content = append(content, string(buf))
		}
	}
	return content, nil
}

//func write(writer io.Writer) error {
//
//}

func main() {
	flag.Parse()
	var content []string
	var reader io.Reader
	//var writer io.Writer
	if inputFile := flag.Arg(0); inputFile != "" {
		if _, err := os.Open(inputFile); err != nil {
			fmt.Errorf("File %s is not found", inputFile)
			os.Exit(-1)
		}
	} else {
		reader = os.Stdin
	}
	content, status := read(reader)
	fmt.Println(content)
	fmt.Println(status)
}
