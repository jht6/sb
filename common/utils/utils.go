package utils

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

func SHA256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

func Dig(prevHash string, infoList []string, zeroCount int) int {
	pow := 0
	str := prevHash + "||" + strings.Join(infoList, "||")
	for {
		t := str + "||" + strconv.Itoa(pow)
		hex := SHA256(t)
		// fmt.Printf("pow: %v, hex: %v\n", pow, hex)

		ok := true
		for i := 0; i < zeroCount; i++ {
			if hex[i] != '0' {
				ok = false
				break
			}
		}

		if ok {
			return pow
		}

		pow++
	}
}
