package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Must be at least 3 arguments")
		return
	}
	dir := os.Args[1]
	env, err := ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd := os.Args[2:]
	// for i := 0; i < len(cmd); i++ {
	// 	cmd[i] = strings.ReplaceAll(cmd[i], "$(pwd)/", "")
	// }
	os.Exit(RunCmd(cmd, env))
}
