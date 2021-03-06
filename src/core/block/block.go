package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Meta *Meta       `json:"meta"`
	Data interface{} `json:"data"`
}

func New(index int, prevHash string, data interface{}) *Block {
	meta := &Meta{
		Index:     index,
		Timestamp: int(time.Now().Unix()),
		PrevHash:  prevHash,
	}
	block := &Block{
		Meta: meta,
		Data: data,
	}

	block.Meta.Hash = block.CreateHash()

	return block
}

func (b Block) CreateHash() string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d%d%s%d", b.Meta.Index, b.Meta.Timestamp, b.Meta.PrevHash, b.Meta.Nonce)))

	return hex.EncodeToString(h.Sum(nil))
}

func (b *Block) ProofOfWork(difficulty int) {
	if difficulty <= 0 {
		difficulty = 1
	}

	for b.Meta.Hash[:difficulty] != strings.Repeat("0", difficulty) {
		b.Meta.Nonce++
		b.Meta.Hash = b.CreateHash()
	}
}
