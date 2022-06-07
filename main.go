package main

import (
	"github.com/akvlad/flog/generator"
	"math/rand"
	"time"

	"github.com/mingrammer/cfmt"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	opts := ParseOptions()
	if err := generator.Run(opts); err != nil {
		cfmt.Warningln(err.Error())
	}
}
