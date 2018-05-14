package main

import (
	"github.com/ynozue/lib_sample/subpackage_A"
	"github.com/ynozue/lib_sample/subpackage_B"
)

func main() {
	a := subpackage_A.SubA{Name: "hoge"}
	a.PrintA()
	b := subpackage_B.SubB{Name: "hoge"}
	b.PrintA()
}
