package deployer

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

// DeployContract deploy a smart contract to simulated chain to get out the contract's code and state
func DeployContract(deployCallback func(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (common.Address, error)) (code []byte, storage map[common.Hash]common.Hash, err error) {
	// Generate a new random account and a funded simulator
	prvKey, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(prvKey)
	auth.GasLimit = 12344321
	sim := backends.NewSimulatedBackend(core.GenesisAlloc{auth.From: {Balance: new(big.Int).Lsh(big.NewInt(1), 256-7)}}, auth.GasLimit)
	address, err := deployCallback(sim, auth)
	if err != nil {
		fmt.Println("Can't deploy nexty governance smart contract")
	}
	sim.Commit()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	code, _ = sim.CodeAt(ctx, address, nil)
	storage = make(map[common.Hash]common.Hash)
	sim.ForEachStorageAt(ctx, address, nil, func(key, val common.Hash) bool {
		storage[key] = val
		log.Info("DecodeBytes", "key", key, "value", storage[key])
		return true
	})
	return code, storage, err
}
