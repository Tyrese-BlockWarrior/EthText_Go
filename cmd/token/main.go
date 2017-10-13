//go:generate abigen --sol ./contracts/token.sol --pkg contracts --out ./contracts/token.go

package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/toba/ethtest/contracts"
)

// http://ethdocs.org/en/latest/network/test-networks.html#morden-s-genesis-json
func main() {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("/Users/jason-abbott/Library/Ethereum/testnet/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Test with Ethereum "unicorn" contract https://ethereum.org/donate
	token, err := contracts.NewToken(common.HexToAddress("0x21e6fc92f93c8a1bb41e2be64b4e1f88a54d3576"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	name, err := token.Name(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	fmt.Println("Token name:", name)
}
