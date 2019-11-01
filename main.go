// sha-print prints the SHA256, SHA384, or SHA521 of the received input.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var s = flag.String("s", "256", "specify hash strength")

func main() {
	flag.Parse()

	c := flag.Arg(0)

	switch *s {
	case "384":
		fmt.Printf("%x\n", get384(c))
	case "512":
		fmt.Printf("%x\n", get512(c))
	default:
		fmt.Printf("%x\n", get256(c))
	}
}

func get256(c string) [32]byte {
	return sha256.Sum256([]byte(c))
}

func get384(c string) [48]byte {
	return sha512.Sum384([]byte(c))
}

func get512(c string) [64]byte {
	return sha512.Sum512([]byte(c))
}
