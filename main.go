package main

import (
	"fmt"
	"log"

	md5utils "github.com/mortalglitch/hashwood/internal/md5_utils"
)

func main() {
	fmt.Println("Starting up and checking example directory")

	hashResults, err := md5utils.ParseDirectory("./testing/")
	if err != nil {
		log.Fatal(err)
	}

	for _, hash := range hashResults {
		fmt.Printf("%s - %x\n", hash.Filename, hash.Hash)
	}
}
