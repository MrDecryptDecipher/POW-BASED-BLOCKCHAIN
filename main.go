package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"
)

//Used to store blocks to connect to a blockchain
var blockchain []block

//mining difficulty value
var diffNum uint = 17

type block struct {
	//Hash of the previous block
	Lasthash string
	//Hash of this block
	Hash string
	//The data stored in the block (for example, the Bitcoin UTXO model can be used to store transactions here)
	Data string
	//time stamp
	Timestamp string
	// block height
	Height int
	//difficulty value
	DiffNum uint
	//random number
	Nonce int64
}

//Block mining (calculate hash by incrementing the nonce value by itself)
func mine(data string) block {
	if len(blockchain) < 1 {
		log.Panic("还未生成创世区块！")
	}
	lastBlock := blockchain[len(blockchain)-1]
	// create a new block
	newBlock := new(block)
	newBlock.Lasthash = lastBlock.Hash
	newBlock.Timestamp = time.Now().String()
	newBlock.Height = lastBlock.Height + 1
	newBlock.DiffNum = diffNum
	newBlock.Data = data
	var nonce int64 = 0
	//A large number calculated according to the mining difficulty value
	newBigint := big.NewInt(1)
	newBigint.Lsh(newBigint, 256-diffNum) //Equivalent to left shift 1<<256-diffNum
	for {
		newBlock.Nonce = nonce
		newBlock.getHash()
		hashInt := big.Int{}
		hashBytes, _ := hex.DecodeString(newBlock.Hash)
		hashInt.SetBytes(hashBytes) //Convert the hash value of this block into a string of numbers
		//If the hash is less than a large number calculated by the mining difficulty value, it means that the mining is successful
		if hashInt.Cmp(newBigint) == -1 {
			break
		} else {
			nonce++ // If the condition is not met, the random number will be incremented continuously until the hash value of this block is less than the specified large number
		}
	}
	return *newBlock
}

func (b *block) serialize() []byte {
	bytes, err := json.Marshal(b)
	if err != nil {
		log.Panic(err)
	}
	return bytes
}

func (b *block) getHash() {
	result := sha256.Sum256(b.serialize())
	b.Hash = hex.EncodeToString(result[:])
}

func main() {
	// Create a genesis block
	genesisBlock := new(block)
	genesisBlock.Timestamp = time.Now().String()
	genesisBlock.Data = "I am the genesis block！"
	genesisBlock.Lasthash = "0000000000000000000000000000000000000000000000000000000000000000"
	genesisBlock.Height = 1
	genesisBlock.Nonce = 0
	genesisBlock.DiffNum = 0
	genesisBlock.getHash()
	fmt.Println(*genesisBlock)
	//Add the genesis block to the blockchain
	blockchain = append(blockchain, *genesisBlock)
	for i := 0; i < 10; i++ {
		newBlock := mine("Sandeep made this blockchain" + strconv.Itoa(i))
		blockchain = append(blockchain, newBlock)
		fmt.Println(newBlock)
	}
}
