package main

import (
	"app/packageone"
)

var myvar = "package level variable"

func main() {
	var blockvar = "block level variable"
	packageone.Printme(blockvar, myvar)
}
