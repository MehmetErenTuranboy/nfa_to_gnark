package to_gnark

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/MehmetErenTuranboy/thompsons-construction_golang/tools"
)

var content string

func ConstructCircuit(regextext string) {

	fmt.Println("Before addConcatOperators:", regextext)
	WildcardDetected(regextext)
	regextext = tools.AddConcatOperators(regextext)
	postfixVal := tools.InfixToPostfix(regextext)

	fmt.Println("Postfix: ", postfixVal) // Changed from 'input' to 'postfixVal'

	automataRes := tools.Compile(postfixVal)
	visited := make(map[*tools.State]bool)

	tools.PrintTransition(automataRes.InitialState, visited)

}

func WildcardDetected(regextext string) {

	// Content to write to the new file
	content = `package main

	import (
		"fmt"
		"log"
		"math/big"
	
		"github.com/consensys/gnark-crypto/ecc"
		"github.com/consensys/gnark/backend/groth16"
		"github.com/consensys/gnark/frontend"
		"github.com/consensys/gnark/frontend/cs/r1cs"
	)
	
	const (
		stringLength    = 10
		substringLenght = `
	content += strconv.Itoa(len(regextext))
	content += `
	)
	
	type charEqualityCircuit struct {
		A [stringLength]frontend.Variable    ` + "`gnark:\",secret\"`" + `
		B [substringLenght]frontend.Variable ` + "`gnark:\",public\"`" + `
	}
	
	func (circuit *charEqualityCircuit) Define(api frontend.API) error {
		matchedFront := frontend.Variable(1)
		result := frontend.Variable(0)
		regexSize := frontend.Variable(substringLenght)
		pivotA := 0
		// Initialize a variable to store if any comparison was successful
		for i := 0; i < len(circuit.A); i++ {
			pivotA = i
			for j := 0; j < len(circuit.B); j++ {
				diff := api.Sub(circuit.A[pivotA], circuit.B[j])
				isEqual := api.IsZero(diff)
				flag := api.IsZero(api.Sub(regexSize, result))
				matchedFront = api.Select(flag, 0, matchedFront)
				result = api.Select(api.Or(isEqual, flag), api.Add(result, matchedFront), 0)
				api.Println("Matching result ", result, circuit.A[pivotA], circuit.B[j])
				if pivotA < len(circuit.A)-1 {
					pivotA++
				} else {
					break
				}
			}
		}
	
		api.AssertIsEqual(result, regexSize)
	
		return nil
	}
	`

	content += `func main() {
	var circuit charEqualityCircuit

	// Secret values
	a := make([]*big.Int, stringLength)
	inputString := "HELLOWORLD"
	for i, char := range inputString {
		if i < stringLength {
			a[i] = big.NewInt(int64(char))
		}
	}

	b := make([]*big.Int, substringLenght)
	regexPattern := `
	content += "\"" + regextext + "\""
	content += `
	for i, char := range regexPattern {
		if i < substringLenght {
			b[i] = big.NewInt(int64(char))
		}
	}

	// Compile the circuit into a set of constraints
	ccs, err := frontend.Compile(ecc.BN254, r1cs.NewBuilder, &circuit, frontend.IgnoreUnconstrainedInputs())
	if err != nil {
		log.Fatalf("Failed to compile the circuit: %v", err)
	}

	// Setup the Proving and Verifying keys
	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		log.Fatalf("Failed to setup the proving and verifying keys: %v", err)
	}

	assignment := charEqualityCircuit{
		A: [stringLength]frontend.Variable{},
		B: [substringLenght]frontend.Variable{},
	}

	for i := 0; i < stringLength; i++ {
		assignment.A[i] = frontend.Variable(a[i])
	}

	for i := 0; i < substringLenght; i++ {
		assignment.B[i] = frontend.Variable(b[i])
	}

	// Create a witness from the assignment
	witness, err := frontend.NewWitness(&assignment, ecc.BN254)
	if err != nil {
		log.Fatalf("Failed to create a witness: %v", err)
	}

	// Extract the public part of the witness
	publicWitness, err := witness.Public()
	if err != nil {
		log.Fatalf("Failed to extract the public witness: %v", err)
	}

	// Prove the witness
	proof, err := groth16.Prove(ccs, pk, witness)
	if err != nil {
		log.Fatalf("Failed to prove the witness: %v", err)
	}

	// Verify the proof
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		fmt.Println("Verification Result: Failed")
		log.Fatalf("Failed to verify the proof: %v", err)
	} else {
		fmt.Println("Verification Result: Success")
	}
}`

	// Use os.WriteFile to create the circuit.go file
	err := os.WriteFile("output/circuit.go", []byte(content), 0644)
	if err != nil {
		log.Fatalf("Failed to write to circuit.go: %s", err)
	}

}
