/* Implementation of a rudimentary blockchain in Go */

package main

import (
	"fmt"
	"os"
	"bufio"
	"Blockchain"
)

func main () {
	// Create a new blockchain
	var blgocks []Blockchain.Block
	blgockchain := Blockchain.Blockchain{ 
		Chain: blgocks,
		Difficulty: 3 }

	// Create first block on the blockchain
	blgockchain = blgockchain.CreateGenesisBlock()
	blgockchain.ShowChain()

	// Handle user input from commandline
	var transactionData string
	scanner := bufio.NewScanner(os.Stdin)

	// Loop until user keys 'exit'
	for {
		// Get user's data to add to the blockchain
		fmt.Printf("\nEnter transaction data to add to the blockchain:")
		scanner.Scan()
		transactionData = scanner.Text()		

        if transactionData == "exit" {
			os.Exit(0)
		} else {
			// Create and add new block to the blockchain
			blgockchain = blgockchain.CreateBlock(transactionData)
			
			// Display current transactions on the blockchain and verify it's integrity
			fmt.Printf("\nTransaction successfully added to the blockchain:\n")
			blgockchain.ShowChain()
			fmt.Printf("\nBlockchain integrity check successful: %t\n", blgockchain.VerifyIntegrity())

		}
	}
}