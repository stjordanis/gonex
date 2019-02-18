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
const NextyGovernanceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isWithdrawable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getCoinbase\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isCoinbase\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"coinbases\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"getSealer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_coinbase\",\"type\":\"address\"}],\"name\":\"join\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LOCK_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isBanned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"sealers\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"coinbase\",\"type\":\"address\"},{\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_NTF_AMOUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"leave\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUnlockTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_sealers\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_coinbase\",\"type\":\"address\"}],\"name\":\"Joined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_coinbase\",\"type\":\"address\"}],\"name\":\"Left\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"}],\"name\":\"Banned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sealer\",\"type\":\"address\"}],\"name\":\"Unbanned\",\"type\":\"event\"}]"

// NextyGovernanceBin is the compiled bytecode used for deploying new contracts.
const NextyGovernanceBin = `0x60806040523480156200001157600080fd5b506040516200132838038062001328833981018060405260408110156200003757600080fd5b8151602083018051919392830192916401000000008111156200005957600080fd5b820160208101848111156200006d57600080fd5b81518560208202830111640100000000821117156200008b57600080fd5b505060008054600160a060020a031916600160a060020a0387161781559093509150505b81518110156200029b5760028282815181101515620000ca57fe5b6020908102919091018101518254600180820185556000948552928420018054600160a060020a031916600160a060020a03909216919091179055835190916003918590859081106200011957fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff191691151591909117905581518290829081106200015a57fe5b906020019060200201516004600084848151811015156200017757fe5b90602001906020020151600160a060020a0316600160a060020a0316815260200190815260200160002060006101000a815481600160a060020a030219169083600160a060020a031602179055508181815181101515620001d457fe5b90602001906020020151600160008484815181101515620001f157fe5b90602001906020020151600160a060020a0316600160a060020a0316815260200190815260200160002060020160006101000a815481600160a060020a030219169083600160a060020a03160217905550600180600084848151811015156200025657fe5b6020908102909101810151600160a060020a03168252810191909152604001600020805460ff191660018360048111156200028d57fe5b0217905550600101620000af565b50505061107a80620002ae6000396000f3fe608060405234801561001057600080fd5b5060043610610128576000357c010000000000000000000000000000000000000000000000000000000090048063485d3834116100bf578063b6b55f251161008e578063b6b55f251461031f578063d66d9e191461033c578063dbed1e8814610344578063f8b2cb4f1461036a578063fc0c546a1461039057610128565b8063485d38341461027857806397f735d514610280578063aa61aef3146102a6578063b22ff4541461031757610128565b80632041d33f116100fb5780632041d33f146101ec57806328ffe6c81461021257806330ccebb5146102385780633ccfd60b1461027057610128565b806302a184841461012d5780630a9a3f07146101675780630d8754ac146101a95780630ebf0de7146101cf575b600080fd5b6101536004803603602081101561014357600080fd5b5035600160a060020a0316610398565b604080519115158252519081900360200190f35b61018d6004803603602081101561017d57600080fd5b5035600160a060020a0316610426565b60408051600160a060020a039092168252519081900360200190f35b610153600480360360208110156101bf57600080fd5b5035600160a060020a0316610447565b61018d600480360360208110156101e557600080fd5b503561045c565b61018d6004803603602081101561020257600080fd5b5035600160a060020a0316610484565b6101536004803603602081101561022857600080fd5b5035600160a060020a031661049f565b61025e6004803603602081101561024e57600080fd5b5035600160a060020a0316610861565b60408051918252519081900360200190f35b610153610886565b61025e610a60565b6101536004803603602081101561029657600080fd5b5035600160a060020a0316610a66565b6102cc600480360360208110156102bc57600080fd5b5035600160a060020a0316610a98565b604051808560048111156102dc57fe5b60ff16815260200184815260200183600160a060020a0316600160a060020a0316815260200182815260200194505050505060405180910390f35b61025e610acf565b6101536004803603602081101561033557600080fd5b5035610ad4565b610153610bf1565b61025e6004803603602081101561035a57600080fd5b5035600160a060020a0316610dc9565b61025e6004803603602081101561038057600080fd5b5035600160a060020a0316610de7565b61018d610e06565b60006001600160a060020a03831660009081526001602052604090205460ff1660048111156103c357fe5b141580156103f857506004600160a060020a03831660009081526001602052604090205460ff1660048111156103f557fe5b14155b801561041e5750600160a060020a03821660009081526001602052604090206003015442115b90505b919050565b600160a060020a039081166000908152600160205260409020600201541690565b60036020526000908152604090205460ff1681565b600280548290811061046a57fe5b600091825260209091200154600160a060020a0316905081565b600460205260009081526040902054600160a060020a031681565b600060043360009081526001602052604090205460ff1660048111156104c157fe5b1415610517576040805160e560020a62461bcd02815260206004820152600760248201527f62616e6e65642000000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b60013360009081526001602052604090205460ff16600481111561053757fe5b141561058d576040805160e560020a62461bcd02815260206004820152600f60248201527f616c7265616479206a6f696e6564200000000000000000000000000000000000604482015290519081900360640190fd5b3360009081526001602081905260409091200154606411156105f9576040805160e560020a62461bcd02815260206004820152600e60248201527f6e6f7420656e6f756768206e7466000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a038216600090815260036020526040902054829060ff161561066c576040805160e560020a62461bcd02815260206004820152601560248201527f636f696e6261736520616c726561647920757365640000000000000000000000604482015290519081900360640190fd5b600160a060020a03811615156106cc576040805160e560020a62461bcd02815260206004820152600d60248201527f636f696e62617365207a65726f00000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a03811630141561072d576040805160e560020a62461bcd02815260206004820152601760248201527f73616d6520636f6e747261637427732061646472657373000000000000000000604482015290519081900360640190fd5b600160a060020a03811633141561078e576040805160e560020a62461bcd02815260206004820152601560248201527f73616d652073656e646572277320616464726573730000000000000000000000604482015290519081900360640190fd5b33600090815260016020819052604090912060028101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038716179055805460ff191682800217905550600160a060020a0383166000908152600460205260409020805473ffffffffffffffffffffffffffffffffffffffff19163317905561081583610e15565b60408051338152600160a060020a038516602082015281517f7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea929181900390910190a150600192915050565b600160a060020a03811660009081526001602052604081205461041e9060ff16610e88565b600060043360009081526001602052604090205460ff1660048111156108a857fe5b14156108fe576040805160e560020a62461bcd02815260206004820152600760248201527f62616e6e65642000000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b61090733610398565b151561095d576040805160e560020a62461bcd02815260206004820181905260248201527f756e61626c6520746f20776974686472617720617420746865206d6f6d656e74604482015290519081900360640190fd5b336000818152600160208181526040808420928301805490859055835460ff1916600317909355835481517fa9059cbb00000000000000000000000000000000000000000000000000000000815260048101969096526024860184905290519294600160a060020a039091169363a9059cbb93604480840194939192918390030190829087803b1580156109f057600080fd5b505af1158015610a04573d6000803e3d6000fd5b505050506040513d6020811015610a1a57600080fd5b5050604080513381526020810183905281517f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5929181900390910190a160019150505b90565b61012c81565b60006004600160a060020a03831660009081526001602052604090205460ff166004811115610a9157fe5b1492915050565b6001602081905260009182526040909120805491810154600282015460039092015460ff909316929091600160a060020a03169084565b606481565b60008054604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529051600160a060020a03909216916323b872dd9160648082019260209290919082900301818787803b158015610b4757600080fd5b505af1158015610b5b573d6000803e3d6000fd5b505050506040513d6020811015610b7157600080fd5b50503360009081526001602081905260409091200154610b97908363ffffffff610f0116565b3360008181526001602081815260409283902090910193909355805191825291810184905281517f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4929181900390910190a1506001919050565b600060043360009081526001602052604090205460ff166004811115610c1357fe5b1415610c69576040805160e560020a62461bcd02815260206004820152600760248201527f62616e6e65642000000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b60013360009081526001602052604090205460ff166004811115610c8957fe5b14610cde576040805160e560020a62461bcd02815260206004820152600b60248201527f6e6f74206a6f696e656420000000000000000000000000000000000000000000604482015290519081900360640190fd5b3360009081526001602052604090206002808201805473ffffffffffffffffffffffffffffffffffffffff198116909155825460ff1916909117909155600160a060020a0316610d3061012c42610f01565b33600090815260016020908152604080832060030193909355600160a060020a03841682526004905220805473ffffffffffffffffffffffffffffffffffffffff19169055610d7e81610f1a565b60408051338152600160a060020a038316602082015281517f4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5929181900390910190a1600191505090565b600160a060020a031660009081526001602052604090206003015490565b600160a060020a03166000908152600160208190526040909120015490565b600054600160a060020a031681565b600160a060020a03166000818152600360205260408120805460ff191660019081179091556002805491820181559091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01805473ffffffffffffffffffffffffffffffffffffffff19169091179055565b600080826004811115610e9757fe5b1415610ea557506000610421565b6001826004811115610eb357fe5b1415610ec157506001610421565b6002826004811115610ecf57fe5b1415610edd57506002610421565b6003826004811115610eeb57fe5b1415610ef957506003610421565b50607f919050565b600082820183811015610f1357600080fd5b9392505050565b600160a060020a0381166000908152600360205260408120805460ff191690555b600254811015611002576002805482908110610f5357fe5b600091825260209091200154600160a060020a0383811691161415610ffa57600280546000198101908110610f8457fe5b60009182526020909120015460028054600160a060020a039092169183908110610faa57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790556002805490610ff3906000198301611007565b5050611004565b600101610f3b565b505b50565b81548183558181111561102b5760008381526020902061102b918101908301611030565b505050565b610a5d91905b8082111561104a5760008155600101611036565b509056fea165627a7a7230582034ab97cfca2d205c93e00c70f86cd94239db785e54ff3893222af0523ad0a8a50029`

