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

// NextyGovernanceABI is the input ABI used to generate the binding from.
const NextyGovernanceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isWithdrawable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getCoinbase\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"join\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUnlockHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"account\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"unlockHeight\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerCoinbase\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isBanned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeLockHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeRequire\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"leave\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_stakeRequire\",\"type\":\"uint256\"},{\"name\":\"_stakeLockHeight\",\"type\":\"uint256\"},{\"name\":\"_signers\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"Joined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"Left\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"}],\"name\":\"Banned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"}],\"name\":\"Unbanned\",\"type\":\"event\"}]"

// NextyGovernanceBin is the compiled bytecode used for deploying new contracts.
const NextyGovernanceBin = `0x608060405234801561001057600080fd5b506040516112753803806112758339810180604052608081101561003357600080fd5b815160208301516040840151606085018051939592949193918301929164010000000081111561006257600080fd5b8201602081018481111561007557600080fd5b815185602082028301116401000000008211171561009257600080fd5b505060058054600160a060020a031916600160a060020a03891617905560038690556004859055925060009150505b815181101561025f57600082828151811015156100da57fe5b6020908102919091018101518254600181018455600093845291909220018054600160a060020a031916600160a060020a03909216919091179055815182908290811061012357fe5b9060200190602002015160016000848481518110151561013f57fe5b90602001906020020151600160a060020a0316600160a060020a0316815260200190815260200160002060006101000a815481600160a060020a030219169083600160a060020a03160217905550818181518110151561019b57fe5b906020019060200201516002600084848151811015156101b757fe5b90602001906020020151600160a060020a0316600160a060020a0316815260200190815260200160002060020160006101000a815481600160a060020a030219169083600160a060020a03160217905550600160026000848481518110151561021c57fe5b6020908102909101810151600160a060020a03168252810191909152604001600020805460ff1916600183600481111561025257fe5b02179055506001016100c1565b5050505050611002806102736000396000f3fe608060405234801561001057600080fd5b506004361061011d576000357c0100000000000000000000000000000000000000000000000000000000900480638943b62c116100b4578063c690873911610083578063c690873914610329578063d66d9e1914610331578063f8b2cb4f14610339578063fc0c546a1461035f5761011d565b80638943b62c146102b857806397f735d5146102de5780639b212d0114610304578063b6b55f251461030c5761011d565b806330ccebb5116100f057806330ccebb5146101e15780633ccfd60b1461021957806355235d471461022157806373b9aa91146102475761011d565b806302a18484146101225780630a9a3f071461015c5780632079fb9a1461019e57806328ffe6c8146101bb575b600080fd5b6101486004803603602081101561013857600080fd5b5035600160a060020a0316610367565b604080519115158252519081900360200190f35b6101826004803603602081101561017257600080fd5b5035600160a060020a03166103f5565b60408051600160a060020a039092168252519081900360200190f35b610182600480360360208110156101b457600080fd5b5035610417565b610148600480360360208110156101d157600080fd5b5035600160a060020a031661043f565b610207600480360360208110156101f757600080fd5b5035600160a060020a0316610804565b60408051918252519081900360200190f35b610148610829565b6102076004803603602081101561023757600080fd5b5035600160a060020a0316610a05565b61026d6004803603602081101561025d57600080fd5b5035600160a060020a0316610a23565b6040518085600481111561027d57fe5b60ff16815260200184815260200183600160a060020a0316600160a060020a0316815260200182815260200194505050505060405180910390f35b610182600480360360208110156102ce57600080fd5b5035600160a060020a0316610a59565b610148600480360360208110156102f457600080fd5b5035600160a060020a0316610a74565b610207610aa6565b6101486004803603602081101561032257600080fd5b5035610aac565b610207610bc7565b610148610bcd565b6102076004803603602081101561034f57600080fd5b5035600160a060020a0316610da9565b610182610dc7565b60006001600160a060020a03831660009081526002602052604090205460ff16600481111561039257fe5b141580156103c757506004600160a060020a03831660009081526002602052604090205460ff1660048111156103c457fe5b14155b80156103ed5750600160a060020a03821660009081526002602052604090206003015443115b90505b919050565b600160a060020a03908116600090815260026020819052604090912001541690565b600080548290811061042557fe5b600091825260209091200154600160a060020a0316905081565b600060043360009081526002602052604090205460ff16600481111561046157fe5b14156104b7576040805160e560020a62461bcd02815260206004820152600760248201527f62616e6e65642000000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b60013360009081526002602052604090205460ff1660048111156104d757fe5b141561052d576040805160e560020a62461bcd02815260206004820152600f60248201527f616c7265616479206a6f696e6564200000000000000000000000000000000000604482015290519081900360640190fd5b600354336000908152600260205260409020600101541015610599576040805160e560020a62461bcd02815260206004820152600e60248201527f6e6f7420656e6f756768206e7466000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a03808316600090815260016020526040902054839116151561060c576040805160e560020a62461bcd02815260206004820152601360248201527f7369676e657220616c7265616479207573656400000000000000000000000000604482015290519081900360640190fd5b600160a060020a038116151561066c576040805160e560020a62461bcd02815260206004820152600b60248201527f7369676e6572207a65726f000000000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a0381163014156106cd576040805160e560020a62461bcd02815260206004820152601760248201527f73616d6520636f6e747261637427732061646472657373000000000000000000604482015290519081900360640190fd5b600160a060020a03811633141561072e576040805160e560020a62461bcd02815260206004820152601560248201527f73616d652073656e646572277320616464726573730000000000000000000000604482015290519081900360640190fd5b336000908152600260208190526040909120908101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03861617905580546001919060ff191682800217905550600160a060020a0383166000908152600160205260409020805473ffffffffffffffffffffffffffffffffffffffff1916331790556107b883610dd6565b60408051338152600160a060020a038516602082015281517f7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea929181900390910190a150600192915050565b600160a060020a0381166000908152600260205260408120546103ed9060ff16610e32565b600060043360009081526002602052604090205460ff16600481111561084b57fe5b14156108a1576040805160e560020a62461bcd02815260206004820152600760248201527f62616e6e65642000000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6108aa33610367565b1515610900576040805160e560020a62461bcd02815260206004820181905260248201527f756e61626c6520746f20776974686472617720617420746865206d6f6d656e74604482015290519081900360640190fd5b33600081815260026020908152604080832060018101805490859055815460ff191660031790915560055482517fa9059cbb00000000000000000000000000000000000000000000000000000000815260048101969096526024860182905291519094600160a060020a039092169363a9059cbb93604480850194919392918390030190829087803b15801561099557600080fd5b505af11580156109a9573d6000803e3d6000fd5b505050506040513d60208110156109bf57600080fd5b5050604080513381526020810183905281517f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5929181900390910190a160019150505b90565b600160a060020a031660009081526002602052604090206003015490565b6002602081905260009182526040909120805460018201549282015460039092015460ff9091169291600160a060020a03169084565b600160205260009081526040902054600160a060020a031681565b60006004600160a060020a03831660009081526002602052604090205460ff166004811115610a9f57fe5b1492915050565b60045481565b600554604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490529051600092600160a060020a0316916323b872dd91606480830192602092919082900301818787803b158015610b1e57600080fd5b505af1158015610b32573d6000803e3d6000fd5b505050506040513d6020811015610b4857600080fd5b505033600090815260026020526040902060010154610b6d908363ffffffff610eab16565b3360008181526002602090815260409182902060010193909355805191825291810184905281517f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4929181900390910190a1506001919050565b60035481565b600060043360009081526002602052604090205460ff166004811115610bef57fe5b1415610c45576040805160e560020a62461bcd02815260206004820152600760248201527f62616e6e65642000000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b60013360009081526002602052604090205460ff166004811115610c6557fe5b14610cba576040805160e560020a62461bcd02815260206004820152600b60248201527f6e6f74206a6f696e656420000000000000000000000000000000000000000000604482015290519081900360640190fd5b336000908152600260208190526040909120808201805473ffffffffffffffffffffffffffffffffffffffff198116909155815460ff19169092179055600454600160a060020a0390911690610d109043610eab565b33600090815260026020908152604080832060030193909355600160a060020a03841682526001905220805473ffffffffffffffffffffffffffffffffffffffff19169055610d5e81610ec4565b60408051338152600160a060020a038316602082015281517f4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5929181900390910190a1600191505090565b600160a060020a031660009081526002602052604090206001015490565b600554600160a060020a031681565b600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600080826004811115610e4157fe5b1415610e4f575060006103f0565b6001826004811115610e5d57fe5b1415610e6b575060016103f0565b6002826004811115610e7957fe5b1415610e87575060026103f0565b6003826004811115610e9557fe5b1415610ea3575060036103f0565b50607f919050565b600082820183811015610ebd57600080fd5b9392505050565b60005b600054811015610f8a576000805482908110610edf57fe5b600091825260209091200154600160a060020a0383811691161415610f8257600080546000198101908110610f1057fe5b60009182526020822001548154600160a060020a03909116919083908110610f3457fe5b60009182526020822001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039390931692909217909155805490610f7b906000198301610f8f565b5050610f8c565b600101610ec7565b505b50565b815481835581811115610fb357600083815260209020610fb3918101908301610fb8565b505050565b610a0291905b80821115610fd25760008155600101610fbe565b509056fea165627a7a7230582056a374290c1efcbc416f13e7d57b9b0c89826ae44caf0e19e0497e3536a180180029`

