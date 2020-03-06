# A Basic Blockchain in Go

A rudimentary implementation of a blockchain in Go. 

## Introduction

There's certainly plenty of "Introduction to Blockchain" articles out there and blockchain technology explained through various metaphors. The best way to learn is by doing. 

This is my humble attempt to implement a basic blockchain and experiment with the Go language at the same time. 

For a good understanding of cryptographic hash functions, my favourite references [here](https://www.schneier.com/books/applied_cryptography/toc.html) and [here](http://www.unixwiz.net/techtips/iguide-crypto-hashes.html).

## Prerequisites

Following standard Go packages are used:

```
import (
    "time"
    "crypto/sha256"
    "crypto/rand"
    "math/big"
)
```
Import the ```Blockchain``` package contained in ```blockchain.go```. 

## Core Components

A block is defined as a ```struct``` data structure.

```
type Block struct {
    index int
    timestamp int64
    data string
    preceedingHash string
    hash string 
}
```

A blockchain is simply defined as another ```struct``` containing a slice array of blocks forming the **chain** of blocks.

```
type Blockchain struct {
    chain []Block
    difficulty int
}
```

Go's built-in crypto package is used to implement the [SHA-256 hashes](https://golang.org/pkg/crypto/sha256/) of each block in the blockchain.

## Building a Blockchain

Create a new blockchain object and add the first block to the blockchain:

```
var blgocks []Blockchain.Blocks 
blgockchain := Blockchain.Blockchain {
    chain: bglocks,
    difficulty: 3 }

// Create first block on the blockchain
blgockchain = blgockchain.CreateGenesisBlock()
```

Create and add a new block to the blockchain with some given ```transactionData```:

```
newBlock := blgockchain.CreateBlock(transactionData)
```
   			
### Proof of Work

The function ```ProofOfWork()``` and the ```difficulty``` parameter are used to implement the **Proof of Work** concept, important security features within blockchains and part of the original consensus algorithm in a blockchain network.

The ```difficulty``` parameter determines how many leading zeros much exist for a SHA-256 hash to be accepted, thus increasing the cost-complexity of adding a block to the blockchain.

To keep things simple, ```createNonce()``` function uses ```crypto/rand``` to generate six digit random numbers to act as nonce values within each block added to the blockchain.   

The ```ProofOfWork()``` function continues to recompute SHA-256 hashes, testing different possible nonce values, until a match is found with the preceding block's hash. A new block can then be added to the blockchain having successfully gained a Proof of Work. 

## Test Example

Build the Bglockchain with it's initial block and add two new blocks:

```
>$ go run blgockchain.go
...
Enter transaction data to add to the blockchain:Sender: Son, Receiver: Father, Message: agreed!              

Transaction successfully added to the blockchain:
-----------------------------------------------------------------------------------
| Block Index: 0
| Timestamp: 1583163252557868826
| Data: Welcome to the Blgockchain!
| Hash: 000970a802576e20d573e4e32b4d0163798837edf82cbd808e1fc600d8e509c7
| Preceding Hash: 0
-----------------------------------------------------------------------------------
                                        ||                                         
-----------------------------------------------------------------------------------
| Block Index: 1
| Timestamp: 1583163256664408686
| Data: Sender: Father, Receiver: Son, Message: eat vegetables then watch cartoons
| Hash: 000a40612265f0e9ab598d322cbc5579dc39bf7006821f2cb3cbde85e63fca13
| Preceding Hash: 000970a802576e20d573e4e32b4d0163798837edf82cbd808e1fc600d8e509c7
-----------------------------------------------------------------------------------
                                        ||                                         
-----------------------------------------------------------------------------------
| Block Index: 2
| Timestamp: 1583163286493328911
| Data: Sender: Son, Receiver: Father, Message: agreed!
| Hash: 000418f03664ff322e701c3ee4a381496968bb5e30c8c280d0afee8d27ddf1fd
| Preceding Hash: 000a40612265f0e9ab598d322cbc5579dc39bf7006821f2cb3cbde85e63fca13
-----------------------------------------------------------------------------------
                                        ||                                         

Blockchain integrity check successful: true

Enter transaction data to add to the blockchain: exit
```

## Built with

* [Go](https://golang.org)

## Authors

Initial work contributed by Andrew Houlbrook - [andrewhoulbrook](https://github.com/andrewhoulbrook)