// DeployNextyGovernance deploys a new Ethereum contract, binding an instance of NextyGovernance to it.
func DeployNextyGovernance(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _sealers []common.Address) (common.Address, *types.Transaction, *NextyGovernance, error) {
	parsed, err := abi.JSON(strings.NewReader(NextyGovernanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NextyGovernanceBin), backend, _token, _sealers)
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

// LOCKDURATION is a free data retrieval call binding the contract method 0x485d3834.
//
// Solidity: function LOCK_DURATION() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCaller) LOCKDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "LOCK_DURATION")
	return *ret0, err
}

// LOCKDURATION is a free data retrieval call binding the contract method 0x485d3834.
//
// Solidity: function LOCK_DURATION() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) LOCKDURATION() (*big.Int, error) {
	return _NextyGovernance.Contract.LOCKDURATION(&_NextyGovernance.CallOpts)
}

// LOCKDURATION is a free data retrieval call binding the contract method 0x485d3834.
//
// Solidity: function LOCK_DURATION() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCallerSession) LOCKDURATION() (*big.Int, error) {
	return _NextyGovernance.Contract.LOCKDURATION(&_NextyGovernance.CallOpts)
}

// MINNTFAMOUNT is a free data retrieval call binding the contract method 0xb22ff454.
//
// Solidity: function MIN_NTF_AMOUNT() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCaller) MINNTFAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "MIN_NTF_AMOUNT")
	return *ret0, err
}

