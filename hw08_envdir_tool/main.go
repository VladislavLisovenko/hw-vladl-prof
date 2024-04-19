package main

import "fmt"

func main() {
	env, err := ReadDir("testdata\\env")
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd := []string{"test", "a1", "a2"}
	resp := RunCmd(cmd, env)
	fmt.Println("resp", resp)
}