// DeployNextyGovernance deploys a new Ethereum contract, binding an instance of NextyGovernance to it.
func DeployNextyGovernance(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _stakeRequire *big.Int, _stakeLockHeight *big.Int, _signers []common.Address) (common.Address, *types.Transaction, *NextyGovernance, error) {
	parsed, err := abi.JSON(strings.NewReader(NextyGovernanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NextyGovernanceBin), backend, _token, _stakeRequire, _stakeLockHeight, _signers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NextyGovernance{NextyGovernanceCaller: NextyGovernanceCaller{contract: contract}, NextyGovernanceTransactor: NextyGovernanceTransactor{contract: contract}, NextyGovernanceFilterer: NextyGovernanceFilterer{contract: contract}}, nil
}

// NextyGovernance is an auto generated Go binding around an Ethereum contract.
type NextyGovernance struct {
	NextyGovernanceCaller     // Read-only binding to the contract
	NextyGovernanceTransactor // Write-only binding to the contract
	NextyGovernanceFilterer   // Log filterer for contract events
}

// NextyGovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type NextyGovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NextyGovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NextyGovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NextyGovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NextyGovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NextyGovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NextyGovernanceSession struct {
	Contract     *NextyGovernance  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NextyGovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NextyGovernanceCallerSession struct {
	Contract *NextyGovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NextyGovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NextyGovernanceTransactorSession struct {
	Contract     *NextyGovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NextyGovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type NextyGovernanceRaw struct {
	Contract *NextyGovernance // Generic contract binding to access the raw methods on
}

// NextyGovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NextyGovernanceCallerRaw struct {
	Contract *NextyGovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// NextyGovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NextyGovernanceTransactorRaw struct {
	Contract *NextyGovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNextyGovernance creates a new instance of NextyGovernance, bound to a specific deployed contract.
func NewNextyGovernance(address common.Address, backend bind.ContractBackend) (*NextyGovernance, error) {
	contract, err := bindNextyGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NextyGovernance{NextyGovernanceCaller: NextyGovernanceCaller{contract: contract}, NextyGovernanceTransactor: NextyGovernanceTransactor{contract: contract}, NextyGovernanceFilterer: NextyGovernanceFilterer{contract: contract}}, nil
}

// NewNextyGovernanceCaller creates a new read-only instance of NextyGovernance, bound to a specific deployed contract.
func NewNextyGovernanceCaller(address common.Address, caller bind.ContractCaller) (*NextyGovernanceCaller, error) {
	contract, err := bindNextyGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceCaller{contract: contract}, nil
}

// NewNextyGovernanceTransactor creates a new write-only instance of NextyGovernance, bound to a specific deployed contract.
func NewNextyGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*NextyGovernanceTransactor, error) {
	contract, err := bindNextyGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceTransactor{contract: contract}, nil
}

// NewNextyGovernanceFilterer creates a new log filterer instance of NextyGovernance, bound to a specific deployed contract.
func NewNextyGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*NextyGovernanceFilterer, error) {
	contract, err := bindNextyGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceFilterer{contract: contract}, nil
}

// bindNextyGovernance binds a generic wrapper to an already deployed contract.
func bindNextyGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NextyGovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NextyGovernance *NextyGovernanceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NextyGovernance.Contract.NextyGovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NextyGovernance *NextyGovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.Contract.NextyGovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NextyGovernance *NextyGovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NextyGovernance.Contract.NextyGovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NextyGovernance *NextyGovernanceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NextyGovernance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NextyGovernance *NextyGovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NextyGovernance *NextyGovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NextyGovernance.Contract.contract.Transact(opts, method, params...)
}

