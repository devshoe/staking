// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"loaner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"}],\"name\":\"Loan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pending\",\"type\":\"uint256\"}],\"name\":\"Payout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"loaner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pending\",\"type\":\"uint256\"}],\"name\":\"Repayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yieldRate\",\"type\":\"uint256\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accruementCycleLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newrate\",\"type\":\"uint256\"}],\"name\":\"adjustInterestRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newrate\",\"type\":\"uint256\"}],\"name\":\"adjustYieldRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentAccruementLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentYieldRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405262278d00600155600160025534801561001c57600080fd5b50600080546001600160a01b031916331790556102588061003e6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80636999ac93116100665780636999ac93146101025780637c3a00fd1461010b5780638614a39114610114578063d880a0e11461011c578063ff4c7e311461012557600080fd5b806302d05d3f1461009857806330ebe02f146100c85780633456c80c146100dd57806342669f13146100f0575b600080fd5b6000546100ab906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100db6100d6366004610209565b61012d565b005b6100db6100eb366004610209565b61019f565b6001545b6040519081526020016100bf565b6100f460025481565b6100f460035481565b6002546100f4565b6100f460015481565b6003546100f4565b6000546001600160a01b0316331461019a5760405162461bcd60e51b815260206004820152602560248201527f6f6e6c792063726561746f722063616e2061646a75737420696e746572657374604482015264207261746560d81b60648201526084015b60405180910390fd5b600355565b6000546001600160a01b031633146102045760405162461bcd60e51b815260206004820152602260248201527f6f6e6c792063726561746f722063616e2061646a757374207969656c64207261604482015261746560f01b6064820152608401610191565b600255565b60006020828403121561021b57600080fd5b503591905056fea264697066735822122097123bb047f844989b6e5a7824aa4f0a07ae0ad626ea444f530b34f2e42f0e3c64736f6c63430008120033",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// MainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MainMetaData.Bin instead.
var MainBin = MainMetaData.Bin

// DeployMain deploys a new Ethereum contract, binding an instance of Main to it.
func DeployMain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Main, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// AccruementCycleLength is a free data retrieval call binding the contract method 0xd880a0e1.
//
// Solidity: function accruementCycleLength() view returns(uint256)
func (_Main *MainCaller) AccruementCycleLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "accruementCycleLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruementCycleLength is a free data retrieval call binding the contract method 0xd880a0e1.
//
// Solidity: function accruementCycleLength() view returns(uint256)
func (_Main *MainSession) AccruementCycleLength() (*big.Int, error) {
	return _Main.Contract.AccruementCycleLength(&_Main.CallOpts)
}