// MINNTFAMOUNT is a free data retrieval call binding the contract method 0xb22ff454.
//
// Solidity: function MIN_NTF_AMOUNT() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) MINNTFAMOUNT() (*big.Int, error) {
	return _NextyGovernance.Contract.MINNTFAMOUNT(&_NextyGovernance.CallOpts)
}

// MINNTFAMOUNT is a free data retrieval call binding the contract method 0xb22ff454.
//
// Solidity: function MIN_NTF_AMOUNT() constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCallerSession) MINNTFAMOUNT() (*big.Int, error) {
	return _NextyGovernance.Contract.MINNTFAMOUNT(&_NextyGovernance.CallOpts)
}

// Coinbases is a free data retrieval call binding the contract method 0x0ebf0de7.
//
// Solidity: function coinbases( uint256) constant returns(address)
func (_NextyGovernance *NextyGovernanceCaller) Coinbases(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "coinbases", arg0)
	return *ret0, err
}

// Coinbases is a free data retrieval call binding the contract method 0x0ebf0de7.
//
// Solidity: function coinbases( uint256) constant returns(address)
func (_NextyGovernance *NextyGovernanceSession) Coinbases(arg0 *big.Int) (common.Address, error) {
	return _NextyGovernance.Contract.Coinbases(&_NextyGovernance.CallOpts, arg0)
}