// Account is a free data retrieval call binding the contract method 0x73b9aa91.
//
// Solidity: function account( address) constant returns(status uint8, balance uint256, signer address, unlockHeight uint256)
func (_NextyGovernance *NextyGovernanceCaller) Account(opts *bind.CallOpts, arg0 common.Address) (struct {
	Status       uint8
	Balance      *big.Int
	Signer       common.Address
	UnlockHeight *big.Int
}, error) {
	ret := new(struct {
		Status       uint8
		Balance      *big.Int
		Signer       common.Address
		UnlockHeight *big.Int
	})
	out := ret
	err := _NextyGovernance.contract.Call(opts, out, "account", arg0)
	return *ret, err
}

// Account is a free data retrieval call binding the contract method 0x73b9aa91.
//
// Solidity: function account( address) constant returns(status uint8, balance uint256, signer address, unlockHeight uint256)
func (_NextyGovernance *NextyGovernanceSession) Account(arg0 common.Address) (struct {
	Status       uint8
	Balance      *big.Int
	Signer       common.Address
	UnlockHeight *big.Int
}, error) {
	return _NextyGovernance.Contract.Account(&_NextyGovernance.CallOpts, arg0)
}

// Account is a free data retrieval call binding the contract method 0x73b9aa91.
//
// Solidity: function account( address) constant returns(status uint8, balance uint256, signer address, unlockHeight uint256)
func (_NextyGovernance *NextyGovernanceCallerSession) Account(arg0 common.Address) (struct {
	Status       uint8
	Balance      *big.Int
	Signer       common.Address
	UnlockHeight *big.Int
}, error) {
	return _NextyGovernance.Contract.Account(&_NextyGovernance.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(_address address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCaller) GetBalance(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "getBalance", _address)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(_address address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) GetBalance(_address common.Address) (*big.Int, error) {
	return _NextyGovernance.Contract.GetBalance(&_NextyGovernance.CallOpts, _address)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(_address address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCallerSession) GetBalance(_address common.Address) (*big.Int, error) {
	return _NextyGovernance.Contract.GetBalance(&_NextyGovernance.CallOpts, _address)
}

// GetCoinbase is a free data retrieval call binding the contract method 0x0a9a3f07.
//
// Solidity: function getCoinbase(_address address) constant returns(address)
func (_NextyGovernance *NextyGovernanceCaller) GetCoinbase(opts *bind.CallOpts, _address common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "getCoinbase", _address)
	return *ret0, err
}

// GetCoinbase is a free data retrieval call binding the contract method 0x0a9a3f07.
//
// Solidity: function getCoinbase(_address address) constant returns(address)
func (_NextyGovernance *NextyGovernanceSession) GetCoinbase(_address common.Address) (common.Address, error) {
	return _NextyGovernance.Contract.GetCoinbase(&_NextyGovernance.CallOpts, _address)
}

// GetCoinbase is a free data retrieval call binding the contract method 0x0a9a3f07.
//
// Solidity: function getCoinbase(_address address) constant returns(address)
func (_NextyGovernance *NextyGovernanceCallerSession) GetCoinbase(_address common.Address) (common.Address, error) {
	return _NextyGovernance.Contract.GetCoinbase(&_NextyGovernance.CallOpts, _address)
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(_address address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCaller) GetStatus(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "getStatus", _address)
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(_address address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) GetStatus(_address common.Address) (*big.Int, error) {
	return _NextyGovernance.Contract.GetStatus(&_NextyGovernance.CallOpts, _address)
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(_address address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCallerSession) GetStatus(_address common.Address) (*big.Int, error) {
	return _NextyGovernance.Contract.GetStatus(&_NextyGovernance.CallOpts, _address)
}

// GetUnlockHeight is a free data retrieval call binding the contract method 0x55235d47.
//
// Solidity: function getUnlockHeight(_address address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCaller) GetUnlockHeight(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "getUnlockHeight", _address)
	return *ret0, err
}

// GetUnlockHeight is a free data retrieval call binding the contract method 0x55235d47.
//
// Solidity: function getUnlockHeight(_address address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) GetUnlockHeight(_address common.Address) (*big.Int, error) {
	return _NextyGovernance.Contract.GetUnlockHeight(&_NextyGovernance.CallOpts, _address)
}

// GetUnlockHeight is a free data retrieval call binding the contract method 0x55235d47.
//
// Solidity: function getUnlockHeight(_address address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCallerSession) GetUnlockHeight(_address common.Address) (*big.Int, error) {
	return _NextyGovernance.Contract.GetUnlockHeight(&_NextyGovernance.CallOpts, _address)
}

// IsBanned is a free data retrieval call binding the contract method 0x97f735d5.
//
// Solidity: function isBanned(_address address) constant returns(bool)
func (_NextyGovernance *NextyGovernanceCaller) IsBanned(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "isBanned", _address)
	return *ret0, err
}

