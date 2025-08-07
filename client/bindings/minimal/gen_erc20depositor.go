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

// IERC20DepositorMetaData contains all meta data concerning the IERC20Depositor contract.
var IERC20DepositorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"InsufficientInitialDeposit\",\"inputs\":[{\"name\":\"deposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredBond\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroTokenAddress\",\"inputs\":[]}]",
}

// IERC20DepositorABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20DepositorMetaData.ABI instead.
var IERC20DepositorABI = IERC20DepositorMetaData.ABI

// IERC20Depositor is an auto generated Go binding around an Ethereum contract.
type IERC20Depositor struct {
	IERC20DepositorCaller     // Read-only binding to the contract
	IERC20DepositorTransactor // Write-only binding to the contract
	IERC20DepositorFilterer   // Log filterer for contract events
}

// IERC20DepositorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20DepositorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20DepositorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20DepositorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20DepositorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20DepositorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20DepositorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20DepositorSession struct {
	Contract     *IERC20Depositor  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20DepositorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20DepositorCallerSession struct {
	Contract *IERC20DepositorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IERC20DepositorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20DepositorTransactorSession struct {
	Contract     *IERC20DepositorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC20DepositorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20DepositorRaw struct {
	Contract *IERC20Depositor // Generic contract binding to access the raw methods on
}

// IERC20DepositorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20DepositorCallerRaw struct {
	Contract *IERC20DepositorCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20DepositorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20DepositorTransactorRaw struct {
	Contract *IERC20DepositorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Depositor creates a new instance of IERC20Depositor, bound to a specific deployed contract.
func NewIERC20Depositor(address common.Address, backend bind.ContractBackend) (*IERC20Depositor, error) {
	contract, err := bindIERC20Depositor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Depositor{IERC20DepositorCaller: IERC20DepositorCaller{contract: contract}, IERC20DepositorTransactor: IERC20DepositorTransactor{contract: contract}, IERC20DepositorFilterer: IERC20DepositorFilterer{contract: contract}}, nil
}

// NewIERC20DepositorCaller creates a new read-only instance of IERC20Depositor, bound to a specific deployed contract.
func NewIERC20DepositorCaller(address common.Address, caller bind.ContractCaller) (*IERC20DepositorCaller, error) {
	contract, err := bindIERC20Depositor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20DepositorCaller{contract: contract}, nil
}

// NewIERC20DepositorTransactor creates a new write-only instance of IERC20Depositor, bound to a specific deployed contract.
func NewIERC20DepositorTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20DepositorTransactor, error) {
	contract, err := bindIERC20Depositor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20DepositorTransactor{contract: contract}, nil
}

// NewIERC20DepositorFilterer creates a new log filterer instance of IERC20Depositor, bound to a specific deployed contract.
func NewIERC20DepositorFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20DepositorFilterer, error) {
	contract, err := bindIERC20Depositor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20DepositorFilterer{contract: contract}, nil
}

// bindIERC20Depositor binds a generic wrapper to an already deployed contract.
func bindIERC20Depositor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20DepositorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Depositor *IERC20DepositorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Depositor.Contract.IERC20DepositorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Depositor *IERC20DepositorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Depositor.Contract.IERC20DepositorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Depositor *IERC20DepositorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Depositor.Contract.IERC20DepositorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Depositor *IERC20DepositorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Depositor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Depositor *IERC20DepositorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Depositor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Depositor *IERC20DepositorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Depositor.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_IERC20Depositor *IERC20DepositorTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Depositor.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_IERC20Depositor *IERC20DepositorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _IERC20Depositor.Contract.Deposit(&_IERC20Depositor.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_IERC20Depositor *IERC20DepositorTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _IERC20Depositor.Contract.Deposit(&_IERC20Depositor.TransactOpts, amount)
}
