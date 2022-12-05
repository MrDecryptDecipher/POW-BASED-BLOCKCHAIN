SIMPLE PROOF OF WORK BASED BLOCKCHAIN WHICH ALSO PRINTS DATA.
<br>



>The core of the proof-of-work mechanism is to constantly hash the block itself, and compare the hash value with a series of large numbers calculated according to the difficulty value. If the self-hash is smaller than the large number, it means that the mining is successful, otherwise change its own random number and recalculate. And the program will dynamically adjust the difficulty value according to the block interval time (such as Bitcoin)

<br>

block structure
```go
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
```
Mining function:
Use the math/big package to calculate a series of large numbers newBigint for actual comparison according to the difficulty value diffNum of the global variable, and at the same time convert the block hash into a large number hashInt and compare the two large numbers numerically. If hashInt is smaller than newBigint then Represents successful mining

```go
//Block mining (calculate hash by incrementing the nonce value by itself)
func mine(data string) block {
	if len(blockchain) < 1 {
		log.Panic("The genesis block has not been generated yet！")
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
```

```go
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
		newBlock := mine("nice weather"+strconv.Itoa(i))
		blockchain = append(blockchain, newBlock)
		fmt.Println(newBlock)
	}
```



