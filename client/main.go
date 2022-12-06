package main

import (
	"fmt"
	"log"
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
	wethAddr         = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	uniswapTokenAddr = "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"
)

func main() {
	var reply float64
	var pair = Pair{
		TokenA: wethAddr,
		TokenB: uniswapTokenAddr,
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
