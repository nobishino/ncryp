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
	keygen  bool
	rsa     bool
)

func main() {
	switch {
	case rsa:
		os.Exit(execRSAEncryption(key, payload))
	case keygen:
		os.Exit(execGenKey())
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

func execGenKey() int {
	k1, k2 := ncryp.GenKeyPair()
	fmt.Println(k1.Payload(), k2.Payload())
	return 0
}

func execRSAEncryption(keyStr, payloadStr string) int {
	key, err := ncryp.NewPayload(keyStr)
	if err != nil {
		log.Println(err)
		return 1
	}
	payload, err := ncryp.NewPayload(payloadStr)
	if err != nil {
		log.Println(err)
		return 1
	}
	if len(key) != 16 {
		log.Printf("expect key string to have length of 16 but got %d", len(key))
		return 1
	}
	rsaKey := ncryp.RSAKey{
		Mod:   ncryp.Payload(key[:8]).Uint64(),
		Power: ncryp.Payload(key[8:]).Uint64(),
	}
	payloadNum := payload.Uint64() % rsaKey.Mod
	cipher := ncryp.ModPow(payloadNum, rsaKey.Power, rsaKey.Mod)
	fmt.Println(ncryp.PayloadUint64(cipher))
	return 0
}

func init() {
	flag.StringVar(&key, "key", "00", "1 byte key for encryption. e.g. 3E")
	flag.StringVar(&payload, "payload", "", "byte sequence to encrypt")
	flag.BoolVar(&naive, "naive", false, "naive encryption mode")
	flag.BoolVar(&keygen, "keygen", false, "RSA key pair generation mode")
	flag.BoolVar(&rsa, "rsa", false, "RSA encryption mode")
	flag.Parse()
}
