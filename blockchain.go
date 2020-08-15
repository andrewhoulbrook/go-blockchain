/* Implementation of a rudimentary blockchain in Go */

package Blockchain

import (
	"fmt"
	"time"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

// Blockchain declared as slice of 'block' type structs
type Blockchain struct {
	Chain []Block
	Difficulty int
}

// Blocks on the blockchain declared as structs
type Block struct {
	index int
	timestamp int64
	data string
	precedingHash string
	hash string
}

// Function to compute a SHA256 cryptographic hash
func (b Block) computeHash(nonce int64) Block {
	var blockString string
	if nonce == 0 {
		blockString = fmt.Sprintf("%d%d%s%d%d", b.index, b.timestamp, b.data, b.precedingHash)
	} else {
		blockString = fmt.Sprintf("%s%d", b.hash, nonce)
	}
	hashed := sha256.Sum256([]byte(blockString))
	b.hash = fmt.Sprintf("%x", hashed)
	return b
}

// Function to create new block for the blockchain
func (bc Blockchain) CreateBlock(blockData string) Blockchain {
	// Complete Proof of Work before adding new block	
	lb := bc.GetLastBlock()
	lb.ProofOfWork()

	// Build new block
	b := Block{
		index: lb.index + 1,
		timestamp: time.Now().UnixNano(),
		data: blockData,
		precedingHash: lb.hash,
		}
	return bc.AddBlock(b)
}

// Function to create the Blockchain's initial block
func (bc Blockchain) CreateGenesisBlock() Blockchain {
	gb := Block{
		index: 0,
		timestamp: time.Now().UnixNano(),
		data: "Welcome to the Blgockchain!",
		precedingHash: "0",
		}
	return bc.AddBlock(gb)
}
		
// Function to get last block added to the blockchain
func (bc Blockchain) GetLastBlock() Block {
	return bc.Chain[len(bc.Chain) - 1]
}

// Function to add a new block to the blockchain
func (bc Blockchain) AddBlock(b Block) Blockchain {
	bHashed := b.computeHash(0) 
	bReHashed := bHashed.computeHash(createNonce())

	// Create string of zeros for pattern matching
	var prefixPattern string
	for i := 0; i < bc.Difficulty; i++ {
		prefixPattern = fmt.Sprintf("%s%s", prefixPattern, "0")
	} 

	// Loop and recompute hash until leading digits match pattern
	for { 
		if bReHashed.hash[:bc.Difficulty] != prefixPattern {
			bReHashed = bHashed.computeHash(createNonce())
		} else {
			break
		}
	}
	bc.Chain = append(bc.Chain, bReHashed)
	return bc
}

// Function to generate a 6-digit random number
func createNonce() int64 {
	rnd, _ := rand.Int(rand.Reader, big.NewInt(999999 - 100000))
	return rnd.Int64() + 100000
}

// Function to add proof of work mechanism to the blockchain
func (b Block) ProofOfWork() {
	lb := b
	findNonce := int64(100000)
	lbHashed := lb.computeHash(0)
	lbRehashed := lbHashed.computeHash(findNonce)

	// Loop nonce values and recompute hash until it matches the last block's 
	for {
		lbRehashed = lbHashed.computeHash(findNonce)
		if lbRehashed.hash != b.hash {
			findNonce++
		} else {
			return
		} 
	}
}

// Trival function to check integrity of hashes and timestamps for each block in the chain
func (bc Blockchain) VerifyIntegrity() bool {
	for i := 1; i < len(bc.Chain); i++ {
		if bc.Chain[i].precedingHash != bc.Chain[i-1].hash {
			if bc.Chain[i].timestamp < bc.Chain[i-1].timestamp {
				return false
			}
		}
	}
	return true
}

// Print out all blocks in the blockchain
func (bc Blockchain) ShowChain() {
	for i := 0; i < len(bc.Chain); i++ {
		fmt.Printf("-----------------------------------------------------------------------------------\n")
		fmt.Printf("| Block Index: %d\n", bc.Chain[i].index)
		fmt.Printf("| Timestamp: %d\n", bc.Chain[i].timestamp)
		fmt.Printf("| Data: %s\n", bc.Chain[i].data)
		fmt.Printf("| Hash: %s\n", bc.Chain[i].hash)
		fmt.Printf("| Preceding Hash: %s\n", bc.Chain[i].precedingHash)
		fmt.Printf("-----------------------------------------------------------------------------------\n")
		fmt.Printf("                                        ||                                         \n")
	}
}
