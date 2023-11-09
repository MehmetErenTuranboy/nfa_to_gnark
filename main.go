package main

import (
	"fmt"

	"github.com/MehmetErenTuranboy/nfa_to_gnark/to_gnark"
)

func main() {
	to_gnark.ConstructCircuit("abcd")
	fmt.Println("%b", to_gnark.WildcardDetected())

}