// Coinbases is a free data retrieval call binding the contract method 0x0ebf0de7.
//
// Solidity: function coinbases( uint256) constant returns(address)
func (_NextyGovernance *NextyGovernanceCallerSession) Coinbases(arg0 *big.Int) (common.Address, error) {
	return _NextyGovernance.Contract.Coinbases(&_NextyGovernance.CallOpts, arg0)
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

// GetSealer is a free data retrieval call binding the contract method 0x2041d33f.
//
// Solidity: function getSealer( address) constant returns(address)
func (_NextyGovernance *NextyGovernanceCaller) GetSealer(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "getSealer", arg0)
	return *ret0, err
}

// GetSealer is a free data retrieval call binding the contract method 0x2041d33f.
//
// Solidity: function getSealer( address) constant returns(address)
func (_NextyGovernance *NextyGovernanceSession) GetSealer(arg0 common.Address) (common.Address, error) {
	return _NextyGovernance.Contract.GetSealer(&_NextyGovernance.CallOpts, arg0)
}

// GetSealer is a free data retrieval call binding the contract method 0x2041d33f.
//
// Solidity: function getSealer( address) constant returns(address)
func (_NextyGovernance *NextyGovernanceCallerSession) GetSealer(arg0 common.Address) (common.Address, error) {
	return _NextyGovernance.Contract.GetSealer(&_NextyGovernance.CallOpts, arg0)
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

// GetUnlockTime is a free data retrieval call binding the contract method 0xdbed1e88.
//
// Solidity: function getUnlockTime(_address address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCaller) GetUnlockTime(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "getUnlockTime", _address)
	return *ret0, err
}

// GetUnlockTime is a free data retrieval call binding the contract method 0xdbed1e88.
//
// Solidity: function getUnlockTime(_address address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) GetUnlockTime(_address common.Address) (*big.Int, error) {
	return _NextyGovernance.Contract.GetUnlockTime(&_NextyGovernance.CallOpts, _address)
}

// GetUnlockTime is a free data retrieval call binding the contract method 0xdbed1e88.
//
// Solidity: function getUnlockTime(_address address) constant returns(uint256)
func (_NextyGovernance *NextyGovernanceCallerSession) GetUnlockTime(_address common.Address) (*big.Int, error) {
	return _NextyGovernance.Contract.GetUnlockTime(&_NextyGovernance.CallOpts, _address)
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

// IsCoinbase is a free data retrieval call binding the contract method 0x0d8754ac.
//
// Solidity: function isCoinbase( address) constant returns(bool)
func (_NextyGovernance *NextyGovernanceCaller) IsCoinbase(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NextyGovernance.contract.Call(opts, out, "isCoinbase", arg0)
	return *ret0, err
}

// IsCoinbase is a free data retrieval call binding the contract method 0x0d8754ac.
//
// Solidity: function isCoinbase( address) constant returns(bool)
func (_NextyGovernance *NextyGovernanceSession) IsCoinbase(arg0 common.Address) (bool, error) {
	return _NextyGovernance.Contract.IsCoinbase(&_NextyGovernance.CallOpts, arg0)
}

// IsCoinbase is a free data retrieval call binding the contract method 0x0d8754ac.
//
// Solidity: function isCoinbase( address) constant returns(bool)
func (_NextyGovernance *NextyGovernanceCallerSession) IsCoinbase(arg0 common.Address) (bool, error) {
	return _NextyGovernance.Contract.IsCoinbase(&_NextyGovernance.CallOpts, arg0)
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

// Sealers is a free data retrieval call binding the contract method 0xaa61aef3.
//
// Solidity: function sealers( address) constant returns(status uint8, balance uint256, coinbase address, unlockTime uint256)
func (_NextyGovernance *NextyGovernanceCaller) Sealers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Status     uint8
	Balance    *big.Int
	Coinbase   common.Address
	UnlockTime *big.Int
}, error) {
	ret := new(struct {
		Status     uint8
		Balance    *big.Int
		Coinbase   common.Address
		UnlockTime *big.Int
	})
	out := ret
	err := _NextyGovernance.contract.Call(opts, out, "sealers", arg0)
	return *ret, err
}

