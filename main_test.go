package main

import (
	"os"
	"testing"

	"github.com/rs/zerolog/log"
	"gotest.tools/assert"
)

func TestNewRPCService(t *testing.T) {

	var uniswapAddr = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
	var alchemyBaseUrl = "https://eth-mainnet.ws.alchemyapi.io/v2/"
	var alchemyConnection = alchemyBaseUrl + os.Getenv("ALCHEMY_API_KEY")

	t.Run("Should create a new RPC service", func(t *testing.T) {
		alchemyService, err := NewRPC(alchemyConnection, uniswapAddr)
		if err != nil {
			log.Error().Err(err)
		}

		assert.Equal(t, alchemyService.UniswapAddr, uniswapAddr)
	})

}
