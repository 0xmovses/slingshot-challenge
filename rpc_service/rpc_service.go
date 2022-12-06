package rpc_service

import (
	"context"
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

type Route struct {
	route []common.Address
	rate  *big.Int
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
		log.Error().Msgf(
			"Error getting the exchange rate, a pool for this pair may not exist: %v",
			err,
		)
		return err
	}

	bigIntRate := new(big.Int).Div(amountsOut[1], amountsOut[0])
	*rate = float64(bigIntRate.Int64())
	return nil
}

// GetRoute returns the best route and rate for the given token pair
func (r *RPC) GetRoute(ctx context.Context, tokenA common.Address, tokenB common.Address) (*Route, error) {
	uniswap, err := factory.NewUniswapFactory(
		common.HexToAddress(r.UniswapAddr),
		r.EthClient,
	)
	if err != nil {
		return nil, err
	}

	totalCount, err := uniswap.AllPairsLength(nil)
	if err != nil {
		return nil, err
	}

	routes := make(map[common.Address][]common.Address)

	for i := 0; i < int(totalCount.Int64()); i++ {
		_, err := uniswap.AllPairs(nil, big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}
	}

	uniswap.AllPairs(nil, big.NewInt(totalCount.Int64()))

	routes[tokenA] = []common.Address{tokenA}

	bestRates := make(map[common.Address]*big.Int)
	bestRates[tokenA] = big.NewInt(1)

	bestRoute := Route{route: []common.Address{tokenA}, rate: big.NewInt(1)}

	return &bestRoute, nil
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
