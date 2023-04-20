package block

import (
	"fmt"
	"sb/common/utils"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Height    int
	Timestamp int64
	PrevHash  string
	InfoList  []string
	ZeroCnt   int
	PoW       int
}

func New(height int, prevHash string, infoList []string, zeroCnt int, pow int) *Block {
	return &Block{
		Height:    height,
		Timestamp: time.Now().Unix(),
		PrevHash:  prevHash,
		InfoList:  infoList,
		ZeroCnt:   zeroCnt,
		PoW:       pow,
	}
}

func NewGenesisBlock() *Block {
	hash := utils.SHA256("tsy666")
	info := []string{
		"2023-04-20 is a good day",
		"tsy666",
	}
	return &Block{
		Height:    0,
		Timestamp: 1681920000000,
		PrevHash:  hash,
		InfoList:  info,
		ZeroCnt:   6,
		PoW:       13395002,
	}
}

// Get block's sha256 hash hex string.
func (blk *Block) Hash() string {
	str := fmt.Sprintf(
		"%v||%v||%v||%v",
		blk.Height,
		blk.Timestamp,
		blk.PrevHash,
		strings.Join(blk.InfoList, "__"),
	)
	return utils.SHA256(str)
}

func (blk *Block) Dig() int {
	pow := 0
	hash := blk.Hash()

	for {
		str := hash + "||" + strconv.Itoa(pow)
		hex := utils.SHA256(str)
		fmt.Printf("str: %v\n", str)
		if utils.HasLeadingZero(hex, blk.ZeroCnt) {
			fmt.Println(hex)
			return pow
		}
		pow++
	}
}

// Verify() verifies block's PoW
func (blk *Block) Verify() bool {
	hash := blk.Hash()
	// str=9ac1927269436217dd8ea7856fd28e12177fc1ef3d28cba4eb3ccf882fa9968d||13395002
	str := hash + "||" + strconv.Itoa(blk.PoW)
	hex := utils.SHA256(str)
	return utils.HasLeadingZero(hex, blk.ZeroCnt)
}

func (block *Block) SaveToChain(filepath string) {

}
