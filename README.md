###slinghsot-challenge

An RPC Service for interacting with the Uniswap V2 Router, written in Go

Go must be installed on your machine to run this project

Instructions to run 
- create a .env file at the root of the project
- create an Environment Variable `ALCHEMY_API_KEY` and set the value to your api key
- Run `go mod tidy`
- Then `go run main.go`
The RPC service is now running and listening on port localhost:4040

To run the tests run `go test -v` at the root of the project

An example client has been written inside `/client`
To run the client, which calls the RPC service methods navigate to `/client` and run `go run main.go`
