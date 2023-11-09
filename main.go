package main

import (
	"fmt"
	"nfa2regex/to_gnark"

	"github.com/MehmetErenTuranboy/thompsons-construction_golang/tools"
)

func main() {

	input := "abcd"
	fmt.Println("Before addConcatOperators:", input)
	input = tools.AddConcatOperators(input)
	postfixVal := tools.InfixToPostfix(input)

	fmt.Println("Postfix: ", postfixVal) // Changed from 'input' to 'postfixVal'

	automataRes := tools.Compile(postfixVal)
	visited := make(map[*tools.State]bool)

	tools.PrintTransition(automataRes.InitialState, visited)
	fmt.Println("%b", to_gnark.WildcardDetected())

}