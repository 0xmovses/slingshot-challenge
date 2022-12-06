package rpc_service

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	factory "github.com/rvmelkonian/slingshot-challenge/UniswapFactory"
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

// GetRate returns the exchange rate for the giventoken pair
func (r *RPC) GetRate(pair Pair, reply *big.Int) error {
	fmt.Println("GetRate called")
	if err := validateAddresses(pair.TokenA, pair.TokenB); err != nil {
		return err
	}

	factoryInstance, err := factory.NewUniswapFactoryCaller(
		common.HexToAddress(r.UniswapAddr),
		r.EthClient,
	)
	if err != nil {
		log.Error()
		return err
	}

	fmt.Println("factory initialized")

	exchangeAddr, err := factoryInstance.GetPair(nil,
		common.HexToAddress(pair.TokenA),
		common.HexToAddress(pair.TokenB),
	)
	if err != nil {
		fmt.Printf("Error : %s", err)
		log.Error()
		return err
	}

	fmt.Printf("exchange Address : %v", exchangeAddr)

	routerInstance, err := router.NewUniswapRouterCaller(
		exchangeAddr,
		r.EthClient,
	)
	if err != nil {
		log.Error()
		return err
	}

	fmt.Println("router initialized")

	path := []common.Address{
		common.HexToAddress(pair.TokenA),
		common.HexToAddress(pair.TokenB),
	}

	amountsOut, err := routerInstance.GetAmountsOut(nil, big.NewInt(1), path)
	if err != nil {
		log.Error()
		return err
	}

	fmt.Println("AmountsOut fetched from contract")

	rate := new(big.Int).Div(amountsOut[0], amountsOut[1])
	reply = rate
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
