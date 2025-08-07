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

// IETHBridgeETHDeposit is an auto generated low-level Go binding around an user-defined struct.
type IETHBridgeETHDeposit struct {
	Nonce    *big.Int
	From     common.Address
	To       common.Address
	Amount   *big.Int
	Data     []byte
	Context  []byte
	Canceler common.Address
}

// IETHBridgeMetaData contains all meta data concerning the IETHBridge contract.
var IETHBridgeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"cancelDeposit\",\"inputs\":[{\"name\":\"ethDeposit\",\"type\":\"tuple\",\"internalType\":\"structIETHBridge.ETHDeposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"claimee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimDeposit\",\"inputs\":[{\"name\":\"ethDeposit\",\"type\":\"tuple\",\"internalType\":\"structIETHBridge.ETHDeposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getDepositId\",\"inputs\":[{\"name\":\"ethDeposit\",\"type\":\"tuple\",\"internalType\":\"structIETHBridge.ETHDeposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DepositCancelled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"claimee\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositClaimed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"deposit\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIETHBridge.ETHDeposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositMade\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"deposit\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIETHBridge.ETHDeposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyProcessed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyCounterpart\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptySignalService\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyTrustedPublisher\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedClaim\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCanceler\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroReceiver\",\"inputs\":[]}]",
}

// IETHBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use IETHBridgeMetaData.ABI instead.
var IETHBridgeABI = IETHBridgeMetaData.ABI

// IETHBridge is an auto generated Go binding around an Ethereum contract.
type IETHBridge struct {
	IETHBridgeCaller     // Read-only binding to the contract
	IETHBridgeTransactor // Write-only binding to the contract
	IETHBridgeFilterer   // Log filterer for contract events
}

// IETHBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IETHBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IETHBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IETHBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IETHBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IETHBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IETHBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IETHBridgeSession struct {
	Contract     *IETHBridge       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IETHBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IETHBridgeCallerSession struct {
	Contract *IETHBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IETHBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IETHBridgeTransactorSession struct {
	Contract     *IETHBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IETHBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IETHBridgeRaw struct {
	Contract *IETHBridge // Generic contract binding to access the raw methods on
}

// IETHBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IETHBridgeCallerRaw struct {
	Contract *IETHBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IETHBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IETHBridgeTransactorRaw struct {
	Contract *IETHBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIETHBridge creates a new instance of IETHBridge, bound to a specific deployed contract.
func NewIETHBridge(address common.Address, backend bind.ContractBackend) (*IETHBridge, error) {
	contract, err := bindIETHBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IETHBridge{IETHBridgeCaller: IETHBridgeCaller{contract: contract}, IETHBridgeTransactor: IETHBridgeTransactor{contract: contract}, IETHBridgeFilterer: IETHBridgeFilterer{contract: contract}}, nil
}

// NewIETHBridgeCaller creates a new read-only instance of IETHBridge, bound to a specific deployed contract.
func NewIETHBridgeCaller(address common.Address, caller bind.ContractCaller) (*IETHBridgeCaller, error) {
	contract, err := bindIETHBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IETHBridgeCaller{contract: contract}, nil
}

// NewIETHBridgeTransactor creates a new write-only instance of IETHBridge, bound to a specific deployed contract.
func NewIETHBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IETHBridgeTransactor, error) {
	contract, err := bindIETHBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IETHBridgeTransactor{contract: contract}, nil
}

// NewIETHBridgeFilterer creates a new log filterer instance of IETHBridge, bound to a specific deployed contract.
func NewIETHBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IETHBridgeFilterer, error) {
	contract, err := bindIETHBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IETHBridgeFilterer{contract: contract}, nil
}

// bindIETHBridge binds a generic wrapper to an already deployed contract.
func bindIETHBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IETHBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IETHBridge *IETHBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IETHBridge.Contract.IETHBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IETHBridge *IETHBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IETHBridge.Contract.IETHBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IETHBridge *IETHBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IETHBridge.Contract.IETHBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IETHBridge *IETHBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IETHBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IETHBridge *IETHBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IETHBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IETHBridge *IETHBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IETHBridge.Contract.contract.Transact(opts, method, params...)
}

