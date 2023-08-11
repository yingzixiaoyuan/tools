package main

import (
	"log"

	"github.com/yingzixiaoyuan/tools/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
