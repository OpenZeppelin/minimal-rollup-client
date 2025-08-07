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

// ILookaheadMetaData contains all meta data concerning the ILookahead contract.
var ILookaheadMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"isCurrentPreconfer\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"}]",
}

// ILookaheadABI is the input ABI used to generate the binding from.
// Deprecated: Use ILookaheadMetaData.ABI instead.
var ILookaheadABI = ILookaheadMetaData.ABI

// ILookahead is an auto generated Go binding around an Ethereum contract.
type ILookahead struct {
	ILookaheadCaller     // Read-only binding to the contract
	ILookaheadTransactor // Write-only binding to the contract
	ILookaheadFilterer   // Log filterer for contract events
}

// ILookaheadCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILookaheadCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILookaheadTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILookaheadTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILookaheadFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILookaheadFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILookaheadSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILookaheadSession struct {
	Contract     *ILookahead       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILookaheadCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILookaheadCallerSession struct {
	Contract *ILookaheadCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ILookaheadTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILookaheadTransactorSession struct {
	Contract     *ILookaheadTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ILookaheadRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILookaheadRaw struct {
	Contract *ILookahead // Generic contract binding to access the raw methods on
}

// ILookaheadCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILookaheadCallerRaw struct {
	Contract *ILookaheadCaller // Generic read-only contract binding to access the raw methods on
}

// ILookaheadTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILookaheadTransactorRaw struct {
	Contract *ILookaheadTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILookahead creates a new instance of ILookahead, bound to a specific deployed contract.
func NewILookahead(address common.Address, backend bind.ContractBackend) (*ILookahead, error) {
	contract, err := bindILookahead(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILookahead{ILookaheadCaller: ILookaheadCaller{contract: contract}, ILookaheadTransactor: ILookaheadTransactor{contract: contract}, ILookaheadFilterer: ILookaheadFilterer{contract: contract}}, nil
}

// NewILookaheadCaller creates a new read-only instance of ILookahead, bound to a specific deployed contract.
func NewILookaheadCaller(address common.Address, caller bind.ContractCaller) (*ILookaheadCaller, error) {
	contract, err := bindILookahead(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILookaheadCaller{contract: contract}, nil
}

// NewILookaheadTransactor creates a new write-only instance of ILookahead, bound to a specific deployed contract.
func NewILookaheadTransactor(address common.Address, transactor bind.ContractTransactor) (*ILookaheadTransactor, error) {
	contract, err := bindILookahead(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILookaheadTransactor{contract: contract}, nil
}

// NewILookaheadFilterer creates a new log filterer instance of ILookahead, bound to a specific deployed contract.
func NewILookaheadFilterer(address common.Address, filterer bind.ContractFilterer) (*ILookaheadFilterer, error) {
	contract, err := bindILookahead(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILookaheadFilterer{contract: contract}, nil
}

// bindILookahead binds a generic wrapper to an already deployed contract.
func bindILookahead(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ILookaheadMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILookahead *ILookaheadRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILookahead.Contract.ILookaheadCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILookahead *ILookaheadRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILookahead.Contract.ILookaheadTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILookahead *ILookaheadRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILookahead.Contract.ILookaheadTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILookahead *ILookaheadCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILookahead.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILookahead *ILookaheadTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILookahead.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILookahead *ILookaheadTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILookahead.Contract.contract.Transact(opts, method, params...)
}

// IsCurrentPreconfer is a paid mutator transaction binding the contract method 0x551c50c2.
//
// Solidity: function isCurrentPreconfer(address addr) returns(bool)
func (_ILookahead *ILookaheadTransactor) IsCurrentPreconfer(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ILookahead.contract.Transact(opts, "isCurrentPreconfer", addr)
}

// IsCurrentPreconfer is a paid mutator transaction binding the contract method 0x551c50c2.
//
// Solidity: function isCurrentPreconfer(address addr) returns(bool)
func (_ILookahead *ILookaheadSession) IsCurrentPreconfer(addr common.Address) (*types.Transaction, error) {
	return _ILookahead.Contract.IsCurrentPreconfer(&_ILookahead.TransactOpts, addr)
}

// IsCurrentPreconfer is a paid mutator transaction binding the contract method 0x551c50c2.
//
// Solidity: function isCurrentPreconfer(address addr) returns(bool)
func (_ILookahead *ILookaheadTransactorSession) IsCurrentPreconfer(addr common.Address) (*types.Transaction, error) {
	return _ILookahead.Contract.IsCurrentPreconfer(&_ILookahead.TransactOpts, addr)
}