// Sealers is a free data retrieval call binding the contract method 0xaa61aef3.
//
// Solidity: function sealers( address) constant returns(status uint8, balance uint256, coinbase address, unlockTime uint256)
func (_NextyGovernance *NextyGovernanceSession) Sealers(arg0 common.Address) (struct {
	Status     uint8
	Balance    *big.Int
	Coinbase   common.Address
	UnlockTime *big.Int
}, error) {
	return _NextyGovernance.Contract.Sealers(&_NextyGovernance.CallOpts, arg0)
}

// Sealers is a free data retrieval call binding the contract method 0xaa61aef3.
//
// Solidity: function sealers( address) constant returns(status uint8, balance uint256, coinbase address, unlockTime uint256)
func (_NextyGovernance *NextyGovernanceCallerSession) Sealers(arg0 common.Address) (struct {
	Status     uint8
	Balance    *big.Int
	Coinbase   common.Address
	UnlockTime *big.Int
}, error) {
	return _NextyGovernance.Contract.Sealers(&_NextyGovernance.CallOpts, arg0)
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
// Solidity: function join(_coinbase address) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactor) Join(opts *bind.TransactOpts, _coinbase common.Address) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "join", _coinbase)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(_coinbase address) returns(bool)
func (_NextyGovernance *NextyGovernanceSession) Join(_coinbase common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Join(&_NextyGovernance.TransactOpts, _coinbase)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(_coinbase address) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactorSession) Join(_coinbase common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Join(&_NextyGovernance.TransactOpts, _coinbase)
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
	Sealer   common.Address
	Coinbase common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterJoined is a free log retrieval operation binding the contract event 0x7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea.
//
// Solidity: e Joined(_sealer address, _coinbase address)
func (_NextyGovernance *NextyGovernanceFilterer) FilterJoined(opts *bind.FilterOpts) (*NextyGovernanceJoinedIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Joined")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceJoinedIterator{contract: _NextyGovernance.contract, event: "Joined", logs: logs, sub: sub}, nil
}

// WatchJoined is a free log subscription operation binding the contract event 0x7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea.
//
// Solidity: e Joined(_sealer address, _coinbase address)
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
	Sealer   common.Address
	Coinbase common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLeft is a free log retrieval operation binding the contract event 0x4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5.
//
// Solidity: e Left(_sealer address, _coinbase address)
func (_NextyGovernance *NextyGovernanceFilterer) FilterLeft(opts *bind.FilterOpts) (*NextyGovernanceLeftIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Left")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceLeftIterator{contract: _NextyGovernance.contract, event: "Left", logs: logs, sub: sub}, nil
}

// WatchLeft is a free log subscription operation binding the contract event 0x4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5.
//
// Solidity: e Left(_sealer address, _coinbase address)
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
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820be508e32da7c78d10a71311f9952845d5d4d9713aa81be2f8be9b007b922e5540029`

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
