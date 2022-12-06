package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/rvmelkonian/slingshot-challenge/rpc_service"
)

const uniswapAddr = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("cannot load env")
	}

	var alchemyBaseUrl = "https://eth-mainnet.g.alchemy.com/v2/"
	var alchemyConnection = alchemyBaseUrl + os.Getenv("ALCHEMY_API_KEY")

	alchemyService, err := NewRPC(alchemyConnection, uniswapAddr)
	if err != nil {
		log.Fatal(err)
	}

	if err := rpc.Register(alchemyService); err != nil {
		log.Fatal("Error registering the alchemy service", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listener error", err)
	}

	fmt.Println("listening to RPC over localhost:4040 ... ")
	http.Serve(listener, nil)
}

func NewRPC(rpcURL string, uniswapAddr string) (*rpc_service.RPC, error) {
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	return &rpc_service.RPC{
		EthClient:   ethClient,
		UniswapAddr: uniswapAddr,
	}, nil
}
