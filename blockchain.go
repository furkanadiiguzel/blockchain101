package main

import (
	"crypto/sha256"

	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func (t Transaction) Hash() string {
	data := t.Sender + t.Receiver + fmt.Sprintf("%.2f", t.Amount)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

type Block struct {
	Version         int
	PreviousHash    string
	Difficulty      int
	Timestamp       int64
	MerkleRoot      string
	Nonce           int
	NumTransactions int
	Transactions    []Transaction
	Hash            string
}

type Blockchain struct {
	chain      []Block
	difficulty int
}

type user struct {
	userInput int
	sender    string
	receiver  string
	amount    float64
}

func (b *Block) calculateHash() string {
	blockData := strconv.Itoa(b.Version) + b.PreviousHash + strconv.Itoa(b.Difficulty) + strconv.FormatInt(b.Timestamp, 10) + b.MerkleRoot + strconv.Itoa(b.Nonce) + strconv.Itoa(b.NumTransactions)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine() {
	targetPrefix := strings.Repeat("0", b.Difficulty)
	for !strings.HasPrefix(b.Hash, targetPrefix) {
		b.Nonce++
		b.Hash = b.calculateHash()
	}
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		Version:         1,
		PreviousHash:    "0",
		Difficulty:      difficulty,
		Timestamp:       time.Now().Unix(),
		MerkleRoot:      "",
		Nonce:           0,
		NumTransactions: 0,
		Transactions:    []Transaction{},
	}
	return Blockchain{
		chain:      []Block{genesisBlock},
		difficulty: difficulty,
	}
}

func (b *Blockchain) addBlock(transactions []Transaction) {
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		Version:         1,
		PreviousHash:    lastBlock.Hash,
		Difficulty:      b.difficulty,
		Timestamp:       time.Now().Unix(),
		MerkleRoot:      "",
		Nonce:           0,
		NumTransactions: len(transactions),
		Transactions:    transactions,
	}
	newBlock.mine()
	b.chain = append(b.chain, newBlock)
}

type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
}

func calculateMerkleRoot(transactions []Transaction) string {
	if len(transactions) == 0 {
		return ""
	} else if len(transactions) == 1 {
		return transactions[0].Hash()
	}

	var hashedTransactions []string
	for _, tx := range transactions {
		hashedTransactions = append(hashedTransactions, tx.Hash())
	}

	for len(hashedTransactions) > 1 {
		if len(hashedTransactions)%2 != 0 {
			hashedTransactions = append(hashedTransactions, hashedTransactions[len(hashedTransactions)-1])
		}

		var newHashedTransactions []string
		for i := 0; i < len(hashedTransactions); i += 2 {
			combined := hashedTransactions[i] + hashedTransactions[i+1]
			hash := sha256.Sum256([]byte(combined))
			newHashedTransactions = append(newHashedTransactions, fmt.Sprintf("%x", hash))
		}
		hashedTransactions = newHashedTransactions
	}

	return hashedTransactions[0]
}

func main() {
	var User user

	fmt.Print("Enter Difficulty: ")
	fmt.Scanln(&User.userInput)
	blockchain := CreateBlockchain(User.userInput)

	transactions := []Transaction{
		{"Alice", "Bob", 2.5},
		{"Jack", "Charlie", 1.0},
		{"Charlie", "Furkan", 0.5},
		{"Sefa", "Charlie", 1.2},
		{"Bob", "Ali", 0.8},
		{"Charlie", "Mevlüt", 0.3},
		{"Furkan", "Semih", 0.5},
	}

	rand.Seed(time.Now().UnixNano())
	blocks := len(transactions) / 2
	usedIndices := make(map[int]bool)

	for i := 0; i < blocks; i++ {
		transactionsForBlock := make([]Transaction, 0)
		for len(transactionsForBlock) < 2 {
			index := rand.Intn(len(transactions))
			if !usedIndices[index] {
				transactionsForBlock = append(transactionsForBlock, transactions[index])
				usedIndices[index] = true
			}
		}
		blockchain.addBlock(transactionsForBlock)
	}

	merkleRoot := calculateMerkleRoot(transactions)
	fmt.Println("Merkle Root:", merkleRoot)
	/*
		for _, block := range blockchain.chain {
			fmt.Println("Block Version: ", block.Version)
			fmt.Println("Block Previous Hash: ", block.PreviousHash)
			fmt.Println("Block Difficulty: ", block.Difficulty)
			fmt.Println("Block Timestamp: ", block.Timestamp)
			fmt.Println("Block Merkle Root: ", block.MerkleRoot)
			fmt.Println("Block Nonce: ", block.Nonce)
			fmt.Println("Number of Transactions: ", block.NumTransactions)
			fmt.Println("Block Transactions: ", block.Transactions)
			fmt.Println("Block Hash: ", block.Hash)
			fmt.Println("-------------------------")
		}
	*/
}
