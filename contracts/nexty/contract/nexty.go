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
const NextyABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"join\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"leave\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_sealers\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// NextyBin is the compiled bytecode used for deploying new contracts.
const NextyBin = `0x608060405234801561001057600080fd5b506040516103d23803806103d28339810180604052602081101561003357600080fd5b81019080805164010000000081111561004b57600080fd5b8201602081018481111561005e57600080fd5b815185602082028301116401000000008211171561007b57600080fd5b509093506000925050505b815181101561015b57818181518110151561009d57fe5b9060200190602002015160008084848151811015156100b857fe5b90602001906020020151600160a060020a0316600160a060020a0316815260200190815260200160002060006101000a815481600160a060020a030219169083600160a060020a031602179055506001828281518110151561011657fe5b6020908102919091018101518254600180820185556000948552929093209092018054600160a060020a031916600160a060020a039093169290921790915501610086565b50506102668061016c6000396000f3fe608060405260043610610045577c01000000000000000000000000000000000000000000000000000000006000350463b688a363811461004a578063d66d9e1914610061575b600080fd5b34801561005657600080fd5b5061005f610076565b005b34801561006d57600080fd5b5061005f6100c6565b6001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601805473ffffffffffffffffffffffffffffffffffffffff191633179055565b60005b6001548110156101ec5760018054829081106100e157fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff163314156101e45760018054600019810190811061011c57fe5b6000918252602090912001546001805473ffffffffffffffffffffffffffffffffffffffff909216918390811061014f57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff929092169190911790556001805460001981019081106101a457fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916905560018054906101dd9060001983016101f0565b50506101ee565b6001016100c9565b505b565b81548183558181111561021457600083815260209020610214918101908301610219565b505050565b61023791905b80821115610233576000815560010161021f565b5090565b9056fea165627a7a72305820d300a2dce13c180c0dff4b5e71b33acc44c4023e865483557ef84e68a7f70cf10029`

// DeployNexty deploys a new Ethereum contract, binding an instance of Nexty to it.
func DeployNexty(auth *bind.TransactOpts, backend bind.ContractBackend, _sealers []common.Address) (common.Address, *types.Transaction, *Nexty, error) {
	parsed, err := abi.JSON(strings.NewReader(NextyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NextyBin), backend, _sealers)
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
