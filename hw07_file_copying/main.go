package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func Usage() {
	fmt.Println("Usage:")
	fmt.Println(os.Args[0], " -from <source-path> -to <dest-path> [-limit N] [-offset M] - copy file content from <source-path> to <dest-path> skipping first M bytes, limiting write by N bytes")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if from == "" || to == "" {
		Usage()
		return
	}

	err := Copy(from, to, offset, limit)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Done!")
}
