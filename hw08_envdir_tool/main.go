package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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

	for k, v := range env {
		fmt.Printf("%s is (%s)\n", k, v.Value)
	}

	if len(os.Args) > 3 {
		fmt.Print("arguments are ", strings.Join(os.Args[4:], " "), "\n")
	}

	cmd := os.Args[2:]
	resp := RunCmd(cmd, env)
	fmt.Println("resp", resp)
}
