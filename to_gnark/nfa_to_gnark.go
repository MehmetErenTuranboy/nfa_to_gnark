package to_gnark

import (
	"log"
	"os"
)

func WildcardDetected() bool {
	// Content to write to the new file
	content := `package main


import "fmt"

func curcuit() {
	fmt.Println("This is a simple circuit.")
}
`

	// Use os.WriteFile to create the circuit.go file
	err := os.WriteFile("circuit.go", []byte(content), 0644)
	if err != nil {
		log.Fatalf("Failed to write to circuit.go: %s", err)
	}
	return true
}