// GetDepositId is a free data retrieval call binding the contract method 0x0f94f524.
//
// Solidity: function getDepositId((uint256,address,address,uint256,bytes,bytes,address) ethDeposit) view returns(bytes32 id)
func (_IETHBridge *IETHBridgeCaller) GetDepositId(opts *bind.CallOpts, ethDeposit IETHBridgeETHDeposit) ([32]byte, error) {
	var out []interface{}
	err := _IETHBridge.contract.Call(opts, &out, "getDepositId", ethDeposit)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDepositId is a free data retrieval call binding the contract method 0x0f94f524.
//
// Solidity: function getDepositId((uint256,address,address,uint256,bytes,bytes,address) ethDeposit) view returns(bytes32 id)
func (_IETHBridge *IETHBridgeSession) GetDepositId(ethDeposit IETHBridgeETHDeposit) ([32]byte, error) {
	return _IETHBridge.Contract.GetDepositId(&_IETHBridge.CallOpts, ethDeposit)
}

// GetDepositId is a free data retrieval call binding the contract method 0x0f94f524.
//
// Solidity: function getDepositId((uint256,address,address,uint256,bytes,bytes,address) ethDeposit) view returns(bytes32 id)
func (_IETHBridge *IETHBridgeCallerSession) GetDepositId(ethDeposit IETHBridgeETHDeposit) ([32]byte, error) {
	return _IETHBridge.Contract.GetDepositId(&_IETHBridge.CallOpts, ethDeposit)
}

// Processed is a free data retrieval call binding the contract method 0xc1f0808a.
//
// Solidity: function processed(bytes32 id) view returns(bool)
func (_IETHBridge *IETHBridgeCaller) Processed(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _IETHBridge.contract.Call(opts, &out, "processed", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Processed is a free data retrieval call binding the contract method 0xc1f0808a.
//
// Solidity: function processed(bytes32 id) view returns(bool)
func (_IETHBridge *IETHBridgeSession) Processed(id [32]byte) (bool, error) {
	return _IETHBridge.Contract.Processed(&_IETHBridge.CallOpts, id)
}

// Processed is a free data retrieval call binding the contract method 0xc1f0808a.
//
// Solidity: function processed(bytes32 id) view returns(bool)
func (_IETHBridge *IETHBridgeCallerSession) Processed(id [32]byte) (bool, error) {
	return _IETHBridge.Contract.Processed(&_IETHBridge.CallOpts, id)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x70eac629.
//
// Solidity: function cancelDeposit((uint256,address,address,uint256,bytes,bytes,address) ethDeposit, address claimee, bytes data, uint256 height, bytes proof) returns()
func (_IETHBridge *IETHBridgeTransactor) CancelDeposit(opts *bind.TransactOpts, ethDeposit IETHBridgeETHDeposit, claimee common.Address, data []byte, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IETHBridge.contract.Transact(opts, "cancelDeposit", ethDeposit, claimee, data, height, proof)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x70eac629.
//
// Solidity: function cancelDeposit((uint256,address,address,uint256,bytes,bytes,address) ethDeposit, address claimee, bytes data, uint256 height, bytes proof) returns()
func (_IETHBridge *IETHBridgeSession) CancelDeposit(ethDeposit IETHBridgeETHDeposit, claimee common.Address, data []byte, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IETHBridge.Contract.CancelDeposit(&_IETHBridge.TransactOpts, ethDeposit, claimee, data, height, proof)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x70eac629.
//
// Solidity: function cancelDeposit((uint256,address,address,uint256,bytes,bytes,address) ethDeposit, address claimee, bytes data, uint256 height, bytes proof) returns()
func (_IETHBridge *IETHBridgeTransactorSession) CancelDeposit(ethDeposit IETHBridgeETHDeposit, claimee common.Address, data []byte, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IETHBridge.Contract.CancelDeposit(&_IETHBridge.TransactOpts, ethDeposit, claimee, data, height, proof)
}

// ClaimDeposit is a paid mutator transaction binding the contract method 0x2b6ab0d3.
//
// Solidity: function claimDeposit((uint256,address,address,uint256,bytes,bytes,address) ethDeposit, uint256 height, bytes proof) returns()
func (_IETHBridge *IETHBridgeTransactor) ClaimDeposit(opts *bind.TransactOpts, ethDeposit IETHBridgeETHDeposit, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IETHBridge.contract.Transact(opts, "claimDeposit", ethDeposit, height, proof)
}

// ClaimDeposit is a paid mutator transaction binding the contract method 0x2b6ab0d3.
//
// Solidity: function claimDeposit((uint256,address,address,uint256,bytes,bytes,address) ethDeposit, uint256 height, bytes proof) returns()
func (_IETHBridge *IETHBridgeSession) ClaimDeposit(ethDeposit IETHBridgeETHDeposit, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IETHBridge.Contract.ClaimDeposit(&_IETHBridge.TransactOpts, ethDeposit, height, proof)
}

// ClaimDeposit is a paid mutator transaction binding the contract method 0x2b6ab0d3.
//
// Solidity: function claimDeposit((uint256,address,address,uint256,bytes,bytes,address) ethDeposit, uint256 height, bytes proof) returns()
func (_IETHBridge *IETHBridgeTransactorSession) ClaimDeposit(ethDeposit IETHBridgeETHDeposit, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IETHBridge.Contract.ClaimDeposit(&_IETHBridge.TransactOpts, ethDeposit, height, proof)
}

// Deposit is a paid mutator transaction binding the contract method 0x0a5f93a6.
//
// Solidity: function deposit(address to, bytes data, bytes context, address canceler) payable returns(bytes32 id)
func (_IETHBridge *IETHBridgeTransactor) Deposit(opts *bind.TransactOpts, to common.Address, data []byte, context []byte, canceler common.Address) (*types.Transaction, error) {
	return _IETHBridge.contract.Transact(opts, "deposit", to, data, context, canceler)
}

// Deposit is a paid mutator transaction binding the contract method 0x0a5f93a6.
//
// Solidity: function deposit(address to, bytes data, bytes context, address canceler) payable returns(bytes32 id)
func (_IETHBridge *IETHBridgeSession) Deposit(to common.Address, data []byte, context []byte, canceler common.Address) (*types.Transaction, error) {
	return _IETHBridge.Contract.Deposit(&_IETHBridge.TransactOpts, to, data, context, canceler)
}

// Deposit is a paid mutator transaction binding the contract method 0x0a5f93a6.
//
// Solidity: function deposit(address to, bytes data, bytes context, address canceler) payable returns(bytes32 id)
func (_IETHBridge *IETHBridgeTransactorSession) Deposit(to common.Address, data []byte, context []byte, canceler common.Address) (*types.Transaction, error) {
	return _IETHBridge.Contract.Deposit(&_IETHBridge.TransactOpts, to, data, context, canceler)
}

// IETHBridgeDepositCancelledIterator is returned from FilterDepositCancelled and is used to iterate over the raw logs and unpacked data for DepositCancelled events raised by the IETHBridge contract.
type IETHBridgeDepositCancelledIterator struct {
	Event *IETHBridgeDepositCancelled // Event containing the contract specifics and raw log

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
func (it *IETHBridgeDepositCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IETHBridgeDepositCancelled)
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
		it.Event = new(IETHBridgeDepositCancelled)
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
func (it *IETHBridgeDepositCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IETHBridgeDepositCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IETHBridgeDepositCancelled represents a DepositCancelled event raised by the IETHBridge contract.
type IETHBridgeDepositCancelled struct {
	Id      [32]byte
	Claimee common.Address
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositCancelled is a free log retrieval operation binding the contract event 0x27935aeb5a2922f0bd79b3897d1fb539bc8f77c281c68b3b1b67edc9a0beb808.
//
// Solidity: event DepositCancelled(bytes32 indexed id, address claimee, bytes data)
func (_IETHBridge *IETHBridgeFilterer) FilterDepositCancelled(opts *bind.FilterOpts, id [][32]byte) (*IETHBridgeDepositCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IETHBridge.contract.FilterLogs(opts, "DepositCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &IETHBridgeDepositCancelledIterator{contract: _IETHBridge.contract, event: "DepositCancelled", logs: logs, sub: sub}, nil
}

// WatchDepositCancelled is a free log subscription operation binding the contract event 0x27935aeb5a2922f0bd79b3897d1fb539bc8f77c281c68b3b1b67edc9a0beb808.
//
// Solidity: event DepositCancelled(bytes32 indexed id, address claimee, bytes data)
func (_IETHBridge *IETHBridgeFilterer) WatchDepositCancelled(opts *bind.WatchOpts, sink chan<- *IETHBridgeDepositCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IETHBridge.contract.WatchLogs(opts, "DepositCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IETHBridgeDepositCancelled)
				if err := _IETHBridge.contract.UnpackLog(event, "DepositCancelled", log); err != nil {
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

// ParseDepositCancelled is a log parse operation binding the contract event 0x27935aeb5a2922f0bd79b3897d1fb539bc8f77c281c68b3b1b67edc9a0beb808.
//
// Solidity: event DepositCancelled(bytes32 indexed id, address claimee, bytes data)
func (_IETHBridge *IETHBridgeFilterer) ParseDepositCancelled(log types.Log) (*IETHBridgeDepositCancelled, error) {
	event := new(IETHBridgeDepositCancelled)
	if err := _IETHBridge.contract.UnpackLog(event, "DepositCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IETHBridgeDepositClaimedIterator is returned from FilterDepositClaimed and is used to iterate over the raw logs and unpacked data for DepositClaimed events raised by the IETHBridge contract.
type IETHBridgeDepositClaimedIterator struct {
	Event *IETHBridgeDepositClaimed // Event containing the contract specifics and raw log

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
func (it *IETHBridgeDepositClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IETHBridgeDepositClaimed)
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
		it.Event = new(IETHBridgeDepositClaimed)
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
func (it *IETHBridgeDepositClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IETHBridgeDepositClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IETHBridgeDepositClaimed represents a DepositClaimed event raised by the IETHBridge contract.
type IETHBridgeDepositClaimed struct {
	Id      [32]byte
	Deposit IETHBridgeETHDeposit
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositClaimed is a free log retrieval operation binding the contract event 0x652f8b73a098034852d56898180b2a5302862923641e9e5c312a79c9e1dcc025.
//
// Solidity: event DepositClaimed(bytes32 indexed id, (uint256,address,address,uint256,bytes,bytes,address) deposit)
func (_IETHBridge *IETHBridgeFilterer) FilterDepositClaimed(opts *bind.FilterOpts, id [][32]byte) (*IETHBridgeDepositClaimedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IETHBridge.contract.FilterLogs(opts, "DepositClaimed", idRule)
	if err != nil {
		return nil, err
	}
	return &IETHBridgeDepositClaimedIterator{contract: _IETHBridge.contract, event: "DepositClaimed", logs: logs, sub: sub}, nil
}

// WatchDepositClaimed is a free log subscription operation binding the contract event 0x652f8b73a098034852d56898180b2a5302862923641e9e5c312a79c9e1dcc025.
//
// Solidity: event DepositClaimed(bytes32 indexed id, (uint256,address,address,uint256,bytes,bytes,address) deposit)
func (_IETHBridge *IETHBridgeFilterer) WatchDepositClaimed(opts *bind.WatchOpts, sink chan<- *IETHBridgeDepositClaimed, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IETHBridge.contract.WatchLogs(opts, "DepositClaimed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IETHBridgeDepositClaimed)
				if err := _IETHBridge.contract.UnpackLog(event, "DepositClaimed", log); err != nil {
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

// ParseDepositClaimed is a log parse operation binding the contract event 0x652f8b73a098034852d56898180b2a5302862923641e9e5c312a79c9e1dcc025.
//
// Solidity: event DepositClaimed(bytes32 indexed id, (uint256,address,address,uint256,bytes,bytes,address) deposit)
func (_IETHBridge *IETHBridgeFilterer) ParseDepositClaimed(log types.Log) (*IETHBridgeDepositClaimed, error) {
	event := new(IETHBridgeDepositClaimed)
	if err := _IETHBridge.contract.UnpackLog(event, "DepositClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IETHBridgeDepositMadeIterator is returned from FilterDepositMade and is used to iterate over the raw logs and unpacked data for DepositMade events raised by the IETHBridge contract.
type IETHBridgeDepositMadeIterator struct {
	Event *IETHBridgeDepositMade // Event containing the contract specifics and raw log

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
func (it *IETHBridgeDepositMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IETHBridgeDepositMade)
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
		it.Event = new(IETHBridgeDepositMade)
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
func (it *IETHBridgeDepositMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IETHBridgeDepositMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IETHBridgeDepositMade represents a DepositMade event raised by the IETHBridge contract.
type IETHBridgeDepositMade struct {
	Id      [32]byte
	Deposit IETHBridgeETHDeposit
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositMade is a free log retrieval operation binding the contract event 0x3354939929a6683d1d35c363a5bfa3f418ad78287919a969d99b89e496b17b74.
//
// Solidity: event DepositMade(bytes32 indexed id, (uint256,address,address,uint256,bytes,bytes,address) deposit)
func (_IETHBridge *IETHBridgeFilterer) FilterDepositMade(opts *bind.FilterOpts, id [][32]byte) (*IETHBridgeDepositMadeIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IETHBridge.contract.FilterLogs(opts, "DepositMade", idRule)
	if err != nil {
		return nil, err
	}
	return &IETHBridgeDepositMadeIterator{contract: _IETHBridge.contract, event: "DepositMade", logs: logs, sub: sub}, nil
}

// WatchDepositMade is a free log subscription operation binding the contract event 0x3354939929a6683d1d35c363a5bfa3f418ad78287919a969d99b89e496b17b74.
//
// Solidity: event DepositMade(bytes32 indexed id, (uint256,address,address,uint256,bytes,bytes,address) deposit)
func (_IETHBridge *IETHBridgeFilterer) WatchDepositMade(opts *bind.WatchOpts, sink chan<- *IETHBridgeDepositMade, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IETHBridge.contract.WatchLogs(opts, "DepositMade", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IETHBridgeDepositMade)
				if err := _IETHBridge.contract.UnpackLog(event, "DepositMade", log); err != nil {
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

// ParseDepositMade is a log parse operation binding the contract event 0x3354939929a6683d1d35c363a5bfa3f418ad78287919a969d99b89e496b17b74.
//
// Solidity: event DepositMade(bytes32 indexed id, (uint256,address,address,uint256,bytes,bytes,address) deposit)
func (_IETHBridge *IETHBridgeFilterer) ParseDepositMade(log types.Log) (*IETHBridgeDepositMade, error) {
	event := new(IETHBridgeDepositMade)
	if err := _IETHBridge.contract.UnpackLog(event, "DepositMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