// IsBanned is a free data retrieval call binding the contract method 0x97f735d5.
//
// Solidity: function isBanned(_address address) constant returns(bool)
func (_NextyGovernance *NextyGovernanceSession) IsBanned(_address common.Address) (bool, error) {
	return _NextyGovernance.Contract.IsBanned(&_NextyGovernance.CallOpts, _address)
}

// IsBanned is a free data retrieval call binding the contract method 0x97f735d5.
//
// Solidity: function isBanned(_address address) constant returns(bool)
func (_NextyGovernance *NextyGovernanceCallerSession) IsBanned(_address common.Address) (bool, error) {
	return _NextyGovernance.Contract.IsBanned(&_NextyGovernance.CallOpts, _address)
}

// IsWithdrawable is a free data retrieval call binding the contract method 0x02a18484.
//
// Solidity: function isWithdrawable(_address address) constant returns(bool)
func (_NextyGovernance *NextyGovernanceCaller) IsWithdrawable(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "isWithdrawable", _address)
	return *ret0, err
}

// IsWithdrawable is a free data retrieval call binding the contract method 0x02a18484.
//
// Solidity: function isWithdrawable(_address address) constant returns(bool)
func (_NextyGovernance *NextyGovernanceSession) IsWithdrawable(_address common.Address) (bool, error) {
	return _NextyGovernance.Contract.IsWithdrawable(&_NextyGovernance.CallOpts, _address)
}

