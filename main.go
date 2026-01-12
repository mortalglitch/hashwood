package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting up and checking example directory")

	files, err := os.ReadDir("./testing")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		f, err := os.Open("./testing/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s - %x\n", f.Name(), h.Sum(nil))
	}
}
