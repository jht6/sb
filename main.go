package main

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	initBlockchain()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		h := sha256.Sum256([]byte("a"))
		h1 := sha256.Sum256([]byte("a"))
		fmt.Printf("%x\n", h)
		fmt.Printf("%x\n", h1)
	})

	r.Run(":8080")
}

func initBlockchain() {
	// 初始化区块链
	// 建立一个创世区块
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
	mainFilePath := "./" + chainDir + "/" + mainFile
	_, err = os.Stat(mainFilePath)
	if err == nil {
		os.Remove(mainFilePath)
	}
	_, err = os.Create(mainFilePath)
	if err != nil {
		fmt.Printf("creating [%v] main file fails, process exits: %v\n", mainFilePath, err)
		os.Exit(1)
	}

}

// func main() {
// 	start := time.Now()
// 	utils.Dig(
// 		utils.SHA256("tsy666"),
// 		[]string{
// 			"2023-04-20 is a good day",
// 			"tsy666",
// 		},
// 		6,
// 	)
// 	fmt.Printf("cost time: %v\n", time.Since(start))
// }
