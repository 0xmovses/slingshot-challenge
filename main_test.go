package main

import (
	"log"
	"net/rpc"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var alchemyBaseUrl = "https://eth-mainnet.ws.alchemyapi.io/v2/"
var alchemyConnection = alchemyBaseUrl + os.Getenv("ALCHEMY_API_KEY")

type Pair struct {
	TokenA string
	TokenB string
}

func TestRPCService(t *testing.T) {

	t.Run("Should create a new RPC service", func(t *testing.T) {
		alchemyService, err := NewRPC(alchemyConnection, uniswapAddr)
		if err != nil {
			log.Fatal("Error creating the alchemy service", err)
		}

		assert.Equal(t, alchemyService.UniswapAddr, uniswapAddr)
	})

	t.Run("Should return an error if the connection string is invalid", func(t *testing.T) {
		_, err := NewRPC("invalid", uniswapAddr)
		assert.NotNil(t, err)
	})
}

func TestGetRate(t *testing.T) {

	t.Run("Should return the exchange rate between two tokens as float64", func(t *testing.T) {
		var rate float64
		var wethAddr = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
		var uniswapAddr = "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"

		var pair = Pair{
			TokenA: wethAddr,
			TokenB: uniswapAddr,
		}

		client, err := rpc.DialHTTP("tcp", "localhost:4040")
		if err != nil {
			log.Fatal("Connection error: ", err)
		}

		if err = client.Call("RPC.GetRate", pair, &rate); err != nil {
			log.Fatal("Error when calling GetRate() : ", err)
		}

		assert.IsType(t, rate, float64(0))
	})
}
