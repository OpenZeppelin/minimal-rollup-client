// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package minimal

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

// IETHDepositorMetaData contains all meta data concerning the IETHDepositor contract.
var IETHDepositorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"error\",\"name\":\"InsufficientETHDeposit\",\"inputs\":[{\"name\":\"deposited\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"WithdrawFailed\",\"inputs\":[]}]",
}

// IETHDepositorABI is the input ABI used to generate the binding from.
// Deprecated: Use IETHDepositorMetaData.ABI instead.
var IETHDepositorABI = IETHDepositorMetaData.ABI

// IETHDepositor is an auto generated Go binding around an Ethereum contract.
type IETHDepositor struct {
	IETHDepositorCaller     // Read-only binding to the contract
	IETHDepositorTransactor // Write-only binding to the contract
	IETHDepositorFilterer   // Log filterer for contract events
}

// IETHDepositorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IETHDepositorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IETHDepositorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IETHDepositorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IETHDepositorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IETHDepositorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IETHDepositorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IETHDepositorSession struct {
	Contract     *IETHDepositor    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IETHDepositorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IETHDepositorCallerSession struct {
	Contract *IETHDepositorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IETHDepositorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IETHDepositorTransactorSession struct {
	Contract     *IETHDepositorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IETHDepositorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IETHDepositorRaw struct {
	Contract *IETHDepositor // Generic contract binding to access the raw methods on
}

// IETHDepositorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IETHDepositorCallerRaw struct {
	Contract *IETHDepositorCaller // Generic read-only contract binding to access the raw methods on
}

// IETHDepositorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IETHDepositorTransactorRaw struct {
	Contract *IETHDepositorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIETHDepositor creates a new instance of IETHDepositor, bound to a specific deployed contract.
func NewIETHDepositor(address common.Address, backend bind.ContractBackend) (*IETHDepositor, error) {
	contract, err := bindIETHDepositor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IETHDepositor{IETHDepositorCaller: IETHDepositorCaller{contract: contract}, IETHDepositorTransactor: IETHDepositorTransactor{contract: contract}, IETHDepositorFilterer: IETHDepositorFilterer{contract: contract}}, nil
}

// NewIETHDepositorCaller creates a new read-only instance of IETHDepositor, bound to a specific deployed contract.
func NewIETHDepositorCaller(address common.Address, caller bind.ContractCaller) (*IETHDepositorCaller, error) {
	contract, err := bindIETHDepositor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IETHDepositorCaller{contract: contract}, nil
}

// NewIETHDepositorTransactor creates a new write-only instance of IETHDepositor, bound to a specific deployed contract.
func NewIETHDepositorTransactor(address common.Address, transactor bind.ContractTransactor) (*IETHDepositorTransactor, error) {
	contract, err := bindIETHDepositor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IETHDepositorTransactor{contract: contract}, nil
}

// NewIETHDepositorFilterer creates a new log filterer instance of IETHDepositor, bound to a specific deployed contract.
func NewIETHDepositorFilterer(address common.Address, filterer bind.ContractFilterer) (*IETHDepositorFilterer, error) {
	contract, err := bindIETHDepositor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IETHDepositorFilterer{contract: contract}, nil
}

// bindIETHDepositor binds a generic wrapper to an already deployed contract.
func bindIETHDepositor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IETHDepositorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IETHDepositor *IETHDepositorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IETHDepositor.Contract.IETHDepositorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IETHDepositor *IETHDepositorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IETHDepositor.Contract.IETHDepositorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IETHDepositor *IETHDepositorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IETHDepositor.Contract.IETHDepositorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IETHDepositor *IETHDepositorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IETHDepositor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IETHDepositor *IETHDepositorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IETHDepositor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IETHDepositor *IETHDepositorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IETHDepositor.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IETHDepositor *IETHDepositorTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IETHDepositor.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IETHDepositor *IETHDepositorSession) Deposit() (*types.Transaction, error) {
	return _IETHDepositor.Contract.Deposit(&_IETHDepositor.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IETHDepositor *IETHDepositorTransactorSession) Deposit() (*types.Transaction, error) {
	return _IETHDepositor.Contract.Deposit(&_IETHDepositor.TransactOpts)
}
