## slinghsot-challenge

An RPC Service for interacting with the Uniswap V2 Router, written in Go

### Question One:

Go must be installed on your machine to run this project

Instructions to run 
- create a .env file at the root of the project
- create an Environment Variable `ALCHEMY_API_KEY` and set the value to your api key
- Run `go mod tidy`
- Then `go run main.go`
The RPC service is now running and listening on port localhost:4040

To run the tests run `go test -v` at the root of the project

N.B In a real world scenario many more tests would be written, perhaps using the table test method, 
for exploring more edge cases.

An example client has been written inside `/client`
To run the client, which calls the RPC service methods navigate to `/client` and run `go run main.go`

###  Question Two:

A full working solution for question two has not yet been implemented yet, as I ran out of time,
however the work-in-progress solution can be checked out at branch `get-route`

As I have not yet been able to test this function, its difficult to say what needs improving. 
My general approach was to get all the pair pools, find matching pools that contain either TokenA and TokenB, then call `GetRate` on them; append the results into a slice, sort the slice, find the best rate. This may be a flawed approach, but I wasn't able to dig in deep with this yet


