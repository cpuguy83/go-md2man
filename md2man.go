package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cpuguy83/go-md2man/md2man"
)

var Version string

var inFilePath = flag.String("in", "", "Path to file to be processed (default: stdin)")
var outFilePath = flag.String("out", "", "Path to output processed file (default: stdout)")
var version = flag.Bool("v", false, "prints current md2man version")

func main() {
	var err error
	flag.Parse()

	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}

	inFile := os.Stdin
	if *inFilePath != "" {
		inFile, err = os.Open(*inFilePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	defer inFile.Close() // nolint: errcheck

	doc, err := ioutil.ReadAll(inFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(doc) == 0 {
		fmt.Println(os.Stderr, "Input must be non-zero in length")
		os.Exit(1)
	}

	out := md2man.Render(doc)

	outFile := os.Stdout
	if *outFilePath != "" {
		outFile, err = os.Create(*outFilePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer outFile.Close() // nolint: errcheck
	}
	_, err = outFile.Write(out)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
