package main

import (
	"encoding/json"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"golang.org/x/crypto/sha3"
)

type Function struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Inputs     []Input  `json:"inputs"`
	Outputs    []Output `json:"outputs"`
	Payable    bool     `json:"payable"`
	StateMutability string `json:"stateMutability"`
}

type Input struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Output struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func main() {
	// Read the ABI file
	abiData, err := ioutil.ReadFile("./abi/Counter.abi")
	if err != nil {
		log.Fatal(err)
	}

	// Define a slice to store the function selectors
	var selectors []string

	// Unmarshal the ABI data into a slice of Function structs
	var functions []Function
	err = json.Unmarshal(abiData, &functions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over the functions and extract the function selectors
	for _, function := range functions {
		// Calculate the function selector
		selector := calculateSelector(function.Name, function.Inputs)

		// Append the selector to the slice
		selectors = append(selectors, selector)
	}

	// Print the function selectors
	for _, selector := range selectors {
		fmt.Println(selector)
	}
}

// Function to calculate the function selector
func calculateSelector(name string, inputs []Input) string {
	// Concatenate the function signature
	signature := name + "("
	for i, input := range inputs {
		if i > 0 {
			signature += ","
		}
		signature += input.Type
	}
	signature += ")"

	// Calculate the keccak256 hash
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(signature))
	hashed := hash.Sum(nil)

	// Extract the first 4 bytes as the function selector
	selector := hex.EncodeToString(hashed[:4])

	return selector
}
