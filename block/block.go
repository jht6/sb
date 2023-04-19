package block

import (
	"sb/common/utils"
	"time"
)

type Block struct {
	Height    int
	Timestamp int64
	PrevHash  string
	InfoList  []string
	PoW       int
}

func NewFirstBlock() *Block {
	hash := utils.SHA256("tsy666")
	info := []string{
		"2023-04-20 is a good day",
		"tsy666",
	}
	return New(0, hash, info, 11189882)
}

func New(height int, prevHash string, infoList []string, pow int) *Block {
	return &Block{
		Height:    height,
		Timestamp: time.Now().Unix(),
		PrevHash:  prevHash,
		InfoList:  infoList,
		PoW:       pow,
	}
}
