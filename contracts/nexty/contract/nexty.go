// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NextyABI is the input ABI used to generate the binding from.
const NextyABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"join\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"leave\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// NextyBin is the compiled bytecode used for deploying new contracts.
const NextyBin = `0x608060405234801561001057600080fd5b50600060208190527f2874ff5f94a33fde51343733d248eb2a96767424898545c0c2d26675e52d40798054600160a060020a031990811673391730a41d2c27279b2b37e5e9209d5682c6d09a179091557f1d7c5114599e86d6147844633dd89acf1d376cde4cb7abdce9d84b2ead224f918054821673766ea022b264cd64fb346b2b88409137790354b51790557f149f473e5237dfb9fae6f38765f20cf018e89b38f42b028be810d8c58c66fd3f80548216734df4f0232bd4ac6808c98871959850eac389013d179055600180548082018255928190527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf692830180548316731dee87fc224a41c8699a3d8a8557918be7e4795a17905580548082018255830180548316738e2a6a6690156004185457c0ee70675bc4c1b1a51790558054808201909155909101805490911673d88d2aeb036d162db43e1b18ebd820cbd99d91b2179055610266806101836000396000f3fe608060405260043610610045577c01000000000000000000000000000000000000000000000000000000006000350463b688a363811461004a578063d66d9e1914610061575b600080fd5b34801561005657600080fd5b5061005f610076565b005b34801561006d57600080fd5b5061005f6100c6565b6001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601805473ffffffffffffffffffffffffffffffffffffffff191633179055565b60005b6001548110156101ec5760018054829081106100e157fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff163314156101e45760018054600019810190811061011c57fe5b6000918252602090912001546001805473ffffffffffffffffffffffffffffffffffffffff909216918390811061014f57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff929092169190911790556001805460001981019081106101a457fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916905560018054906101dd9060001983016101f0565b50506101ee565b6001016100c9565b505b565b81548183558181111561021457600083815260209020610214918101908301610219565b505050565b61023791905b80821115610233576000815560010161021f565b5090565b9056fea165627a7a723058201d70f1c32a866a9554a8d113a9ab9641d844798b391dbdf8062e67bc640efb850029`

// DeployNexty deploys a new Ethereum contract, binding an instance of Nexty to it.
func DeployNexty(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Nexty, error) {
	parsed, err := abi.JSON(strings.NewReader(NextyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NextyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Nexty{NextyCaller: NextyCaller{contract: contract}, NextyTransactor: NextyTransactor{contract: contract}, NextyFilterer: NextyFilterer{contract: contract}}, nil
}

// Nexty is an auto generated Go binding around an Ethereum contract.
type Nexty struct {
	NextyCaller     // Read-only binding to the contract
	NextyTransactor // Write-only binding to the contract
	NextyFilterer   // Log filterer for contract events
}

// NextyCaller is an auto generated read-only Go binding around an Ethereum contract.
type NextyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NextyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NextyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NextyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NextyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NextySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NextySession struct {
	Contract     *Nexty            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NextyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NextyCallerSession struct {
	Contract *NextyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NextyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NextyTransactorSession struct {
	Contract     *NextyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NextyRaw is an auto generated low-level Go binding around an Ethereum contract.
type NextyRaw struct {
	Contract *Nexty // Generic contract binding to access the raw methods on
}

// NextyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NextyCallerRaw struct {
	Contract *NextyCaller // Generic read-only contract binding to access the raw methods on
}

// NextyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NextyTransactorRaw struct {
	Contract *NextyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNexty creates a new instance of Nexty, bound to a specific deployed contract.
func NewNexty(address common.Address, backend bind.ContractBackend) (*Nexty, error) {
	contract, err := bindNexty(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nexty{NextyCaller: NextyCaller{contract: contract}, NextyTransactor: NextyTransactor{contract: contract}, NextyFilterer: NextyFilterer{contract: contract}}, nil
}

// NewNextyCaller creates a new read-only instance of Nexty, bound to a specific deployed contract.
func NewNextyCaller(address common.Address, caller bind.ContractCaller) (*NextyCaller, error) {
	contract, err := bindNexty(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NextyCaller{contract: contract}, nil
}

// NewNextyTransactor creates a new write-only instance of Nexty, bound to a specific deployed contract.
func NewNextyTransactor(address common.Address, transactor bind.ContractTransactor) (*NextyTransactor, error) {
	contract, err := bindNexty(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NextyTransactor{contract: contract}, nil
}

// NewNextyFilterer creates a new log filterer instance of Nexty, bound to a specific deployed contract.
func NewNextyFilterer(address common.Address, filterer bind.ContractFilterer) (*NextyFilterer, error) {
	contract, err := bindNexty(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NextyFilterer{contract: contract}, nil
}

// bindNexty binds a generic wrapper to an already deployed contract.
func bindNexty(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NextyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nexty *NextyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Nexty.Contract.NextyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nexty *NextyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nexty.Contract.NextyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nexty *NextyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nexty.Contract.NextyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nexty *NextyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Nexty.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nexty *NextyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nexty.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nexty *NextyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nexty.Contract.contract.Transact(opts, method, params...)
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_Nexty *NextyTransactor) Join(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nexty.contract.Transact(opts, "join")
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_Nexty *NextySession) Join() (*types.Transaction, error) {
	return _Nexty.Contract.Join(&_Nexty.TransactOpts)
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_Nexty *NextyTransactorSession) Join() (*types.Transaction, error) {
	return _Nexty.Contract.Join(&_Nexty.TransactOpts)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_Nexty *NextyTransactor) Leave(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nexty.contract.Transact(opts, "leave")
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_Nexty *NextySession) Leave() (*types.Transaction, error) {
	return _Nexty.Contract.Leave(&_Nexty.TransactOpts)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns()
func (_Nexty *NextyTransactorSession) Leave() (*types.Transaction, error) {
	return _Nexty.Contract.Leave(&_Nexty.TransactOpts)
}
