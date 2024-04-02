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
	flag.StringVar(&from, "from", "testdata/largefile.txt", "file to read from")
	flag.StringVar(&to, "to", "sampleoutput.txt", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()
	err := Copy(from, to, offset, limit)
	if err != nil {
		fmt.Println(err)
	}
}
