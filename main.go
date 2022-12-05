package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	blockchain "github.com/rvmelkonian/slingshot-challenge/UniswapRouter"
	factory "github.com/rvmelkonian/slingshot-challenge/UniswapRouter/UniswapFactory"
)

const (
	alchemyConnection = "https://eth-mainnet.g.alchemy.com/v2/W0GZ_LZ8Pbuu-c4c44IFEZqStmNbbgh_"
	uniswapAddr       = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
)

// RPCService is the struct that holds the Ethereum client and Uniswap contract addresses
type RPCService struct {
	ethClient   *ethclient.Client
	uniswapAddr string
}

// NewRPCService creates a new instance of the RPCService
func NewRPCService(rpcURL string, uniswapAddr string) (*RPCService, error) {
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
func (r *RPCService) GetRate(ctx context.Context, tokenA, tokenB string) (*big.Int, error) {
	// create a new instance of the Uniswap contract
	uniswap, err := factory.NewUniswapRouter(common.HexToAddress(r.uniswapAddr), r.ethClient)
	if err != nil {
		return nil, err
	}

	// get the exchange address for the given token pair
	exchangeAddr, err := uniswap.GetExchange(nil, tokenA, tokenB)
	if err != nil {
		return nil, err
	}

	// create a new instance of the exchange contract
	exchange, err := blockchain.NewUniswapV2Router02(exchangeAddr, r.ethClient)
	if err != nil {
		return nil, err
	}

	// call the 'getAmountsOut' function on the exchange contract to get the token amounts
	amountsOut, err := exchange.GetAmountsOut(nil, big.NewInt(1), tokenA, tokenB)
	if err != nil {
		return nil, err
	}

	// calculate the exchange rate using the token amounts
	rate := new(big.Int).Div(amountsOut[0], amountsOut[1])

	return rate, nil
}

func main() {

	rpcService, err := NewRPCService(alchemyConnection, uniswapAddr)
	if err != nil {
		log.Fatal(err)
	}

	// get the exchange rate for the ETH-DAI token pair
	ethAddr := common.HexToAddress("0xEeeeeEeeeEeEeeEe")
	fmt.Printf(rpcService, ethAddr)
}
