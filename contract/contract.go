package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetContract(client *ethclient.Client, hexContractAddress string, abi abi.ABI) *bind.BoundContract {
	contractAddress := common.HexToAddress(hexContractAddress)
	return bind.NewBoundContract(
		contractAddress,
		abi,
		client,
		client,
		client,
	)
}
