package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nobishino/ncryp"
)

var (
	key     string
	payload string
	naive   bool
)

func main() {
	switch {
	case naive:
		os.Exit(execNaive(payload, key))
	default:
		fmt.Println("helly, ncryp")
	}
}

func execNaive(payload, key string) int {
	p, err := ncryp.NewPayload(payload)
	if err != nil {
		log.Fatal(err)
	}
	k, err := ncryp.NewPayload(key)
	if err != nil {
		log.Fatal(err)
	}
	key16 := ncryp.Key16{k[0], k[1]}

	e := ncryp.NaiveSymCryp(p, key16)

	fmt.Println(e)
	return 0
}

func init() {
	flag.StringVar(&key, "key", "00", "1 byte key for encryption. e.g. 3E")
	flag.StringVar(&payload, "payload", "", "byte sequence to encrypt")
	flag.BoolVar(&naive, "naive", false, "naive encryption mode")
	flag.Parse()
}
