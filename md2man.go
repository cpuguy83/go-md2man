package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cpuguy83/go-md2man/v2/md2man"
)

var (
	inFilePath  = flag.String("in", "", "Path to file to be processed (default: stdin)")
	outFilePath = flag.String("out", "", "Path to output processed file (default: stdout)")
)

func main() {
	retcode := 0
	var err error
	flag.Parse()
	inFile := os.Stdin

	defer func() {
		inFile.Close()
		os.Exit(retcode)
	}()

	if *inFilePath != "" {
		inFile, err = os.Open(*inFilePath)
		if err != nil {
			fmt.Println(err)
			retcode = 1
			return
		}
	}

	doc, err := ioutil.ReadAll(inFile)
	if err != nil {
		fmt.Println(err)
		retcode = 1
		return
	}

	out := md2man.Render(doc)

	outFile := os.Stdout
	if *outFilePath != "" {
		outFile, err = os.Create(*outFilePath)
		if err != nil {
			fmt.Println(err)
			retcode = 1
			return
		}
	}
	_, err = outFile.Write(out)
	if err != nil {
		fmt.Println(err)
		retcode = 1
		return
	}
}
