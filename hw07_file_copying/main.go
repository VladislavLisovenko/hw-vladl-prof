package main

import (
	"flag"
	"fmt"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "testdata/input.txt", "file to read from")
	flag.StringVar(&to, "to", "sampleoutput.txt", "file to write to")
	flag.Int64Var(&limit, "limit", -600, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", -1000, "offset in input file")
}

func main() {
	flag.Parse()
	err := Copy(from, to, offset, limit)
	if err != nil {
		fmt.Println(err)
	}
}
