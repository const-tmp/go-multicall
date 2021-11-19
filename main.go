package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"multicallETH/contract"
	"multicallETH/erc20"
	"multicallETH/multicall"
	"os"
)

const multicallAddress = "0xeefba1e63905ef1d7acba5a8513c70307c1ce441"

var erc20Tokens = []string{
	"0xB8c77482e45F1F44dE1745F52C74426C631bDD52",
	"0xdac17f958d2ee523a2206206994597c13d831ec7",
	"0x2b591e99afe9f32eaa6214f7b7629768c40eeb39",
	"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
	"0x2260fac5e5542a773aa44fbcfedf7c193bc2c599",
	"0xa0b73e1ff0b80914ab6fe0444e65848c4c34450b",
	"0x4fabb145d64652a948d72533023f6e7a623c7c53",
	"0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0",
	"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984",
	"0x514910771af9ca656af840dff83e8264ecf986ca",
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: multicall [RPC URL] [holder address]")
		return
	}
	rpcUrl := os.Args[1]

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		fmt.Printf("Client connection error: %s\n", err)
		return
	}

	holderAddress := common.HexToAddress(os.Args[2])

	multicallContract := contract.GetContract(client, multicallAddress, multicall.Abi)

	type Call struct {
		Target   common.Address
		CallData []byte
	}

	var calls []Call

	for _, addr := range erc20Tokens {
		callData, err := erc20.Abi.Pack("balanceOf", holderAddress)
		if err != nil {
			fmt.Printf("ERC20 argument packing error: %s\n", err)
			return
		}
		target := common.HexToAddress(addr)

		calls = append(calls, Call{Target: target, CallData: callData})
	}
	//TODO: implement correct unpacking
	results := new([]interface{})
	err = multicallContract.Call(nil, results, "aggregate", calls)
	if err != nil {
		fmt.Printf("Contract call error: %s\n", err)
		return
	}
	block := (*results)[0].(*big.Int)
	fmt.Println("block:", block)
	for i, r := range (*results)[1].([][]uint8) {
		fmt.Println(i, common.Bytes2Hex(r))
	}
}
