package main

import (
	"github.com/go-programming-tour/tour/cmd"
	"log"
)

func main() {
	err := cmd.Execute()

	if err != nil{
		log.Fatalf("cmd.Execute failed: %v", err)
	}
}
