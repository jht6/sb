package block

import (
	"bufio"
	"fmt"
	"os"
	"sb/common/utils"
	"strconv"
	"strings"
	"time"

	"github.com/goccy/go-json"
)

type BlockOpt struct {
	Height    int
	Timestamp int64
	PrevHash  string
	InfoList  []string
	ZeroCnt   int
}
type Block struct {
	Height    int   // block height
	Timestamp int64 // second level timestamp
	PrevHash  string
	InfoList  []string
	ZeroCnt   int

	Nonce int
}

func New(opt *BlockOpt) *Block {
	var ts int64
	if opt.Timestamp > 0 {
		ts = opt.Timestamp
	} else {
		ts = time.Now().Unix()
	}

	return &Block{
		Height:    opt.Height,
		Timestamp: ts,
		PrevHash:  opt.PrevHash,
		InfoList:  opt.InfoList,
		ZeroCnt:   opt.ZeroCnt,

		Nonce: -1,
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
		Timestamp: 1681920000,
		PrevHash:  hash,
		InfoList:  info,
		ZeroCnt:   5,
		Nonce:     85084,
	}
}

// Get block's sha256 hash hex string.
func (blk *Block) Hash() string {
	str := fmt.Sprintf(
		"%v||%v||%v||%v",
		blk.Height,
		blk.Timestamp,
		blk.PrevHash,
		blk.InfoHash(),
	)
	return utils.SHA256(str)
}

func (blk *Block) InfoHash() string {
	return utils.SHA256(strings.Join(blk.InfoList, "__"))
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

// Verify() verifies if block's Nonce meets requirement
func (blk *Block) Verify() (bool, error) {
	hash := blk.Hash()
	str := hash + "||" + strconv.Itoa(blk.Nonce)
	hex := utils.SHA256(str)
	if utils.HasLeadingZero(hex, blk.ZeroCnt) {
		return true, nil
	}
	err := fmt.Errorf("Block verifies fail:\nhex:%v\nzeroCnt:%v", hex, blk.ZeroCnt)
	return false, err
}

// Verify block's Nonce, if it's ok, transform Block
// to JSON and save it to disk.
func (blk *Block) SaveToChain(filepath string) error {
	if _, err := blk.Verify(); err != nil {
		return err
	}

	jsonByte, _ := json.Marshal(blk)
	jsonStr := string(jsonByte)

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(jsonStr + "\n")
	err = writer.Flush()
	if err != nil {
		fmt.Println("write block to file err:", err)
		return err
	}

	return nil
}
