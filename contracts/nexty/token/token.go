// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// IERC20Bin is the compiled bytecode used for deploying new contracts.
const IERC20Bin = `0x`

// DeployIERC20 deploys a new Ethereum contract, binding an instance of IERC20 to it.
func DeployIERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MultiOwnableABI is the input ABI used to generate the binding from.
const MultiOwnableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownerByAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"setOwners\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"SetOwners\",\"type\":\"event\"}]"

// MultiOwnableBin is the compiled bytecode used for deploying new contracts.
const MultiOwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a031916331790556104d7806100326000396000f3fe608060405234801561001057600080fd5b5060043610610073577c01000000000000000000000000000000000000000000000000000000006000350463025e7c278114610078578063481c6a75146100b1578063a0e67e2b146100b9578063eb6b192f14610111578063fa4d36981461014b575b600080fd5b6100956004803603602081101561008e57600080fd5b50356101f0565b60408051600160a060020a039092168252519081900360200190f35b610095610218565b6100c1610227565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100fd5781810151838201526020016100e5565b505050509050019250505060405180910390f35b6101376004803603602081101561012757600080fd5b5035600160a060020a031661028a565b604080519115158252519081900360200190f35b6101ee6004803603602081101561016157600080fd5b81019060208101813564010000000081111561017c57600080fd5b82018360208201111561018e57600080fd5b803590602001918460208302840111640100000000831117156101b057600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061029f945050505050565b005b60018054829081106101fe57fe5b600091825260209091200154600160a060020a0316905081565b600054600160a060020a031681565b6060600180548060200260200160405190810160405280929190818152602001828054801561027f57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610261575b505050505090505b90565b60026020526000908152604090205460ff1681565b600054600160a060020a031633146102b657600080fd5b6102bf816102c2565b50565b60005b600154811015610320576000600260006001848154811015156102e457fe5b600091825260208083209190910154600160a060020a031683528201929092526040019020805460ff19169115159190911790556001016102c5565b5060005b815181101561037b57600160026000848481518110151561034157fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff1916911515919091179055600101610324565b50805161038f906001906020840190610408565b507f9465cd279c2de393c5568ae444599e3644e3d1864ca2c05ced8a654df2aea3cb816040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156103f25781810151838201526020016103da565b505050509050019250505060405180910390a150565b82805482825590600052602060002090810192821561046a579160200282015b8281111561046a578251825473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909116178255602090920191600190910190610428565b5061047692915061047a565b5090565b61028791905b8082111561047657805473ffffffffffffffffffffffffffffffffffffffff1916815560010161048056fea165627a7a7230582008039274668331380a12d06070648f2262b4a70074444a40006f31d8d3d261f30029`

// DeployMultiOwnable deploys a new Ethereum contract, binding an instance of MultiOwnable to it.
func DeployMultiOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MultiOwnable, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiOwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MultiOwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiOwnable{MultiOwnableCaller: MultiOwnableCaller{contract: contract}, MultiOwnableTransactor: MultiOwnableTransactor{contract: contract}, MultiOwnableFilterer: MultiOwnableFilterer{contract: contract}}, nil
}

// MultiOwnable is an auto generated Go binding around an Ethereum contract.
type MultiOwnable struct {
	MultiOwnableCaller     // Read-only binding to the contract
	MultiOwnableTransactor // Write-only binding to the contract
	MultiOwnableFilterer   // Log filterer for contract events
}

// MultiOwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiOwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiOwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiOwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiOwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiOwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiOwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiOwnableSession struct {
	Contract     *MultiOwnable     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiOwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiOwnableCallerSession struct {
	Contract *MultiOwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MultiOwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiOwnableTransactorSession struct {
	Contract     *MultiOwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MultiOwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiOwnableRaw struct {
	Contract *MultiOwnable // Generic contract binding to access the raw methods on
}

// MultiOwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiOwnableCallerRaw struct {
	Contract *MultiOwnableCaller // Generic read-only contract binding to access the raw methods on
}

// MultiOwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiOwnableTransactorRaw struct {
	Contract *MultiOwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiOwnable creates a new instance of MultiOwnable, bound to a specific deployed contract.
func NewMultiOwnable(address common.Address, backend bind.ContractBackend) (*MultiOwnable, error) {
	contract, err := bindMultiOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiOwnable{MultiOwnableCaller: MultiOwnableCaller{contract: contract}, MultiOwnableTransactor: MultiOwnableTransactor{contract: contract}, MultiOwnableFilterer: MultiOwnableFilterer{contract: contract}}, nil
}

// NewMultiOwnableCaller creates a new read-only instance of MultiOwnable, bound to a specific deployed contract.
func NewMultiOwnableCaller(address common.Address, caller bind.ContractCaller) (*MultiOwnableCaller, error) {
	contract, err := bindMultiOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiOwnableCaller{contract: contract}, nil
}

// NewMultiOwnableTransactor creates a new write-only instance of MultiOwnable, bound to a specific deployed contract.
func NewMultiOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiOwnableTransactor, error) {
	contract, err := bindMultiOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiOwnableTransactor{contract: contract}, nil
}

// NewMultiOwnableFilterer creates a new log filterer instance of MultiOwnable, bound to a specific deployed contract.
func NewMultiOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiOwnableFilterer, error) {
	contract, err := bindMultiOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiOwnableFilterer{contract: contract}, nil
}

// bindMultiOwnable binds a generic wrapper to an already deployed contract.
func bindMultiOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiOwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiOwnable *MultiOwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiOwnable.Contract.MultiOwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiOwnable *MultiOwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiOwnable.Contract.MultiOwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiOwnable *MultiOwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiOwnable.Contract.MultiOwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiOwnable *MultiOwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiOwnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiOwnable *MultiOwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiOwnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiOwnable *MultiOwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiOwnable.Contract.contract.Transact(opts, method, params...)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MultiOwnable *MultiOwnableCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MultiOwnable.contract.Call(opts, out, "getOwners")
	return *ret0, err
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MultiOwnable *MultiOwnableSession) GetOwners() ([]common.Address, error) {
	return _MultiOwnable.Contract.GetOwners(&_MultiOwnable.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MultiOwnable *MultiOwnableCallerSession) GetOwners() ([]common.Address, error) {
	return _MultiOwnable.Contract.GetOwners(&_MultiOwnable.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_MultiOwnable *MultiOwnableCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultiOwnable.contract.Call(opts, out, "manager")
	return *ret0, err
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_MultiOwnable *MultiOwnableSession) Manager() (common.Address, error) {
	return _MultiOwnable.Contract.Manager(&_MultiOwnable.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_MultiOwnable *MultiOwnableCallerSession) Manager() (common.Address, error) {
	return _MultiOwnable.Contract.Manager(&_MultiOwnable.CallOpts)
}

// OwnerByAddress is a free data retrieval call binding the contract method 0xeb6b192f.
//
// Solidity: function ownerByAddress( address) constant returns(bool)
func (_MultiOwnable *MultiOwnableCaller) OwnerByAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiOwnable.contract.Call(opts, out, "ownerByAddress", arg0)
	return *ret0, err
}

// OwnerByAddress is a free data retrieval call binding the contract method 0xeb6b192f.
//
// Solidity: function ownerByAddress( address) constant returns(bool)
func (_MultiOwnable *MultiOwnableSession) OwnerByAddress(arg0 common.Address) (bool, error) {
	return _MultiOwnable.Contract.OwnerByAddress(&_MultiOwnable.CallOpts, arg0)
}

// OwnerByAddress is a free data retrieval call binding the contract method 0xeb6b192f.
//
// Solidity: function ownerByAddress( address) constant returns(bool)
func (_MultiOwnable *MultiOwnableCallerSession) OwnerByAddress(arg0 common.Address) (bool, error) {
	return _MultiOwnable.Contract.OwnerByAddress(&_MultiOwnable.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_MultiOwnable *MultiOwnableCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultiOwnable.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_MultiOwnable *MultiOwnableSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiOwnable.Contract.Owners(&_MultiOwnable.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_MultiOwnable *MultiOwnableCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiOwnable.Contract.Owners(&_MultiOwnable.CallOpts, arg0)
}

// SetOwners is a paid mutator transaction binding the contract method 0xfa4d3698.
//
// Solidity: function setOwners(_owners address[]) returns()
func (_MultiOwnable *MultiOwnableTransactor) SetOwners(opts *bind.TransactOpts, _owners []common.Address) (*types.Transaction, error) {
	return _MultiOwnable.contract.Transact(opts, "setOwners", _owners)
}

// SetOwners is a paid mutator transaction binding the contract method 0xfa4d3698.
//
// Solidity: function setOwners(_owners address[]) returns()
func (_MultiOwnable *MultiOwnableSession) SetOwners(_owners []common.Address) (*types.Transaction, error) {
	return _MultiOwnable.Contract.SetOwners(&_MultiOwnable.TransactOpts, _owners)
}

// SetOwners is a paid mutator transaction binding the contract method 0xfa4d3698.
//
// Solidity: function setOwners(_owners address[]) returns()
func (_MultiOwnable *MultiOwnableTransactorSession) SetOwners(_owners []common.Address) (*types.Transaction, error) {
	return _MultiOwnable.Contract.SetOwners(&_MultiOwnable.TransactOpts, _owners)
}

// MultiOwnableSetOwnersIterator is returned from FilterSetOwners and is used to iterate over the raw logs and unpacked data for SetOwners events raised by the MultiOwnable contract.
type MultiOwnableSetOwnersIterator struct {
	Event *MultiOwnableSetOwners // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiOwnableSetOwnersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiOwnableSetOwners)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiOwnableSetOwners)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiOwnableSetOwnersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiOwnableSetOwnersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiOwnableSetOwners represents a SetOwners event raised by the MultiOwnable contract.
type MultiOwnableSetOwners struct {
	Owners []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetOwners is a free log retrieval operation binding the contract event 0x9465cd279c2de393c5568ae444599e3644e3d1864ca2c05ced8a654df2aea3cb.
//
// Solidity: e SetOwners(owners address[])
func (_MultiOwnable *MultiOwnableFilterer) FilterSetOwners(opts *bind.FilterOpts) (*MultiOwnableSetOwnersIterator, error) {

	logs, sub, err := _MultiOwnable.contract.FilterLogs(opts, "SetOwners")
	if err != nil {
		return nil, err
	}
	return &MultiOwnableSetOwnersIterator{contract: _MultiOwnable.contract, event: "SetOwners", logs: logs, sub: sub}, nil
}

// WatchSetOwners is a free log subscription operation binding the contract event 0x9465cd279c2de393c5568ae444599e3644e3d1864ca2c05ced8a654df2aea3cb.
//
// Solidity: e SetOwners(owners address[])
func (_MultiOwnable *MultiOwnableFilterer) WatchSetOwners(opts *bind.WatchOpts, sink chan<- *MultiOwnableSetOwners) (event.Subscription, error) {

	logs, sub, err := _MultiOwnable.contract.WatchLogs(opts, "SetOwners")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiOwnableSetOwners)
				if err := _MultiOwnable.contract.UnpackLog(event, "SetOwners", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// NtfTokenABI is the input ABI used to generate the binding from.
const NtfTokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sendTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownerByAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"setOwners\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenOwner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"SetOwners\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// NtfTokenBin is the compiled bytecode used for deploying new contracts.
const NtfTokenBin = `0x60806040526004805460ff1916905534801561001a57600080fd5b506040516020806110468339810180604052602081101561003a57600080fd5b505160018054600160a060020a031990811633179091556a084595161401484a000000600081815560078054600160a060020a038087169190951617808255841682526006602090815260408084209490945590549093168082528282205483519081529251909391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92908290030190a350610f69806100dd6000396000f3fe608060405234801561001057600080fd5b506004361061015f576000357c01000000000000000000000000000000000000000000000000000000009004806366188463116100d5578063a9059cbb11610099578063a9059cbb146103ba578063bff99c6c146103e6578063d73dd623146103ee578063dd62ed3e1461041a578063eb6b192f14610448578063fa4d36981461046e5761015f565b8063661884631461030057806370a082311461032c5780638456cb591461035257806395d89b411461035a578063a0e67e2b146103625761015f565b806323b872dd1161012757806323b872dd146102a25780632ff2e9dc146102d8578063313ce567146102e05780633f4ba83a146102e8578063481c6a75146102f05780635c975abb146102f85761015f565b8063025e7c271461016457806305ab421d1461019d57806306fdde03146101cb578063095ea7b31461024857806318160ddd14610288575b600080fd5b6101816004803603602081101561017a57600080fd5b5035610511565b60408051600160a060020a039092168252519081900360200190f35b6101c9600480360360408110156101b357600080fd5b50600160a060020a038135169060200135610539565b005b6101d3610641565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561020d5781810151838201526020016101f5565b50505050905090810190601f16801561023a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102746004803603604081101561025e57600080fd5b50600160a060020a038135169060200135610678565b604080519115158252519081900360200190f35b6102906106de565b60408051918252519081900360200190f35b610274600480360360608110156102b857600080fd5b50600160a060020a038135811691602081013590911690604001356106e4565b610290610849565b610290610858565b61027461085d565b6101816108c9565b6102746108d8565b6102746004803603604081101561031657600080fd5b50600160a060020a0381351690602001356108e1565b6102906004803603602081101561034257600080fd5b5035600160a060020a03166109d1565b6102746109ec565b6101d3610a59565b61036a610a90565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103a657818101518382015260200161038e565b505050509050019250505060405180910390f35b610274600480360360408110156103d057600080fd5b50600160a060020a038135169060200135610af2565b610181610c1b565b6102746004803603604081101561040457600080fd5b50600160a060020a038135169060200135610c2a565b6102906004803603604081101561043057600080fd5b50600160a060020a0381358116916020013516610cc3565b6102746004803603602081101561045e57600080fd5b5035600160a060020a0316610cee565b6101c96004803603602081101561048457600080fd5b81019060208101813564010000000081111561049f57600080fd5b8201836020820111156104b157600080fd5b803590602001918460208302840111640100000000831117156104d357600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610d03945050505050565b600280548290811061051f57fe5b600091825260209091200154600160a060020a0316905081565b3360009081526003602052604090205460ff16151560011461055a57600080fd5b600754600160a060020a031660009081526006602052604090205481111561058157600080fd5b600754600160a060020a03166000908152600660205260409020546105ac908263ffffffff610d2616565b600754600160a060020a0390811660009081526006602052604080822093909355908416815220546105e4908263ffffffff610d3b16565b600160a060020a03808416600081815260066020908152604091829020949094556007548151868152915192949316927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a35050565b60408051808201909152601681527f4e6578747920546f6b656e20466f756e646174696f6e00000000000000000000602082015281565b336000818152600560209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60005481565b60045460009060ff16156106f757600080fd5b600160a060020a038316151561070c57600080fd5b600160a060020a038416600090815260056020908152604080832033845290915290205482111561073c57600080fd5b600160a060020a038416600090815260066020526040902054610765908363ffffffff610d2616565b600160a060020a03808616600090815260066020526040808220939093559085168152205461079a908363ffffffff610d3b16565b600160a060020a0380851660009081526006602090815260408083209490945591871681526005825282812033825290915220546107de908363ffffffff610d2616565b600160a060020a03808616600081815260056020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b6a084595161401484a00000081565b601281565b3360009081526003602052604081205460ff16151560011461087e57600080fd5b60045460ff16151561088f57600080fd5b6004805460ff191690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a15060015b90565b600154600160a060020a031681565b60045460ff1681565b336000908152600560209081526040808320600160a060020a03861684529091528120548083111561093657336000908152600560209081526040808320600160a060020a038816845290915281205561096b565b610946818463ffffffff610d2616565b336000908152600560209081526040808320600160a060020a03891684529091529020555b336000818152600560209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a031660009081526006602052604090205490565b3360009081526003602052604081205460ff161515600114610a0d57600080fd5b60045460ff1615610a1d57600080fd5b6004805460ff191660011790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a150600190565b60408051808201909152600381527f4e54460000000000000000000000000000000000000000000000000000000000602082015281565b60606002805480602002602001604051908101604052809291908181526020018280548015610ae857602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610aca575b5050505050905090565b60045460009060ff1615610b0557600080fd5b600160a060020a0383161515610b1a57600080fd5b336000908152600660205260409020548211801590610b605750600160a060020a038316600090815260066020526040902054610b5d818463ffffffff610d3b16565b10155b1515610b6b57600080fd5b33600090815260066020526040902054610b8b908363ffffffff610d2616565b3360009081526006602052604080822092909255600160a060020a03851681522054610bbd908363ffffffff610d3b16565b600160a060020a0384166000818152600660209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b600754600160a060020a031681565b336000908152600560209081526040808320600160a060020a0386168452909152812054610c5e908363ffffffff610d3b16565b336000818152600560209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03918216600090815260056020908152604080832093909416825291909152205490565b60036020526000908152604090205460ff1681565b600154600160a060020a03163314610d1a57600080fd5b610d2381610d54565b50565b600082821115610d3557600080fd5b50900390565b600082820183811015610d4d57600080fd5b9392505050565b60005b600254811015610db257600060036000600284815481101515610d7657fe5b600091825260208083209190910154600160a060020a031683528201929092526040019020805460ff1916911515919091179055600101610d57565b5060005b8151811015610e0d576001600360008484815181101515610dd357fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff1916911515919091179055600101610db6565b508051610e21906002906020840190610e9a565b507f9465cd279c2de393c5568ae444599e3644e3d1864ca2c05ced8a654df2aea3cb816040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610e84578181015183820152602001610e6c565b505050509050019250505060405180910390a150565b828054828255906000526020600020908101928215610efc579160200282015b82811115610efc578251825473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909116178255602090920191600190910190610eba565b50610f08929150610f0c565b5090565b6108c691905b80821115610f0857805473ffffffffffffffffffffffffffffffffffffffff19168155600101610f1256fea165627a7a72305820d724e9b6e70410a35f6c2d0340666172e08b468c5d3f9ac8f7cc055823257fc00029`

// DeployNtfToken deploys a new Ethereum contract, binding an instance of NtfToken to it.
func DeployNtfToken(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenOwner common.Address) (common.Address, *types.Transaction, *NtfToken, error) {
	parsed, err := abi.JSON(strings.NewReader(NtfTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NtfTokenBin), backend, _tokenOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NtfToken{NtfTokenCaller: NtfTokenCaller{contract: contract}, NtfTokenTransactor: NtfTokenTransactor{contract: contract}, NtfTokenFilterer: NtfTokenFilterer{contract: contract}}, nil
}

// NtfToken is an auto generated Go binding around an Ethereum contract.
type NtfToken struct {
	NtfTokenCaller     // Read-only binding to the contract
	NtfTokenTransactor // Write-only binding to the contract
	NtfTokenFilterer   // Log filterer for contract events
}

// NtfTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type NtfTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NtfTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NtfTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NtfTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NtfTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NtfTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NtfTokenSession struct {
	Contract     *NtfToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NtfTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NtfTokenCallerSession struct {
	Contract *NtfTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// NtfTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NtfTokenTransactorSession struct {
	Contract     *NtfTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// NtfTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type NtfTokenRaw struct {
	Contract *NtfToken // Generic contract binding to access the raw methods on
}

// NtfTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NtfTokenCallerRaw struct {
	Contract *NtfTokenCaller // Generic read-only contract binding to access the raw methods on
}

// NtfTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NtfTokenTransactorRaw struct {
	Contract *NtfTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNtfToken creates a new instance of NtfToken, bound to a specific deployed contract.
func NewNtfToken(address common.Address, backend bind.ContractBackend) (*NtfToken, error) {
	contract, err := bindNtfToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NtfToken{NtfTokenCaller: NtfTokenCaller{contract: contract}, NtfTokenTransactor: NtfTokenTransactor{contract: contract}, NtfTokenFilterer: NtfTokenFilterer{contract: contract}}, nil
}

// NewNtfTokenCaller creates a new read-only instance of NtfToken, bound to a specific deployed contract.
func NewNtfTokenCaller(address common.Address, caller bind.ContractCaller) (*NtfTokenCaller, error) {
	contract, err := bindNtfToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NtfTokenCaller{contract: contract}, nil
}

// NewNtfTokenTransactor creates a new write-only instance of NtfToken, bound to a specific deployed contract.
func NewNtfTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*NtfTokenTransactor, error) {
	contract, err := bindNtfToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NtfTokenTransactor{contract: contract}, nil
}

// NewNtfTokenFilterer creates a new log filterer instance of NtfToken, bound to a specific deployed contract.
func NewNtfTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*NtfTokenFilterer, error) {
	contract, err := bindNtfToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NtfTokenFilterer{contract: contract}, nil
}

// bindNtfToken binds a generic wrapper to an already deployed contract.
func bindNtfToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NtfTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NtfToken *NtfTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NtfToken.Contract.NtfTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NtfToken *NtfTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NtfToken.Contract.NtfTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NtfToken *NtfTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NtfToken.Contract.NtfTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NtfToken *NtfTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NtfToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NtfToken *NtfTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NtfToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NtfToken *NtfTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NtfToken.Contract.contract.Transact(opts, method, params...)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_NtfToken *NtfTokenCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NtfToken.contract.Call(opts, out, "INITIAL_SUPPLY")
	return *ret0, err
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_NtfToken *NtfTokenSession) INITIALSUPPLY() (*big.Int, error) {
	return _NtfToken.Contract.INITIALSUPPLY(&_NtfToken.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_NtfToken *NtfTokenCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _NtfToken.Contract.INITIALSUPPLY(&_NtfToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_NtfToken *NtfTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NtfToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_NtfToken *NtfTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _NtfToken.Contract.Allowance(&_NtfToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_NtfToken *NtfTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _NtfToken.Contract.Allowance(&_NtfToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_NtfToken *NtfTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NtfToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_NtfToken *NtfTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _NtfToken.Contract.BalanceOf(&_NtfToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_NtfToken *NtfTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _NtfToken.Contract.BalanceOf(&_NtfToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_NtfToken *NtfTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NtfToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_NtfToken *NtfTokenSession) Decimals() (*big.Int, error) {
	return _NtfToken.Contract.Decimals(&_NtfToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_NtfToken *NtfTokenCallerSession) Decimals() (*big.Int, error) {
	return _NtfToken.Contract.Decimals(&_NtfToken.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_NtfToken *NtfTokenCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _NtfToken.contract.Call(opts, out, "getOwners")
	return *ret0, err
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_NtfToken *NtfTokenSession) GetOwners() ([]common.Address, error) {
	return _NtfToken.Contract.GetOwners(&_NtfToken.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_NtfToken *NtfTokenCallerSession) GetOwners() ([]common.Address, error) {
	return _NtfToken.Contract.GetOwners(&_NtfToken.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_NtfToken *NtfTokenCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NtfToken.contract.Call(opts, out, "manager")
	return *ret0, err
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_NtfToken *NtfTokenSession) Manager() (common.Address, error) {
	return _NtfToken.Contract.Manager(&_NtfToken.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_NtfToken *NtfTokenCallerSession) Manager() (common.Address, error) {
	return _NtfToken.Contract.Manager(&_NtfToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_NtfToken *NtfTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NtfToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_NtfToken *NtfTokenSession) Name() (string, error) {
	return _NtfToken.Contract.Name(&_NtfToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_NtfToken *NtfTokenCallerSession) Name() (string, error) {
	return _NtfToken.Contract.Name(&_NtfToken.CallOpts)
}

// OwnerByAddress is a free data retrieval call binding the contract method 0xeb6b192f.
//
// Solidity: function ownerByAddress( address) constant returns(bool)
func (_NtfToken *NtfTokenCaller) OwnerByAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NtfToken.contract.Call(opts, out, "ownerByAddress", arg0)
	return *ret0, err
}

// OwnerByAddress is a free data retrieval call binding the contract method 0xeb6b192f.
//
// Solidity: function ownerByAddress( address) constant returns(bool)
func (_NtfToken *NtfTokenSession) OwnerByAddress(arg0 common.Address) (bool, error) {
	return _NtfToken.Contract.OwnerByAddress(&_NtfToken.CallOpts, arg0)
}

// OwnerByAddress is a free data retrieval call binding the contract method 0xeb6b192f.
//
// Solidity: function ownerByAddress( address) constant returns(bool)
func (_NtfToken *NtfTokenCallerSession) OwnerByAddress(arg0 common.Address) (bool, error) {
	return _NtfToken.Contract.OwnerByAddress(&_NtfToken.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_NtfToken *NtfTokenCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NtfToken.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_NtfToken *NtfTokenSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _NtfToken.Contract.Owners(&_NtfToken.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_NtfToken *NtfTokenCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _NtfToken.Contract.Owners(&_NtfToken.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_NtfToken *NtfTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NtfToken.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_NtfToken *NtfTokenSession) Paused() (bool, error) {
	return _NtfToken.Contract.Paused(&_NtfToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_NtfToken *NtfTokenCallerSession) Paused() (bool, error) {
	return _NtfToken.Contract.Paused(&_NtfToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_NtfToken *NtfTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NtfToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_NtfToken *NtfTokenSession) Symbol() (string, error) {
	return _NtfToken.Contract.Symbol(&_NtfToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_NtfToken *NtfTokenCallerSession) Symbol() (string, error) {
	return _NtfToken.Contract.Symbol(&_NtfToken.CallOpts)
}

// TokenWallet is a free data retrieval call binding the contract method 0xbff99c6c.
//
// Solidity: function tokenWallet() constant returns(address)
func (_NtfToken *NtfTokenCaller) TokenWallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NtfToken.contract.Call(opts, out, "tokenWallet")
	return *ret0, err
}

// TokenWallet is a free data retrieval call binding the contract method 0xbff99c6c.
//
// Solidity: function tokenWallet() constant returns(address)
func (_NtfToken *NtfTokenSession) TokenWallet() (common.Address, error) {
	return _NtfToken.Contract.TokenWallet(&_NtfToken.CallOpts)
}

// TokenWallet is a free data retrieval call binding the contract method 0xbff99c6c.
//
// Solidity: function tokenWallet() constant returns(address)
func (_NtfToken *NtfTokenCallerSession) TokenWallet() (common.Address, error) {
	return _NtfToken.Contract.TokenWallet(&_NtfToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_NtfToken *NtfTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NtfToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_NtfToken *NtfTokenSession) TotalSupply() (*big.Int, error) {
	return _NtfToken.Contract.TotalSupply(&_NtfToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_NtfToken *NtfTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _NtfToken.Contract.TotalSupply(&_NtfToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_NtfToken *NtfTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NtfToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_NtfToken *NtfTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NtfToken.Contract.Approve(&_NtfToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_NtfToken *NtfTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NtfToken.Contract.Approve(&_NtfToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_NtfToken *NtfTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _NtfToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_NtfToken *NtfTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _NtfToken.Contract.DecreaseApproval(&_NtfToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_NtfToken *NtfTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _NtfToken.Contract.DecreaseApproval(&_NtfToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_NtfToken *NtfTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _NtfToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_NtfToken *NtfTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _NtfToken.Contract.IncreaseApproval(&_NtfToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_NtfToken *NtfTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _NtfToken.Contract.IncreaseApproval(&_NtfToken.TransactOpts, _spender, _addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_NtfToken *NtfTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NtfToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_NtfToken *NtfTokenSession) Pause() (*types.Transaction, error) {
	return _NtfToken.Contract.Pause(&_NtfToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_NtfToken *NtfTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _NtfToken.Contract.Pause(&_NtfToken.TransactOpts)
}

// SendTokens is a paid mutator transaction binding the contract method 0x05ab421d.
//
// Solidity: function sendTokens(_to address, _amount uint256) returns()
func (_NtfToken *NtfTokenTransactor) SendTokens(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _NtfToken.contract.Transact(opts, "sendTokens", _to, _amount)
}

// SendTokens is a paid mutator transaction binding the contract method 0x05ab421d.
//
// Solidity: function sendTokens(_to address, _amount uint256) returns()
func (_NtfToken *NtfTokenSession) SendTokens(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _NtfToken.Contract.SendTokens(&_NtfToken.TransactOpts, _to, _amount)
}

// SendTokens is a paid mutator transaction binding the contract method 0x05ab421d.
//
// Solidity: function sendTokens(_to address, _amount uint256) returns()
func (_NtfToken *NtfTokenTransactorSession) SendTokens(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _NtfToken.Contract.SendTokens(&_NtfToken.TransactOpts, _to, _amount)
}

// SetOwners is a paid mutator transaction binding the contract method 0xfa4d3698.
//
// Solidity: function setOwners(_owners address[]) returns()
func (_NtfToken *NtfTokenTransactor) SetOwners(opts *bind.TransactOpts, _owners []common.Address) (*types.Transaction, error) {
	return _NtfToken.contract.Transact(opts, "setOwners", _owners)
}

// SetOwners is a paid mutator transaction binding the contract method 0xfa4d3698.
//
// Solidity: function setOwners(_owners address[]) returns()
func (_NtfToken *NtfTokenSession) SetOwners(_owners []common.Address) (*types.Transaction, error) {
	return _NtfToken.Contract.SetOwners(&_NtfToken.TransactOpts, _owners)
}

// SetOwners is a paid mutator transaction binding the contract method 0xfa4d3698.
//
// Solidity: function setOwners(_owners address[]) returns()
func (_NtfToken *NtfTokenTransactorSession) SetOwners(_owners []common.Address) (*types.Transaction, error) {
	return _NtfToken.Contract.SetOwners(&_NtfToken.TransactOpts, _owners)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_NtfToken *NtfTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NtfToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_NtfToken *NtfTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NtfToken.Contract.Transfer(&_NtfToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_NtfToken *NtfTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NtfToken.Contract.Transfer(&_NtfToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_NtfToken *NtfTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NtfToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_NtfToken *NtfTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NtfToken.Contract.TransferFrom(&_NtfToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_NtfToken *NtfTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _NtfToken.Contract.TransferFrom(&_NtfToken.TransactOpts, _from, _to, _value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_NtfToken *NtfTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NtfToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_NtfToken *NtfTokenSession) Unpause() (*types.Transaction, error) {
	return _NtfToken.Contract.Unpause(&_NtfToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_NtfToken *NtfTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _NtfToken.Contract.Unpause(&_NtfToken.TransactOpts)
}

// NtfTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NtfToken contract.
type NtfTokenApprovalIterator struct {
	Event *NtfTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NtfTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NtfTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NtfTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NtfTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NtfTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NtfTokenApproval represents a Approval event raised by the NtfToken contract.
type NtfTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_NtfToken *NtfTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NtfTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NtfToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NtfTokenApprovalIterator{contract: _NtfToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_NtfToken *NtfTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NtfTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NtfToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NtfTokenApproval)
				if err := _NtfToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// NtfTokenPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the NtfToken contract.
type NtfTokenPauseIterator struct {
	Event *NtfTokenPause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NtfTokenPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NtfTokenPause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NtfTokenPause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NtfTokenPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NtfTokenPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NtfTokenPause represents a Pause event raised by the NtfToken contract.
type NtfTokenPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_NtfToken *NtfTokenFilterer) FilterPause(opts *bind.FilterOpts) (*NtfTokenPauseIterator, error) {

	logs, sub, err := _NtfToken.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &NtfTokenPauseIterator{contract: _NtfToken.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_NtfToken *NtfTokenFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *NtfTokenPause) (event.Subscription, error) {

	logs, sub, err := _NtfToken.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NtfTokenPause)
				if err := _NtfToken.contract.UnpackLog(event, "Pause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// NtfTokenSetOwnersIterator is returned from FilterSetOwners and is used to iterate over the raw logs and unpacked data for SetOwners events raised by the NtfToken contract.
type NtfTokenSetOwnersIterator struct {
	Event *NtfTokenSetOwners // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NtfTokenSetOwnersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NtfTokenSetOwners)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NtfTokenSetOwners)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NtfTokenSetOwnersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NtfTokenSetOwnersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NtfTokenSetOwners represents a SetOwners event raised by the NtfToken contract.
type NtfTokenSetOwners struct {
	Owners []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetOwners is a free log retrieval operation binding the contract event 0x9465cd279c2de393c5568ae444599e3644e3d1864ca2c05ced8a654df2aea3cb.
//
// Solidity: e SetOwners(owners address[])
func (_NtfToken *NtfTokenFilterer) FilterSetOwners(opts *bind.FilterOpts) (*NtfTokenSetOwnersIterator, error) {

	logs, sub, err := _NtfToken.contract.FilterLogs(opts, "SetOwners")
	if err != nil {
		return nil, err
	}
	return &NtfTokenSetOwnersIterator{contract: _NtfToken.contract, event: "SetOwners", logs: logs, sub: sub}, nil
}

// WatchSetOwners is a free log subscription operation binding the contract event 0x9465cd279c2de393c5568ae444599e3644e3d1864ca2c05ced8a654df2aea3cb.
//
// Solidity: e SetOwners(owners address[])
func (_NtfToken *NtfTokenFilterer) WatchSetOwners(opts *bind.WatchOpts, sink chan<- *NtfTokenSetOwners) (event.Subscription, error) {

	logs, sub, err := _NtfToken.contract.WatchLogs(opts, "SetOwners")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NtfTokenSetOwners)
				if err := _NtfToken.contract.UnpackLog(event, "SetOwners", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// NtfTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NtfToken contract.
type NtfTokenTransferIterator struct {
	Event *NtfTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NtfTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NtfTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NtfTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NtfTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NtfTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NtfTokenTransfer represents a Transfer event raised by the NtfToken contract.
type NtfTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_NtfToken *NtfTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NtfTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NtfToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NtfTokenTransferIterator{contract: _NtfToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_NtfToken *NtfTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NtfTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NtfToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NtfTokenTransfer)
				if err := _NtfToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// NtfTokenUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the NtfToken contract.
type NtfTokenUnpauseIterator struct {
	Event *NtfTokenUnpause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NtfTokenUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NtfTokenUnpause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NtfTokenUnpause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NtfTokenUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NtfTokenUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NtfTokenUnpause represents a Unpause event raised by the NtfToken contract.
type NtfTokenUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_NtfToken *NtfTokenFilterer) FilterUnpause(opts *bind.FilterOpts) (*NtfTokenUnpauseIterator, error) {

	logs, sub, err := _NtfToken.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &NtfTokenUnpauseIterator{contract: _NtfToken.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_NtfToken *NtfTokenFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *NtfTokenUnpause) (event.Subscription, error) {

	logs, sub, err := _NtfToken.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NtfTokenUnpause)
				if err := _NtfToken.contract.UnpackLog(event, "Unpause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownerByAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"setOwners\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"SetOwners\",\"type\":\"event\"}]"

// PausableBin is the compiled bytecode used for deploying new contracts.
const PausableBin = `0x60806040526003805460ff1916905560008054600160a060020a031916331790556106028061002f6000396000f3fe608060405234801561001057600080fd5b50600436106100a5576000357c0100000000000000000000000000000000000000000000000000000000900480638456cb59116100785780638456cb591461010f578063a0e67e2b14610117578063eb6b192f1461016f578063fa4d369814610195576100a5565b8063025e7c27146100aa5780633f4ba83a146100e3578063481c6a75146100ff5780635c975abb14610107575b600080fd5b6100c7600480360360208110156100c057600080fd5b503561023a565b60408051600160a060020a039092168252519081900360200190f35b6100eb610262565b604080519115158252519081900360200190f35b6100c76102ce565b6100eb6102dd565b6100eb6102e6565b61011f610353565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561015b578181015183820152602001610143565b505050509050019250505060405180910390f35b6100eb6004803603602081101561018557600080fd5b5035600160a060020a03166103b5565b610238600480360360208110156101ab57600080fd5b8101906020810181356401000000008111156101c657600080fd5b8201836020820111156101d857600080fd5b803590602001918460208302840111640100000000831117156101fa57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506103ca945050505050565b005b600180548290811061024857fe5b600091825260209091200154600160a060020a0316905081565b3360009081526002602052604081205460ff16151560011461028357600080fd5b60035460ff16151561029457600080fd5b6003805460ff191690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a15060015b90565b600054600160a060020a031681565b60035460ff1681565b3360009081526002602052604081205460ff16151560011461030757600080fd5b60035460ff161561031757600080fd5b6003805460ff191660011790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a150600190565b606060018054806020026020016040519081016040528092919081815260200182805480156103ab57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161038d575b5050505050905090565b60026020526000908152604090205460ff1681565b600054600160a060020a031633146103e157600080fd5b6103ea816103ed565b50565b60005b60015481101561044b5760006002600060018481548110151561040f57fe5b600091825260208083209190910154600160a060020a031683528201929092526040019020805460ff19169115159190911790556001016103f0565b5060005b81518110156104a657600160026000848481518110151561046c57fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff191691151591909117905560010161044f565b5080516104ba906001906020840190610533565b507f9465cd279c2de393c5568ae444599e3644e3d1864ca2c05ced8a654df2aea3cb816040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561051d578181015183820152602001610505565b505050509050019250505060405180910390a150565b828054828255906000526020600020908101928215610595579160200282015b82811115610595578251825473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909116178255602090920191600190910190610553565b506105a19291506105a5565b5090565b6102cb91905b808211156105a157805473ffffffffffffffffffffffffffffffffffffffff191681556001016105ab56fea165627a7a7230582034f24d389c795864914a387b22d1bbe095a95b07e0818b72a1a6a08f61a0a89a0029`

// DeployPausable deploys a new Ethereum contract, binding an instance of Pausable to it.
func DeployPausable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pausable, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_Pausable *PausableCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "getOwners")
	return *ret0, err
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_Pausable *PausableSession) GetOwners() ([]common.Address, error) {
	return _Pausable.Contract.GetOwners(&_Pausable.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_Pausable *PausableCallerSession) GetOwners() ([]common.Address, error) {
	return _Pausable.Contract.GetOwners(&_Pausable.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Pausable *PausableCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "manager")
	return *ret0, err
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Pausable *PausableSession) Manager() (common.Address, error) {
	return _Pausable.Contract.Manager(&_Pausable.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_Pausable *PausableCallerSession) Manager() (common.Address, error) {
	return _Pausable.Contract.Manager(&_Pausable.CallOpts)
}

// OwnerByAddress is a free data retrieval call binding the contract method 0xeb6b192f.
//
// Solidity: function ownerByAddress( address) constant returns(bool)
func (_Pausable *PausableCaller) OwnerByAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "ownerByAddress", arg0)
	return *ret0, err
}

// OwnerByAddress is a free data retrieval call binding the contract method 0xeb6b192f.
//
// Solidity: function ownerByAddress( address) constant returns(bool)
func (_Pausable *PausableSession) OwnerByAddress(arg0 common.Address) (bool, error) {
	return _Pausable.Contract.OwnerByAddress(&_Pausable.CallOpts, arg0)
}

// OwnerByAddress is a free data retrieval call binding the contract method 0xeb6b192f.
//
// Solidity: function ownerByAddress( address) constant returns(bool)
func (_Pausable *PausableCallerSession) OwnerByAddress(arg0 common.Address) (bool, error) {
	return _Pausable.Contract.OwnerByAddress(&_Pausable.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_Pausable *PausableCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_Pausable *PausableSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Pausable.Contract.Owners(&_Pausable.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_Pausable *PausableCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Pausable.Contract.Owners(&_Pausable.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_Pausable *PausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_Pausable *PausableSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_Pausable *PausableTransactorSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// SetOwners is a paid mutator transaction binding the contract method 0xfa4d3698.
//
// Solidity: function setOwners(_owners address[]) returns()
func (_Pausable *PausableTransactor) SetOwners(opts *bind.TransactOpts, _owners []common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "setOwners", _owners)
}

// SetOwners is a paid mutator transaction binding the contract method 0xfa4d3698.
//
// Solidity: function setOwners(_owners address[]) returns()
func (_Pausable *PausableSession) SetOwners(_owners []common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.SetOwners(&_Pausable.TransactOpts, _owners)
}

// SetOwners is a paid mutator transaction binding the contract method 0xfa4d3698.
//
// Solidity: function setOwners(_owners address[]) returns()
func (_Pausable *PausableTransactorSession) SetOwners(_owners []common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.SetOwners(&_Pausable.TransactOpts, _owners)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_Pausable *PausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_Pausable *PausableSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_Pausable *PausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// PausablePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Pausable contract.
type PausablePauseIterator struct {
	Event *PausablePause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PausablePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PausablePause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PausablePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePause represents a Pause event raised by the Pausable contract.
type PausablePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Pausable *PausableFilterer) FilterPause(opts *bind.FilterOpts) (*PausablePauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausablePauseIterator{contract: _Pausable.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Pausable *PausableFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausablePause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePause)
				if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PausableSetOwnersIterator is returned from FilterSetOwners and is used to iterate over the raw logs and unpacked data for SetOwners events raised by the Pausable contract.
type PausableSetOwnersIterator struct {
	Event *PausableSetOwners // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PausableSetOwnersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableSetOwners)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PausableSetOwners)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PausableSetOwnersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableSetOwnersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableSetOwners represents a SetOwners event raised by the Pausable contract.
type PausableSetOwners struct {
	Owners []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetOwners is a free log retrieval operation binding the contract event 0x9465cd279c2de393c5568ae444599e3644e3d1864ca2c05ced8a654df2aea3cb.
//
// Solidity: e SetOwners(owners address[])
func (_Pausable *PausableFilterer) FilterSetOwners(opts *bind.FilterOpts) (*PausableSetOwnersIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "SetOwners")
	if err != nil {
		return nil, err
	}
	return &PausableSetOwnersIterator{contract: _Pausable.contract, event: "SetOwners", logs: logs, sub: sub}, nil
}

// WatchSetOwners is a free log subscription operation binding the contract event 0x9465cd279c2de393c5568ae444599e3644e3d1864ca2c05ced8a654df2aea3cb.
//
// Solidity: e SetOwners(owners address[])
func (_Pausable *PausableFilterer) WatchSetOwners(opts *bind.WatchOpts, sink chan<- *PausableSetOwners) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "SetOwners")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableSetOwners)
				if err := _Pausable.contract.UnpackLog(event, "SetOwners", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PausableUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Pausable contract.
type PausableUnpauseIterator struct {
	Event *PausableUnpause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PausableUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PausableUnpause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PausableUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpause represents a Unpause event raised by the Pausable contract.
type PausableUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Pausable *PausableFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableUnpauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableUnpauseIterator{contract: _Pausable.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Pausable *PausableFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableUnpause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpause)
				if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058206965b015309963ced9c7240e03d655d9681b2485065843e26584782c410478450029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownerByAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"setOwners\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"SetOwners\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x60806040526004805460ff1916905560018054600160a060020a03191633179055610cc58061002f6000396000f3fe608060405234801561001057600080fd5b506004361061011d576000357c01000000000000000000000000000000000000000000000000000000009004806370a08231116100b4578063d73dd62311610083578063d73dd623146102e1578063dd62ed3e1461030d578063eb6b192f1461033b578063fa4d3698146103615761011d565b806370a082311461022f5780638456cb5914610255578063a0e67e2b1461025d578063a9059cbb146102b55761011d565b80633f4ba83a116100f05780633f4ba83a146101eb578063481c6a75146101f35780635c975abb146101fb57806366188463146102035761011d565b8063025e7c2714610122578063095ea7b31461015b57806318160ddd1461019b57806323b872dd146101b5575b600080fd5b61013f6004803603602081101561013857600080fd5b5035610406565b60408051600160a060020a039092168252519081900360200190f35b6101876004803603604081101561017157600080fd5b50600160a060020a03813516906020013561042e565b604080519115158252519081900360200190f35b6101a3610494565b60408051918252519081900360200190f35b610187600480360360608110156101cb57600080fd5b50600160a060020a0381358116916020810135909116906040013561049a565b6101876105ff565b61013f61066b565b61018761067a565b6101876004803603604081101561021957600080fd5b50600160a060020a038135169060200135610683565b6101a36004803603602081101561024557600080fd5b5035600160a060020a0316610773565b61018761078e565b6102656107fb565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156102a1578181015183820152602001610289565b505050509050019250505060405180910390f35b610187600480360360408110156102cb57600080fd5b50600160a060020a03813516906020013561085d565b610187600480360360408110156102f757600080fd5b50600160a060020a038135169060200135610986565b6101a36004803603604081101561032357600080fd5b50600160a060020a0381358116916020013516610a1f565b6101876004803603602081101561035157600080fd5b5035600160a060020a0316610a4a565b6104046004803603602081101561037757600080fd5b81019060208101813564010000000081111561039257600080fd5b8201836020820111156103a457600080fd5b803590602001918460208302840111640100000000831117156103c657600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610a5f945050505050565b005b600280548290811061041457fe5b600091825260209091200154600160a060020a0316905081565b336000818152600560209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60005481565b60045460009060ff16156104ad57600080fd5b600160a060020a03831615156104c257600080fd5b600160a060020a03841660009081526005602090815260408083203384529091529020548211156104f257600080fd5b600160a060020a03841660009081526006602052604090205461051b908363ffffffff610a8216565b600160a060020a038086166000908152600660205260408082209390935590851681522054610550908363ffffffff610a9716565b600160a060020a038085166000908152600660209081526040808320949094559187168152600582528281203382529091522054610594908363ffffffff610a8216565b600160a060020a03808616600081815260056020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b3360009081526003602052604081205460ff16151560011461062057600080fd5b60045460ff16151561063157600080fd5b6004805460ff191690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a15060015b90565b600154600160a060020a031681565b60045460ff1681565b336000908152600560209081526040808320600160a060020a0386168452909152812054808311156106d857336000908152600560209081526040808320600160a060020a038816845290915281205561070d565b6106e8818463ffffffff610a8216565b336000908152600560209081526040808320600160a060020a03891684529091529020555b336000818152600560209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a031660009081526006602052604090205490565b3360009081526003602052604081205460ff1615156001146107af57600080fd5b60045460ff16156107bf57600080fd5b6004805460ff191660011790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a150600190565b6060600280548060200260200160405190810160405280929190818152602001828054801561085357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610835575b5050505050905090565b60045460009060ff161561087057600080fd5b600160a060020a038316151561088557600080fd5b3360009081526006602052604090205482118015906108cb5750600160a060020a0383166000908152600660205260409020546108c8818463ffffffff610a9716565b10155b15156108d657600080fd5b336000908152600660205260409020546108f6908363ffffffff610a8216565b3360009081526006602052604080822092909255600160a060020a03851681522054610928908363ffffffff610a9716565b600160a060020a0384166000818152600660209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b336000908152600560209081526040808320600160a060020a03861684529091528120546109ba908363ffffffff610a9716565b336000818152600560209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03918216600090815260056020908152604080832093909416825291909152205490565b60036020526000908152604090205460ff1681565b600154600160a060020a03163314610a7657600080fd5b610a7f81610ab0565b50565b600082821115610a9157600080fd5b50900390565b600082820183811015610aa957600080fd5b9392505050565b60005b600254811015610b0e57600060036000600284815481101515610ad257fe5b600091825260208083209190910154600160a060020a031683528201929092526040019020805460ff1916911515919091179055600101610ab3565b5060005b8151811015610b69576001600360008484815181101515610b2f57fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff1916911515919091179055600101610b12565b508051610b7d906002906020840190610bf6565b507f9465cd279c2de393c5568ae444599e3644e3d1864ca2c05ced8a654df2aea3cb816040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610be0578181015183820152602001610bc8565b505050509050019250505060405180910390a150565b828054828255906000526020600020908101928215610c58579160200282015b82811115610c58578251825473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909116178255602090920191600190910190610c16565b50610c64929150610c68565b5090565b61066891905b80821115610c6457805473ffffffffffffffffffffffffffffffffffffffff19168155600101610c6e56fea165627a7a72305820439f34d8511fce17388c0af4b769748dd75349d22d4d4efe94667e378975b9460029`

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
	StandardTokenFilterer   // Log filterer for contract events
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// NewStandardTokenFilterer creates a new log filterer instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenFilterer, error) {
	contract, err := bindStandardToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFilterer{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_StandardToken *StandardTokenCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "getOwners")
	return *ret0, err
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_StandardToken *StandardTokenSession) GetOwners() ([]common.Address, error) {
	return _StandardToken.Contract.GetOwners(&_StandardToken.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_StandardToken *StandardTokenCallerSession) GetOwners() ([]common.Address, error) {
	return _StandardToken.Contract.GetOwners(&_StandardToken.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_StandardToken *StandardTokenCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "manager")
	return *ret0, err
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_StandardToken *StandardTokenSession) Manager() (common.Address, error) {
	return _StandardToken.Contract.Manager(&_StandardToken.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_StandardToken *StandardTokenCallerSession) Manager() (common.Address, error) {
	return _StandardToken.Contract.Manager(&_StandardToken.CallOpts)
}

// OwnerByAddress is a free data retrieval call binding the contract method 0xeb6b192f.
//
// Solidity: function ownerByAddress( address) constant returns(bool)
func (_StandardToken *StandardTokenCaller) OwnerByAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "ownerByAddress", arg0)
	return *ret0, err
}

