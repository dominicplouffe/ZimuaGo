package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/pkg/profile"
)

var wrt = bufio.NewWriter(os.Stdout)

func main() {

	args := os.Args[1:]

	if len(args) > 1 && args[1] == "-profile" {
		defer profile.Start().Stop()
	}
	rand.Seed(time.Now().UnixNano())

	f, err := os.OpenFile("zimua.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	if len(args) >= 1 && args[0] == "-uci" {
		xBoard()
	} else if len(args) >= 1 && args[0] == "-cpu" {
		computerVSComputer()
	} else if len(args) >= 1 && args[0] == "-human" {
		computerVSHuman()
	} else {
		// fmt.Println("Usage: ./engine.go [-uci|-cpu|-human] [-profile]")
		computerVSComputer()
	}

}