// IsWithdrawable is a free data retrieval call binding the contract method 0x02a18484.
//
// Solidity: function isWithdrawable(_address address) constant returns(bool)
func (_NextyGovernance *NextyGovernanceCallerSession) IsWithdrawable(_address common.Address) (bool, error) {
	return _NextyGovernance.Contract.IsWithdrawable(&_NextyGovernance.CallOpts, _address)
}

// SignerCoinbase is a free data retrieval call binding the contract method 0x8943b62c.
//
// Solidity: function signerCoinbase( address) constant returns(address)
func (_NextyGovernance *NextyGovernanceCaller) SignerCoinbase(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "signerCoinbase", arg0)
	return *ret0, err
}

// SignerCoinbase is a free data retrieval call binding the contract method 0x8943b62c.
//
// Solidity: function signerCoinbase( address) constant returns(address)
func (_NextyGovernance *NextyGovernanceSession) SignerCoinbase(arg0 common.Address) (common.Address, error) {
	return _NextyGovernance.Contract.SignerCoinbase(&_NextyGovernance.CallOpts, arg0)
}

// SignerCoinbase is a free data retrieval call binding the contract method 0x8943b62c.
//
// Solidity: function signerCoinbase( address) constant returns(address)
func (_NextyGovernance *NextyGovernanceCallerSession) SignerCoinbase(arg0 common.Address) (common.Address, error) {
	return _NextyGovernance.Contract.SignerCoinbase(&_NextyGovernance.CallOpts, arg0)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers( uint256) constant returns(address)
func (_NextyGovernance *NextyGovernanceCaller) Signers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "signers", arg0)
	return *ret0, err
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers( uint256) constant returns(address)
func (_NextyGovernance *NextyGovernanceSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _NextyGovernance.Contract.Signers(&_NextyGovernance.CallOpts, arg0)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers( uint256) constant returns(address)
func (_NextyGovernance *NextyGovernanceCallerSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _NextyGovernance.Contract.Signers(&_NextyGovernance.CallOpts, arg0)
}

// StakeLockHeight is a free data retrieval call binding the contract method 0x9b212d01.
//
// Solidity: function stakeLockHeight() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCaller) StakeLockHeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "stakeLockHeight")
	return *ret0, err
}

