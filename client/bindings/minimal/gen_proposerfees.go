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

// IProposerFeesMetaData contains all meta data concerning the IProposerFees contract.
var IProposerFeesMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getCurrentFees\",\"inputs\":[],\"outputs\":[{\"name\":\"fee\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"delayedFee\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payPublicationFee\",\"inputs\":[{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isDelayed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IProposerFeesABI is the input ABI used to generate the binding from.
// Deprecated: Use IProposerFeesMetaData.ABI instead.
var IProposerFeesABI = IProposerFeesMetaData.ABI

// IProposerFees is an auto generated Go binding around an Ethereum contract.
type IProposerFees struct {
	IProposerFeesCaller     // Read-only binding to the contract
	IProposerFeesTransactor // Write-only binding to the contract
	IProposerFeesFilterer   // Log filterer for contract events
}

// IProposerFeesCaller is an auto generated read-only Go binding around an Ethereum contract.
type IProposerFeesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProposerFeesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IProposerFeesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProposerFeesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IProposerFeesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProposerFeesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IProposerFeesSession struct {
	Contract     *IProposerFees    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IProposerFeesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IProposerFeesCallerSession struct {
	Contract *IProposerFeesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IProposerFeesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IProposerFeesTransactorSession struct {
	Contract     *IProposerFeesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IProposerFeesRaw is an auto generated low-level Go binding around an Ethereum contract.
type IProposerFeesRaw struct {
	Contract *IProposerFees // Generic contract binding to access the raw methods on
}

// IProposerFeesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IProposerFeesCallerRaw struct {
	Contract *IProposerFeesCaller // Generic read-only contract binding to access the raw methods on
}

// IProposerFeesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IProposerFeesTransactorRaw struct {
	Contract *IProposerFeesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIProposerFees creates a new instance of IProposerFees, bound to a specific deployed contract.
func NewIProposerFees(address common.Address, backend bind.ContractBackend) (*IProposerFees, error) {
	contract, err := bindIProposerFees(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IProposerFees{IProposerFeesCaller: IProposerFeesCaller{contract: contract}, IProposerFeesTransactor: IProposerFeesTransactor{contract: contract}, IProposerFeesFilterer: IProposerFeesFilterer{contract: contract}}, nil
}

// NewIProposerFeesCaller creates a new read-only instance of IProposerFees, bound to a specific deployed contract.
func NewIProposerFeesCaller(address common.Address, caller bind.ContractCaller) (*IProposerFeesCaller, error) {
	contract, err := bindIProposerFees(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IProposerFeesCaller{contract: contract}, nil
}

// NewIProposerFeesTransactor creates a new write-only instance of IProposerFees, bound to a specific deployed contract.
func NewIProposerFeesTransactor(address common.Address, transactor bind.ContractTransactor) (*IProposerFeesTransactor, error) {
	contract, err := bindIProposerFees(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IProposerFeesTransactor{contract: contract}, nil
}

// NewIProposerFeesFilterer creates a new log filterer instance of IProposerFees, bound to a specific deployed contract.
func NewIProposerFeesFilterer(address common.Address, filterer bind.ContractFilterer) (*IProposerFeesFilterer, error) {
	contract, err := bindIProposerFees(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IProposerFeesFilterer{contract: contract}, nil
}

// bindIProposerFees binds a generic wrapper to an already deployed contract.
func bindIProposerFees(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IProposerFeesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProposerFees *IProposerFeesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProposerFees.Contract.IProposerFeesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProposerFees *IProposerFeesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProposerFees.Contract.IProposerFeesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProposerFees *IProposerFeesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProposerFees.Contract.IProposerFeesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProposerFees *IProposerFeesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProposerFees.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProposerFees *IProposerFeesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProposerFees.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProposerFees *IProposerFeesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProposerFees.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentFees is a free data retrieval call binding the contract method 0x71908a03.
//
// Solidity: function getCurrentFees() view returns(uint96 fee, uint96 delayedFee)
func (_IProposerFees *IProposerFeesCaller) GetCurrentFees(opts *bind.CallOpts) (struct {
	Fee        *big.Int
	DelayedFee *big.Int
}, error) {
	var out []interface{}
	err := _IProposerFees.contract.Call(opts, &out, "getCurrentFees")

	outstruct := new(struct {
		Fee        *big.Int
		DelayedFee *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DelayedFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCurrentFees is a free data retrieval call binding the contract method 0x71908a03.
//
// Solidity: function getCurrentFees() view returns(uint96 fee, uint96 delayedFee)
func (_IProposerFees *IProposerFeesSession) GetCurrentFees() (struct {
	Fee        *big.Int
	DelayedFee *big.Int
}, error) {
	return _IProposerFees.Contract.GetCurrentFees(&_IProposerFees.CallOpts)
}

// GetCurrentFees is a free data retrieval call binding the contract method 0x71908a03.
//
// Solidity: function getCurrentFees() view returns(uint96 fee, uint96 delayedFee)
func (_IProposerFees *IProposerFeesCallerSession) GetCurrentFees() (struct {
	Fee        *big.Int
	DelayedFee *big.Int
}, error) {
	return _IProposerFees.Contract.GetCurrentFees(&_IProposerFees.CallOpts)
}

// PayPublicationFee is a paid mutator transaction binding the contract method 0xd51ca5c7.
//
// Solidity: function payPublicationFee(address proposer, bool isDelayed) returns()
func (_IProposerFees *IProposerFeesTransactor) PayPublicationFee(opts *bind.TransactOpts, proposer common.Address, isDelayed bool) (*types.Transaction, error) {
	return _IProposerFees.contract.Transact(opts, "payPublicationFee", proposer, isDelayed)
}

// PayPublicationFee is a paid mutator transaction binding the contract method 0xd51ca5c7.
//
// Solidity: function payPublicationFee(address proposer, bool isDelayed) returns()
func (_IProposerFees *IProposerFeesSession) PayPublicationFee(proposer common.Address, isDelayed bool) (*types.Transaction, error) {
	return _IProposerFees.Contract.PayPublicationFee(&_IProposerFees.TransactOpts, proposer, isDelayed)
}

// PayPublicationFee is a paid mutator transaction binding the contract method 0xd51ca5c7.
//
// Solidity: function payPublicationFee(address proposer, bool isDelayed) returns()
func (_IProposerFees *IProposerFeesTransactorSession) PayPublicationFee(proposer common.Address, isDelayed bool) (*types.Transaction, error) {
	return _IProposerFees.Contract.PayPublicationFee(&_IProposerFees.TransactOpts, proposer, isDelayed)
}
