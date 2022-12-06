package rpc_service

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	router "github.com/rvmelkonian/slingshot-challenge/UniswapRouter"
)

type RPC struct {
	EthClient   *ethclient.Client
	UniswapAddr string
}

type Pair struct {
	TokenA string
	TokenB string
}

const uniswapV2RouterAddr = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"

// GetRate returns the exchange rate for the giventoken pair
func (r *RPC) GetRate(pair Pair, rate *float64) error {
	if err := validateAddresses(pair.TokenA, pair.TokenB); err != nil {
		return err
	}

	routerInstance, err := router.NewUniswapRouter(
		common.HexToAddress(uniswapV2RouterAddr),
		r.EthClient,
	)
	if err != nil {
		log.Error()
		return err
	}

	path := []common.Address{
		common.HexToAddress(pair.TokenA),
		common.HexToAddress(pair.TokenB),
	}

	amountsOut, err := routerInstance.GetAmountsOut(nil, big.NewInt(1), path)
	if err != nil {
		log.Error()
		return err
	}

	for _, amount := range amountsOut {
		fmt.Println(amount)
	}

	bigIntRate := new(big.Int).Div(amountsOut[1], amountsOut[0])
	*rate = float64(bigIntRate.Int64())
	return nil
}

func validateAddresses(tokenA, tokenB string) error {
	isValid := common.IsHexAddress(tokenA)
	if !isValid {
		err := new(error)
		return *err
	}

	isValid = common.IsHexAddress(tokenB)
	if isValid {
		err := new(error)
		return *err
	}

	return nil
}