// StakeLockHeight is a free data retrieval call binding the contract method 0x9b212d01.
//
// Solidity: function stakeLockHeight() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) StakeLockHeight() (*big.Int, error) {
	return _NextyGovernance.Contract.StakeLockHeight(&_NextyGovernance.CallOpts)
}

// StakeLockHeight is a free data retrieval call binding the contract method 0x9b212d01.
//
// Solidity: function stakeLockHeight() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCallerSession) StakeLockHeight() (*big.Int, error) {
	return _NextyGovernance.Contract.StakeLockHeight(&_NextyGovernance.CallOpts)
}

// StakeRequire is a free data retrieval call binding the contract method 0xc6908739.
//
// Solidity: function stakeRequire() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCaller) StakeRequire(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "stakeRequire")
	return *ret0, err
}

// StakeRequire is a free data retrieval call binding the contract method 0xc6908739.
//
// Solidity: function stakeRequire() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) StakeRequire() (*big.Int, error) {
	return _NextyGovernance.Contract.StakeRequire(&_NextyGovernance.CallOpts)
}

// StakeRequire is a free data retrieval call binding the contract method 0xc6908739.
//
// Solidity: function stakeRequire() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCallerSession) StakeRequire() (*big.Int, error) {
	return _NextyGovernance.Contract.StakeRequire(&_NextyGovernance.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_NextyGovernance *NextyGovernanceCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_NextyGovernance *NextyGovernanceSession) Token() (common.Address, error) {
	return _NextyGovernance.Contract.Token(&_NextyGovernance.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_NextyGovernance *NextyGovernanceCallerSession) Token() (common.Address, error) {
	return _NextyGovernance.Contract.Token(&_NextyGovernance.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(_amount uint256) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(_amount uint256) returns(bool)
func (_NextyGovernance *NextyGovernanceSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Deposit(&_NextyGovernance.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(_amount uint256) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Deposit(&_NextyGovernance.TransactOpts, _amount)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(_signer address) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactor) Join(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "join", _signer)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(_signer address) returns(bool)
func (_NextyGovernance *NextyGovernanceSession) Join(_signer common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Join(&_NextyGovernance.TransactOpts, _signer)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(_signer address) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactorSession) Join(_signer common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Join(&_NextyGovernance.TransactOpts, _signer)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns(bool)
func (_NextyGovernance *NextyGovernanceTransactor) Leave(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "leave")
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns(bool)
func (_NextyGovernance *NextyGovernanceSession) Leave() (*types.Transaction, error) {
	return _NextyGovernance.Contract.Leave(&_NextyGovernance.TransactOpts)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns(bool)
func (_NextyGovernance *NextyGovernanceTransactorSession) Leave() (*types.Transaction, error) {
	return _NextyGovernance.Contract.Leave(&_NextyGovernance.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_NextyGovernance *NextyGovernanceTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_NextyGovernance *NextyGovernanceSession) Withdraw() (*types.Transaction, error) {
	return _NextyGovernance.Contract.Withdraw(&_NextyGovernance.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_NextyGovernance *NextyGovernanceTransactorSession) Withdraw() (*types.Transaction, error) {
	return _NextyGovernance.Contract.Withdraw(&_NextyGovernance.TransactOpts)
}

// NextyGovernanceBannedIterator is returned from FilterBanned and is used to iterate over the raw logs and unpacked data for Banned events raised by the NextyGovernance contract.
type NextyGovernanceBannedIterator struct {
	Event *NextyGovernanceBanned // Event containing the contract specifics and raw log

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
func (it *NextyGovernanceBannedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceBanned)
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
		it.Event = new(NextyGovernanceBanned)
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
func (it *NextyGovernanceBannedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceBannedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceBanned represents a Banned event raised by the NextyGovernance contract.
type NextyGovernanceBanned struct {
	Sealer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBanned is a free log retrieval operation binding the contract event 0x30d1df1214d91553408ca5384ce29e10e5866af8423c628be22860e41fb81005.
//
// Solidity: e Banned(_sealer address)
func (_NextyGovernance *NextyGovernanceFilterer) FilterBanned(opts *bind.FilterOpts) (*NextyGovernanceBannedIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Banned")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceBannedIterator{contract: _NextyGovernance.contract, event: "Banned", logs: logs, sub: sub}, nil
}

// WatchBanned is a free log subscription operation binding the contract event 0x30d1df1214d91553408ca5384ce29e10e5866af8423c628be22860e41fb81005.
//
// Solidity: e Banned(_sealer address)
func (_NextyGovernance *NextyGovernanceFilterer) WatchBanned(opts *bind.WatchOpts, sink chan<- *NextyGovernanceBanned) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Banned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceBanned)
				if err := _NextyGovernance.contract.UnpackLog(event, "Banned", log); err != nil {
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

// NextyGovernanceDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the NextyGovernance contract.
type NextyGovernanceDepositedIterator struct {
	Event *NextyGovernanceDeposited // Event containing the contract specifics and raw log

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
func (it *NextyGovernanceDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceDeposited)
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
		it.Event = new(NextyGovernanceDeposited)
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
func (it *NextyGovernanceDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceDeposited represents a Deposited event raised by the NextyGovernance contract.
type NextyGovernanceDeposited struct {
	Sealer common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: e Deposited(_sealer address, _amount uint256)
func (_NextyGovernance *NextyGovernanceFilterer) FilterDeposited(opts *bind.FilterOpts) (*NextyGovernanceDepositedIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceDepositedIterator{contract: _NextyGovernance.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: e Deposited(_sealer address, _amount uint256)
func (_NextyGovernance *NextyGovernanceFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *NextyGovernanceDeposited) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceDeposited)
				if err := _NextyGovernance.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// NextyGovernanceJoinedIterator is returned from FilterJoined and is used to iterate over the raw logs and unpacked data for Joined events raised by the NextyGovernance contract.
type NextyGovernanceJoinedIterator struct {
	Event *NextyGovernanceJoined // Event containing the contract specifics and raw log

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
func (it *NextyGovernanceJoinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceJoined)
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
		it.Event = new(NextyGovernanceJoined)
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
func (it *NextyGovernanceJoinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceJoinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceJoined represents a Joined event raised by the NextyGovernance contract.
type NextyGovernanceJoined struct {
	Sealer common.Address
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterJoined is a free log retrieval operation binding the contract event 0x7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea.
//
// Solidity: e Joined(_sealer address, _signer address)
func (_NextyGovernance *NextyGovernanceFilterer) FilterJoined(opts *bind.FilterOpts) (*NextyGovernanceJoinedIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Joined")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceJoinedIterator{contract: _NextyGovernance.contract, event: "Joined", logs: logs, sub: sub}, nil
}

// WatchJoined is a free log subscription operation binding the contract event 0x7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea.
//
// Solidity: e Joined(_sealer address, _signer address)
func (_NextyGovernance *NextyGovernanceFilterer) WatchJoined(opts *bind.WatchOpts, sink chan<- *NextyGovernanceJoined) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Joined")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceJoined)
				if err := _NextyGovernance.contract.UnpackLog(event, "Joined", log); err != nil {
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

// NextyGovernanceLeftIterator is returned from FilterLeft and is used to iterate over the raw logs and unpacked data for Left events raised by the NextyGovernance contract.
type NextyGovernanceLeftIterator struct {
	Event *NextyGovernanceLeft // Event containing the contract specifics and raw log

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
func (it *NextyGovernanceLeftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceLeft)
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
		it.Event = new(NextyGovernanceLeft)
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
func (it *NextyGovernanceLeftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceLeftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceLeft represents a Left event raised by the NextyGovernance contract.
type NextyGovernanceLeft struct {
	Sealer common.Address
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLeft is a free log retrieval operation binding the contract event 0x4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5.
//
// Solidity: e Left(_sealer address, _signer address)
func (_NextyGovernance *NextyGovernanceFilterer) FilterLeft(opts *bind.FilterOpts) (*NextyGovernanceLeftIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Left")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceLeftIterator{contract: _NextyGovernance.contract, event: "Left", logs: logs, sub: sub}, nil
}

// WatchLeft is a free log subscription operation binding the contract event 0x4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5.
//
// Solidity: e Left(_sealer address, _signer address)
func (_NextyGovernance *NextyGovernanceFilterer) WatchLeft(opts *bind.WatchOpts, sink chan<- *NextyGovernanceLeft) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Left")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceLeft)
				if err := _NextyGovernance.contract.UnpackLog(event, "Left", log); err != nil {
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

// NextyGovernanceUnbannedIterator is returned from FilterUnbanned and is used to iterate over the raw logs and unpacked data for Unbanned events raised by the NextyGovernance contract.
type NextyGovernanceUnbannedIterator struct {
	Event *NextyGovernanceUnbanned // Event containing the contract specifics and raw log

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
func (it *NextyGovernanceUnbannedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceUnbanned)
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
		it.Event = new(NextyGovernanceUnbanned)
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
func (it *NextyGovernanceUnbannedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceUnbannedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceUnbanned represents a Unbanned event raised by the NextyGovernance contract.
type NextyGovernanceUnbanned struct {
	Sealer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnbanned is a free log retrieval operation binding the contract event 0x2ab91b53354938415bb6962c4322231cd4cb2c84930f1a4b9abbedc2fe8abe72.
//
// Solidity: e Unbanned(_sealer address)
func (_NextyGovernance *NextyGovernanceFilterer) FilterUnbanned(opts *bind.FilterOpts) (*NextyGovernanceUnbannedIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Unbanned")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceUnbannedIterator{contract: _NextyGovernance.contract, event: "Unbanned", logs: logs, sub: sub}, nil
}

// WatchUnbanned is a free log subscription operation binding the contract event 0x2ab91b53354938415bb6962c4322231cd4cb2c84930f1a4b9abbedc2fe8abe72.
//
// Solidity: e Unbanned(_sealer address)
func (_NextyGovernance *NextyGovernanceFilterer) WatchUnbanned(opts *bind.WatchOpts, sink chan<- *NextyGovernanceUnbanned) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Unbanned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceUnbanned)
				if err := _NextyGovernance.contract.UnpackLog(event, "Unbanned", log); err != nil {
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

// NextyGovernanceWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the NextyGovernance contract.
type NextyGovernanceWithdrawnIterator struct {
	Event *NextyGovernanceWithdrawn // Event containing the contract specifics and raw log

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
func (it *NextyGovernanceWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceWithdrawn)
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
		it.Event = new(NextyGovernanceWithdrawn)
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
func (it *NextyGovernanceWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceWithdrawn represents a Withdrawn event raised by the NextyGovernance contract.
type NextyGovernanceWithdrawn struct {
	Sealer common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: e Withdrawn(_sealer address, _amount uint256)
func (_NextyGovernance *NextyGovernanceFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*NextyGovernanceWithdrawnIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceWithdrawnIterator{contract: _NextyGovernance.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: e Withdrawn(_sealer address, _amount uint256)
func (_NextyGovernance *NextyGovernanceFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *NextyGovernanceWithdrawn) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceWithdrawn)
				if err := _NextyGovernance.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820138581c5c923b6c9f179b19fa4e4e4affe758a205c47a5c664d688f9de80709c0029`

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

