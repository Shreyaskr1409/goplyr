package main

import (
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "goplyr: ", log.LstdFlags)
	l.Println("Logging starts...")
}
