package rpc_service

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	factory "github.com/rvmelkonian/slingshot-challenge/UniswapFactory"
	router "github.com/rvmelkonian/slingshot-challenge/UniswapRouter"
)

type RPCService struct {
	ethClient   *ethclient.Client
	uniswapAddr string
}

func NewRPC(rpcURL string, uniswapAddr string) (*RPCService, error) {
	// create a new Ethereum client
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	return &RPCService{
		ethClient:   ethClient,
		uniswapAddr: uniswapAddr,
	}, nil
}

// GetRate returns the exchange rate for the giventoken pair
func (r *RPCService) GetRate(
	ctx context.Context,
	tokenA,
	tokenB string,
) (
	*big.Int,
	error,
) {
	if err := validateAddresses(tokenA, tokenB); err != nil {
		return big.NewInt(0), err
	}

	factoryInstance, err := factory.NewUniswapFactoryCaller(
		common.HexToAddress(r.uniswapAddr),
		r.ethClient,
	)
	if err != nil {
		return nil, err
	}

	exchangeAddr, err := factoryInstance.GetPair(nil,
		common.HexToAddress(tokenA),
		common.HexToAddress(tokenB),
	)
	if err != nil {
		return nil, err
	}

	routerInstance, err := router.NewUniswapRouterCaller(
		exchangeAddr,
		r.ethClient,
	)
	if err != nil {
		return nil, err
	}

	pair := []common.Address{
		common.HexToAddress(tokenA),
		common.HexToAddress(tokenB),
	}

	amountsOut, err := routerInstance.GetAmountsOut(nil, big.NewInt(1), pair)
	if err != nil {
		return nil, err
	}

	rate := new(big.Int).Div(amountsOut[0], amountsOut[1])
	return rate, nil
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
