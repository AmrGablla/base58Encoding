package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"

	"github.com/mr-tron/base58"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Base58 Check Encoding")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		hash := sha256.Sum256([]byte(text))
		hash = sha256.Sum256(hash[:])
		if strings.Compare("exit", text) == 0 {
			fmt.Printf("bye, see u.")
			fmt.Printf("\n")
			os.Exit(0)
		}

		fmt.Printf("Double hash: ")
		fmt.Printf("%x", hash)
		fmt.Printf("\n")

		fmt.Printf("Checksum : ")
		fmt.Printf("%x", hash[:2])
		fmt.Printf("\n")

		base58Encoded := base58.Encode(hash[:])
		fmt.Printf("base 58 Encode: ")
		fmt.Printf("%x", base58Encoded)
		fmt.Printf("\n")

	}

}