// AccruementCycleLength is a free data retrieval call binding the contract method 0xd880a0e1.
//
// Solidity: function accruementCycleLength() view returns(uint256)
func (_Main *MainCallerSession) AccruementCycleLength() (*big.Int, error) {
	return _Main.Contract.AccruementCycleLength(&_Main.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Main *MainCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Main *MainSession) Creator() (common.Address, error) {
	return _Main.Contract.Creator(&_Main.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Main *MainCallerSession) Creator() (common.Address, error) {
	return _Main.Contract.Creator(&_Main.CallOpts)
}

// CurrentAccruementLength is a free data retrieval call binding the contract method 0x42669f13.
//
// Solidity: function currentAccruementLength() view returns(uint256)
func (_Main *MainCaller) CurrentAccruementLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "currentAccruementLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentAccruementLength is a free data retrieval call binding the contract method 0x42669f13.
//
// Solidity: function currentAccruementLength() view returns(uint256)
func (_Main *MainSession) CurrentAccruementLength() (*big.Int, error) {
	return _Main.Contract.CurrentAccruementLength(&_Main.CallOpts)
}

// CurrentAccruementLength is a free data retrieval call binding the contract method 0x42669f13.
//
// Solidity: function currentAccruementLength() view returns(uint256)
func (_Main *MainCallerSession) CurrentAccruementLength() (*big.Int, error) {
	return _Main.Contract.CurrentAccruementLength(&_Main.CallOpts)
}

// CurrentInterestRate is a free data retrieval call binding the contract method 0xff4c7e31.
//
// Solidity: function currentInterestRate() view returns(uint256)
func (_Main *MainCaller) CurrentInterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "currentInterestRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentInterestRate is a free data retrieval call binding the contract method 0xff4c7e31.
//
// Solidity: function currentInterestRate() view returns(uint256)
func (_Main *MainSession) CurrentInterestRate() (*big.Int, error) {
	return _Main.Contract.CurrentInterestRate(&_Main.CallOpts)
}

// CurrentInterestRate is a free data retrieval call binding the contract method 0xff4c7e31.
//
// Solidity: function currentInterestRate() view returns(uint256)
func (_Main *MainCallerSession) CurrentInterestRate() (*big.Int, error) {
	return _Main.Contract.CurrentInterestRate(&_Main.CallOpts)
}

// CurrentYieldRate is a free data retrieval call binding the contract method 0x8614a391.
//
// Solidity: function currentYieldRate() view returns(uint256)
func (_Main *MainCaller) CurrentYieldRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "currentYieldRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentYieldRate is a free data retrieval call binding the contract method 0x8614a391.
//
// Solidity: function currentYieldRate() view returns(uint256)
func (_Main *MainSession) CurrentYieldRate() (*big.Int, error) {
	return _Main.Contract.CurrentYieldRate(&_Main.CallOpts)
}

// CurrentYieldRate is a free data retrieval call binding the contract method 0x8614a391.
//
// Solidity: function currentYieldRate() view returns(uint256)
func (_Main *MainCallerSession) CurrentYieldRate() (*big.Int, error) {
	return _Main.Contract.CurrentYieldRate(&_Main.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_Main *MainCaller) InterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "interestRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_Main *MainSession) InterestRate() (*big.Int, error) {
	return _Main.Contract.InterestRate(&_Main.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_Main *MainCallerSession) InterestRate() (*big.Int, error) {
	return _Main.Contract.InterestRate(&_Main.CallOpts)
}

// YieldRate is a free data retrieval call binding the contract method 0x6999ac93.
//
// Solidity: function yieldRate() view returns(uint256)
func (_Main *MainCaller) YieldRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "yieldRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YieldRate is a free data retrieval call binding the contract method 0x6999ac93.
//
// Solidity: function yieldRate() view returns(uint256)
func (_Main *MainSession) YieldRate() (*big.Int, error) {
	return _Main.Contract.YieldRate(&_Main.CallOpts)
}

// YieldRate is a free data retrieval call binding the contract method 0x6999ac93.
//
// Solidity: function yieldRate() view returns(uint256)
func (_Main *MainCallerSession) YieldRate() (*big.Int, error) {
	return _Main.Contract.YieldRate(&_Main.CallOpts)
}

// AdjustInterestRate is a paid mutator transaction binding the contract method 0x30ebe02f.
//
// Solidity: function adjustInterestRate(uint256 newrate) returns()
func (_Main *MainTransactor) AdjustInterestRate(opts *bind.TransactOpts, newrate *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "adjustInterestRate", newrate)
}

// AdjustInterestRate is a paid mutator transaction binding the contract method 0x30ebe02f.
//
// Solidity: function adjustInterestRate(uint256 newrate) returns()
func (_Main *MainSession) AdjustInterestRate(newrate *big.Int) (*types.Transaction, error) {
	return _Main.Contract.AdjustInterestRate(&_Main.TransactOpts, newrate)
}

// AdjustInterestRate is a paid mutator transaction binding the contract method 0x30ebe02f.
//
// Solidity: function adjustInterestRate(uint256 newrate) returns()
func (_Main *MainTransactorSession) AdjustInterestRate(newrate *big.Int) (*types.Transaction, error) {
	return _Main.Contract.AdjustInterestRate(&_Main.TransactOpts, newrate)
}

// AdjustYieldRate is a paid mutator transaction binding the contract method 0x3456c80c.
//
// Solidity: function adjustYieldRate(uint256 newrate) returns()
func (_Main *MainTransactor) AdjustYieldRate(opts *bind.TransactOpts, newrate *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "adjustYieldRate", newrate)
}

// AdjustYieldRate is a paid mutator transaction binding the contract method 0x3456c80c.
//
// Solidity: function adjustYieldRate(uint256 newrate) returns()
func (_Main *MainSession) AdjustYieldRate(newrate *big.Int) (*types.Transaction, error) {
	return _Main.Contract.AdjustYieldRate(&_Main.TransactOpts, newrate)
}

// AdjustYieldRate is a paid mutator transaction binding the contract method 0x3456c80c.
//
// Solidity: function adjustYieldRate(uint256 newrate) returns()
func (_Main *MainTransactorSession) AdjustYieldRate(newrate *big.Int) (*types.Transaction, error) {
	return _Main.Contract.AdjustYieldRate(&_Main.TransactOpts, newrate)
}

// MainLoanIterator is returned from FilterLoan and is used to iterate over the raw logs and unpacked data for Loan events raised by the Main contract.
type MainLoanIterator struct {
	Event *MainLoan // Event containing the contract specifics and raw log

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
func (it *MainLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainLoan)
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
		it.Event = new(MainLoan)
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
func (it *MainLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainLoan represents a Loan event raised by the Main contract.
type MainLoan struct {
	Loaner       common.Address
	Amount       *big.Int
	InterestRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLoan is a free log retrieval operation binding the contract event 0xe73e31007d63465eceaabfae515099de1ee3152614bd2824a7e75ce65517ae16.
//
// Solidity: event Loan(address indexed loaner, uint256 amount, uint256 interestRate)
func (_Main *MainFilterer) FilterLoan(opts *bind.FilterOpts, loaner []common.Address) (*MainLoanIterator, error) {

	var loanerRule []interface{}
	for _, loanerItem := range loaner {
		loanerRule = append(loanerRule, loanerItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Loan", loanerRule)
	if err != nil {
		return nil, err
	}
	return &MainLoanIterator{contract: _Main.contract, event: "Loan", logs: logs, sub: sub}, nil
}

// WatchLoan is a free log subscription operation binding the contract event 0xe73e31007d63465eceaabfae515099de1ee3152614bd2824a7e75ce65517ae16.
//
// Solidity: event Loan(address indexed loaner, uint256 amount, uint256 interestRate)
func (_Main *MainFilterer) WatchLoan(opts *bind.WatchOpts, sink chan<- *MainLoan, loaner []common.Address) (event.Subscription, error) {

	var loanerRule []interface{}
	for _, loanerItem := range loaner {
		loanerRule = append(loanerRule, loanerItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Loan", loanerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainLoan)
				if err := _Main.contract.UnpackLog(event, "Loan", log); err != nil {
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

// ParseLoan is a log parse operation binding the contract event 0xe73e31007d63465eceaabfae515099de1ee3152614bd2824a7e75ce65517ae16.
//
// Solidity: event Loan(address indexed loaner, uint256 amount, uint256 interestRate)
func (_Main *MainFilterer) ParseLoan(log types.Log) (*MainLoan, error) {
	event := new(MainLoan)
	if err := _Main.contract.UnpackLog(event, "Loan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainPayoutIterator is returned from FilterPayout and is used to iterate over the raw logs and unpacked data for Payout events raised by the Main contract.
type MainPayoutIterator struct {
	Event *MainPayout // Event containing the contract specifics and raw log

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
func (it *MainPayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainPayout)
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
		it.Event = new(MainPayout)
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
func (it *MainPayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainPayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainPayout represents a Payout event raised by the Main contract.
type MainPayout struct {
	Depositor common.Address
	Amount    *big.Int
	Pending   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPayout is a free log retrieval operation binding the contract event 0x634235fcf5af0adbca1a405ec65f6f6c08f55e1f379c2c45cd10f23cb29e0e31.
//
// Solidity: event Payout(address indexed depositor, uint256 amount, uint256 pending)
func (_Main *MainFilterer) FilterPayout(opts *bind.FilterOpts, depositor []common.Address) (*MainPayoutIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Payout", depositorRule)
	if err != nil {
		return nil, err
	}
	return &MainPayoutIterator{contract: _Main.contract, event: "Payout", logs: logs, sub: sub}, nil
}

// WatchPayout is a free log subscription operation binding the contract event 0x634235fcf5af0adbca1a405ec65f6f6c08f55e1f379c2c45cd10f23cb29e0e31.
//
// Solidity: event Payout(address indexed depositor, uint256 amount, uint256 pending)
func (_Main *MainFilterer) WatchPayout(opts *bind.WatchOpts, sink chan<- *MainPayout, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Payout", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainPayout)
				if err := _Main.contract.UnpackLog(event, "Payout", log); err != nil {
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

// ParsePayout is a log parse operation binding the contract event 0x634235fcf5af0adbca1a405ec65f6f6c08f55e1f379c2c45cd10f23cb29e0e31.
//
// Solidity: event Payout(address indexed depositor, uint256 amount, uint256 pending)
func (_Main *MainFilterer) ParsePayout(log types.Log) (*MainPayout, error) {
	event := new(MainPayout)
	if err := _Main.contract.UnpackLog(event, "Payout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRepaymentIterator is returned from FilterRepayment and is used to iterate over the raw logs and unpacked data for Repayment events raised by the Main contract.
type MainRepaymentIterator struct {
	Event *MainRepayment // Event containing the contract specifics and raw log

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
func (it *MainRepaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRepayment)
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
		it.Event = new(MainRepayment)
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
func (it *MainRepaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRepaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRepayment represents a Repayment event raised by the Main contract.
type MainRepayment struct {
	Loaner  common.Address
	Amount  *big.Int
	Pending *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRepayment is a free log retrieval operation binding the contract event 0x24fcca58a997b1b2eff6db8107e860458544c09ddd3693b3b779e1df6c0d6c5d.
//
// Solidity: event Repayment(address indexed loaner, uint256 amount, uint256 pending)
func (_Main *MainFilterer) FilterRepayment(opts *bind.FilterOpts, loaner []common.Address) (*MainRepaymentIterator, error) {

	var loanerRule []interface{}
	for _, loanerItem := range loaner {
		loanerRule = append(loanerRule, loanerItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Repayment", loanerRule)
	if err != nil {
		return nil, err
	}
	return &MainRepaymentIterator{contract: _Main.contract, event: "Repayment", logs: logs, sub: sub}, nil
}

// WatchRepayment is a free log subscription operation binding the contract event 0x24fcca58a997b1b2eff6db8107e860458544c09ddd3693b3b779e1df6c0d6c5d.
//
// Solidity: event Repayment(address indexed loaner, uint256 amount, uint256 pending)
func (_Main *MainFilterer) WatchRepayment(opts *bind.WatchOpts, sink chan<- *MainRepayment, loaner []common.Address) (event.Subscription, error) {

	var loanerRule []interface{}
	for _, loanerItem := range loaner {
		loanerRule = append(loanerRule, loanerItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Repayment", loanerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRepayment)
				if err := _Main.contract.UnpackLog(event, "Repayment", log); err != nil {
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

// ParseRepayment is a log parse operation binding the contract event 0x24fcca58a997b1b2eff6db8107e860458544c09ddd3693b3b779e1df6c0d6c5d.
//
// Solidity: event Repayment(address indexed loaner, uint256 amount, uint256 pending)
func (_Main *MainFilterer) ParseRepayment(log types.Log) (*MainRepayment, error) {
	event := new(MainRepayment)
	if err := _Main.contract.UnpackLog(event, "Repayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the Main contract.
type MainStakeIterator struct {
	Event *MainStake // Event containing the contract specifics and raw log

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
func (it *MainStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainStake)
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
		it.Event = new(MainStake)
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
func (it *MainStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainStake represents a Stake event raised by the Main contract.
type MainStake struct {
	Depositor common.Address
	Amount    *big.Int
	YieldRate *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0x5af417134f72a9d41143ace85b0a26dce6f550f894f2cbc1eeee8810603d91b6.
//
// Solidity: event Stake(address indexed depositor, uint256 amount, uint256 yieldRate)
func (_Main *MainFilterer) FilterStake(opts *bind.FilterOpts, depositor []common.Address) (*MainStakeIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Stake", depositorRule)
	if err != nil {
		return nil, err
	}
	return &MainStakeIterator{contract: _Main.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0x5af417134f72a9d41143ace85b0a26dce6f550f894f2cbc1eeee8810603d91b6.
//
// Solidity: event Stake(address indexed depositor, uint256 amount, uint256 yieldRate)
func (_Main *MainFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *MainStake, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Stake", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainStake)
				if err := _Main.contract.UnpackLog(event, "Stake", log); err != nil {
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

// ParseStake is a log parse operation binding the contract event 0x5af417134f72a9d41143ace85b0a26dce6f550f894f2cbc1eeee8810603d91b6.
//
// Solidity: event Stake(address indexed depositor, uint256 amount, uint256 yieldRate)
func (_Main *MainFilterer) ParseStake(log types.Log) (*MainStake, error) {
	event := new(MainStake)
	if err := _Main.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