// OwnerByAddress is a free data retrieval call binding the contract method 0xeb6b192f.
//
// Solidity: function ownerByAddress( address) constant returns(bool)
func (_StandardToken *StandardTokenSession) OwnerByAddress(arg0 common.Address) (bool, error) {
	return _StandardToken.Contract.OwnerByAddress(&_StandardToken.CallOpts, arg0)
}

// OwnerByAddress is a free data retrieval call binding the contract method 0xeb6b192f.
//
// Solidity: function ownerByAddress( address) constant returns(bool)
func (_StandardToken *StandardTokenCallerSession) OwnerByAddress(arg0 common.Address) (bool, error) {
	return _StandardToken.Contract.OwnerByAddress(&_StandardToken.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_StandardToken *StandardTokenCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_StandardToken *StandardTokenSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _StandardToken.Contract.Owners(&_StandardToken.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_StandardToken *StandardTokenCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _StandardToken.Contract.Owners(&_StandardToken.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StandardToken *StandardTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StandardToken *StandardTokenSession) Paused() (bool, error) {
	return _StandardToken.Contract.Paused(&_StandardToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_StandardToken *StandardTokenCallerSession) Paused() (bool, error) {
	return _StandardToken.Contract.Paused(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_StandardToken *StandardTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_StandardToken *StandardTokenSession) Pause() (*types.Transaction, error) {
	return _StandardToken.Contract.Pause(&_StandardToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _StandardToken.Contract.Pause(&_StandardToken.TransactOpts)
}

// SetOwners is a paid mutator transaction binding the contract method 0xfa4d3698.
//
// Solidity: function setOwners(_owners address[]) returns()
func (_StandardToken *StandardTokenTransactor) SetOwners(opts *bind.TransactOpts, _owners []common.Address) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "setOwners", _owners)
}

// SetOwners is a paid mutator transaction binding the contract method 0xfa4d3698.
//
// Solidity: function setOwners(_owners address[]) returns()
func (_StandardToken *StandardTokenSession) SetOwners(_owners []common.Address) (*types.Transaction, error) {
	return _StandardToken.Contract.SetOwners(&_StandardToken.TransactOpts, _owners)
}

// SetOwners is a paid mutator transaction binding the contract method 0xfa4d3698.
//
// Solidity: function setOwners(_owners address[]) returns()
func (_StandardToken *StandardTokenTransactorSession) SetOwners(_owners []common.Address) (*types.Transaction, error) {
	return _StandardToken.Contract.SetOwners(&_StandardToken.TransactOpts, _owners)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_StandardToken *StandardTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_StandardToken *StandardTokenSession) Unpause() (*types.Transaction, error) {
	return _StandardToken.Contract.Unpause(&_StandardToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _StandardToken.Contract.Unpause(&_StandardToken.TransactOpts)
}

// StandardTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardToken contract.
type StandardTokenApprovalIterator struct {
	Event *StandardTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StandardTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StandardTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenApproval represents a Approval event raised by the StandardToken contract.
type StandardTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandardTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenApprovalIterator{contract: _StandardToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenApproval)
				if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StandardTokenPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the StandardToken contract.
type StandardTokenPauseIterator struct {
	Event *StandardTokenPause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenPause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StandardTokenPause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StandardTokenPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenPause represents a Pause event raised by the StandardToken contract.
type StandardTokenPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_StandardToken *StandardTokenFilterer) FilterPause(opts *bind.FilterOpts) (*StandardTokenPauseIterator, error) {

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &StandardTokenPauseIterator{contract: _StandardToken.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_StandardToken *StandardTokenFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *StandardTokenPause) (event.Subscription, error) {

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenPause)
				if err := _StandardToken.contract.UnpackLog(event, "Pause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StandardTokenSetOwnersIterator is returned from FilterSetOwners and is used to iterate over the raw logs and unpacked data for SetOwners events raised by the StandardToken contract.
type StandardTokenSetOwnersIterator struct {
	Event *StandardTokenSetOwners // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenSetOwnersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenSetOwners)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StandardTokenSetOwners)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StandardTokenSetOwnersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenSetOwnersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenSetOwners represents a SetOwners event raised by the StandardToken contract.
type StandardTokenSetOwners struct {
	Owners []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetOwners is a free log retrieval operation binding the contract event 0x9465cd279c2de393c5568ae444599e3644e3d1864ca2c05ced8a654df2aea3cb.
//
// Solidity: e SetOwners(owners address[])
func (_StandardToken *StandardTokenFilterer) FilterSetOwners(opts *bind.FilterOpts) (*StandardTokenSetOwnersIterator, error) {

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "SetOwners")
	if err != nil {
		return nil, err
	}
	return &StandardTokenSetOwnersIterator{contract: _StandardToken.contract, event: "SetOwners", logs: logs, sub: sub}, nil
}

// WatchSetOwners is a free log subscription operation binding the contract event 0x9465cd279c2de393c5568ae444599e3644e3d1864ca2c05ced8a654df2aea3cb.
//
// Solidity: e SetOwners(owners address[])
func (_StandardToken *StandardTokenFilterer) WatchSetOwners(opts *bind.WatchOpts, sink chan<- *StandardTokenSetOwners) (event.Subscription, error) {

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "SetOwners")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenSetOwners)
				if err := _StandardToken.contract.UnpackLog(event, "SetOwners", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StandardTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardToken contract.
type StandardTokenTransferIterator struct {
	Event *StandardTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StandardTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StandardTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenTransfer represents a Transfer event raised by the StandardToken contract.
type StandardTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandardTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransferIterator{contract: _StandardToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenTransfer)
				if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StandardTokenUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the StandardToken contract.
type StandardTokenUnpauseIterator struct {
	Event *StandardTokenUnpause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenUnpause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StandardTokenUnpause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StandardTokenUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenUnpause represents a Unpause event raised by the StandardToken contract.
type StandardTokenUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_StandardToken *StandardTokenFilterer) FilterUnpause(opts *bind.FilterOpts) (*StandardTokenUnpauseIterator, error) {

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &StandardTokenUnpauseIterator{contract: _StandardToken.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_StandardToken *StandardTokenFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *StandardTokenUnpause) (event.Subscription, error) {

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenUnpause)
				if err := _StandardToken.contract.UnpackLog(event, "Unpause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_Token *TokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_Token *TokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_Token *TokenSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_Token *TokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_Token *TokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_Token *TokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_Token *TokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_Token *TokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_Token *TokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, value)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
