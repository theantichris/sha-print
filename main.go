// sha-print prints the SHA256, SHA384, or SHA521 of the received input.
package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	c := flag.Arg(0)
	hash := getHash(c)
	fmt.Printf("%x\n", hash)
}

func getHash(c string) [32]byte {
	return sha256.Sum256([]byte(c))
}
