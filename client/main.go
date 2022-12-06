package main

import (
	"fmt"
	"log"
	"math/big"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

type Pair struct {
	TokenA string
	TokenB string
}

const (
	wethAddress = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	usdtAddress = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
)

func main() {
	var reply big.Int
	var pair = Pair{
		TokenA: wethAddress,
		TokenB: usdtAddress,
	}

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	if err = client.Call("RPC.GetRate", pair, &reply); err != nil {
		log.Fatal("Error when calling GetRate() : ", err)
	}
	fmt.Printf("Rate : %v", reply)
}
