package main

import (
	"fmt"
	"os"
	"sb/block"
)

func main() {
	initBlockchain()

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	h := sha256.Sum256([]byte("a"))
	// 	h1 := sha256.Sum256([]byte("a"))
	// 	fmt.Printf("%x\n", h)
	// 	fmt.Printf("%x\n", h1)
	// })

	// r.Run(":8080")
}

func initBlockchain() {
	chainDir := ".blockchain"
	mainFile := "main"

	// if .blockchain dir doesn't exist，create it.
	_, err := os.Stat("./" + chainDir)
	if err != nil {
		err := os.Mkdir(chainDir, os.ModePerm)
		if err != nil {
			fmt.Printf("creating [%v] dir fails, process exits: %v\n", chainDir, err)
			os.Exit(1)
		}
	}

	// if main chain file exists, clear it.
	cwd, _ := os.Getwd()
	mainFilePath := cwd + "/" + chainDir + "/" + mainFile
	_, err = os.Stat(mainFilePath)
	if err == nil {
		os.Remove(mainFilePath)
	}
	_, err = os.Create(mainFilePath)
	if err != nil {
		fmt.Printf("creating [%v] main file fails, process exits: %v\n", mainFilePath, err)
		os.Exit(1)
	}

	genesisBlock := block.NewGenesisBlock()
	fmt.Println("genesisBlock created.")
	err = genesisBlock.SaveToChain(mainFilePath)
	if err != nil {
		fmt.Println("save genesis block err:", err)
	}
}

// func main() {
// 	start := time.Now()
// 	blk := block.NewGenesisBlock()
// 	blk.Dig()
// 	fmt.Printf("cost time: %v\n", time.Since(start))

// 	blk := block.NewGenesisBlock()
// 	fmt.Println(blk.Verify())
// }
