package main

import (
	"fmt"
	"time"

	"github.com/goombaio/namegenerator"
)

func main() {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	name := nameGenerator.Generate()

	fmt.Println(name)
}
