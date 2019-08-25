package main

import (
	_ "github.com/joke111/go-in-action/chapter2/matchers"
	"github.com/joke111/go-in-action/chapter2/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("Trumpet")
}