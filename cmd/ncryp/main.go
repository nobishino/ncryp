package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nobishino/ncryp"
)

var (
	key     string
	payload string
	naive   bool
)

func main() {
	fmt.Println("hello, ncryp!")
	key = *flag.String("key", "00", "1 byte key for encryption. e.g. 3E")
	payload = *flag.String("payload", "", "byte sequence to encrypt")
	naive = *flag.Bool("naive", false, "naive encryption mode")

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

}
