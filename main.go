package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"

	"github.com/rvmelkonian/slingshot-challenge/rpc_service"
)

const uniswapAddr = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"

var alchemyConnection = "https://eth-mainnet.ws.alchemyapi.io/v2/" + os.Getenv("ALCHEMY_API_KEY")

func main() {
	alchemyService, err := rpc_service.NewRPC(alchemyConnection, uniswapAddr)
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
	log.Printf("serving rpc on port %d", 4040)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}
}
