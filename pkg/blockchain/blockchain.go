package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Timestamp    int64
	Data         string
	PrevHash     string
	Hash         string
	Nonce        int
	Difficulty   int
}

type Blockchain struct {
	Blocks []*Block
}

func (b *Block) CalculateHash() string {
	record := fmt.Sprintf("%d%s%s%d%d", b.Timestamp, b.Data, b.PrevHash, b.Nonce, b.Difficulty)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (b *Block) MineBlock(difficulty int) {
	target := string(make([]byte, difficulty))
	for i := 0; i < difficulty; i++ {
		target = "0" + target
	}

	for {
		b.Hash = b.CalculateHash()
		if b.Hash[:difficulty] == target[:difficulty] {
			break
		}
		b.Nonce++
	}
}

func CreateBlock(data string, prevHash string, difficulty int) *Block {
	block := &Block{
		Timestamp:  time.Now().Unix(),
		Data:       data,
		PrevHash:   prevHash,
		Difficulty: difficulty,
		Nonce:      0,
	}
	block.MineBlock(difficulty)
	return block
}

func NewBlockchain() *Blockchain {
	genesisBlock := CreateBlock("Genesis Block", "", 2)
	return &Blockchain{[]*Block{genesisBlock}}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash, 2)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		prevBlock := bc.Blocks[i-1]

		if currentBlock.Hash != currentBlock.CalculateHash() {
			return false
		}

		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}
	}
	return true
}
